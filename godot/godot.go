package godot

import (
	"fmt"
	"github.com/pinzolo/casee"
	"github.com/shadowapex/godot-go/gdnative"
	"log"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

var debug = false

// EnableDebug will enable debug logging of the godot library.
func EnableDebug() {
	debug = true
}

// Init is a special Go function that will be called upon library initialization.
// This is also the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes.
func init() {
	// Configure GDNative to use our own NativeScript init function.
	gdnative.SetNativeScriptInit(
		configureLogging,
		registerClasses,
		autoRegisterClasses,
	)
}

// configureLogging will set up the Go logger to output to the Godot console log.
func configureLogging() {
	// Configure logging.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
}

// registerClasses will loop through all classes to register and register them
// with Godot. It is the caller's responsibility for registering properties and
// methods on this class.
func registerClasses() {
	for classString, constructor := range godotConstructorsToRegister {
		if debug {
			log.Println("Registering class:", classString)
		}
		// Call the constructor to get the BaseClass
		class := constructor()

		// Get the type of the class so we can register it.
		classType := reflect.TypeOf(class)

		// Create a registered class structure that will hold information about the
		// cass and its methods.
		regClass := newRegisteredClass(classType)

		// Register our class in our Go registry.
		classRegistry[classString] = regClass

		// Set up our constructor and destructor function structs.
		createFunc := createConstructor(classString, constructor)
		destroyFunc := createDestructor(classString)

		// Register our class with Godot.
		gdnative.NativeScript.RegisterClass(classString, class.BaseClass(), createFunc, destroyFunc)
	}
}

// autoRegisterClasses will automatically inspect all given classes to register
// and use reflection to find exported methods and properties and register them
// with Godot.
func autoRegisterClasses() {
	log.Println("Discovering classes to register with Godot...")

	// Loop through our registered classes and register them with the Godot API.
	for _, constructor := range godotConstructorsToAutoRegister {
		// Use the constructor to build a class to inspect the given structure.
		if debug {
			log.Println("Calling constructor to inspect object with reflection...")
		}
		class := constructor()

		// Get the type of the given struct, and get its name as a string
		classType := reflect.TypeOf(class)
		classString := strings.Replace(classType.String(), "*", "", 1)
		if debug {
			log.Println("Registering class:", classString)
		}

		// Create a registered class structure that will hold information about the
		// cass and its methods.
		regClass := newRegisteredClass(classType)

		// Call the "BaseClass" method on the class to get the base class.
		baseClass := class.(Class).BaseClass()
		if debug {
			log.Println("  Using Base Class:", baseClass)
		}

		// Set up our constructor and destructor function structs.
		createFunc := createConstructor(classString, constructor)
		destroyFunc := createDestructor(classString)

		// Register our class with Godot.
		gdnative.NativeScript.RegisterClass(classString, baseClass, createFunc, destroyFunc)

		// Loop through our class's struct fields. We do this to register properties as well
		// as find the embedded parent struct to ensure we don't register those methods.
		if debug {
			log.Println("  Looking at struct fields:")
			log.Println("    Found", classType.Elem().NumField(), "struct fields.")
		}
		for i := 0; i < classType.Elem().NumField(); i++ {
			classField := classType.Elem().Field(i)

			// Skip anonymously embedded fields
			if classField.Anonymous {
				continue
			}

			// Check if the field is private
			firstChar := []rune(classField.Name)[0]
			if unicode.ToLower(firstChar) == firstChar {
				continue
			}

			if debug {
				log.Println("  Found field:", classField.Name)
				log.Println("    Type:", classField.Type.String())
				log.Println("    Anonymous:", classField.Anonymous)
				log.Println("    Package Path:", classField.PkgPath)
			}

			// Check to see if this field is a signal.
			if classField.Type.String() == "godot.Signal" {
				if debug {
					log.Println("  Field is a signal type. Registering as signal.")
				}

				// Get the signal field
				classValue := reflect.ValueOf(class)
				signalField := classValue.Elem().FieldByName(classField.Name)
				signalValue := signalField.Interface().(Signal)

				// Construct our GDNative Signal struct from the discovered
				// signal object.
				signal := &gdnative.Signal{}
				signal.Name = gdnative.String(signalValue.Name)
				signal.NumArgs = gdnative.Int(len(signalValue.Args))
				signal.NumDefaultArgs = gdnative.Int(len(signalValue.DefaultArgs))
				signal.Args = []gdnative.SignalArgument{}
				signal.DefaultArgs = []gdnative.Variant{}

				// Construct the arguments for our GDNative Signal struct
				for _, argValue := range signalValue.Args {
					var arg gdnative.SignalArgument
					arg.Name = argValue.Name
					arg.Hint = argValue.Hint
					arg.HintString = argValue.HintString
					arg.Usage = argValue.Usage

					// Get the type of the arg value
					argType := reflect.TypeOf(argValue.Value)
					if isGodotClass(argType) {
						arg.Type = gdnative.Int(gdnative.VariantTypeObject)
					} else {
						arg.Type = gdnative.Int(VariantTypeToConstant(argType))
					}

					signal.Args = append(signal.Args, arg)
				}

				// Construct the default arguments for our GDNative Signal struct
				for _, argValue := range signalValue.DefaultArgs {
					var variant gdnative.Variant

					// Get the type of the arg value
					argType := reflect.TypeOf(argValue)
					if isGodotClass(argType) {
						obj := argValue.(Class).GetBaseObject()
						variant = gdnative.NewVariantObject(obj)
					} else {
						variant = GoTypeToVariant(reflect.ValueOf(argValue))
					}

					signal.DefaultArgs = append(signal.DefaultArgs, variant)
				}

				gdnative.NativeScript.RegisterSignal(classString, signal)
				continue
			}

			// Create our property getter/setter structs that we will register with Godot.
			setPropertyFunc := createPropertySetter(classString, classField.Name, classField.Type)
			getPropertyFunc := createPropertyGetter(classString, classField.Name, classField.Type)
			propertyAttrs := createPropertyAttributes(classField)

			// Register the public property with Godot.
			gdnative.NativeScript.RegisterProperty(
				classString,
				classField.Name,
				propertyAttrs,
				setPropertyFunc,
				getPropertyFunc,
			)
		}

		// Loop through our class's methods that are attached to it.
		if debug {
			log.Println("  Looking at methods:")
			log.Println("    Found", classType.NumMethod(), "methods")
		}
		for i := 0; i < classType.NumMethod(); i++ {
			classMethod := classType.Method(i)

			// TODO: For now we are only checking if the given method is embedded or
			// not. If the method comes from an embedded structure, skip it.
			// We need to figure this shit out, so we can allow embedding of non-godot
			// types.
			runMethod := runtime.FuncForPC(classMethod.Func.Pointer())
			filename, _ := runMethod.FileLine(classMethod.Func.Pointer())
			if strings.Contains(filename, "<autogenerated>") {
				continue
			}

			if debug {
				log.Println("  Found method:", classMethod.Name)
				log.Println("    Method in package path:", classMethod.PkgPath)
				log.Println("    Type package path:", classMethod.Type.PkgPath())
				log.Println("    Type:", classMethod.Type.String())
				log.Println("    Kind:", classMethod.Type.Kind().String())
			}

			// Construct a registered method structure that inspects all of the
			// arguments and return types.
			regMethod := newRegisteredMethod(classMethod)
			regClass.addMethod(classMethod.Name, regMethod)
			if debug {
				log.Println("    Method Arguments:", len(regMethod.arguments))
				log.Println("    Method Arguments:", regMethod.arguments)
				log.Println("    Method Returns:", len(regMethod.returns))
				log.Println("    Method Returns:", regMethod.returns)
			}

			// Skip the method if its a Class interface method
			skip := false
			for _, exclude := range []string{"BaseClass", "GetBaseObject", "SetBaseObject"} {
				if classMethod.Name == exclude {
					skip = true
				}
			}
			if skip {
				continue
			}

			// Look at the method name to see if it starts with "X_". If it does, we need to
			// replace it with an underscore. This is required because Go method visibility
			// is done through case sensitivity. Since Godot private methods start with an
			// "_" character.
			goMethodName := classMethod.Name
			godotMethodName := toGodotMethodName(goMethodName)

			// Set up our method structure
			method := createMethod(classString, goMethodName)
			attributes := &gdnative.MethodAttributes{
				RPCType: gdnative.MethodRpcModeDisabled,
			}

			// Register the method.
			gdnative.NativeScript.RegisterMethod(classString, godotMethodName, attributes, method)
		}

		// Register our class in our Go registry.
		classRegistry[classString] = regClass

	}
}

// CreateConstructor will create the InstanceCreateFunc structure with the given class name
// and constructor. This structure can be used when registering a class with Godot.
func createConstructor(classString string, constructor ClassConstructor) *gdnative.InstanceCreateFunc {
	var createFunc gdnative.InstanceCreateFunc
	createFunc.CreateFunc = func(object gdnative.Object, methodData string) string {
		// Create a new instance of the object.
		class := constructor()
		if debug {
			log.Println("Created new object instance:", class, "with instance address:", object.ID())
		}

		// Add the Godot object pointer to the class structure.
		class.SetBaseObject(object)

		// Add the instance to our instance registry.
		InstanceRegistry.Add(object.ID(), class)

		// Return the instance string. This will be passed to the method function as userData, so we
		// can look up the instance in our registry.
		return object.ID()

	}
	createFunc.MethodData = classString
	createFunc.FreeFunc = func(methodData string) {}

	return &createFunc
}

// CreateDestructor will create the InstanceDestroyFunc structure with the given class name.
// This structure can be used when registering a class with Godot.
func createDestructor(classString string) *gdnative.InstanceDestroyFunc {
	var destroyFunc gdnative.InstanceDestroyFunc
	destroyFunc.DestroyFunc = func(object gdnative.Object, className, instanceID string) {
		if debug {
			log.Println("Destroying object instance:", className, "with instance address:", object.ID())
		}

		// Unregister it from our InstanceRegistry so it can be garbage collected.
		InstanceRegistry.Delete(instanceID)
	}
	destroyFunc.MethodData = classString
	destroyFunc.FreeFunc = func(methodData string) {}

	return &destroyFunc
}

// CreateMethod will create the InstanceMethod structure. This will be called whenever
// a Godot method is called.
func createMethod(classString, methodString string) *gdnative.InstanceMethod {
	var methodFunc gdnative.InstanceMethod
	methodFunc.Method = func(object gdnative.Object, classMethod, instanceString string, numArgs int, args []gdnative.Variant) gdnative.Variant {
		var ret gdnative.Variant

		// Get the object instance based on the instance string given in userData.
		class, ok := InstanceRegistry.Get(instanceString)
		if !ok {
			panic("Method " + classMethod + " was called on instance (" + instanceString + "), but does not exist in the instance registry!")
		}
		classValue := reflect.ValueOf(class)

		if debug {
			log.Println("Method was called!")
			log.Println("  godotObject:", object)
			log.Println("  numArgs:", numArgs)
			log.Println("  args:", args)
			log.Println("  instance:", class)
			log.Println("  methodString (methodData):", classMethod)
			log.Println("  instanceString (userData):", instanceString)
		}

		// Create a slice of Godot Variant arguments
		goArgsSlice := []reflect.Value{}

		// If we have arguments, append the first argument.
		for _, arg := range args {
			// Convert the variant into its base type
			goArgsSlice = append(goArgsSlice, VariantToGoType(arg))
		}

		// Use the method string to get the class name and method name.
		if debug {
			log.Println("  Getting class name and method name...")
		}
		classMethodSlice := strings.Split(classMethod, "::")
		className := classMethodSlice[0]
		methodName := classMethodSlice[1]
		if debug {
			log.Println("    Class Name: ", className)
			log.Println("    Method Name: ", methodName)
		}

		// Look up the registered class so we can find out how many arguments it takes
		// and their types.
		if debug {
			log.Println("  Look up the registered class and its method")
			log.Println("  Registered classes:", classRegistry)
		}
		regClass := classRegistry[className]
		if regClass == nil {
			log.Fatal("  This class has not been registered! Class name: ", className, " Method name: ", methodName)
		}
		if debug {
			log.Println("  Looked up class:", regClass)
			log.Println("  Methods in class:", regClass.methods)
		}
		regMethod := regClass.methods[methodName]

		if debug {
			log.Println("  Registered method arguments:", regMethod.arguments)
			log.Println("  Arguments to pass:", goArgsSlice)
		}

		// Check to ensure the method has the same number of arguments we expect
		if len(regMethod.arguments)-1 != int(numArgs) {
			gdnative.Log.Error("Invalid number of arguments. Expected ", numArgs, " arguments. (Got ", len(regMethod.arguments), ")")
			panic("Invalid number of arguments.")
		}

		// Get the value of the class, so we can call methods on it.
		method := classValue.MethodByName(methodName)
		rawRet := method.Call(goArgsSlice)
		if debug {
			log.Println("Got raw return value after method call:", rawRet)
		}

		// Check to see if this returns anything.
		if len(rawRet) == 0 {
			nilReturn := gdnative.NewVariantNil()
			return nilReturn
		}

		// Convert our returned value into a Godot Variant.
		rawRetInterface := rawRet[0].Interface()
		switch regMethod.returns[0].String() {
		case "bool":
			base := gdnative.Bool(rawRetInterface.(bool))
			variant := gdnative.NewVariantBool(base)
			ret = variant
		case "int64":
			base := gdnative.Int64T(rawRetInterface.(int64))
			variant := gdnative.NewVariantInt(base)
			ret = variant
		case "int32":
			base := gdnative.Int64T(rawRetInterface.(int32))
			variant := gdnative.NewVariantInt(base)
			ret = variant
		case "int":
			base := gdnative.Int64T(rawRetInterface.(int))
			variant := gdnative.NewVariantInt(base)
			ret = variant
		case "uint64":
			base := gdnative.Uint64T(rawRetInterface.(uint64))
			variant := gdnative.NewVariantUint(base)
			ret = variant
		case "uint32":
			base := gdnative.Uint64T(rawRetInterface.(uint32))
			variant := gdnative.NewVariantUint(base)
			ret = variant
		case "uint":
			base := gdnative.Uint64T(rawRetInterface.(uint))
			variant := gdnative.NewVariantUint(base)
			ret = variant
		case "float64":
			base := gdnative.Double(rawRetInterface.(float64))
			variant := gdnative.NewVariantReal(base)
			ret = variant
		case "string":
			base := gdnative.WcharT(rawRetInterface.(string))
			baseStr := base.AsString()
			variant := gdnative.NewVariantString(baseStr)
			ret = variant
		default:
			panic("The return was not valid. Should be Godot Variant or built-in Go type. Received: " + regMethod.returns[0].String())
		}

		return ret
	}
	methodFunc.MethodData = classString + "::" + methodString
	methodFunc.FreeFunc = func(methodData string) {}

	return &methodFunc
}

// CreatePropertySetter will create the InstancePropertySetFunc structure. This will be called whenever
// Godot needs to set a property on an instance.
func createPropertySetter(classString, propertyString string, propertyType reflect.Type) *gdnative.InstancePropertySet {
	var propertySetFunc gdnative.InstancePropertySet
	propertySetFunc.SetFunc = func(object gdnative.Object, classProperty, instanceString string, property gdnative.Variant) {
		// Get the object instance based on the instance string given in userData.
		class, ok := InstanceRegistry.Get(instanceString)
		if !ok {
			panic("Set property " + classProperty + " was called on instance (" + instanceString + "), but does not exist in the instance registry!")
		}

		// Get the actual class value and the struct field of the property.
		classValue := reflect.ValueOf(class)
		propertyField := classValue.Elem().FieldByName(propertyString)

		// Check to see what kind of type this is. If it is a Godot class,
		// we need to convert our variant into an object.
		if isGodotClass(propertyType) {

			// Create the variant as an object so we can get its class type.
			propAsObject := &Object{}
			propAsObject.SetBaseObject(property.AsObject())
			typeString := propAsObject.GetClass()

			// Get the actual class object.
			obj := getActualClass(typeString, propAsObject.GetBaseObject())
			if debug {
				log.Println("Setting property '" + classString + "." + propertyString + "' (" + instanceString + ") with type '" + string(typeString) + "' as Object with ID (" + obj.GetBaseObject().ID() + ")")
			}
			propertyField.Set(reflect.ValueOf(obj))
			return
		}

		// Otherwise, this should be a base variant type.
		if debug {
			log.Println("Setting property '" + classString + "." + propertyString + "' as a Variant on instance (" + instanceString + ")")
		}
		value := VariantToGoType(property)
		propertyField.Set(value)
	}

	// Set the method data. This will be passed to the SetFunc function we defined above
	// as the "classProperty"
	propertySetFunc.MethodData = classString + "::" + propertyString
	propertySetFunc.FreeFunc = func(methodData string) {}

	return &propertySetFunc
}

// CreatePropertyGetter will create the InstancePropertyGet structure. This will be called whenever
// Godot needs to get a property on an instance.
func createPropertyGetter(classString, propertyString string, propertyType reflect.Type) *gdnative.InstancePropertyGet {
	var propertyGetFunc gdnative.InstancePropertyGet
	propertyGetFunc.GetFunc = func(object gdnative.Object, classProperty, instanceString string) gdnative.Variant {
		// Get the object instance based on the instance string given in userData.
		class, ok := InstanceRegistry.Get(instanceString)
		if !ok {
			panic("Get property " + classProperty + " was called on instance (" + instanceString + "), but does not exist in the instance registry!")
		}
		classValue := reflect.ValueOf(class)
		propertyField := classValue.Elem().FieldByName(classProperty)

		// Check to see what kind of type this is. If it is a Godot class,
		// we need to convert our object into a variant.
		if isGodotClass(propertyType) {
			if debug {
				log.Println("Getting property '" + classString + "." + propertyString + "' as Object on instance (" + instanceString + ")")
			}

			// Get the value of the field as an interface.
			property := propertyField.Interface()

			object := property.(ObjectImplementer).GetBaseObject()
			return gdnative.NewVariantObject(object)
		}

		// Otherwise, convert the property to a base variant.
		if debug {
			log.Println("Getting property '" + classString + "." + propertyString + "' as a Variant on instance (" + instanceString + ")")
		}
		return GoTypeToVariant(propertyField)
	}
	propertyGetFunc.MethodData = classString + "::" + propertyString
	propertyGetFunc.FreeFunc = func(methodData string) {}

	return &propertyGetFunc
}

func createPropertyAttributes(field reflect.StructField) *gdnative.PropertyAttributes {
	// Create our property attributes struct that we will fill.
	var propertyAttrs gdnative.PropertyAttributes

	// Set the default value to nil.
	// TODO: Figure out a way to let the user define a default value.
	propertyAttrs.DefaultValue = gdnative.NewVariantNil()

	// Inspect the struct field for any tags. We will use this for setting the
	// usage, hint, hint string, etc. If none are found, defaults will be used.
	if hintStr, ok := field.Tag.Lookup("hint_string"); ok {
		propertyAttrs.HintString = gdnative.String(hintStr)
	} else {
		propertyAttrs.HintString = ""
	}

	// rset_type
	if rsetType, ok := field.Tag.Lookup("rset_type"); ok {
		rsetTypeStr := "MethodRpcMode" + rsetType
		if propertyAttrs.RsetType, ok = gdnative.MethodRpcModeLookupMap[rsetTypeStr]; !ok {
			validTypes := ""
			for key, _ := range gdnative.MethodRpcModeLookupMap {
				validTypes += " " + strings.Replace(key, "MethodRpcMode", "", 1)
			}
			panic("rset_type must be one of the following: " + validTypes)
		}
	} else {
		propertyAttrs.RsetType = gdnative.MethodRpcModeDisabled
	}

	// usage
	if usage, ok := field.Tag.Lookup("usage"); ok {
		usageStr := "PropertyUsage" + usage
		if propertyAttrs.Usage, ok = gdnative.PropertyUsageFlagsLookupMap[usageStr]; !ok {
			validTypes := ""
			for key, _ := range gdnative.PropertyUsageFlagsLookupMap {
				validTypes += " " + strings.Replace(key, "PropertyUsage", "", 1)
			}
			panic("usage must be one of the following: " + validTypes)
		}
	} else {
		propertyAttrs.Usage = gdnative.PropertyUsageDefault
	}

	// hint
	if hint, ok := field.Tag.Lookup("hint"); ok {
		hintStr := "PropertyHint" + hint
		if propertyAttrs.Hint, ok = gdnative.PropertyHintLookupMap[hintStr]; !ok {
			validTypes := ""
			for key, _ := range gdnative.PropertyHintLookupMap {
				validTypes += " " + strings.Replace(key, "PropertyHint", "", 1)
			}
			panic("hint must be one of the following: " + validTypes)
		}
	} else {
		propertyAttrs.Hint = gdnative.PropertyHintNone
	}

	// Check to see if this field is a Godot object.
	if isGodotClass(field.Type) {
		propertyAttrs.Type = gdnative.Int(gdnative.VariantTypeObject)
		return &propertyAttrs
	}

	// Otherwise, this should be a godot variant type
	propertyAttrs.Type = gdnative.Int(VariantTypeToConstant(field.Type))

	return &propertyAttrs
}

// VariantToGoType will check the given variant type and convert it to its
// actual type. The value is returned as a reflect.Value.
func VariantToGoType(variant gdnative.Variant) reflect.Value {
	switch variant.GetType() {
	case gdnative.VariantTypeNil:
		return reflect.ValueOf(nil)
	case gdnative.VariantTypeBool:
		return reflect.ValueOf(variant.AsBool())
	case gdnative.VariantTypeInt:
		return reflect.ValueOf(variant.AsInt())
	case gdnative.VariantTypeReal:
		return reflect.ValueOf(gdnative.Real(variant.AsReal()))
	case gdnative.VariantTypeString:
		return reflect.ValueOf(variant.AsString())
	case gdnative.VariantTypeVector2:
		return reflect.ValueOf(variant.AsVector2())
	case gdnative.VariantTypeRect2:
		return reflect.ValueOf(variant.AsRect2())
	case gdnative.VariantTypeVector3:
		return reflect.ValueOf(variant.AsVector3())
	case gdnative.VariantTypeTransform2D:
		return reflect.ValueOf(variant.AsTransform2D())
	case gdnative.VariantTypePlane:
		return reflect.ValueOf(variant.AsPlane())
	case gdnative.VariantTypeQuat:
		return reflect.ValueOf(variant.AsQuat())
	case gdnative.VariantTypeAabb:
		return reflect.ValueOf(variant.AsAabb())
	case gdnative.VariantTypeBasis:
		return reflect.ValueOf(variant.AsBasis())
	case gdnative.VariantTypeTransform:
		return reflect.ValueOf(variant.AsTransform())
	case gdnative.VariantTypeColor:
		return reflect.ValueOf(variant.AsColor())
	case gdnative.VariantTypeNodePath:
		return reflect.ValueOf(variant.AsNodePath())
	case gdnative.VariantTypeRid:
		return reflect.ValueOf(variant.AsRid())
	case gdnative.VariantTypeObject:
		return reflect.ValueOf(variant.AsObject())
	case gdnative.VariantTypeDictionary:
		return reflect.ValueOf(variant.AsDictionary())
	case gdnative.VariantTypeArray:
		return reflect.ValueOf(variant.AsArray())
	case gdnative.VariantTypePoolByteArray:
		return reflect.ValueOf(variant.AsPoolByteArray())
	case gdnative.VariantTypePoolIntArray:
		return reflect.ValueOf(variant.AsPoolIntArray())
	case gdnative.VariantTypePoolRealArray:
		return reflect.ValueOf(variant.AsPoolRealArray())
	case gdnative.VariantTypePoolStringArray:
		return reflect.ValueOf(variant.AsPoolStringArray())
	case gdnative.VariantTypePoolVector2Array:
		return reflect.ValueOf(variant.AsPoolVector2Array())
	case gdnative.VariantTypePoolVector3Array:
		return reflect.ValueOf(variant.AsPoolVector3Array())
	case gdnative.VariantTypePoolColorArray:
		return reflect.ValueOf(variant.AsPoolColorArray())
	}
	panic("Unknown type of godot variant argument: " + fmt.Sprint(variant.GetType()))
}

// GoTypeToVariant will check the given Go type and convert it to its
// Variant type. The value is returned as a gdnative.Variant.
func GoTypeToVariant(value reflect.Value) gdnative.Variant {
	valueInterface := value.Interface()
	switch v := valueInterface.(type) {
	case gdnative.Bool:
		return gdnative.NewVariantBool(v)
	case gdnative.Int:
		return gdnative.NewVariantInt(gdnative.Int64T(v))
	case gdnative.Int64T:
		return gdnative.NewVariantInt(v)
	case gdnative.Double:
		return gdnative.NewVariantReal(v)
	case gdnative.Real:
		return gdnative.NewVariantReal(gdnative.Double(v))
	case gdnative.String:
		return gdnative.NewVariantString(v)
	case gdnative.Vector2:
		return gdnative.NewVariantVector2(v)
	case gdnative.Rect2:
		return gdnative.NewVariantRect2(v)
	case gdnative.Vector3:
		return gdnative.NewVariantVector3(v)
	case gdnative.Transform2D:
		return gdnative.NewVariantTransform2D(v)
	case gdnative.Plane:
		return gdnative.NewVariantPlane(v)
	case gdnative.Quat:
		return gdnative.NewVariantQuat(v)
	case gdnative.Aabb:
		return gdnative.NewVariantAabb(v)
	case gdnative.Basis:
		return gdnative.NewVariantBasis(v)
	case gdnative.Transform:
		return gdnative.NewVariantTransform(v)
	case gdnative.Color:
		return gdnative.NewVariantColor(v)
	case gdnative.NodePath:
		return gdnative.NewVariantNodePath(v)
	case gdnative.Rid:
		return gdnative.NewVariantRid(v)
	case gdnative.Object:
		return gdnative.NewVariantObject(v)
	case gdnative.Dictionary:
		return gdnative.NewVariantDictionary(v)
	case gdnative.Array:
		return gdnative.NewVariantArray(v)
	case gdnative.PoolByteArray:
		return gdnative.NewVariantPoolByteArray(v)
	case gdnative.PoolIntArray:
		return gdnative.NewVariantPoolIntArray(v)
	case gdnative.PoolRealArray:
		return gdnative.NewVariantPoolRealArray(v)
	case gdnative.PoolStringArray:
		return gdnative.NewVariantPoolStringArray(v)
	case gdnative.PoolVector2Array:
		return gdnative.NewVariantPoolVector2Array(v)
	case gdnative.PoolVector3Array:
		return gdnative.NewVariantPoolVector3Array(v)
	case gdnative.PoolColorArray:
		return gdnative.NewVariantPoolColorArray(v)
	}
	panic("Unknown type of godot argument: " + value.String())
}

// VariantTypeToConstant will check the given field to see what kind of variant
// it is and return its type as a VariantType int. This will panic if the type
// is not a valid Godot type.
func VariantTypeToConstant(t reflect.Type) gdnative.VariantType {
	switch t.String() {
	case "gdnative.Bool":
		return gdnative.VariantTypeBool
	case "gdnative.Int":
		return gdnative.VariantTypeInt
	case "gdnative.Real":
		return gdnative.VariantTypeReal
	case "gdnative.String":
		return gdnative.VariantTypeString
	case "gdnative.Vector2":
		return gdnative.VariantTypeVector2
	case "gdnative.Rect2":
		return gdnative.VariantTypeRect2
	case "gdnative.Vector3":
		return gdnative.VariantTypeVector3
	case "gdnative.Transform2D":
		return gdnative.VariantTypeTransform2D
	case "gdnative.Plane":
		return gdnative.VariantTypePlane
	case "gdnative.Quat":
		return gdnative.VariantTypeQuat
	case "gdnative.Aabb":
		return gdnative.VariantTypeAabb
	case "gdnative.Basis":
		return gdnative.VariantTypeBasis
	case "gdnative.Transform":
		return gdnative.VariantTypeTransform
	case "gdnative.Color":
		return gdnative.VariantTypeColor
	case "gdnative.NodePath":
		return gdnative.VariantTypeNodePath
	case "gdnative.Rid":
		return gdnative.VariantTypeRid
	case "gdnative.Object":
		return gdnative.VariantTypeObject
	case "gdnative.Dictionary":
		return gdnative.VariantTypeDictionary
	case "gdnative.Array":
		return gdnative.VariantTypeArray
	case "gdnative.PoolByteArray":
		return gdnative.VariantTypePoolByteArray
	case "gdnative.PoolIntArray":
		return gdnative.VariantTypePoolIntArray
	case "gdnative.PoolRealArray":
		return gdnative.VariantTypePoolRealArray
	case "gdnative.PoolStringArray":
		return gdnative.VariantTypePoolStringArray
	case "gdnative.PoolVector2Array":
		return gdnative.VariantTypePoolVector2Array
	case "gdnative.PoolVector3Array":
		return gdnative.VariantTypePoolVector3Array
	case "gdnative.PoolColorArray":
		return gdnative.VariantTypePoolColorArray
	}
	if strings.HasPrefix(t.String(), "godot.") {
		panic("Unknown type of exported godot field: " + t.String() + ". You probably need to use *" + t.String() + " or " + t.String() + "Implementer for this field.")
	}
	panic("Unknown type of exported godot field: " + t.String())
}

// isGodotClass will check to see if the given type implements the ObjectImplementer
// interface.
func isGodotClass(t reflect.Type) bool {
	objectImplType := reflect.TypeOf((*ObjectImplementer)(nil)).Elem()
	if t.Implements(objectImplType) {
		return true
	}
	return false
}

// toGoMethodName will take the given Godot method name in snake_case and convert it
// to a CamelCase method name.
func toGoMethodName(methodName string) string {
	goMethodName := casee.ToPascalCase(methodName)
	if strings.HasPrefix(methodName, "_") {
		goMethodName = "X_" + goMethodName
	}

	if debug {
		log.Println("    Godot method name:", methodName)
		log.Println("    Go mapped method name:", goMethodName)
	}

	return goMethodName
}

// toGodotMethodName will take the given Go method name in CamelCase and convert it
// to a snake_case method name.
func toGodotMethodName(goMethodName string) string {
	methodName := goMethodName
	privatePrefix := ""
	if strings.HasPrefix(goMethodName, "X_") {
		methodName = strings.Replace(methodName, "X_", "_", 1)
		privatePrefix = "_"
	}
	methodName = casee.ToSnakeCase(methodName)
	methodName = privatePrefix + methodName

	// handle signal binding functions that start with _on_
	if strings.HasPrefix(methodName, "_on_") {
		// Grab the character after _on_
		runes := []rune(methodName)
		runes[4] = unicode.ToUpper(runes[4])
		methodName = string(runes)
	}

	if debug {
		log.Println("    Go method name:", goMethodName)
		log.Println("    Godot mapped method name:", methodName)
	}

	return methodName
}

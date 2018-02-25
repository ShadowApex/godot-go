// Package gdnative provides a wrapper around Godot's nativescript
// extension. It exists to provide a way to use Go as an alternative scripting
// language from GDScript.
package gdnative

/*
#include <nativescript/godot_nativescript.h>
#include "nativescript.gen.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

type PropertyUsageFlags int

func (p PropertyUsageFlags) getBase() C.godot_property_usage_flags {
	return C.godot_property_usage_flags(p)
}

const (
	PropertyUsageStorage             PropertyUsageFlags = 1
	PropertyUsageEditor              PropertyUsageFlags = 2
	PropertyUsageNetwork             PropertyUsageFlags = 4
	PropertyUsageEditorHelper        PropertyUsageFlags = 8
	PropertyUsageCheckable           PropertyUsageFlags = 16  //used for editing global variables
	PropertyUsageChecked             PropertyUsageFlags = 32  //used for editing global variables
	PropertyUsageInternationalized   PropertyUsageFlags = 64  //hint for internationalized strings
	PropertyUsageGroup               PropertyUsageFlags = 128 //used for grouping props in the editor
	PropertyUsageCategory            PropertyUsageFlags = 256
	PropertyUsageStoreIfNonZero      PropertyUsageFlags = 512  //only store if nonzero
	PropertyUsageStoreIfNonOne       PropertyUsageFlags = 1024 //only store if false
	PropertyUsageNoInstanceState     PropertyUsageFlags = 2048
	PropertyUsageRestartIfChanged    PropertyUsageFlags = 4096
	PropertyUsageScriptVariable      PropertyUsageFlags = 8192
	PropertyUsageStoreIfNull         PropertyUsageFlags = 16384
	PropertyUsageAnimateAsTrigger    PropertyUsageFlags = 32768
	PropertyUsageUpdateAllIfModified PropertyUsageFlags = 65536

	PropertyUsageDefault     PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork
	PropertyUsageDefaultIntl PropertyUsageFlags = PropertyUsageStorage | PropertyUsageEditor | PropertyUsageNetwork | PropertyUsageInternationalized
	PropertyUsageNoEditor    PropertyUsageFlags = PropertyUsageStorage | PropertyUsageNetwork
)

// CreateFunc will be called when we need to create a new instance of a class.
// Takes the instance object, user data.
type CreateFunc func(Object, interface{})

// DestroyFunc will be called when the object is destroyed. Takes the instance
// object, method data, user data.
type DestroyFunc func(Object, interface{}, interface{})

// FreeFunc will be called when we should free the instance from memory.
type FreeFunc func(interface{})

type InstanceCreateFunc struct {
	base       C.godot_instance_create_func
	CreateFunc CreateFunc
	MethodData interface{}
	FreeFunc   FreeFunc
}

func (i *InstanceCreateFunc) getBase() C.godot_instance_create_func {
	return i.base
}

type InstanceDestroyFunc struct {
	base        C.godot_instance_destroy_func
	DestroyFunc DestroyFunc
	MethodData  interface{}
	FreeFunc    FreeFunc
}

func (i *InstanceDestroyFunc) getBase() C.godot_instance_destroy_func {
	return i.base
}

// NativeScript is a wrapper for the NativeScriptAPI.
var NativeScript = &nativeScript{}

// nativeScript is a structure that wraps the NativeScriptAPI methods.
type nativeScript struct {
	api *C.godot_gdnative_ext_nativescript_api_struct
}

// RegisterClass will register the given class with Godot. This will make it
// available to be attached to a Node in Godot. The name of the class that you
// provide here will be the name that you specify when you attach a NativeScript
// script to a Node. The base parameter that you specify will be what the class
// should be inheriting from (e.g. Node2D, Node, etc.).
//
// The create and destroy function parameters are C structs that contain
// function pointers to C methods that will be called when the object is created
// or destroyed. Because of the pointer passing rules, Go code can not pass a
// function value directly to C, so a gateway function should be used. More
// information on using a gateway function can be found here:
// https://github.com/golang/go/wiki/cgo#function-variables
func (n *nativeScript) RegisterClass(name, base string, createFunc C.godot_instance_create_func, destroyFunc C.godot_instance_destroy_func) {
	C.go_godot_nativescript_register_class(
		n.api,
		handle,
		C.CString(name),
		C.CString(base),
		createFunc,
		destroyFunc,
	)
}

// RegisterToolClass will register the given class with Godot as a tool. Refer to
// the 'RegisterClass' method for more information on how to use this.
func (n *nativeScript) RegisterToolClass(name, base string, createFunc C.godot_instance_create_func, destroyFunc C.godot_instance_destroy_func) {
	C.go_godot_nativescript_register_tool_class(
		n.api,
		handle,
		C.CString(name),
		C.CString(base),
		createFunc,
		destroyFunc,
	)
}

// RegisterMethod will register the given method with Godot and associate it with
// the given class name. The name parameter is the name of the class you wish to
// attach this method to. The funcName parameter is the name of the function you
// want to register. The attributes and method are what will actually be called
// when Godot calls the method on the object.
func (n *nativeScript) RegisterMethod(name, funcName string, attributes C.godot_method_attributes, method C.godot_instance_method) {
	C.go_godot_nativescript_register_method(
		n.api,
		handle,
		C.CString(name),
		C.CString(funcName),
		attributes,
		method,
	)
}

// RegisterProperty will register the given property with Godot and associate it
// with the given class name. The name parameter is the name of the class you wish
// to attach this property to. The path is the name of the property itself. The
// attributes and setter/getter methods are what will be called when Godot gets
// or sets this property.
func (n *nativeScript) RegisterProperty(name, path string, attributes *C.godot_property_attributes, setFunc C.godot_property_set_func, getFunc C.godot_property_get_func) {
	C.go_godot_nativescript_register_property(
		n.api,
		handle,
		C.CString(name),
		C.CString(path),
		attributes,
		setFunc,
		getFunc,
	)
}

// RegisterSignal will register the given signal with Godot.
func (n *nativeScript) RegisterSignal(name string, signal *C.godot_signal) {
	C.go_godot_nativescript_register_signal(
		n.api,
		handle,
		C.CString(name),
		signal,
	)
}

// Handle is a pointer to the nativescript handler. It must be passed to any
// Godot nativescript functions. This will be populated when 'godot_nativescript_init'
// is called by Godot upon script initialization.
var handle unsafe.Pointer

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes,
// etc. The `unsafe.Pointer` type is used to represent a void C pointer.
//export godot_nativescript_init
func godot_nativescript_init(hdl unsafe.Pointer) {
	log.Println("Initializing NativeScript")
	handle = hdl
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_create_func
//func go_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) unsafe.Pointer {
//	// Cast our pointer to a *char, so we can cast it to a Go string.
//	className := unsafeToGoString(methodData)
//	log.Println("Create function called for:", className)
//
//	// Look up our class constructor by its class name in the registry.
//	constructor := constructorRegistry[className]
//
//	// Create a new instance of the object.
//	class := constructor()
//	instanceString := fmt.Sprintf("%p", &class)
//	instanceCString := C.CString(instanceString)
//	log.Println("Created new object instance:", class, "with instance address:", instanceString)
//
//	// Add the Godot object pointer to the class structure.
//	class.setOwner(godotObject)
//
//	// Add the instance to our instance registry.
//	//instanceRegistry[instanceString] = class
//
//	// Return the instance string. This will be passed to the method function as userData, so we
//	// can look up the instance in our registry. Normally you would pass a pointer to the instance
//	// itself, but we should never pass Go structures to C, as the Go garbage collector might
//	// reap the object prematurely.
//	return unsafe.Pointer(instanceCString)
//}
//
//// This is a native Go function that is callable from C. It is called by the
//// gateway functions defined in cgateway.go. It will use the userData passed to it,
//// which is a CString of the instance id, which we can use to delete the instance
//// from the instance registry. This will make the instance available to be garbage
//// collected.
////export go_destroy_func
//func go_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
//	log.Println("Destroy function called!")
//
//	// Look up the instance based on the userData string.
//	instanceString := unsafeToGoString(userData)
//
//	// Remove it from our instanceRegistry so it can be garbage collected.
//	log.Println("Destroying instance:", instanceString)
//	//delete(instanceRegistry, instanceString)
//}
//
////export go_free_func
//func go_free_func(methodData unsafe.Pointer) {
//	log.Println("Free function called!")
//}
//
////godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
////export go_method_func
//func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
//
//	// Get the object instance based on the instance string given in userData.
//	methodString := unsafeToGoString(methodData)
//	instanceString := unsafeToGoString(userData)
//	class := instanceRegistry[instanceString]
//	classValue := reflect.ValueOf(class)
//
//	log.Println("Method was called!")
//	log.Println("  godotObject:", godotObject)
//	log.Println("  numArgs:", int(numArgs))
//	log.Println("  args:", args)
//	log.Println("  instance:", class)
//	log.Println("  instanceString (userData):", instanceString)
//	log.Println("  methodString (methodData):", methodString)
//
//	// Create a slice of godot_variant arguments
//	goArgsSlice := []reflect.Value{}
//
//	// Get the size of each godot_variant object pointer.
//	log.Println("  Getting size of argument pointer")
//	size := unsafe.Sizeof(*args)
//
//	// Panic if something's wrong.
//	if int(numArgs) > 50 {
//		panic("Too many arguments. Invalid method.")
//	}
//
//	// If we have arguments, append the first argument.
//	log.Println("  Checking if method had arguments")
//	if int(numArgs) > 0 {
//		log.Println("    It does!")
//		arg := *args
//		// Loop through all our arguments.
//		for i := 0; i < int(numArgs); i++ {
//			// Check to see what type this variant is
//			variantType := C.godot_variant_get_type(arg)
//			log.Println("Argument variant type:", variantType)
//
//			// TODO: Handle all variant types.
//			switch variantType {
//			case C.godot_variant_type(VariantTypeBool):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsBool(arg)))
//			case C.godot_variant_type(VariantTypeInt):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsInt(arg)))
//			case C.godot_variant_type(VariantTypeReal):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsReal(arg)))
//			case C.godot_variant_type(VariantTypeString):
//				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsString(arg)))
//			default:
//				log.Fatal("Unknown type of argument")
//			}
//
//			// Convert the pointer into a uintptr so we can perform artithmetic on it.
//			arrayPtr := uintptr(unsafe.Pointer(args))
//
//			// Add the size of the godot_variant pointer to our array pointer to get the position
//			// of the next argument.
//			arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))
//		}
//	}
//
//	// Use the method string to get the class name and method name.
//	log.Println("  Getting class name and method name...")
//	classMethodSlice := strings.Split(methodString, "::")
//	className := classMethodSlice[0]
//	methodName := classMethodSlice[1]
//	log.Println("    Class Name: ", className)
//	log.Println("    Method Name: ", methodName)
//
//	// Look up the registered class so we can find out how many arguments it takes
//	// and their types.
//	log.Println("  Look up the registered class and its method")
//	regClass := classRegistry[className]
//	if regClass == nil {
//		log.Fatal("  This class has not been registered! Class name: ", className, " Method name: ", methodName)
//	}
//	regMethod := regClass.methods[methodName]
//
//	log.Println("  Registered method arguments:", regMethod.arguments)
//	log.Println("  Arguments to pass:", goArgsSlice)
//
//	// Check to ensure the method has the same number of arguments we expect
//	if len(regMethod.arguments)-1 != int(numArgs) {
//		Log.Error("Invalid number of arguments. Expected ", numArgs, " arguments. (Got ", len(regMethod.arguments), ")")
//		panic("Invalid number of arguments.")
//	}
//
//	// Get the value of the class, so we can call methods on it.
//	method := classValue.MethodByName(methodName)
//	rawRet := method.Call(goArgsSlice)
//	log.Println(method)
//
//	var ret *C.godot_variant
//	var nonptrrtn C.godot_variant
//
//	if len(regMethod.returns) == 0 {
//		C.godot_variant_new_nil(&nonptrrtn)
//		return nonptrrtn
//	} else if len(regMethod.returns) > 1 {
//		panic("Too many return values from method! Methods called from within Godot should only return a single value.")
//	}
//
//	// Convert our returned value into a Godot Variant.
//	rawRetInterface := rawRet[0].Interface()
//	switch regMethod.returns[0].String() {
//
//	case "bool":
//		ret = boolAsVariant(rawRetInterface.(bool))
//
//	case "int64":
//		ret = intAsVariant(rawRetInterface.(int64))
//
//	case "int32":
//		ret = intAsVariant(int64(rawRetInterface.(int32)))
//
//	case "int":
//		ret = intAsVariant(int64(rawRetInterface.(int)))
//
//	case "uint64":
//		ret = uIntAsVariant(rawRetInterface.(uint64))
//
//	case "uint32":
//		ret = uIntAsVariant(uint64(rawRetInterface.(uint32)))
//
//	case "uint":
//		ret = uIntAsVariant(uint64(rawRetInterface.(uint)))
//
//	case "float64":
//		ret = realAsVariant(rawRetInterface.(float64))
//
//	case "string":
//		ret = stringAsVariant(rawRetInterface.(string))
//
//	default:
//		panic("The return was not valid. Should be Godot Variant or built-in Go type. Received: " + regMethod.returns[0].String())
//	}
//
//	return *ret
//}

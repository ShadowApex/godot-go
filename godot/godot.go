// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#include <stddef.h>
#include <stdlib.h>
#include <godot_nativescript.h>

// Type definitions for any function pointers. Some of these are not defined in
// the godot headers when they are part of a typedef struct.
typedef void (*create_func)(godot_object *, void *);
typedef void (*destroy_func)(godot_object *, void *, void *);
typedef void (*free_func)(void *);
typedef godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);

// Forward declarations of gateway functions defined in cgateway.go.
void *go_create_func_cgo(godot_object *, void *); // Forward declaration.
void *go_destroy_func_cgo(godot_object *, void *); // Forward declaration.
void *go_free_func_cgo(void *); // Forward declaration.
godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args);
*/
import "C"

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"unsafe"
)

const (
	tagPrefix        = "godot"
	defaultBaseClass = "Node"
)

// ClassConstructor is any function that will build and return a class to be registered
// with Godot.
type ClassConstructor func() Class

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
	log.Println("Initializing Go library.")

	// Translate the C struct into a Go struct.
	goOptions := &GodotGDNativeInitOptions{
		InEditor:      bool(options.in_editor),
		CoreAPIHash:   uint64(options.core_api_hash),
		EditorAPIHash: uint64(options.editor_api_hash),
		NoAPIHash:     uint64(options.no_api_hash),
	}

	// Call user defined init if it was set.
	if godotGDNativeInit != nil {
		godotGDNativeInit(goOptions)
	}
}

/** Library de-initialization **/
// godot_gdnative_terminate is the library's de-initialization method. When
// Godot unloads the library, this method will be called.
//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	log.Println("De-initializing Go library.")
}

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes
// and stuff. The `unsafe.Pointer` type is used to represent a void C pointer.
//export godot_nativescript_init
func godot_nativescript_init(desc unsafe.Pointer) {
	log.Println("Initializing script")

	// Loop through our registered classes and register them with the Godot API.
	for _, constructor := range godotConstructorsToRegister {
		// Use the constructor to build a class to inspect the given structure.
		class := constructor()

		// Get the type of the given struct, and get its name as a string and C string.
		classType := reflect.TypeOf(class)
		classString := strings.Replace(classType.String(), "*", "", 1)
		classCString := C.CString(classString)
		log.Println("Registering class:", classString)

		// Add the type to our internal type registry. This is used so the constructor
		// function can create the correct kind of struct.
		constructorRegistry[classString] = constructor

		// Create a registered class structure that will hold information about the
		// cass and its methods.
		regClass := newRegisteredClass(classType)

		// Call the "BaseClass" method on the class to get the base class.
		baseClass := class.(Class).BaseClass()
		baseClassCString := C.CString(baseClass)
		log.Println("  Using Base Class:", baseClass)

		// Set up our create function C struct
		var createFunc C.godot_instance_create_func
		var destroyFunc C.godot_instance_destroy_func

		// *** CREATE FUNC ***
		// Use a pointer to our gateway function.
		// GDCALLINGCONV void *(*create_func)(godot_object *, void *);
		createFunc.create_func = (C.create_func)(unsafe.Pointer(C.go_create_func_cgo))
		// void *method_data;
		createFunc.method_data = unsafe.Pointer(classCString)
		// GDCALLINGCONV void (*free_func)(void *);
		createFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

		// *** DESTROY FUNC ***
		// GDCALLINGCONV void (*destroy_func)(godot_object *, void *, void *);
		destroyFunc.destroy_func = (C.destroy_func)(unsafe.Pointer(C.go_destroy_func_cgo))
		// void *method_data;
		destroyFunc.method_data = unsafe.Pointer(classCString)
		// GDCALLINGCONV void (*free_func)(void *);
		destroyFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

		// Register our class.
		C.godot_nativescript_register_class(desc, classCString, baseClassCString, createFunc, destroyFunc)

		// Loop through our class's methods that are attached to it.
		log.Println("  Looking at methods:")
		log.Println("    Found", classType.NumMethod(), "methods")
		for i := 0; i < classType.NumMethod(); i++ {
			classMethod := classType.Method(i)
			log.Println("  Found method:", classMethod.Name)

			// Construct a registered method structure that inspects all of the
			// arguments and return types.
			regMethod := newRegisteredMethod(classMethod)
			regClass.addMethod(classMethod.Name, regMethod)
			log.Println("    Method Arguments:", len(regMethod.arguments))
			log.Println("    Method Arguments:", regMethod.arguments)
			log.Println("    Method Returns:", len(regMethod.returns))
			log.Println("    Method Returns:", regMethod.returns)

			// Look at the method name to see if it starts with "X_". If it does, we need to
			// replace it with an underscore. This is required because Go method visibility
			// is done through case sensitivity. Since Godot private methods start with an
			// "_" character.
			origMethodName := classMethod.Name
			methodString := origMethodName
			if strings.HasPrefix(methodString, "X_") {
				methodString = strings.Replace(methodString, "X_", "_", 1)
			}
			methodString = camelToSnake(methodString)
			methodString = strings.Replace(methodString, "__", "_", 1) // NOTE: Prevent double underscore from camelToSnake lib.
			log.Println("    Original method name:", origMethodName)
			log.Println("    Godot mapped method name:", methodString)

			// Set up the method name
			methodCString := C.CString(methodString)

			// Set up the method data, which will include the Go type name and Go method name.
			classMethodCString := C.CString(classString + "." + origMethodName)

			// Set up registering a method
			var method C.godot_instance_method

			// *** METHOD STRUCTURE ***
			// GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
			method.method = (C.method)(unsafe.Pointer(C.go_method_func_cgo))
			// void *method_data;
			method.method_data = unsafe.Pointer(classMethodCString)
			// GDCALLINGCONV void (*free_func)(void *);
			method.free_func = (C.free_func)(unsafe.Pointer(C.go_free_func_cgo))

			// Set up the method attributes.
			var attr C.godot_method_attributes
			attr.rpc_type = C.GODOT_METHOD_RPC_MODE_DISABLED

			// Register a method.
			C.godot_nativescript_register_method(desc, classCString, methodCString, attr, method)
		}

		// Register our class in our Go registry.
		classRegistry[classString] = regClass

	}
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_create_func
func go_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) unsafe.Pointer {
	// Cast our pointer to a *char, so we can cast it to a Go string.
	className := unsafeToGoString(methodData)
	log.Println("Create function called for:", className)

	// Look up our class constructor by its class name in the registry.
	constructor := constructorRegistry[className]

	// Create a new instance of the object.
	class := constructor()
	instanceString := fmt.Sprintf("%p", &class)
	instanceCString := C.CString(instanceString)
	log.Println("Created new object instance:", class, "with instance address:", instanceString)

	// Add the Godot object pointer to the class structure.
	class.SetOwner(godotObject)

	// Add the instance to our instance registry.
	instanceRegistry[instanceString] = class

	// Return the instance string. This will be passed to the method function as userData, so we
	// can look up the instance in our registry. Normally you would pass a pointer to the instance
	// itself, but we should never pass Go structures to C, as the Go garbage collector might
	// reap the object prematurely.
	return unsafe.Pointer(instanceCString)
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go. It will use the userData passed to it,
// which is a CString of the instance id, which we can use to delete the instance
// from the instance registry. This will make the instance available to be garbage
// collected.
//export go_destroy_func
func go_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
	log.Println("Destroy function called!")

	// Look up the instance based on the userData string.
	instanceString := unsafeToGoString(userData)

	// Remove it from our instanceRegistry so it can be garbage collected.
	log.Println("Destroying instance:", instanceString)
	delete(instanceRegistry, instanceString)
}

//export go_free_func
func go_free_func(methodData unsafe.Pointer) {
	log.Println("Free function called!")
}

//godot_variant go_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
//export go_method_func
func go_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) C.godot_variant {
	// Get the object instance based on the instance string given in userData.
	methodString := unsafeToGoString(methodData)
	instanceString := unsafeToGoString(userData)
	class := instanceRegistry[instanceString]
	classValue := reflect.ValueOf(class)

	log.Println("Method was called!")
	log.Println("  godotObject:", godotObject)
	log.Println("  numArgs:", int(numArgs))
	log.Println("  args:", args)
	log.Println("  instance:", class)
	log.Println("  instanceString (userData):", instanceString)
	log.Println("  methodString (methodData):", methodString)

	// Create a slice of godot_variant arguments
	goArgsSlice := []reflect.Value{}

	// Get the size of each godot_variant object pointer.
	size := unsafe.Sizeof(*args)

	// Panic if something's wrong.
	if int(numArgs) > 500 {
		panic("Wtf, 500+ arguments?")
	}

	// If we have arguments, append the first argument.
	if int(numArgs) > 0 {
		arg := *args
		// Loop through all our arguments.
		for i := 0; i < int(numArgs); i++ {
			// Check to see what type this variant is
			variantType := C.godot_variant_get_type(arg)
			log.Println("Argument variant type:", variantType)

			// TODO: Handle all variant types.
			switch variantType {
			case C.godot_variant_type(VariantTypeBool):
				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsBool(arg)))
			case C.godot_variant_type(VariantTypeInt):
				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsInt(arg)))
			case C.godot_variant_type(VariantTypeReal):
				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsReal(arg)))
			case C.godot_variant_type(VariantTypeString):
				goArgsSlice = append(goArgsSlice, reflect.ValueOf(variantAsString(arg)))
			default:
				log.Fatal("Unknown type of argument")
			}

			// Convert the pointer into a uintptr so we can perform artithmetic on it.
			arrayPtr := uintptr(unsafe.Pointer(args))

			// Add the size of the godot_variant pointer to our array pointer to get the position
			// of the next argument.
			arg = (*C.godot_variant)(unsafe.Pointer(arrayPtr + size))
		}
	}

	// Use the method string to get the class name and method name.
	classMethodSlice := strings.Split(methodString, ".")
	className := classMethodSlice[0]
	methodName := classMethodSlice[1]

	// Look up the registered class so we can find out how many methods and their
	// types we should call.
	regClass := classRegistry[className]
	regMethod := regClass.methods[methodName]

	log.Println("  Registered method arguments:", regMethod.arguments)
	log.Println("  Arguments to pass:", goArgsSlice)

	// Check to ensure the method has the same number of arguments we expect
	if len(regMethod.arguments)-1 != int(numArgs) {
		Log.Error("Invalid number of arguments. Expected ", numArgs, " arguments. (Got ", len(regMethod.arguments), ")")
		panic("Invalid number of arguments.")
	}

	// Get the value of the class, so we can call methods on it.
	method := classValue.MethodByName(methodName)
	method.Call(goArgsSlice)
	log.Println(method)

	var ret C.godot_variant
	C.godot_variant_new_nil(&ret)

	return ret
}

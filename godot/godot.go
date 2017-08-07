// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#include <stddef.h>
#include <godot_nativescript.h>

// Type definitions for any function pointers. Some of these are not defined in
// the godot headers when they are part of a typedef struct.
typedef void (*create_func)(godot_object *, void *);
typedef void (*destroy_func)(godot_object *, void *, void *);
typedef void (*free_func)(void *);
typedef godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);

// Forward declarations of gateway functions defined in cgateway.go.
void *go_godot_instance_create_func_cgo(godot_object *, void *); // Forward declaration.
void *go_godot_instance_destroy_func_cgo(godot_object *, void *); // Forward declaration.
void *go_godot_instance_free_func_cgo(void *); // Forward declaration.
godot_variant go_godot_instance_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args);
*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

const (
	readyMethod      = "Ready"
	readyGodotMethod = "_ready"
)

// GodotGDNativeInit is a function signature for functions that will be called
// upon library initialization.
type GodotGDNativeInit func(options *GodotGDNativeInitOptions)

// godotGDNativeInit is a user defined function that will be set by SetGodotGDNativeInit.
var godotGDNativeInit GodotGDNativeInit

// SetGodotGDNativeInit takes an initialization function that will be called
// when Godot loads the Go library. This can only work if it is run by
// setting a variable. E.g. var myinit = godot.SetGodotGDNativeInit(myfunc)
func SetGodotGDNativeInit(init GodotGDNativeInit) GodotGDNativeInit {
	godotGDNativeInit = init
	return init
}

// godotClassesToRegister is a slice of objects that will be registered as a Godot class
// upon library initialization.
var godotClassesToRegister = []interface{}{}

// RegisterClass will register the given object as a Godot class. It will be available
// inside Godot.
func RegisterClass(object interface{}) {
	godotClassesToRegister = append(godotClassesToRegister, object)
}

// function signature for Godot classes
type GodotMethod func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant)

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	fmt.Println("GO: Initializing Go library.")

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
	fmt.Println("GO: De-initializing Go library.")
}

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes
// and stuff. The `unsafe.Pointer` type is used to represent a null C pointer.
//export godot_nativescript_init
func godot_nativescript_init(desc unsafe.Pointer) {
	fmt.Println("GO: Initializing script")

	// Loop through our registered classes and register them with the Godot API.
	for _, class := range godotClassesToRegister {
		// Get the type of the given struct.
		classType := reflect.TypeOf(class)
		fmt.Println("GO: Registering class:", classType)

		// Set up our create function C struct
		var createFunc C.godot_instance_create_func
		var destroyFunc C.godot_instance_destroy_func

		// *** CREATE FUNC ***
		// Use a pointer to our gateway function.
		// GDCALLINGCONV void *(*create_func)(godot_object *, void *);
		createFunc.create_func = (C.create_func)(unsafe.Pointer(C.go_godot_instance_create_func_cgo))
		// void *method_data;
		createFunc.method_data = unsafe.Pointer(C.CString("Some data"))
		// GDCALLINGCONV void (*free_func)(void *);
		createFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_godot_instance_free_func_cgo))

		// *** DESTROY FUNC ***
		// GDCALLINGCONV void (*destroy_func)(godot_object *, void *, void *);
		destroyFunc.destroy_func = (C.destroy_func)(unsafe.Pointer(C.go_godot_instance_destroy_func_cgo))
		// void *method_data;
		destroyFunc.method_data = unsafe.Pointer(C.CString("Some data"))
		// GDCALLINGCONV void (*free_func)(void *);
		destroyFunc.free_func = (C.free_func)(unsafe.Pointer(C.go_godot_instance_free_func_cgo))

		// Register our class.
		C.godot_nativescript_register_class(desc, C.CString("SimpleClass"), C.CString("Node"), createFunc, destroyFunc)

		// Loop through our class's methods that are attached to it.
		fmt.Println("GO:   Looking at methods:")
		fmt.Println("GO:     Found", classType.NumMethod(), "methods")
		for i := 0; i < classType.NumMethod(); i++ {
			classMethod := classType.Method(i)
			fmt.Println("GO:   Found method:", classMethod.Name)

			// Set up registering a method
			var method C.godot_instance_method

			// *** METHOD STRUCTURE ***
			// GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
			method.method = (C.method)(unsafe.Pointer(C.go_godot_instance_method_func_cgo))
			// void *method_data;
			method.method_data = unsafe.Pointer(C.CString("Some data"))
			// GDCALLINGCONV void (*free_func)(void *);
			method.free_func = (C.free_func)(unsafe.Pointer(C.go_godot_instance_free_func_cgo))

			// Set up the method attributes.
			var attr C.godot_method_attributes
			attr.rpc_type = C.GODOT_METHOD_RPC_MODE_DISABLED

			// Register a method.
			C.godot_nativescript_register_method(desc, C.CString("SimpleClass"), C.CString("_ready"), attr, method)

		}
	}
	/*
		    godot_instance_create_func create_func = {
		        .create_func = &test_constructor,
		                .method_data = 0,
		                .free_func   = 0
		        };

			typedef struct {
			    // instance pointer, method data, user data, num args, args - return result as varaint
			    GDCALLINGCONV godot_variant (*method)(godot_object *, void *, void *, int, godot_variant **);
			    void *method_data;
			    GDCALLINGCONV void (*free_func)(void *);
			} godot_instance_method;

	*/
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_godot_instance_create_func
func go_godot_instance_create_func(godotObject *C.godot_object, methodData unsafe.Pointer) {
	fmt.Println("GO: Create function called!")
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cgateway.go.
//export go_godot_instance_destroy_func
func go_godot_instance_destroy_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer) {
	fmt.Println("GO: Destroy function called!")
}

//export go_godot_instance_free_func
func go_godot_instance_free_func(param unsafe.Pointer) {
	fmt.Println("GO: Free function called!")
}

//godot_variant go_godot_instance_method_func_cgo(godot_object *obj, void *method_data, void *user_data, int num_args, godot_variant **args)
//export go_godot_instance_method_func
func go_godot_instance_method_func(godotObject *C.godot_object, methodData unsafe.Pointer, userData unsafe.Pointer, numArgs C.int, args **C.godot_variant) {
	fmt.Println("GO: Method was called!")
	fmt.Println("GO:   godotObject:", godotObject)
	fmt.Println("GO:   methodData:", methodData)
	fmt.Println("GO:   userData:", userData)
	fmt.Println("GO:   numArgs:", numArgs)
	fmt.Println("GO:   args:", args)
}

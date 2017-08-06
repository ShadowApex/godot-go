// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#include <stddef.h>
#include <godot_nativescript.h>

// Type definitions
typedef void (*create_func)(godot_object *, void *);

// Forward declarations of gateway functions defined in cfuncs.go.
void *go_godot_instance_create_func_cgo(godot_object *, void *); // Forward declaration.
*/
import "C"

import (
	"fmt"
	"unsafe"
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

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	fmt.Println("Initializing Go library.")

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
	fmt.Println("De-initializing Go library.")
}

/** Script entry (Registering all the classes and stuff) **/
// godot_nativescript_init is the script's entrypoint. It is called by Godot
// when a script is loaded. It is responsible for registering all the classes
// and stuff. The `unsafe.Pointer` type is used to represent a null C pointer.
//export godot_nativescript_init
func godot_nativescript_init(desc unsafe.Pointer) {
	fmt.Println("Initializing script")
	// Set up our create function C struct
	var createFunc C.godot_instance_create_func
	var destroyFunc C.godot_instance_destroy_func

	// Use a pointer to our gateway function.
	// GDCALLINGCONV void *(*create_func)(godot_object *, void *);
	createFunc.create_func = (C.create_func)(unsafe.Pointer(C.go_godot_instance_create_func_cgo))
	// void *method_data;
	createFunc.method_data = unsafe.Pointer(C.CString("Some data"))
	// GDCALLINGCONV void (*free_func)(void *);
	//createFunc.free_func = C.CString("hello")

	fmt.Println(createFunc)

	C.godot_nativescript_register_class(desc, C.CString("SimpleClass"), C.CString("Node"), createFunc, destroyFunc)

	/*
			typedef struct {
			    // instance pointer, method_data - return user data
			    GDCALLINGCONV void *(*create_func)(godot_object *, void *);
			    void *method_data;
			    GDCALLINGCONV void (*free_func)(void *);
			} godot_instance_create_func;

		    godot_instance_create_func create_func = {
		        .create_func = &test_constructor,
		    │   │   │   .method_data = 0,
		    │   │   │   .free_func   = 0
		    │   };

	*/
}

// This is a native Go function that is callable from C. It is called by the
// gateway functions defined in cfuncs.go.
//export go_godot_instance_create_func
func go_godot_instance_create_func(godotObject *C.godot_object, param unsafe.Pointer) {
	fmt.Println("Native Go code is being executed here!")
}

// package name: libgodot
package godot

/*
#cgo linux CXXFLAGS: -I/usr/local/include -std=c11
#cgo darwin CFLAGS: -I/usr/local/include
#include <stddef.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
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

	//C.godot_nativescript_register_class(desc, "SimpleClass", "Node", create_func, destroy_func)
	//godot_nativescript_register_class(desc, "SimpleClass", "Node", create_func, destroy_func);
}

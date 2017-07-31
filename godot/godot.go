// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
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

// SetGodotGDNativeInit takes an initialization function that will be called
// when Godot loads the Go library.
func SetGodotGDNativeInit(init GodotGDNativeInit) {
}

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	fmt.Println("Initializing Go library.")
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
}

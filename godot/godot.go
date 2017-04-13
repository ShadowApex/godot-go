// package name: libgodot
package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#include <stddef.h>
#include <godot.h>
*/
import "C"

import (
	"fmt"
)

var go_godot_native_init Init

type Init func(options *InitOptions)

// SetNativeInit is the entry point for the script
func SetNativeInit(init Init) Init {
	go_godot_native_init = init
	return init
}

func newInitOptions(options *C.godot_native_init_options) *InitOptions {
	return &InitOptions{
		InEditor:      bool(options.in_editor),
		CoreAPIHash:   uint64(options.core_api_hash),
		EditorAPIHash: uint64(options.editor_api_hash),
		NoAPIHash:     uint64(options.no_api_hash),
	}
}

// InitOptions are options that are passed from Godot.
type InitOptions struct {
	InEditor      bool
	CoreAPIHash   uint64
	EditorAPIHash uint64
	NoAPIHash     uint64
}

// Godot entry point
//export godot_native_init
func godot_native_init(options *C.godot_native_init_options) {
	if go_godot_native_init == nil {
		panic("You must set the native init function!")
	}
	go_godot_native_init(newInitOptions(options))
	fmt.Println("Hey from go!")
	fmt.Println("Init options:", options)
}

// Godot termination point
//export godot_native_terminate
func godot_native_terminate() {
	fmt.Println("Exiting go!")
}

// void GDAPI godot_script_register_class(const char *p_name, const char *p_base, godot_instance_create_func p_create_func, godot_instance_destroy_func p_destroy_func);

type GodotObject struct {
}

func (g *GodotObject) Blah() {
}

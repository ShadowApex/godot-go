// Package godot provides a wrapper around the Godot GDNative API.
package godot

/*
#cgo CFLAGS: -I../godot_headers -std=c11
#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup

#include <gdnative/gdnative.h>
#include <gdnative_api_struct.gen.h>
#include "util.h"
*/
import "C"

import (
	"log"
	"unsafe"
)

// Set our API to null. This will be set when 'godot_gdnative_init' is called
// by Godot when the library is loaded.
var GDNative = &gdNative{}

// gdNative is a structure that wraps the GDNativeAPI.
type gdNative struct {
	api *C.godot_gdnative_core_api_struct
}

func Null() {
}

// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	// Get the API struct from the init options passed by Godot when this
	// library is loaded. This API struct will have all of the functions
	// to call.
	GDNative.api = (*options).api_struct

	// Configure logging.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
	log.Println("Initializing godot-go library.")

	// Find GDNative extensions that we support.
	for i := 0; i < int(GDNative.api.num_extensions); i++ {
		extension := C.cgo_get_ext(GDNative.api.extensions, C.int(i))
		switch extension._type {
		case C.GDNATIVE_EXT_NATIVESCRIPT:
			log.Println("Found nativescript extension!")
			NativeScript.api = (*C.godot_gdnative_ext_nativescript_api_struct)(unsafe.Pointer(extension))
		}
	}
}

/** Library de-initialization **/
// godot_gdnative_terminate is the library's de-initialization method. When
// Godot unloads the library, this method will be called.
//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	log.Println("De-initializing Go library.")
	GDNative.api = nil
	NativeScript.api = nil
}

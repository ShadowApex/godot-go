// Package gdnative provides a thin wrapper around the Godot GDNative API.
package gdnative

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
var API *C.godot_gdnative_core_api_struct
var NativeScriptAPI *C.godot_gdnative_ext_nativescript_api_struct

func Null() {
}

/** Library entry point **/
// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	// Get the API struct from the init options passed by Godot when this
	// library is loaded. This API struct will have all of the functions
	// to call.
	API = (*options).api_struct

	// Configure logging.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
	log.Println("Initializing godot-go library.")

	// Find GDNative extensions that we support.
	for i := 0; i < int(API.num_extensions); i++ {
		extension := C.cgo_get_ext(API.extensions, C.int(i))
		switch extension._type {
		case C.GDNATIVE_EXT_NATIVESCRIPT:
			log.Println("Found nativescript extension!")
			NativeScriptAPI = (*C.godot_gdnative_ext_nativescript_api_struct)(unsafe.Pointer(extension))
		}
	}
}

/** Library de-initialization **/
// godot_gdnative_terminate is the library's de-initialization method. When
// Godot unloads the library, this method will be called.
//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	log.Println("De-initializing Go library.")
	API = nil
	NativeScriptAPI = nil
}

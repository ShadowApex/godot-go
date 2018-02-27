// Package gdnative provides a wrapper around the Godot GDNative API.
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

	"github.com/vitaminwater/cgo.wchar"
)

// Set our API to null. This will be set when 'godot_gdnative_init' is called
// by Godot when the library is loaded.
var GDNative = &gdNative{}

// gdNative is a structure that wraps the GDNativeAPI.
type gdNative struct {
	api *C.godot_gdnative_core_api_struct
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

type Char string

func (c Char) getBase() *C.char {
	return C.CString(string(c))
}

type Double float64

func (d Double) getBase() C.double {
	return C.double(d)
}

type Int64T int64

func (i Int64T) getBase() C.int64_t {
	return C.int64_t(i)
}

type SignedChar int8

func (s SignedChar) getBase() *C.schar {
	intVal := int8(s)
	return (*C.schar)(unsafe.Pointer(&intVal))
}

// Uint is a Godot C uint wrapper
type Uint uint

func (u Uint) getBase() C.uint {
	return C.uint(u)
}

// Uint8T is a Godot C uint8_t wrapper
type Uint8T uint8

func (u Uint8T) getBase() C.uint8_t {
	return C.uint8_t(u)
}

// Uint32T is a Godot C uint32_t wrapper
type Uint32T uint32

func (u Uint32T) getBase() C.uint32_t {
	return C.uint32_t(u)
}

// Uint64T is a Godot C uint64_t wrapper
type Uint64T uint64

func (u Uint64T) getBase() C.uint64_t {
	return C.uint64_t(u)
}

// newWcharT will convert the given C.wchar_t into a Go string
func newWcharT(str *C.wchar_t) WcharT {
	goStr, err := wchar.WcharStringPtrToGoString(unsafe.Pointer(str))
	if err != nil {
		log.Println("Error converting wchar_t to Go string:", err)
	}
	return WcharT(goStr)
}

// WcharT is a Godot C wchar_t wrapper
type WcharT string

func (w WcharT) getBase() *C.wchar_t {
	wcharString, err := wchar.FromGoString(string(w))
	if err != nil {
		log.Println("Error decoding WcharT:", err)
	}

	return (*C.wchar_t)(wcharString.Pointer())
}

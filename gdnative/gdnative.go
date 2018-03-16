// Package gdnative provides a wrapper around the Godot GDNative API.
package gdnative

/*
#cgo CFLAGS: -I../godot_headers -std=c11
#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all
#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup

#include <gdnative/gdnative.h>
#include <gdnative_api_struct.gen.h>
#include "gdnative.gen.h"
#include "util.h"
*/
import "C"

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/vitaminwater/cgo.wchar"
)

// debug determines whether or not we should log messages
var debug = false

// EnableDebug will enable logging of GDNative debug messages.
func EnableDebug() {
	debug = true
}

// Set our API to null. This will be set when 'godot_gdnative_init' is called
// by Godot when the library is loaded.
var GDNative = &gdNative{}

// gdNative is a structure that wraps the GDNativeAPI.
type gdNative struct {
	api         *C.godot_gdnative_core_api_struct
	initialized bool
}

// IsInitialized will return true if `godot_gdnative_init` was called by Godot.
func (g *gdNative) IsInitialized() bool {
	return g.initialized
}

// CheckInit will check to see if GDNative has initialized. If it is not, it will
// throw a panic.
func (g *gdNative) checkInit() {
	if !g.IsInitialized() {
		panic("GDNative has not initialized! You cannot call this in init()!")
	}
}

// godot_gdnative_init is the library entry point. When the library is loaded
// this method will be called by Godot.
//export godot_gdnative_init
func godot_gdnative_init(options *C.godot_gdnative_init_options) {
	// Get the API struct from the init options passed by Godot when this
	// library is loaded. This API struct will have all of the functions
	// to call.
	GDNative.api = (*options).api_struct
	GDNative.initialized = true

	// Configure logging.
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(Log)
	if debug {
		log.Println("Initializing godot-go library.")
	}

	// Find GDNative extensions that we support.
	for i := 0; i < int(GDNative.api.num_extensions); i++ {
		extension := C.cgo_get_ext(GDNative.api.extensions, C.int(i))
		switch extension._type {
		case C.GDNATIVE_EXT_NATIVESCRIPT:
			if debug {
				log.Println("Found nativescript extension!")
			}
			NativeScript.api = (*C.godot_gdnative_ext_nativescript_api_struct)(unsafe.Pointer(extension))
		}
	}
}

/** Library de-initialization **/
// godot_gdnative_terminate is the library's de-initialization method. When
// Godot unloads the library, this method will be called.
//export godot_gdnative_terminate
func godot_gdnative_terminate(options *C.godot_gdnative_terminate_options) {
	if debug {
		log.Println("De-initializing Go library.")
	}
	GDNative.api = nil
	NativeScript.api = nil
}

func NewEmptyVoid() Pointer {
	var empty C.void
	return Pointer{base: unsafe.Pointer(&empty)}
}

// GetSingleton will return an instance of the given singleton.
func GetSingleton(name String) Object {
	if debug {
		log.Println("Getting singleton:", name)
	}
	GDNative.checkInit()

	// Create a C string from the name argument.
	cName := C.CString(string(name))

	// Call the C method
	obj := C.go_godot_global_get_singleton(GDNative.api, cName)
	return Object{base: (*C.godot_object)(obj)}
}

// NewMethodBind will return a method binding using the given class name and method
// name.
func NewMethodBind(class, method string) MethodBind {
	if debug {
		log.Println("Creating method bind for:", class+"."+method)
	}
	GDNative.checkInit()
	methodBind := C.go_godot_method_bind_get_method(
		GDNative.api,
		C.CString(class),
		C.CString(method),
	)

	return MethodBind{base: methodBind}
}

// MethodBindPtrCall will call the given method on the given Godot Object. Its return
// value is given as a pointer, which can be used to convert it to a variant.
func MethodBindPtrCall(methodBind MethodBind, instance Object, args []Pointer, returns Pointer) Pointer {
	GDNative.checkInit()
	if instance.getBase() == nil {
		panic("Godot object pointer was nil when calling MethodBindPtrCall")
	}

	// Build out our C arguments array
	cArgs := C.go_void_build_array(C.int(len(args)))
	for i, arg := range args {
		C.go_void_add_element(cArgs, arg.getBase(), C.int(i))
	}

	// Print all of the shit we're passing
	for i, arg := range args {
		if debug {
			log.Println("arg", i, ": ", arg.getBase())
		}
	}
	if debug {
		log.Println("args: ", cArgs)
		log.Println("returns: ", returns.getBase())
		log.Println("object: ", unsafe.Pointer(instance.getBase()))
		log.Println("methodbind: ", unsafe.Pointer(methodBind.getBase()))
	}

	// Call the C method
	C.go_godot_method_bind_ptrcall(
		GDNative.api,
		methodBind.getBase(),
		unsafe.Pointer(instance.getBase()),
		cArgs,
		returns.getBase(),
	)
	if debug {
		log.Println("Finished calling method.")
	}

	return returns
}

// Pointer is a pointer to arbitrary underlying data. This is primarily used
// in conjunction with MethodBindPtrCall.
type Pointer struct {
	base unsafe.Pointer
}

func (p *Pointer) getBase() unsafe.Pointer {
	return p.base
}

type Char string

func (c Char) getBase() *C.char {
	return C.CString(string(c))
}

type Double float64

func (d Double) getBase() C.double {
	return C.double(d)
}

// NewPointerFromFloat will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromFloat(obj Float) Pointer {
	base := obj.getBase()
	return Pointer{base: unsafe.Pointer(&base)}
}

// NewFloatFromPointer will return a Float from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewFloatFromPointer(ptr Pointer) Float {
	base := ptr.getBase()
	return Float(*(*C.float)(base))
}

// NewEmptyFloat will return a pointer to an empty
// initialized Float. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyFloat() Pointer {
	var obj C.float
	return Pointer{base: unsafe.Pointer(&obj)}
}

type Float float64

func (f Float) getBase() C.float {
	return C.float(f)
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

// newWcharTWithLength will convert the given C.wchar_t into a Go string
func newWcharTWithLength(str *C.wchar_t, length int) WcharT {
	goStr, err := wchar.WcharStringPtrNToGoString(unsafe.Pointer(str), length)
	if err != nil {
		log.Println("Error converting wchar_t to Go string:", err)
	}
	if len(goStr) != length {
		goStr = truncateString(goStr, length)
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

func (w WcharT) AsString() String {
	return NewStringWithWideString(string(w))
}

// ID will return the Godot object memory address as a string, which can
// be used in an instance registry for registering classes.
func (gdt Object) ID() string {
	return fmt.Sprintf("%p", gdt.base)
}

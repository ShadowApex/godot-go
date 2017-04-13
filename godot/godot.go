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
	//"unsafe"
)

////// Error
const (
	GODOT_OK                        = iota
	GODOT_FAILED                    ///< Generic fail error
	GODOT_ERR_UNAVAILABLE           ///< What is requested is unsupported/unavailable
	GODOT_ERR_UNCONFIGURED          ///< The object being used hasnt been properly set up yet
	GODOT_ERR_UNAUTHORIZED          ///< Missing credentials for requested resource
	GODOT_ERR_PARAMETER_RANGE_ERROR ///< Parameter given out of range (5)
	GODOT_ERR_OUT_OF_MEMORY         ///< Out of memory
	GODOT_ERR_FILE_NOT_FOUND
	GODOT_ERR_FILE_BAD_DRIVE
	GODOT_ERR_FILE_BAD_PATH
	GODOT_ERR_FILE_NO_PERMISSION // (10)
	GODOT_ERR_FILE_ALREADY_IN_USE
	GODOT_ERR_FILE_CANT_OPEN
	GODOT_ERR_FILE_CANT_WRITE
	GODOT_ERR_FILE_CANT_READ
	GODOT_ERR_FILE_UNRECOGNIZED // (15)
	GODOT_ERR_FILE_CORRUPT
	GODOT_ERR_FILE_MISSING_DEPENDENCIES
	GODOT_ERR_FILE_EOF
	GODOT_ERR_CANT_OPEN   ///< Can't open a resource/socket/file
	GODOT_ERR_CANT_CREATE // (20)
	GODOT_ERR_QUERY_FAILED
	GODOT_ERR_ALREADY_IN_USE
	GODOT_ERR_LOCKED ///< resource is locked
	GODOT_ERR_TIMEOUT
	GODOT_ERR_CANT_CONNECT // (25)
	GODOT_ERR_CANT_RESOLVE
	GODOT_ERR_CONNECTION_ERROR
	GODOT_ERR_CANT_AQUIRE_RESOURCE
	GODOT_ERR_CANT_FORK
	GODOT_ERR_INVALID_DATA        ///< Data passed is invalid	(30)
	GODOT_ERR_INVALID_PARAMETER   ///< Parameter passed is invalid
	GODOT_ERR_ALREADY_EXISTS      ///< When adding, item already exists
	GODOT_ERR_DOES_NOT_EXIST      ///< When retrieving/erasing, it item does not exist
	GODOT_ERR_DATABASE_CANT_READ  ///< database is full
	GODOT_ERR_DATABASE_CANT_WRITE ///< database is full	(35)
	GODOT_ERR_COMPILATION_FAILED
	GODOT_ERR_METHOD_NOT_FOUND
	GODOT_ERR_LINK_FAILED
	GODOT_ERR_SCRIPT_FAILED
	GODOT_ERR_CYCLIC_LINK // (40)
	GODOT_ERR_INVALID_DECLARATION
	GODOT_ERR_DUPLICATE_SYMBOL
	GODOT_ERR_PARSE_ERROR
	GODOT_ERR_BUSY
	GODOT_ERR_SKIP                                                              // (45)
	GODOT_ERR_HELP                                                              ///< user requested help!!
	GODOT_ERR_BUG                                                               ///< a bug in the software certainly happened, due to a double check failing or unexpected behavior.
	GODOT_ERR_PRINTER_ON_FIRE                                                   /// the parallel port printer is engulfed in flames
	GODOT_ERR_OMFG_THIS_IS_VERY_VERY_BAD                                        ///< shit happens, has never been used, though
	GODOT_ERR_WTF                        = GODOT_ERR_OMFG_THIS_IS_VERY_VERY_BAD ///< short version of the above
)

// godot_error handles Godot error enums.
func godotError(err int) interface{} {
	return C.godot_error(err)
}

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
// TODO: From Go, we can't expose a `const char*`. Might need to write a wrapper in C in header.
//export go_godot_script_register_class
func go_godot_script_register_class(name string, base string, createFunc C.godot_instance_create_func, destroyFunc C.godot_instance_destroy_func) {
	fmt.Println("GO: godot_script_register_class")
}

type GodotObject struct {
}

func (g *GodotObject) Blah() {
}

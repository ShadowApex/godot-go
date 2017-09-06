package godot

import (
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

// goToGodotConverter is a function that will convert a Go object into
// a Godot object.
type goToGodotConverter func(goObject interface{}) unsafe.Pointer

// goToGodotConversionMap is an internal mapping of Go types to functions that can
// convert to Godot types. This mapping is essentially a more optimal case/switch
// system for converting Go types to Godot types that can be used as arguments to
// a Godot function.
var goToGodotConversionMap = map[string]goToGodotConverter{
	"bool": func(goObject interface{}) unsafe.Pointer {
		boolArg := C.godot_bool(goObject.(bool))
		return unsafe.Pointer(&boolArg)
	},
	"int64": func(goObject interface{}) unsafe.Pointer {
		var intArg C.int64_t
		intArg = C.int64_t(goObject.(int64))
		return unsafe.Pointer(&intArg)
	},
	"uint64": func(goObject interface{}) unsafe.Pointer {
		var intArg C.uint64_t
		intArg = C.uint64_t(goObject.(uint64))
		return unsafe.Pointer(&intArg)
	},
	"float64": func(goObject interface{}) unsafe.Pointer {
		floatArg := C.double(goObject.(float64))
		return unsafe.Pointer(&floatArg)
	},
	"string": func(goObject interface{}) unsafe.Pointer {
		stringArg := stringAsGodotString(goObject.(string))
		return unsafe.Pointer(stringArg)
	},
}

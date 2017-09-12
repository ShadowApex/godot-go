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
	"*Array": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Array)
		return unsafe.Pointer(arg.array)
	},
	"*Basis": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Basis)
		return unsafe.Pointer(arg.basis)
	},
	"*Color": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Color)
		return unsafe.Pointer(arg.color)
	},
	"*Dictionary": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Dictionary)
		return unsafe.Pointer(arg.dictionary)
	},
	"*NodePath": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*NodePath)
		return unsafe.Pointer(arg.nodePath)
	},
	"*Plane": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Plane)
		return unsafe.Pointer(arg.plane)
	},
	"*Quat": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Quat)
		return unsafe.Pointer(arg.quat)
	},
	"*Real": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Real)
		return unsafe.Pointer(arg.real)
	},
	"*Rect2": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Rect2)
		return unsafe.Pointer(arg.rect2)
	},
	"*Rect3": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Rect3)
		return unsafe.Pointer(arg.rect3)
	},
	"*RID": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*RID)
		return unsafe.Pointer(arg.rid)
	},
	"*Transform": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Transform)
		return unsafe.Pointer(arg.transform)
	},
	"*Transform2D": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Transform2D)
		return unsafe.Pointer(arg.transform2d)
	},
	"*Variant": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Variant)
		return unsafe.Pointer(arg.variant)
	},
	"*Vector2": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Vector2)
		return unsafe.Pointer(arg.vector2)
	},
	"*Vector3": func(goObject interface{}) unsafe.Pointer {
		arg := goObject.(*Vector3)
		return unsafe.Pointer(arg.vector3)
	},
}

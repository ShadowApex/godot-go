package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

const (
	VariantTypeNil = iota

	// atomic types
	VariantTypeBool
	VariantTypeInt
	VariantTypeReal
	VariantTypeString

	// math types
	VariantTypeVector2 // 5
	VariantTypeRect2
	VariantTypeVector3
	VariantTypeTransform2d
	VariantTypePlane
	VariantTypeQuat // 10
	VariantTypeRect3
	VariantTypeBasis
	VariantTypeTransform

	// misc types
	VariantTypeColor
	VariantTypenodePath // 15
	VariantTyperId
	VariantTypeObject
	VariantTypeDictionary
	VariantTypeArray // 20

	// ARRAYS
	VariantTypePoolByteArray
	VariantTypePoolIntArray
	VariantTypePoolRealArray
	VariantTypePoolStringArray
	VariantTypePoolVector2Array // 25
	VariantTypePoolVector3Array
	VariantTypePoolColorArray
)

type Variant struct {
	variant *C.godot_variant
}

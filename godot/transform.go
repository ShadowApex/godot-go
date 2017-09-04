package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Matrix3, Vector3
//func NewTransform() *Transform {
//	transform := &Transform{}
//
//	// Create our godot transform object
//	var godotTransform C.godot_transform
//
//	// Create our transform
//	C.godot_transform_new(&godotTransform)
//
//	// Internally set our transform
//	transform.transform = &godotTransform
//
//	return transform
//}

// TODO: Finish implementing this
type Transform struct {
	transform *C.godot_transform
}

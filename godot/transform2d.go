package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//func NewTransform2D() *Transform2D {
//	transform2d := &Transform2D{}
//
//	// Create our godot transform2d object
//	var godotTransform2d C.godot_transform2d
//
//	// Create our transform2d
//	C.godot_transform2d_new(&godotTransform2d)
//
//	// Internally set our transform2d
//	transform2d.transform2d = &godotTransform2d
//
//	return transform2d
//}

// TODO: Finish implementing this
type Transform2D struct {
	transform2d *C.godot_transform2d
}

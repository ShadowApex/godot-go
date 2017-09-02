package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//func NewVector3() *Vector3 {
//	vector3 := &Vector3{}
//
//	// Create our godot vector3 object
//	var godotVector3 C.godot_vector3
//
//	// Create our vector3
//	C.godot_vector3_new(&godotVector3)
//
//	// Internally set our vector3
//	vector3.vector3 = &godotVector3
//
//	return vector3
//}

// TODO: Finish implementing this
type Vector3 struct {
	vector3 *C.godot_vector3
}

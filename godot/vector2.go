package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Built-In types
//func NewVector2() *Vector2 {
//	vector2 := &Vector2{}
//
//	// Create our godot vector2 object
//	var godotVector2 C.godot_vector2
//
//	// Create our vector2
//	C.godot_vector2_new(&godotVector2)
//
//	// Internally set our vector2
//	vector2.vector2 = &godotVector2
//
//	return vector2
//}

// TODO: Finish implementing this
type Vector2 struct {
	vector2 *C.godot_vector2
}

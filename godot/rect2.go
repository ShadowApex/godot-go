package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Vector2
//func NewRect2() *Rect2 {
//	rect2 := &Rect2{}
//
//	// Create our godot rect2 object
//	var godotRect2 C.godot_rect2
//
//	// Create our rect2
//	C.godot_rect2_new(&godotRect2)
//
//	// Internally set our rect2
//	rect2.rect2 = &godotRect2
//
//	return rect2
//}

// TODO: Finish implementing this
type Rect2 struct {
	rect2 *C.godot_rect2
}

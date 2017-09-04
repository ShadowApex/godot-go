package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Vector3
//func NewRect3() *Rect3 {
//	rect3 := &Rect3{}
//
//	// Create our godot rect3 object
//	var godotRect3 C.godot_rect3
//
//	// Create our rect3
//	C.godot_rect3_new(&godotRect3)
//
//	// Internally set our rect3
//	rect3.rect3 = &godotRect3
//
//	return rect3
//}

// TODO: Finish implementing this
type Rect3 struct {
	rect3 *C.godot_rect3
}

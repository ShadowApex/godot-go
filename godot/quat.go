package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Vector3
//func NewQuat() *Quat {
//	quat := &Quat{}
//
//	// Create our godot quat object
//	var godotQuat C.godot_quat
//
//	// Create our quat
//	C.godot_quat_new(&godotQuat)
//
//	// Internally set our quat
//	quat.quat = &godotQuat
//
//	return quat
//}

// TODO: Finish implementing this
type Quat struct {
	quat *C.godot_quat
}

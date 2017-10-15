package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

//DEPENDS: Color, Vector2, Vector3, Real, Array
// TODO: Finish implementing these
type PoolByteArray struct {
	array *C.godot_pool_byte_array
}

// TODO: Finish implementing these
type PoolIntArray struct {
	array *C.godot_pool_int_array
}

// TODO: Finish implementing these
type PoolRealArray struct {
	array *C.godot_pool_real_array
}

// TODO: Finish implementing these
type PoolStringArray struct {
	array *C.godot_pool_string_array
}

// TODO: Finish implementing these
type PoolVector2Array struct {
	array *C.godot_pool_vector2_array
}

// TODO: Finish implementing these
type PoolVector3Array struct {
	array *C.godot_pool_vector3_array
}

// TODO: Finish implementing these
type PoolColorArray struct {
	array *C.godot_pool_color_array
}

package godot

/*
#include <gdnative/gdnative.h>
*/
import "C"

type Object struct {
	godotObj *C.godot_object
}

type Bool struct {
	godotObj C.godot_bool
}

type Int struct {
	godotObj C.godot_int
}

type Float struct {
	godotObj C.godot_real
}

type MethodBind struct {
	godotObj C.godot_method_bind
}

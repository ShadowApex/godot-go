package gdnative

/*
#include <gdnative/array.h>
*/
import "C"

type Array struct {
	base C.godot_array
}

func (a *Array) getBase() C.godot_array {
	return a.base
}

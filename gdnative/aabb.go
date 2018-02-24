package gdnative

/*
#include <gdnative/aabb.h>
*/
import "C"

type AABB struct {
	base C.godot_aabb
}

func (a *AABB) getBase() C.godot_aabb {
	return a.base
}

package gdnative

/*
#include <gdnative/basis.h>
*/
import "C"

type Basis struct {
	base C.godot_basis
}

func (a *Basis) getBase() C.godot_basis {
	return a.base
}

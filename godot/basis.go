package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Quat, Vector3
func NewBasis() *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis
	C.godot_basis_new(&godotBasis)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// TODO: Finish implementing this
type Basis struct {
	basis *C.godot_basis
}

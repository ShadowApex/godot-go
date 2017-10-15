package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

//DEPENDS: Vector3
func NewPlane(a, b, c, d float64) *Plane {
	plane := &Plane{}

	// Create our godot plane object
	var godotPlane C.godot_plane

	// Create our plane
	C.godot_plane_new_with_reals(
		&godotPlane,
		realAsGodotReal(a),
		realAsGodotReal(b),
		realAsGodotReal(c),
		realAsGodotReal(d),
	)

	// Internally set our plane
	plane.plane = &godotPlane

	return plane
}

// TODO: Finish implementing this
type Plane struct {
	plane *C.godot_plane
}

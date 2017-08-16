package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

func NewVariant() *Variant {
	return &Variant{}
}

type Variant struct {
}

func (v *Variant) Value() C.godot_variant {
	var godot_variant C.godot_variant
	if v == nil {
		C.godot_variant_new_nil(&godot_variant)
	}

	return godot_variant
}

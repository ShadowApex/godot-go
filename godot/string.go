package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

// TODO: We might not need a "string" struct. Just convert built-in strings.
func NewString() {

	var godotString C.godot_string
	C.godot_string_new(&godotString)
}

type String struct {
}

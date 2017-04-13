package godot

/*
#cgo CXXFLAGS: -I/usr/local/include -std=c11
#include <stddef.h>
#include <godot/godot_array.h>
*/
import "C"

// NewArray returns a new godot array.
func NewArray() *Array {
	var array C.godot_array
	//C.godot_array_new(&array)

	return &Array{
		array: &array,
	}
}

// Array is a structure for Godot arrays.
type Array struct {
	array *C.godot_array
}

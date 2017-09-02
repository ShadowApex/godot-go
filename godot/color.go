package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

func NewColorRGB(r, g, b float64) *Color {
	color := &Color{}

	// Create our godot color object
	var godotColor C.godot_color

	// Create our color
	C.godot_color_new_rgb(&godotColor, realAsGodotReal(r), realAsGodotReal(g), realAsGodotReal(b))

	// Internally set our color
	color.color = &godotColor

	return color
}

// TODO: Finish implementing this
type Color struct {
	color *C.godot_color
}

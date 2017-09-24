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
	C.godot_color_new_rgb(
		&godotColor,
		realAsGodotReal(r),
		realAsGodotReal(g),
		realAsGodotReal(b),
	)

	// Internally set our color
	color.color = &godotColor

	return color
}
func NewColorRGBA(r, g, b, a float64) *Color {
	color := &Color{}

	// Create our godot color object
	var godotColor C.godot_color

	// Create our color
	C.godot_color_new_rgba(
		&godotColor,
		realAsGodotReal(r),
		realAsGodotReal(g),
		realAsGodotReal(b),
		realAsGodotReal(a),
	)

	// Internally set our color
	color.color = &godotColor

	return color
}

type Color struct {
	color *C.godot_color
}

// AsString
func (c *Color) AsString() string {
	asString := C.godot_color_as_string(c.color)
	return godotStringAsString(&asString)
}

//Blend
func (c *Color) Blend(over Color) *Color {
	return godotColorAsColor(
		C.godot_color_blend(
			c.color,
			over.color,
		),
	)
}

// Contrasted
func (c *Color) Contrasted() *Color {
	return godotColorAsColor(C.godot_color_contrasted(c.color))
}

//LinearInterpolate
func (c *Color) LinearInterpolate(b Color, t float64) *Color {
	return godotColorAsColor(
		C.godot_color_linear_interpolate(
			c.color,
			b.color,
			realAsGodotReal(t),
		),
	)
}

// GetA
func (c *Color) GetA() float64 {
	return godotRealAsReal(C.godot_color_get_a(c.color))
}

// GetB
func (c *Color) GetB() float64 {
	return godotRealAsReal(C.godot_color_get_b(c.color))
}

// GetG
func (c *Color) GetG() float64 {
	return godotRealAsReal(C.godot_color_get_g(c.color))
}

// GetH
func (c *Color) GetH() float64 {
	return godotRealAsReal(C.godot_color_get_h(c.color))
}

// GetR
func (c *Color) GetR() float64 {
	return godotRealAsReal(C.godot_color_get_r(c.color))
}

// GetS
func (c *Color) GetS() float64 {
	return godotRealAsReal(C.godot_color_get_s(c.color))
}

// GetV
func (c *Color) GetV() float64 {
	return godotRealAsReal(C.godot_color_get_v(c.color))
}

// Gray
func (c *Color) Gray() float64 {
	return godotRealAsReal(C.godot_color_gray(c.color))
}

// Inverted
func (c *Color) Inverted() *Color {
	return godotColorAsColor(C.godot_color_inverted(c.color))
}

// OperatorEqual
func (c *Color) OperatorEqual(with Color) bool {
	return godotBoolAsBool(
		C.godot_color_operator_equal(
			c.color,
			with.color,
		),
	)
}

// OperartorGreater
func (c *Color) OperatorGreater(with Color) bool {
	if !c.OperatorEqual(with) && !c.OperatorLess(with) {
		return true
	}
	return false
}

// OperatorLess
func (c *Color) OperatorLess(with Color) bool {
	return godotBoolAsBool(
		C.godot_color_operator_less(
			c.color,
			with.color,
		),
	)
}

// SetA
func (c *Color) SetA(alpha float64) {
	C.godot_color_set_a(c.color, realAsGodotReal(alpha))
}

// SetB
func (c *Color) SetB(blue float64) {
	C.godot_color_set_b(c.color, realAsGodotReal(blue))
}

// SetG
func (c *Color) SetG(green float64) {
	C.godot_color_set_g(c.color, realAsGodotReal(green))
}

// SetR
func (c *Color) SetR(red float64) {
	C.godot_color_set_r(c.color, realAsGodotReal(red))
}

// To32
func (c *Color) To32() int64 {
	return int64(C.godot_color_to_32(c.color))
}

// ToARGB32
func (c *Color) ToARGB32() int64 {
	return int64(C.godot_color_to_ARGB32(c.color))
}

// ToHTML
func (c *Color) ToHTML(withAlpha bool) string {
	toHTML := C.godot_color_to_html(
		c.color,
		boolAsGodotBool(withAlpha),
	)
	return godotStringAsString(&toHTML)
}

package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

// NewRect2 constructs a Rect2 by x, y, width and height.
func NewRect2(x float64, y float64, width float64, height float64) *Rect2 {
	rect2 := &Rect2{}

	// Create our godot rect2 object
	var godotRect2 C.godot_rect2

	// Create our rect2
	C.godot_rect2_new(
		&godotRect2,
		realAsGodotReal(x),
		realAsGodotReal(y),
		realAsGodotReal(width),
		realAsGodotReal(height),
	)

	// Internally set our rect2
	rect2.rect2 = &godotRect2

	return rect2
}

// NewRect2WithPositionAndSize constructs a Rect2 by position and size.
func NewRect2WithPositionAndSize(pos Vector2, size Vector2) *Rect2 {
	rect2 := &Rect2{}

	// Create our godot rect2 object
	var godotRect2 C.godot_rect2

	// Create our rect2
	C.godot_rect2_new_with_position_and_size(
		&godotRect2,
		vec2AsGodotVec2(pos),
		vec2AsGodotVec2(size),
	)

	// Internally set our rect2
	rect2.rect2 = &godotRect2

	return rect2
}

type Rect2 struct {
	rect2 *C.godot_rect2
}

// AsString returns a string of the Rect2.
func (r *Rect2) AsString() string {
	asString := C.godot_rect2_as_string(r.rect2)
	return godotStringAsString(&asString)
}

// Clip returns the intersection of this Rect2 and with.
func (r *Rect2) Clip(with Rect2) *Rect2 {
	return godotRect2AsRect2(
		C.godot_rect2_clip(
			r.rect2,
			with.rect2,
		),
	)
}

// Encloses returns true if r.rect completely surrounds with.rect2.
func (r *Rect2) Encloses(with Rect2) bool {
	return godotBoolAsBool(C.godot_rect2_encloses(r.rect2, with.rect2))
}

// Expand returns r.rect2 expanded to include a given point.
func (r *Rect2) Expand(to Vector2) *Rect2 {
	return godotRect2AsRect2(
		C.godot_rect2_expand(
			r.rect2,
			to.vector2,
		),
	)
}

// GetArea returns the area of r.rect2
func (r *Rect2) GetArea() float64 {
	return godotRealAsReal(C.godot_rect2_get_area(r.rect2))
}

// GetPosition returns the X and Y coordinates of r.rect2
func (r *Rect2) GetPosition() *Vector2 {
	return godotVec2AsVec2(C.godot_rect2_get_position(r.rect2))
}

// GetSize returns the width and height of r.rect2
func (r *Rect2) GetSize() *Vector2 {
	return godotVec2AsVec2(C.godot_rect2_get_size(r.rect2))
}

// Grow returns a copy of r.rect2 grown a given amount of units towards all sides.
func (r *Rect2) Grow(by float64) *Rect2 {
	return godotRect2AsRect2(
		C.godot_rect2_grow(
			r.rect2,
			realAsGodotReal(by),
		),
	)
}

// HasNoArea returns true if r.rect2 is flat or empty.
func (r *Rect2) HasNoArea() bool {
	return godotBoolAsBool(C.godot_rect2_has_no_area(r.rect2))
}

// HasPoint returns true if r.rect2 contains a point.
func (r *Rect2) HasPoint(with Vector2) bool {
	return godotBoolAsBool(C.godot_rect2_has_point(r.rect2, with.vector2))
}

// Intersects returns true if the r.rect2 overlaps with with.rect2
func (r *Rect2) Intersects(with Rect2) bool {
	return godotBoolAsBool(C.godot_rect2_intersects(r.rect2, with.rect2))
}

// Merge returns a larger Rect2 t hat contains both r.rect2 and with.rect2
func (r *Rect2) Merge(with Rect2) *Rect2 {
	return godotRect2AsRect2(
		C.godot_rect2_merge(
			r.rect2,
			with.rect2,
		),
	)
}

// OperatorEqual returns true if r.rect2 == with.rect2
func (r *Rect2) OperatorEqual(with Rect2) bool {
	return godotBoolAsBool(C.godot_rect2_operator_equal(r.rect2, with.rect2))
}

// SetPosition sets the X and Y of r.rect2
func (r *Rect2) SetPosition(pos Vector2) {
	C.godot_rect2_set_position(r.rect2, pos.vector2)
}

// SetSize sets the width and height of r.rect2
func (r *Rect2) SetSize(size Vector2) {
	C.godot_rect2_set_size(r.rect2, size.vector2)
}

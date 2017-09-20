package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

// NewRect3 constructs a Rect3.
func NewRect3(pos Vector3, size Vector3) *Rect3 {
	rect3 := &Rect3{}

	// Create our godot rect3 object
	var godotRect3 C.godot_rect3

	// Create our rect3
	C.godot_rect3_new(
		&godotRect3,
		vec3AsGodotVec3(pos),
		vec3AsGodotVec3(size),
	)

	// Internally set our rect3
	rect3.rect3 = &godotRect3

	return rect3
}

type Rect3 struct {
	rect3 *C.godot_rect3
}

// AsString returns a string of the Rect3.
func (r *Rect3) AsString() string {
	asString := C.godot_rect3_as_string(r.rect3)
	return godotStringAsString(&asString)
}

// Encloses returns true if r.rect completely surrounds with.rect3.
func (r *Rect3) Encloses(with Rect3) bool {
	return godotBoolAsBool(C.godot_rect3_encloses(r.rect3, with.rect3))
}

// Expand returns r.rect3 expanded to include a given point.
func (r *Rect3) Expand(to Vector3) *Rect3 {
	return godotRect3AsRect3(
		C.godot_rect3_expand(
			r.rect3,
			to.vector3,
		),
	)
}

// GetArea returns the area of r.rect3
func (r *Rect3) GetArea() float64 {
	return godotRealAsReal(C.godot_rect3_get_area(r.rect3))
}

// Get Endpoint returns
func (r *Rect3) GetEndpoint(index int64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_rect3_get_endpoint(
			r.rect3,
			intAsGodotInt(index),
		),
	)
}

// GetLongestAxis returns the longest axis of r.rect3
func (r *Rect3) GetLongestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_rect3_get_longest_axis(r.rect3))
}

// GetLongestAxisIndex returns the index of longest axis of r.rect3
func (r *Rect3) GetLongestAxisIndex() int64 {
	return int64(C.godot_rect3_get_longest_axis_index(r.rect3))
}

// GetLongestAxisSize returns the size of longest axis of r.rect3
func (r *Rect3) GetLongestAxisSize() float64 {
	return godotRealAsReal(C.godot_rect3_get_longest_axis_size(r.rect3))
}

// GetPosition returns the X, Y, and Z coordinates of r.rect3
func (r *Rect3) GetPosition() *Vector3 {
	return godotVec3AsVec3(C.godot_rect3_get_position(r.rect3))
}

// GetShortestAxis returns the shortest axis of r.rect3
func (r *Rect3) GetShortestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_rect3_get_shortest_axis(r.rect3))
}

// GetShortestAxisIndex returns the index of shortest axis of r.rect3
func (r *Rect3) GetShortestAxisIndex() int64 {
	return int64(C.godot_rect3_get_shortest_axis_index(r.rect3))
}

// GetShortestAxisSize returns the size of shortest axis of r.rect3
func (r *Rect3) GetShortestAxisSize() float64 {
	return godotRealAsReal(C.godot_rect3_get_shortest_axis_size(r.rect3))
}

// GetSize returns the width, height, and depth of r.rect3
func (r *Rect3) GetSize() *Vector3 {
	return godotVec3AsVec3(C.godot_rect3_get_size(r.rect3))
}

// GetSupport returns
func (r *Rect3) GetSupport(dir Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_rect3_get_support(r.rect3, dir.vector3))
}

// Grow returns a copy of r.rect3 grown a given amount of units towards all sides.
func (r *Rect3) Grow(by float64) *Rect3 {
	return godotRect3AsRect3(
		C.godot_rect3_grow(
			r.rect3,
			realAsGodotReal(by),
		),
	)
}

// HasNoArea returns true if r.rect3 is flat or empty.
func (r *Rect3) HasNoArea() bool {
	return godotBoolAsBool(C.godot_rect3_has_no_area(r.rect3))
}

// HasNoSurface returns true if r.rect3 is flat or empty.
func (r *Rect3) HasNoSurface() bool {
	return godotBoolAsBool(C.godot_rect3_has_no_surface(r.rect3))
}

// HasPoint returns true if r.rect3 contains a point.
func (r *Rect3) HasPoint(with Vector3) bool {
	return godotBoolAsBool(C.godot_rect3_has_point(r.rect3, with.vector3))
}

// Intersection returns
func (r *Rect3) Intersection(with Rect3) *Rect3 {
	return godotRect3AsRect3(
		C.godot_rect3_intersection(
			r.rect3,
			with.rect3,
		),
	)
}

// Intersects returns true if the r.rect3 overlaps with with.rect3
func (r *Rect3) Intersects(with Rect3) bool {
	return godotBoolAsBool(C.godot_rect3_intersects(r.rect3, with.rect3))
}

// IntersectsPlane returns true if the r.rect3 overlaps with with.plane
func (r *Rect3) IntersectsPlane(with Plane) bool {
	return godotBoolAsBool(C.godot_rect3_intersects_plane(r.rect3, with.plane))
}

// IntersectsSegment returns true if the r.rect3 overlaps a line from from to to.
func (r *Rect3) IntersectsSegment(from Vector3, to Vector3) bool {
	return godotBoolAsBool(
		C.godot_rect3_intersects_segment(
			r.rect3,
			from.vector3,
			to.vector3,
		),
	)
}

// Merge returns a larger Rect3 t hat contains both r.rect3 and with.rect3
func (r *Rect3) Merge(with Rect3) *Rect3 {
	return godotRect3AsRect3(
		C.godot_rect3_merge(
			r.rect3,
			with.rect3,
		),
	)
}

// OperatorEqual returns true if r.rect3 == with.rect3
func (r *Rect3) OperatorEqual(with Rect3) bool {
	return godotBoolAsBool(C.godot_rect3_operator_equal(r.rect3, with.rect3))
}

// SetPosition sets the X, Y, and Z of r.rect3
func (r *Rect3) SetPosition(pos Vector3) {
	C.godot_rect3_set_position(r.rect3, pos.vector3)
}

// SetSize sets the width, height, and depth of r.rect3
func (r *Rect3) SetSize(size Vector3) {
	C.godot_rect3_set_size(r.rect3, size.vector3)
}

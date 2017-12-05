package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

// NewAABB constructs a AABB.
func NewAABB(pos Vector3, size Vector3) *AABB {
	aabb := &AABB{}

	// Create our godot aabb object
	var godotAABB C.godot_aabb

	// Create our aabb
	C.godot_aabb_new(
		&godotAABB,
		vec3AsGodotVec3(pos),
		vec3AsGodotVec3(size),
	)

	// Internally set our aabb
	aabb.aabb = &godotAABB

	return aabb
}

type AABB struct {
	aabb *C.godot_aabb
}

// AsString returns a string of the AABB.
func (r *AABB) AsString() string {
	asString := C.godot_aabb_as_string(r.aabb)
	return godotStringAsString(&asString)
}

// Encloses returns true if r.rect completely surrounds with.aabb.
func (r *AABB) Encloses(with AABB) bool {
	return godotBoolAsBool(C.godot_aabb_encloses(r.aabb, with.aabb))
}

// Expand returns r.aabb expanded to include a given point.
func (r *AABB) Expand(to Vector3) *AABB {
	return godotAABBAsAABB(
		C.godot_aabb_expand(
			r.aabb,
			to.vector3,
		),
	)
}

// GetArea returns the area of r.aabb
func (r *AABB) GetArea() float64 {
	return godotRealAsReal(C.godot_aabb_get_area(r.aabb))
}

// Get Endpoint returns
func (r *AABB) GetEndpoint(index int64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_aabb_get_endpoint(
			r.aabb,
			intAsGodotInt(index),
		),
	)
}

// GetLongestAxis returns the longest axis of r.aabb
func (r *AABB) GetLongestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_longest_axis(r.aabb))
}

// GetLongestAxisIndex returns the index of longest axis of r.aabb
func (r *AABB) GetLongestAxisIndex() int64 {
	return int64(C.godot_aabb_get_longest_axis_index(r.aabb))
}

// GetLongestAxisSize returns the size of longest axis of r.aabb
func (r *AABB) GetLongestAxisSize() float64 {
	return godotRealAsReal(C.godot_aabb_get_longest_axis_size(r.aabb))
}

// GetPosition returns the X, Y, and Z coordinates of r.aabb
func (r *AABB) GetPosition() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_position(r.aabb))
}

// GetShortestAxis returns the shortest axis of r.aabb
func (r *AABB) GetShortestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_shortest_axis(r.aabb))
}

// GetShortestAxisIndex returns the index of shortest axis of r.aabb
func (r *AABB) GetShortestAxisIndex() int64 {
	return int64(C.godot_aabb_get_shortest_axis_index(r.aabb))
}

// GetShortestAxisSize returns the size of shortest axis of r.aabb
func (r *AABB) GetShortestAxisSize() float64 {
	return godotRealAsReal(C.godot_aabb_get_shortest_axis_size(r.aabb))
}

// GetSize returns the width, height, and depth of r.aabb
func (r *AABB) GetSize() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_size(r.aabb))
}

// GetSupport returns
func (r *AABB) GetSupport(dir Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_support(r.aabb, dir.vector3))
}

// Grow returns a copy of r.aabb grown a given amount of units towards all sides.
func (r *AABB) Grow(by float64) *AABB {
	return godotAABBAsAABB(
		C.godot_aabb_grow(
			r.aabb,
			realAsGodotReal(by),
		),
	)
}

// HasNoArea returns true if r.aabb is flat or empty.
func (r *AABB) HasNoArea() bool {
	return godotBoolAsBool(C.godot_aabb_has_no_area(r.aabb))
}

// HasNoSurface returns true if r.aabb is flat or empty.
func (r *AABB) HasNoSurface() bool {
	return godotBoolAsBool(C.godot_aabb_has_no_surface(r.aabb))
}

// HasPoint returns true if r.aabb contains a point.
func (r *AABB) HasPoint(with Vector3) bool {
	return godotBoolAsBool(C.godot_aabb_has_point(r.aabb, with.vector3))
}

// Intersection returns
func (r *AABB) Intersection(with AABB) *AABB {
	return godotAABBAsAABB(
		C.godot_aabb_intersection(
			r.aabb,
			with.aabb,
		),
	)
}

// Intersects returns true if the r.aabb overlaps with with.aabb
func (r *AABB) Intersects(with AABB) bool {
	return godotBoolAsBool(C.godot_aabb_intersects(r.aabb, with.aabb))
}

// IntersectsPlane returns true if the r.aabb overlaps with with.plane
func (r *AABB) IntersectsPlane(with Plane) bool {
	return godotBoolAsBool(C.godot_aabb_intersects_plane(r.aabb, with.plane))
}

// IntersectsSegment returns true if the r.aabb overlaps a line from from to to.
func (r *AABB) IntersectsSegment(from Vector3, to Vector3) bool {
	return godotBoolAsBool(
		C.godot_aabb_intersects_segment(
			r.aabb,
			from.vector3,
			to.vector3,
		),
	)
}

// Merge returns a larger AABB t hat contains both r.aabb and with.aabb
func (r *AABB) Merge(with AABB) *AABB {
	return godotAABBAsAABB(
		C.godot_aabb_merge(
			r.aabb,
			with.aabb,
		),
	)
}

// OperatorEqual returns true if r.aabb == with.aabb
func (r *AABB) OperatorEqual(with AABB) bool {
	return godotBoolAsBool(C.godot_aabb_operator_equal(r.aabb, with.aabb))
}

// SetPosition sets the X, Y, and Z of r.aabb
func (r *AABB) SetPosition(pos Vector3) {
	C.godot_aabb_set_position(r.aabb, pos.vector3)
}

// SetSize sets the width, height, and depth of r.aabb
func (r *AABB) SetSize(size Vector3) {
	C.godot_aabb_set_size(r.aabb, size.vector3)
}

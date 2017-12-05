package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

// NewAabb constructs a Aabb.
func NewAabb(pos Vector3, size Vector3) *Aabb {
	aabb := &Aabb{}

	// Create our godot aabb object
	var godotAabb C.godot_aabb

	// Create our aabb
	C.godot_aabb_new(
		&godotAabb,
		vec3AsGodotVec3(pos),
		vec3AsGodotVec3(size),
	)

	// Internally set our aabb
	aabb.aabb = &godotAabb

	return aabb
}

type Aabb struct {
	aabb *C.godot_aabb
}

// AsString returns a string of the Aabb.
func (r *Aabb) AsString() string {
	asString := C.godot_aabb_as_string(r.aabb)
	return godotStringAsString(&asString)
}

// Encloses returns true if r.rect completely surrounds with.aabb.
func (r *Aabb) Encloses(with Aabb) bool {
	return godotBoolAsBool(C.godot_aabb_encloses(r.aabb, with.aabb))
}

// Expand returns r.aabb expanded to include a given point.
func (r *Aabb) Expand(to Vector3) *Aabb {
	return godotAabbAsAabb(
		C.godot_aabb_expand(
			r.aabb,
			to.vector3,
		),
	)
}

// GetArea returns the area of r.aabb
func (r *Aabb) GetArea() float64 {
	return godotRealAsReal(C.godot_aabb_get_area(r.aabb))
}

// Get Endpoint returns
func (r *Aabb) GetEndpoint(index int64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_aabb_get_endpoint(
			r.aabb,
			intAsGodotInt(index),
		),
	)
}

// GetLongestAxis returns the longest axis of r.aabb
func (r *Aabb) GetLongestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_longest_axis(r.aabb))
}

// GetLongestAxisIndex returns the index of longest axis of r.aabb
func (r *Aabb) GetLongestAxisIndex() int64 {
	return int64(C.godot_aabb_get_longest_axis_index(r.aabb))
}

// GetLongestAxisSize returns the size of longest axis of r.aabb
func (r *Aabb) GetLongestAxisSize() float64 {
	return godotRealAsReal(C.godot_aabb_get_longest_axis_size(r.aabb))
}

// GetPosition returns the X, Y, and Z coordinates of r.aabb
func (r *Aabb) GetPosition() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_position(r.aabb))
}

// GetShortestAxis returns the shortest axis of r.aabb
func (r *Aabb) GetShortestAxis() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_shortest_axis(r.aabb))
}

// GetShortestAxisIndex returns the index of shortest axis of r.aabb
func (r *Aabb) GetShortestAxisIndex() int64 {
	return int64(C.godot_aabb_get_shortest_axis_index(r.aabb))
}

// GetShortestAxisSize returns the size of shortest axis of r.aabb
func (r *Aabb) GetShortestAxisSize() float64 {
	return godotRealAsReal(C.godot_aabb_get_shortest_axis_size(r.aabb))
}

// GetSize returns the width, height, and depth of r.aabb
func (r *Aabb) GetSize() *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_size(r.aabb))
}

// GetSupport returns
func (r *Aabb) GetSupport(dir Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_aabb_get_support(r.aabb, dir.vector3))
}

// Grow returns a copy of r.aabb grown a given amount of units towards all sides.
func (r *Aabb) Grow(by float64) *Aabb {
	return godotAabbAsAabb(
		C.godot_aabb_grow(
			r.aabb,
			realAsGodotReal(by),
		),
	)
}

// HasNoArea returns true if r.aabb is flat or empty.
func (r *Aabb) HasNoArea() bool {
	return godotBoolAsBool(C.godot_aabb_has_no_area(r.aabb))
}

// HasNoSurface returns true if r.aabb is flat or empty.
func (r *Aabb) HasNoSurface() bool {
	return godotBoolAsBool(C.godot_aabb_has_no_surface(r.aabb))
}

// HasPoint returns true if r.aabb contains a point.
func (r *Aabb) HasPoint(with Vector3) bool {
	return godotBoolAsBool(C.godot_aabb_has_point(r.aabb, with.vector3))
}

// Intersection returns
func (r *Aabb) Intersection(with Aabb) *Aabb {
	return godotAabbAsAabb(
		C.godot_aabb_intersection(
			r.aabb,
			with.aabb,
		),
	)
}

// Intersects returns true if the r.aabb overlaps with with.aabb
func (r *Aabb) Intersects(with Aabb) bool {
	return godotBoolAsBool(C.godot_aabb_intersects(r.aabb, with.aabb))
}

// IntersectsPlane returns true if the r.aabb overlaps with with.plane
func (r *Aabb) IntersectsPlane(with Plane) bool {
	return godotBoolAsBool(C.godot_aabb_intersects_plane(r.aabb, with.plane))
}

// IntersectsSegment returns true if the r.aabb overlaps a line from from to to.
func (r *Aabb) IntersectsSegment(from Vector3, to Vector3) bool {
	return godotBoolAsBool(
		C.godot_aabb_intersects_segment(
			r.aabb,
			from.vector3,
			to.vector3,
		),
	)
}

// Merge returns a larger Aabb t hat contains both r.aabb and with.aabb
func (r *Aabb) Merge(with Aabb) *Aabb {
	return godotAabbAsAabb(
		C.godot_aabb_merge(
			r.aabb,
			with.aabb,
		),
	)
}

// OperatorEqual returns true if r.aabb == with.aabb
func (r *Aabb) OperatorEqual(with Aabb) bool {
	return godotBoolAsBool(C.godot_aabb_operator_equal(r.aabb, with.aabb))
}

// SetPosition sets the X, Y, and Z of r.aabb
func (r *Aabb) SetPosition(pos Vector3) {
	C.godot_aabb_set_position(r.aabb, pos.vector3)
}

// SetSize sets the width, height, and depth of r.aabb
func (r *Aabb) SetSize(size Vector3) {
	C.godot_aabb_set_size(r.aabb, size.vector3)
}

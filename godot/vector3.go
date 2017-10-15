package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

// Duplication of enumerated constants from godot.
type Vec3Axis int

const (
	VECTOR3_AXIS_X = iota // 0
	VECTOR3_AXIS_Y        // 1
	VECTOR3_AXIS_Z        // 2
)

// NewVector3 constructs a new Vector3 from the given x and y.
func NewVector3(x float64, y float64, z float64) *Vector3 {
	vector3 := &Vector3{}

	// Create our godot vector3 object
	var godotVector3 C.godot_vector3

	// Create our vector3
	C.godot_vector3_new(
		&godotVector3,
		realAsGodotReal(x),
		realAsGodotReal(y),
		realAsGodotReal(z),
	)

	// Internally set our vector3
	vector3.vector3 = &godotVector3

	return vector3
}

type Vector3 struct {
	vector3 *C.godot_vector3
}

// Abs returns the absolute value of the vector.
func (v *Vector3) Abs() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_abs(v.vector3))
}

// AngleTo returns the angle in radians between the two vectors.
func (v *Vector3) AngleTo(to Vector3) float64 {
	angleTo := C.godot_vector3_angle_to(v.vector3, to.vector3)
	return godotRealAsReal(angleTo)
}

// AsString returns a string.
func (v *Vector3) AsString() string {
	asString := C.godot_vector3_as_string(v.vector3)
	return godotStringAsString(&asString)
}

// Bounce is not documented in the godot library yet.
func (v *Vector3) Bounce(with Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_bounce(v.vector3, with.vector3))
}

// Ceil returns a new vector with all components rounded up.
func (v *Vector3) Ceil() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_ceil(v.vector3))
}

// Cross returns the cross product with b.
func (v *Vector3) Cross(b Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_cross(v.vector3, b.vector3))
}

// CubicInterpolate Cubicly interpolates between this Vector and “b”, using “pre_a”
// and “post_b” as handles, and returning the result at position “t”.
func (v *Vector3) CubicInterpolate(b Vector3, preA Vector3, postB Vector3, t float64) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_cubic_interpolate(
		v.vector3,
		b.vector3,
		preA.vector3,
		postB.vector3,
		realAsGodotReal(t),
	),
	)
}

// DistanceSquaredTo returns the squared distance to vector “b”. Prefer this
// function over “distance_to” if you need to sort vectors or need the squared
// distance for some formula.
func (v *Vector3) DistanceSquaredTo(to Vector3) float64 {
	distSqrTo := C.godot_vector3_distance_squared_to(v.vector3, to.vector3)
	return godotRealAsReal(distSqrTo)
}

// DistanceTo returns the distance to vector “b”.
func (v *Vector3) DistanceTo(to Vector3) float64 {
	distTo := C.godot_vector3_distance_to(v.vector3, to.vector3)
	return godotRealAsReal(distTo)
}

// Dot returns the dot product with vector “b”.
func (v *Vector3) Dot(with Vector3) float64 {
	dot := C.godot_vector3_dot(v.vector3, with.vector3)
	return godotRealAsReal(dot)
}

// Floor removes the fractional part of x and y.
func (v *Vector3) Floor() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_floor(v.vector3))
}

// GetAxis returns the value of the supplied axis
func (v *Vector3) GetAxis(axis Vec3Axis) float64 {
	return godotRealAsReal(
		C.godot_vector3_get_axis(
			v.vector3,
			vec3AxisAsGodotVec3Axis(axis),
		),
	)
}

// Inverse returns the inverse of the vector. This is the same as i
// Vector3( 1.0 / x, 1.0 / y, 1.0 / z )
func (v *Vector3) Inverse() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_inverse(v.vector3))
}

// IsNormalized
func (v *Vector3) IsNormalized() bool {
	isNormalized := C.godot_vector3_is_normalized(v.vector3)
	return godotBoolAsBool(isNormalized)
}

// Length returns the length of the vector.
func (v *Vector3) Length() float64 {
	return godotRealAsReal(C.godot_vector3_length(v.vector3))
}

// LengthSquared returns the squared length of the vector. Prefer this function
// over “length” if you need to sort vectors or need the squared length for some formula.
func (v *Vector3) LengthSquared() float64 {
	return godotRealAsReal(C.godot_vector3_length_squared(v.vector3))
}

// LinearInterpolate returns the result of the linear interpolation between this
// vector and “b”, by amount “t”.
func (v *Vector3) LinearInterpolate(b Vector3, t float64) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_linear_interpolate(
		v.vector3,
		b.vector3,
		realAsGodotReal(t),
	),
	)
}

// MaxAxis returns VECTOR3_AXIS_X, VECTOR3_AXIS_Y or VECTOR3_AXIS_Z depending on
// which axis is the largest.
func (v *Vector3) MaxAxis() int64 {
	return godotIntAsInt(C.godot_vector3_max_axis(v.vector3))
}

// MinAxis returns VECTOR3_AXIS_X, VECTOR3_AXIS_Y or VECTOR3_AXIS_Z depending on
// which axis is the smallest.
func (v *Vector3) MinAxis() int64 {
	return godotIntAsInt(C.godot_vector3_min_axis(v.vector3))
}

// Normalized returns a normalized vector to unit length.
func (v *Vector3) Normalized() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_normalized(v.vector3))
}

// OperatorAdd adds the current vector and with.
func (v *Vector3) OperatorAdd(with Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_operator_add(v.vector3, with.vector3))
}

// OperatorDivideScalar divides the current vector by with.
func (v *Vector3) OperatorDivideScalar(with float64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_operator_divide_scalar(
			v.vector3,
			realAsGodotReal(with),
		),
	)
}

// OperatorDivideVector multiplys the current vector and with.
func (v *Vector3) OperatorDivideVector(with Vector3) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_operator_divide_vector(
			v.vector3,
			with.vector3,
		),
	)
}

// OperatorEqual compares the current vector and with. Returns true if equal
func (v *Vector3) OperatorEqual(with Vector3) bool {
	operatorEqual := C.godot_vector3_operator_equal(v.vector3, with.vector3)
	return godotBoolAsBool(operatorEqual)
}

// OperatorLess compares the current vector and with. Returns true if v.vector3 < with.vector
func (v *Vector3) OperatorLess(with Vector3) bool {
	operatorEqual := C.godot_vector3_operator_less(v.vector3, with.vector3)
	return godotBoolAsBool(operatorEqual)
}

// OperatorMultiplyScalar multiplys the current vector by with.
func (v *Vector3) OperatorMultiplyScalar(with float64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_operator_multiply_scalar(
			v.vector3,
			realAsGodotReal(with),
		),
	)
}

// OperatorMultiplyVector multiplys the current vector and with.
func (v *Vector3) OperatorMultiplyVector(with Vector3) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_operator_multiply_vector(
			v.vector3,
			with.vector3,
		),
	)
}

// OperatorNeg returns -v.vector3.
func (v *Vector3) OperatorNeg() *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_operator_neg(v.vector3))
}

// OperatorSubtract subtracts the current vector and with.
func (v *Vector3) OperatorSubtract(with Vector3) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_operator_substract(
			v.vector3,
			with.vector3,
		),
	)
}

// Outer
func (v *Vector3) Outer(with Vector3) *Basis {
	return godotBasisAsBasis(C.godot_vector3_outer(v.vector3, with.vector3))
}

// Reflect is like “slide”, but reflects the Vector instead of continuing along the wall.
func (v *Vector3) Reflect(vec Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_reflect(v.vector3, vec.vector3))
}

// Rotated rotates the vector by “phi” radians.
func (v *Vector3) Rotated(by Vector3, phi float64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_vector3_rotated(
			v.vector3,
			by.vector3,
			realAsGodotReal(phi),
		),
	)
}

// SetAxis changes the specified axis to the specified value.
func (v *Vector3) SetAxis(axis Vec3Axis, val float64) {
	C.godot_vector3_set_axis(
		v.vector3,
		vec3AxisAsGodotVec3Axis(axis),
		realAsGodotReal(val),
	)
}

// Slide slides the vector by the other vector.
func (v *Vector3) Slide(by Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_slide(v.vector3, by.vector3))
}

// Snapped snaps the vector to a grid with the given size.
func (v *Vector3) Snapped(by Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_vector3_snapped(v.vector3, by.vector3))
}

// ToDiagonalMatrix
func (v *Vector3) ToDiagonalMatrix() *Basis {
	return godotBasisAsBasis(C.godot_vector3_to_diagonal_matrix(v.vector3))
}

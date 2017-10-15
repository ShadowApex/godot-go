package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <gdnative/gdnative.h>
#include <nativescript/godot_nativescript.h>
*/
import "C"

//NewVector2 constructs a new Vector2 from the given x and y.
func NewVector2(x float64, y float64) *Vector2 {
	vector2 := &Vector2{}

	// Create our godot vector2 object
	var godotVector2 C.godot_vector2

	// Create our vector2
	C.godot_vector2_new(
		&godotVector2,
		realAsGodotReal(x),
		realAsGodotReal(y),
	)

	// Internally set our vector2
	vector2.vector2 = &godotVector2

	return vector2
}

type Vector2 struct {
	vector2 *C.godot_vector2
}

func (v *Vector2) Abs() *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_abs(v.vector2))
}

//Angle returns the result of atan2 when called with the Vector’s x and y as parameters
//(Math::atan2(x,y)). Be aware that it therefore returns an angle oriented clockwise
//with regard to the (0, 1) unit vector, and not an angle oriented counter-clockwise
//with regard to the (1, 0) unit vector (which would be the typical trigonometric
//representation of the angle when calling Math::atan2(y,x))
func (v *Vector2) Angle() float64 {
	angle := C.godot_vector2_angle(v.vector2)
	return godotRealAsReal(angle)
}

//AngleTo returns the angle in radians between the two vectors.
func (v *Vector2) AngleTo(to Vector2) float64 {
	angleTo := C.godot_vector2_angle_to(v.vector2, to.vector2)
	return godotRealAsReal(angleTo)
}

//AngleToPoint returns the angle in radians between the line connecting the two
//points and the x coordinate.
func (v *Vector2) AngleToPoint(to Vector2) float64 {
	angleToPoint := C.godot_vector2_angle_to_point(v.vector2, to.vector2)
	return godotRealAsReal(angleToPoint)
}

//AsString returns a string.
func (v *Vector2) AsString() string {
	asString := C.godot_vector2_as_string(v.vector2)
	return godotStringAsString(&asString)
}

//Bounce is not documented in the godot library yet.
func (v *Vector2) Bounce(with Vector2) *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_bounce(v.vector2, with.vector2))
}

//Clamped replaces v.vector2 with a new vector2 that is equivalent to v.vector2 * length
func (v *Vector2) Clamped(length float64) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_clamped(
			v.vector2,
			realAsGodotReal(length),
		),
	)
}

//CubicInterpolate Cubicly interpolates between this Vector and “b”, using “pre_a”
//and “post_b” as handles, and returning the result at position “t”.
func (v *Vector2) CubicInterpolate(b Vector2, preA Vector2, postB Vector2, t float64) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_cubic_interpolate(
			v.vector2,
			b.vector2,
			preA.vector2,
			postB.vector2,
			realAsGodotReal(t),
		),
	)
}

//DistanceSquaredTo returns the squared distance to vector “b”. Prefer this
//function over “distance_to” if you need to sort vectors or need the squared
//distance for some formula.
func (v *Vector2) DistanceSquaredTo(to Vector2) float64 {
	distSqrTo := C.godot_vector2_distance_squared_to(v.vector2, to.vector2)
	return godotRealAsReal(distSqrTo)
}

//DistanceTo returns the distance to vector “b”.
func (v *Vector2) DistanceTo(to Vector2) float64 {
	distTo := C.godot_vector2_distance_to(v.vector2, to.vector2)
	return godotRealAsReal(distTo)
}

//Dot returns the dot product with vector “b”.
func (v *Vector2) Dot(with Vector2) float64 {
	dot := C.godot_vector2_dot(v.vector2, with.vector2)
	return godotRealAsReal(dot)
}

//Floor removes the fractional part of x and y.
func (v *Vector2) Floor() *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_floor(v.vector2))
}

//GetAspect returns the ratio of X to Y.
func (v *Vector2) GetAspect() float64 {
	return godotRealAsReal(C.godot_vector2_aspect(v.vector2))
}

//GetX returns v.vector2.x
func (v *Vector2) GetX() float64 {
	return godotRealAsReal(C.godot_vector2_get_x(v.vector2))
}

//GetY returns v.vector2.x
func (v *Vector2) GetY() float64 {
	return godotRealAsReal(C.godot_vector2_get_y(v.vector2))
}

func (v *Vector2) IsNormalized() bool {
	isNormalized := C.godot_vector2_is_normalized(v.vector2)
	return godotBoolAsBool(isNormalized)
}

//Length returns the length of the vector.
func (v *Vector2) Length() float64 {
	return godotRealAsReal(C.godot_vector2_length(v.vector2))
}

//LengthSquared returns the squared length of the vector. Prefer this function
//over “length” if you need to sort vectors or need the squared length for some formula.
func (v *Vector2) LengthSquared() float64 {
	return godotRealAsReal(C.godot_vector2_length_squared(v.vector2))
}

//LinearInterpolate returns the result of the linear interpolation between this
//vector and “b”, by amount “t”.
func (v *Vector2) LinearInterpolate(b Vector2, t float64) *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_linear_interpolate(
		v.vector2,
		b.vector2,
		realAsGodotReal(t),
	),
	)
}

//Normalized returns a normalized vector to unit length.
func (v *Vector2) Normalized() *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_normalized(v.vector2))
}

//OperatorAdd adds the current vector and with.
func (v *Vector2) OperatorAdd(with Vector2) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_add(
			v.vector2,
			with.vector2,
		),
	)
}

//OperatorDivideScalar divides the current vector by with.
func (v *Vector2) OperatorDivideScalar(with float64) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_divide_scalar(
			v.vector2,
			realAsGodotReal(with),
		),
	)
}

//OperatorDivideVector multiplys the current vector and with.
func (v *Vector2) OperatorDivideVector(with Vector2) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_divide_vector(
			v.vector2,
			with.vector2,
		),
	)
}

//OperatorEqual compares the current vector and with. Returns true if v.vector2 == with.vector2
func (v *Vector2) OperatorEqual(with Vector2) bool {
	operatorEqual := C.godot_vector2_operator_equal(v.vector2, with.vector2)
	return godotBoolAsBool(operatorEqual)
}

//OperatorLess compares the current vector and with. Returns true if v.vector2 < with.vector
func (v *Vector2) OperatorLess(with Vector2) bool {
	operatorEqual := C.godot_vector2_operator_less(v.vector2, with.vector2)
	return godotBoolAsBool(operatorEqual)
}

//OperatorMultiplyScalar multiplys the current vector by with.
func (v *Vector2) OperatorMultiplyScalar(with float64) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_multiply_scalar(
			v.vector2,
			realAsGodotReal(with),
		),
	)
}

//OperatorMultiplyVector multiplys the current vector and with.
func (v *Vector2) OperatorMultiplyVector(with Vector2) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_multiply_vector(
			v.vector2,
			with.vector2,
		),
	)
}

//OperatorNeg returns -v.vector2.
func (v *Vector2) OperatorNeg() *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_operator_neg(v.vector2))
}

//OperatorSubtract subtracts the current vector and with.
func (v *Vector2) OperatorSubtract(with Vector2) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_operator_substract(
			v.vector2,
			with.vector2,
		),
	)
}

//Reflect is like “slide”, but reflects the Vector instead of continuing along the wall.
func (v *Vector2) Reflect(vec Vector2) *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_reflect(v.vector2, vec.vector2))
}

//Rotated rotates the vector by “phi” radians.
func (v *Vector2) Rotated(phi float64) *Vector2 {
	return godotVec2AsVec2(
		C.godot_vector2_rotated(
			v.vector2,
			realAsGodotReal(phi),
		),
	)
}

//SetX changes v.vector2.x to the value specified.
func (v *Vector2) SetX(x float64) {
	C.godot_vector2_set_x(v.vector2, realAsGodotReal(x))
}

//SetY changes v.vector2.x to the value specified.
func (v *Vector2) SetY(y float64) {
	C.godot_vector2_set_x(v.vector2, realAsGodotReal(y))
}

//Slide slides the vector by the other vector.
func (v *Vector2) Slide(vec Vector2) *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_slide(v.vector2, vec.vector2))
}

//Snapped snaps the vector to a grid with the given size.
func (v *Vector2) Snapped(by Vector2) *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_snapped(v.vector2, by.vector2))
}

//Tangent returns a perpendicular vector.
func (v *Vector2) Tangent() *Vector2 {
	return godotVec2AsVec2(C.godot_vector2_tangent(v.vector2))
}

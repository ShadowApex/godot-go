package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

func NewQuat(x float64, y float64, z float64, w float64) *Quat {
	quat := &Quat{}

	// Create our godot quat object
	var godotQuat C.godot_quat

	// Create our quat
	C.godot_quat_new(
		&godotQuat,
		realAsGodotReal(x),
		realAsGodotReal(y),
		realAsGodotReal(z),
		realAsGodotReal(w),
	)

	// Internally set our quat
	quat.quat = &godotQuat

	return quat
}

func NewQuatWithAxisAngle(axis Vector3, angle float64) *Quat {
	quat := &Quat{}

	// Create our godot quat object
	var godotQuat C.godot_quat

	// Create our quat
	C.godot_quat_new_with_axis_angle(
		&godotQuat,
		axis.vector3,
		realAsGodotReal(angle),
	)

	// Internally set our quat
	quat.quat = &godotQuat

	return quat
}

type Quat struct {
	quat *C.godot_quat
}

// AsString returns the quat as a string.
func (q *Quat) AsString() string {
	asString := C.godot_quat_as_string(q.quat)
	return godotStringAsString(&asString)
}

// Dot
func (q *Quat) Dot(with Quat) float64 {
	return godotRealAsReal(C.godot_quat_dot(q.quat, with.quat))
}

// GetX returns the x value of the quat.
func (q *Quat) GetX() float64 {
	return godotRealAsReal(C.godot_quat_get_x(q.quat))
}

// GetY returns the y value of the quat.
func (q *Quat) GetY() float64 {
	return godotRealAsReal(C.godot_quat_get_y(q.quat))
}

// GetZ returns the z value of the quat.
func (q *Quat) GetZ() float64 {
	return godotRealAsReal(C.godot_quat_get_z(q.quat))
}

// GetW returns the w value of the quat.
func (q *Quat) GetW() float64 {
	return godotRealAsReal(C.godot_quat_get_w(q.quat))
}

// Inverse returns the inverse of itself.
func (q *Quat) Inverse() *Quat {
	return godotQuatAsQuat(C.godot_quat_inverse(q.quat))
}

// IsNormalized returns true if the quat has been normalized.
func (q *Quat) IsNormalized() bool {
	return godotBoolAsBool(C.godot_quat_is_normalized(q.quat))
}

// Length returns the length of the quat.
func (q *Quat) Length() float64 {
	return godotRealAsReal(C.godot_quat_length(q.quat))
}

// Length returns the squared length of the quat.
func (q *Quat) LengthSquared() float64 {
	return godotRealAsReal(C.godot_quat_length_squared(q.quat))
}

// Normalized returns a normalized quat.
func (q *Quat) Normalized() *Quat {
	return godotQuatAsQuat(C.godot_quat_normalized(q.quat))
}

// OperatorAdd returns the sum of q.quat and with.quat.
func (q *Quat) OperatorAdd(with Quat) *Quat {
	return godotQuatAsQuat(C.godot_quat_operator_add(q.quat, with.quat))
}

// OperatorDivide returns the quotent of q.quat and with.
func (q *Quat) OperatorDivide(with float64) *Quat {
	return godotQuatAsQuat(
		C.godot_quat_operator_divide(
			q.quat,
			realAsGodotReal(with),
		),
	)
}

// OperatorEqual returns true if q.quat = with.quat.
func (q *Quat) OperatorEqual(with Quat) bool {
	return godotBoolAsBool(C.godot_quat_operator_equal(q.quat, with.quat))
}

// OperatorMultiply returns the product of q.quat and with.
func (q *Quat) OperatorMultiply(with float64) *Quat {
	return godotQuatAsQuat(
		C.godot_quat_operator_multiply(
			q.quat,
			realAsGodotReal(with),
		),
	)
}

// OperatorNeg returns -q.quat.
func (q *Quat) OperatorNeg() *Quat {
	return godotQuatAsQuat(C.godot_quat_operator_neg(q.quat))
}

// OperatorSubtract returns the difference of q.quat and with.quat.
func (q *Quat) OperatorSubtract(with Quat) *Quat {
	return godotQuatAsQuat(C.godot_quat_operator_substract(q.quat, with.quat))
}

// SetX changes the x value of the quat to the specified value.
func (q *Quat) SetX(val float64) {
	C.godot_quat_set_x(q.quat, realAsGodotReal(val))
}

// SetY changes the y value of the quat to the specified value.
func (q *Quat) SetY(val float64) {
	C.godot_quat_set_y(q.quat, realAsGodotReal(val))
}

// SetZ changes the z value of the quat to the specified value.
func (q *Quat) SetZ(val float64) {
	C.godot_quat_set_z(q.quat, realAsGodotReal(val))
}

// SetW changes the w value of the quat to the specified value.
func (q *Quat) SetW(val float64) {
	C.godot_quat_set_w(q.quat, realAsGodotReal(val))
}

// Slerp
func (q *Quat) Slerp(with Quat, t float64) *Quat {
	return godotQuatAsQuat(
		C.godot_quat_slerp(
			q.quat,
			with.quat,
			realAsGodotReal(t),
		),
	)
}

// Slerpni
func (q *Quat) Slerpni(with Quat, t float64) *Quat {
	return godotQuatAsQuat(
		C.godot_quat_slerpni(
			q.quat,
			with.quat,
			realAsGodotReal(t),
		),
	)
}

// CubicSlerp
func (q *Quat) CubicSlerp(with Quat, preA Quat, postB Quat, t float64) *Quat {
	return godotQuatAsQuat(
		C.godot_quat_cubic_slerp(
			q.quat,
			with.quat,
			preA.quat,
			postB.quat,
			realAsGodotReal(t),
		),
	)
}

// Xform
func (q *Quat) Xform(with Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_quat_xform(q.quat, with.vector3))
}

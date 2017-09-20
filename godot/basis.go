package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Quat, Vector3
func NewBasis() *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis
	C.godot_basis_new(&godotBasis)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// NewBasisWithAxisAndAngle creates a rotation matrix which rotates around the
// given axis by the specified angle. The axis must be a normalized vector.
func NewBasisWithAxisAndAngle(axis Vector3, phi float64) *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis with axis and angle
	C.godot_basis_new_with_axis_and_angle(
		&godotBasis,
		vec3AsGodotVec3(axis),
		realAsGodotReal(phi),
	)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// NewBasisWithEule creates a rotation matrix (in the YXZ convention: first Z,
// then X, and Y last) from the specified Euler angles, given in the vector
// format as (X-angle, Y-angle, Z-angle).
func NewBasisWithEuler(euler Vector3) *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis with Euler
	C.godot_basis_new_with_euler(
		&godotBasis,
		vec3AsGodotVec3(euler),
	)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// NewBasisWithEulerQuat creates a rotation matrix from the given quaternion.
func NewBasisWithEulerQuat(euler Quat) *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis with Euler Quat
	C.godot_basis_new_with_euler_quat(
		&godotBasis,
		euler.quat,
	)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// NewBasisWithRows creates a rotation matrix which rotates around the given
// axis by the specified angle. The axis must be a normalized vector.
func NewBasisWithRows(xAxis Vector3, yAxis Vector3, zAxis Vector3) *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis with rows
	C.godot_basis_new_with_rows(
		&godotBasis,
		vec3AsGodotVec3(xAxis),
		vec3AsGodotVec3(yAxis),
		vec3AsGodotVec3(zAxis),
	)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}

// TODO: Finish implementing this
type Basis struct {
	basis *C.godot_basis
}

// TODO: SetRotationAxisAngle, SetRotationEuler, and SetScale are not yet implemented
// in the godot engine. They have been commented out until they are implemented.

// AsString returns a string of the basis
func (b *Basis) AsString() string {
	asString := C.godot_basis_as_string(b.basis)
	return godotStringAsString(&asString)
}

// Determinant returns the determinant of the matrix.
func (b *Basis) Determinant() float64 {
	determinant := C.godot_basis_determinant(b.basis)
	return godotRealAsReal(determinant)
}

// GetAxis retuns the normalized vector of the axis specified.
func (b *Basis) GetAxis(axis int64) *Vector3 {
	return godotVec3AsVec3(C.godot_basis_get_axis(b.basis, intAsGodotInt(axis)))
}

// GetElements copies the source Basis first three elements to the elements vector.
// From Headers: elements is a pointer to an array of 3 (!!) vector3?
func (b *Basis) GetElements(dest Vector3) {
	C.godot_basis_get_elements(b.basis, dest.vector3)
}

// GetEuler, assuming that the matrix is a proper rotation matrix (orthonormal
// matrix with determinant +1), returns Euler angles (in the YXZ convention:
// first Z, then X, and Y last). Returned vector contains the rotation angles
// in the format (X-angle, Y-angle, Z-angle).
func (b *Basis) GetEuler() *Vector3 {
	return godotVec3AsVec3(C.godot_basis_get_euler(b.basis))
}

// GetOrthagonalIndex considers a discretization of rotations into 24 points on
// unit sphere, lying along the vectors (x,y,z) with each component being either
// -1,0 or 1, and returns the index of the point best representing the orientation
// of the object. It is mainly used by the grid map editor. For further details,
// refer to Godot source code.
func (b *Basis) GetOrthagonalIndex() int64 {
	return godotIntAsInt(C.godot_basis_get_orthogonal_index(b.basis))
}

// GetRow returns the Vector3 at the specified row.
func (b *Basis) GetRow(row int64) *Vector3 {
	return godotVec3AsVec3(
		C.godot_basis_get_row(
			b.basis,
			intAsGodotInt(row),
		),
	)
}

// GetScale assumes that the matrix is the combination of a rotation and scaling,
// and returns the absolute value of scaling factors along each axis.
func (b *Basis) GetScale() *Vector3 {
	return godotVec3AsVec3(C.godot_basis_get_scale(b.basis))
}

// Inverse returns the inverse of the matrix.
func (b *Basis) Inverse() *Basis {
	return godotBasisAsBasis(C.godot_basis_inverse(b.basis))
}

// OperatorAdd adds the current basis and with.
func (b *Basis) OperatorAdd(with Basis) *Basis {
	return godotBasisAsBasis(C.godot_basis_operator_add(b.basis, with.basis))

}

// OperatorEqual compares the current basis and with. Returns true if equal.
func (b *Basis) OperatorEqual(with Basis) bool {
	return godotBoolAsBool(C.godot_basis_operator_equal(b.basis, with.basis))
}

// OperatorMultiplyScalar multiplys the current basis by with.
func (b *Basis) OperatorMultiplyScalar(with float64) *Basis {
	return godotBasisAsBasis(
		C.godot_basis_operator_multiply_scalar(
			b.basis,
			realAsGodotReal(with),
		),
	)
}

// OperatorMultiplyVector multiplys the current basis and with.
func (b *Basis) OperatorMultiplyVector(with Basis) *Basis {
	return godotBasisAsBasis(
		C.godot_basis_operator_multiply_vector(
			b.basis,
			with.basis,
		),
	)
}

// Orthonormalized returns the orthonormalized version of the matrix (useful to
// call from time to time to avoid rounding error for orthogonal matrices). This
// performs a Gram-Schmidt orthonormalization on the basis of the matrix.
func (b *Basis) Orthonormalized() *Basis {
	return godotBasisAsBasis(C.godot_basis_orthonormalized(b.basis))
}

// Rotated introduces an additional rotation around the given axis by phi
// (radians). Only relevant when the matrix is being used as a part of Transform.
// The axis must be a normalized vector.
func (b *Basis) Rotated(axis Vector3, phi float64) *Basis {
	return godotBasisAsBasis(
		C.godot_basis_rotated(
			b.basis,
			axis.vector3,
			realAsGodotReal(phi),
		),
	)
}

// Scaled introduces an additional scaling specified by the given 3D scaling
// factor. Only relevant when the matrix is being used as a part of Transform.
func (b *Basis) Scaled(scale Vector3) *Basis {
	return godotBasisAsBasis(
		C.godot_basis_scaled(
			b.basis,
			scale.vector3,
		),
	)
}

// SetAxis sets the specified axis to the given value.
func (b *Basis) SetAxis(axis int64, value Vector3) {
	C.godot_basis_set_axis(
		b.basis,
		intAsGodotInt(axis),
		value.vector3,
	)
}

/* C.godot_basis_set_rotation_axis_angle and C.godot_basis_set_rotation_euler
not yet implemented in basis.cpp
// SetRotationAxisAngle sets the specified axis and angle.
func (b *Basis) SetRotationAxisAngle(axis Vector3, angle float64) {
	C.godot_basis_set_rotation_axis_angle(
		b.basis,
		axis.vector3,
		realAsGodotReal(angle),
	)

}

// SetRotationEuler.
func (b *Basis) SetRotationEuler(euler Vector3) {
	C.godot_basis_set_rotation_euler(b.basis, euler.vector3)
}
*/

// SetRow sets the specified row to the given value.
func (b *Basis) SetRow(row int64, value Vector3) {
	C.godot_basis_set_row(
		b.basis,
		intAsGodotInt(row),
		value.vector3,
	)
}

/*C.godot_basis_set_scale not yet implemented in basis.cpp
// SetScale sets the scale to the given value.
func (b *Basis) SetScale(scale Vector3) {
	C.godot_basis_set_scale(b.basis, scale.vector3)
}
*/

// Tdotx transposes dot product with the x axis of the matrix.
func (b *Basis) Tdotx(with Vector3) float64 {
	return godotRealAsReal(C.godot_basis_tdotx(b.basis, with.vector3))
}

// Tdoty transposes dot product with the y axis of the matrix.
func (b *Basis) Tdoty(with Vector3) float64 {
	return godotRealAsReal(C.godot_basis_tdoty(b.basis, with.vector3))
}

// Tdotz transposes dot product with the z axis of the matrix.
func (b *Basis) Tdotz(with Vector3) float64 {
	return godotRealAsReal(C.godot_basis_tdotz(b.basis, with.vector3))
}

// Transposed returns the transposed version of the matrix.
func (b *Basis) Transposed() *Basis {
	return godotBasisAsBasis(C.godot_basis_transposed(b.basis))
}

// Xform returns a vector transformed (multiplied) by the matrix.
func (b *Basis) Xform(with Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_basis_xform(b.basis, with.vector3))
}

// XformInv returns a vector transformed (multiplied) by the transposed matrix.
// Note that this results in a multiplication by the inverse of the matrix only
// if it represents a rotation-reflection.
func (b *Basis) XformInv(with Vector3) *Vector3 {
	return godotVec3AsVec3(C.godot_basis_xform_inv(b.basis, with.vector3))
}

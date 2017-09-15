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

/* TODO: Enable when Quat is complete
func NewBasisWithEulerQuat(euler Quat) *Basis {
	basis := &Basis{}

	// Create our godot basis object
	var godotBasis C.godot_basis

	// Create our basis with Euler Quat
	C.godot_basis_new_with_euler_quat(
		&godotBasis,
		quatAsGodotQuat(euler),
	)

	// Internally set our basis
	basis.basis = &godotBasis

	return basis
}
*/

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

/*
func (b *Basis) AsString() string {
}

func (b *Basis) Determinant() bool {
}

func (b *Basis) GetAxis(axis int64) *Vector3 {
}

func (b *Basis) GetElements(elements Vector3) {
}

func (b *Basis) GetEuler() *Vector3 {
}

func (b *Basis) GetOrthagonalIndex() int64 {
}

func (b *Basis) GetRow(row int64) Vector3 {
}

func (b *Basis) GetScale() *Vector3 {
}

func (b *Basis) Inverse() *Basis {
}

func (b *Basis) OperatorAdd(with Vector3) *Basis {
}

func (b *Basis) OperatorEqual(with Vector3) bool {
}

func (b *Basis) OperatorMultiplyScalar(with float64) *Basis {
}

func (b *Basis) OperatorMultiplyVector(with Vector3) *Basis {
}

func (b *Basis) OperatorSubtract(with Vector3) *Basis {
}

func (b *Basis) Orthonormalized() *Basis {
}

func (b *Basis) Rotated(axis Vector3, phi float64) *Basis {
}

func (b *Basis) Scaled(scale Vector3) *Basis {
}

func (b *Basis) SetAxis(axis int64, value Vector3) {
}

func (b *Basis) SetRotationAxisAngle(axis Vector3, angle float64) {
}

func (b *Basis) SetRotationEuler(euler Vector3) {
}

func (b *Basis) SetRow(row int64, value Vector3) {
}

func (b *Basis) SetScale(scale Vector3) {
}

func (b *Basis) TDotX(with Vector3) float64 {
}

func (b *Basis) TDotY(with Vector3) float64 {
}

func (b *Basis) TDotZ(with Vector3) float64 {
}

func (b *Basis) Transposed() *Basis {
}

func (b *Basis) Xform(vec3 Vector2) *Vector3 {
}

func (b *Basis) XformInv(vec3 Vector2) *Vector3 {
}
*/

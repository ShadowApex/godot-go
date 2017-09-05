package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

const (
	AXIS_X = iota // 0
	AXIS_Y        // 1
	AXIS_Z        // 2
)

//DEPENDS: Built-In types
func NewVector3() *Vector3 {
	vector3 := &Vector3{}

	// Create our godot vector3 object
	var godotVector3 C.godot_vector3

	// Create our vector3
	C.godot_vector3_new(&godotVector3)

	// Internally set our vector3
	vector3.vector3 = &godotVector3

	return vector3
}

// TODO: Finish implementing this
type Vector3 struct {
	vector3 *C.godot_vector3
}

func (v *Vector3) Abs() {

}

//AngleTo returns the angle in radians between the two vectors.
func (v *Vector3) AngleTo(to Vector3) float64 {

}

//Ceil returns a new vector with all components rounded up.
func (v *Vector3) Ceil() Vector3 {

}

//Cross returns the cross product with b.
func (v *Vector3) Cross(b Vector3) Vector3 {

}

//CubicInterpolate Cubicly interpolates between this Vector and “b”, using “pre_a”
//and “post_b” as handles, and returning the result at position “t”.
func (v *Vector3) CubicInterpolate(b Vector3, preA Vector3, postB Vector3, t float64) {

}

//DistanceSquaredTo returns the squared distance to vector “b”. Prefer this
//function over “distance_to” if you need to sort vectors or need the squared
//distance for some formula.
func (v *Vector3) DistanceSquaredTo(to Vector3) float64 {

}

//DistanceTo returns the distance to vector “b”.
func (v *Vector3) DistanceTo(to Vector3) float64 {

}

//Dot returns the dot product with vector “b”.
func (v *Vector3) Dot(with Vector3) float64 {

}

//Floor removes the fractional part of x and y.
func (v *Vector3) Floor() {

}

//Inverse returns the inverse of the vector. This is the same as i
//Vector3( 1.0 / v.x, 1.0 / v.y, 1.0 / v.z )
func (v *Vector3) Inverse() Vector3 {

}

//Length returns the length of the vector.
func (v *Vector3) Length() float64 {

}

//LengthSquared returns the squared length of the vector. Prefer this function
//over “length” if you need to sort vectors or need the squared length for some formula.
func (v *Vector3) LengthSquared() float64 {

}

//LinearInterpolate returns the result of the linear interpolation between this
//vector and “b”, by amount “t”.
func (v *Vector3) LinearInterpolate(b Vector3, t float64) Vector3 {

}

//MaxAxis returns AXIS_X, AXIS_Y or AXIS_Z depending on which axis is the largest.
func (v *Vector3) MaxAxis() int8 {

}

//MinAxis returns AXIS_X, AXIS_Y or AXIS_Z depending on which axis is the smallest.
func (v *Vector3) MinAxis() int8 {

}

//Normalized returns a normalized vector to unit length.
func (v *Vector3) Normalized() Vector3 {

}

//Reflect is like “slide”, but reflects the Vector instead of continuing along the wall.
func (v *Vector3) Reflect(vec Vector3) {

}

//Rotated rotates the vector by “phi” radians.
func (v *Vector3) Rotated(phi float64) {

}

//Slide slides the vector by the other vector.
func (v *Vector3) Slide(vec Vector3) {

}

//Snapped snaps the vector to a grid with the given size.
func (v *Vector3) Snapped(by Vector3) {

}

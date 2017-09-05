package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

//DEPENDS: Built-In types
//NewVector2 constructs a new Vector2 from the given x and y.
func NewVector2() *Vector2 {
	vector2 := &Vector2{}

	// Create our godot vector2 object
	var godotVector2 C.godot_vector2

	// Create our vector2
	C.godot_vector2_new(&godotVector2)

	// Internally set our vector2
	vector2.vector2 = &godotVector2

	return vector2
}

// TODO: Finish implementing this
type Vector2 struct {
	vector2 *C.godot_vector2
}

func (v *Vector2) Abs() {

	/*C++
	return Vector2( fabs(x), fabs(y) );
	*/

}

//Angle returns the result of atan2 when called with the Vector’s x and y as parameters
//(Math::atan2(x,y)). Be aware that it therefore returns an angle oriented clockwise
//with regard to the (0, 1) unit vector, and not an angle oriented counter-clockwise
//with regard to the (1, 0) unit vector (which would be the typical trigonometric
//representation of the angle when calling Math::atan2(y,x))
func (v *Vector2) Angle() float64 {
	/*C++
	return atan2(y, x);
	*/

}

//AngleTo returns the angle in radians between the two vectors.
func (v *Vector2) AngleTo(to Vector2) float64 {
	/*C++
	return atan2(cross(p_vector2), dot(p_vector2));
	*/
}

//AngleToPoint returns the angle in radians between the line connecting the two
//points and the x coordinate.
func (v *Vector2) AngleToPoint(to Vector2) float64 {
	/*C++
	return atan2(y - p_vector2.y, x-p_vector2.x);
	*/
}

func (v *Vector2) Clamped(length float64) Vector2 {

	/*C++
	real_t l = length();
	Vector2 v = *this;
	if (l > 0 && p_len < l) {
		v /= l;
		v *= p_len;
	}
	return v;
	*/

}

//CubicInterpolate Cubicly interpolates between this Vector and “b”, using “pre_a”
//and “post_b” as handles, and returning the result at position “t”.
func (v *Vector2) CubicInterpolate(b Vector2, preA Vector2, postB Vector2, t float64) {

	/*C++
	Vector2 p0=p_pre_a;
	Vector2 p1=*this;
	Vector2 p2=p_b;
	Vector2 p3=p_post_b;

	real_t t = p_t;
	real_t t2 = t * t;
	real_t t3 = t2 * t;

	Vector2 out;
	out = ( ( p1 * 2.0) +
	( -p0 + p2 ) * t +
	( p0 * 2.0 - p1 * 5.0 + p2 * 4 - p3 ) * t2 +
	( -p0 + p1 * 3.0 - p2 * 3.0 + p3 ) * t3 ) * 0.5;

	return out;
	*/

}

//DistanceSquaredTo returns the squared distance to vector “b”. Prefer this
//function over “distance_to” if you need to sort vectors or need the squared
//distance for some formula.
func (v *Vector2) DistanceSquaredTo(to Vector2) float64 {
	/*C++
	return (x - p_vector2.x) * (x - p_vector2.x) + (y - p_vector2.y) * (y - p_vector2.y);
	*/

}

//DistanceTo returns the distance to vector “b”.
func (v *Vector2) DistanceTo(to Vector2) float64 {
	/*C++
	return sqrt((x - p_vector2.x) * (x - p_vector2.x) + (y - p_vector2.y) * (y - p_vector2.y));
	*/

}

//Dot returns the dot product with vector “b”.
func (v *Vector2) Dot(with Vector2) float64 {
	/*C++
	return x * p_other.x + y * p_other.y;
	*/

}

//Floor removes the fractional part of x and y.
func (v *Vector2) Floor() {
	/*C++
	return Vector2(::floor(x), ::floor(y));
	*/
}

//FloorF removes the fractional part of x and y.
func (v *Vector2) FloorF() {
	//?
}

//GetAspect returns the ratio of X to Y.
func (v *Vector2) GetAspect() float64 {
	//?
}

//Length returns the length of the vector.
func (v *Vector2) Length() float64 {
	/*C++
	return sqrt(x*x + y*y);
	*/

}

//LengthSquared returns the squared length of the vector. Prefer this function
//over “length” if you need to sort vectors or need the squared length for some formula.
func (v *Vector2) LengthSquared() float64 {
	/*C++
	return x*x + y*y;
	*/

}

//LinearInterpolate returns the result of the linear interpolation between this
//vector and “b”, by amount “t”.
func (v *Vector2) LinearInterpolate(b Vector2, t float64) Vector2 {
	/*C++
	Vector2 res=*this;
	res.x+= (p_t * (p_b.x-x));
	res.y+= (p_t * (p_b.y-y));
	return res;
	*/

}

//Normalized returns a normalized vector to unit length.
func (v *Vector2) Normalized() Vector2 {
	/*C++
		void Vector2::normalize()
	{
		real_t l = x*x + y*y;
		if (l != 0) {
			l = (l);
			x /= l;
			y /= l;
		}
	}

	Vector2 Vector2::normalized() const
	{
		Vector2 v = *this;
		v.normalize();
		return v;
	}
	*/

}

//Reflect is like “slide”, but reflects the Vector instead of continuing along the wall.
func (v *Vector2) Reflect(vec Vector2) {
	/*C++
	return p_vec - *this * this->dot(p_vec) * 2.0;
	*/

}

//Rotated rotates the vector by “phi” radians.
func (v *Vector2) Rotated(phi float64) {
	/*
		void Vector2::set_rotation(real_t p_radians) {

			x = cosf(p_radians);
			y = sinf(p_radians);
		}


		Vector2 Vector2::rotated(real_t p_by) const {
			Vector2 v;
			v.set_rotation(angle() + p_by);
			v *= length();
			return v;

		}
	*/
}

//Slide slides the vector by the other vector.
func (v *Vector2) Slide(vec Vector2) {
	/*C++
	return p_vec - *this * this->dot(p_vec);
	*/

}

//Snapped snaps the vector to a grid with the given size.
func (v *Vector2) Snapped(by Vector2) {
	/*C++
	return Vector2(
				p_by.x != 0 ? ::floor(x / p_by.x + 0.5) * p_by.x : x,
				p_by.y != 0 ? ::floor(y / p_by.y + 0.5) * p_by.y : y
	);
	*/

}

//Tangent returns a perpendicular vector.
func (v *Vector2) Tangent() {
	/*C++
	return Vector2(y,-x);
	*/

}

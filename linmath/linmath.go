package linmath

import "math"

type Vector3 struct {
	x float64
	y float64
	z float64
}

func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{
		x: x,
		y: y,
		z: z,
	}
}

func (v *Vector3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *Vector3) Normal() *Vector3 {
	invert := 1. / v.Length()

	return v.MultiplyOnScalar(invert)
}

func (v *Vector3) Negative() *Vector3 {
	return NewVector3(
		-v.x,
		-v.y,
		-v.z,
	)
}

func (v *Vector3) Multiply(v2 *Vector3) *Vector3 {
	return NewVector3(
		v.x*v2.x,
		v.y*v2.y,
		v.z*v2.z,
	)
}

func (v *Vector3) MultiplyOnScalar(scalar float64) *Vector3 {
	return NewVector3(
		v.x*scalar,
		v.y*scalar,
		v.z*scalar,
	)
}

func (v *Vector3) Dot(v2 *Vector3) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z
}

func (v *Vector3) Add(v2 *Vector3) *Vector3 {
	return NewVector3(
		v.x+v2.x,
		v.y+v2.y,
		v.z+v2.z,
	)
}

func (v *Vector3) Subtraction(v2 *Vector3) *Vector3 {
	return NewVector3(
		v.x-v2.x,
		v.y-v2.y,
		v.z-v2.z,
	)
}

func (v *Vector3) Divide(v2 *Vector3) *Vector3 {
	return NewVector3(
		v.x/v2.x,
		v.y/v2.y,
		v.z/v2.z,
	)
}

func (v *Vector3) DivideOnScalar(scalar float64) *Vector3 {
	return NewVector3(
		v.x/scalar,
		v.y/scalar,
		v.z/scalar,
	)
}

func (v *Vector3) Clamp(min, max float64) *Vector3 {
	return NewVector3(
		math.Min(math.Max(v.x, min), max),
		math.Min(math.Max(v.y, min), max),
		math.Min(math.Max(v.z, min), max),
	)
}

func (v *Vector3) Power(scalar float64) *Vector3 {
	return NewVector3(
		math.Pow(v.x, scalar),
		math.Pow(v.y, scalar),
		math.Pow(v.z, scalar),
	)
}

func Radians(degrees float64) float64 {
	return math.Pi * degrees / 180.
}

func Splat(scalar float64) *Vector3 {
	return NewVector3(scalar, scalar, scalar)
}

package main

import "math"

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) Add(u *Vec2) {
	v.X += u.X
	v.Y += u.Y
}

// NewVec2 returns a new vector with a given magnitude and direction angle
func NewVec2(mag, dir float64) Vec2 {
	return Vec2{mag * math.Cos(dir), mag * math.Sin(dir)}
}

// Angle returns angle between vectors u and v
func (v *Vec2) Angle(u Vec2) float64 {
	dx := u.X - v.X
	dy := u.Y - v.Y
	return math.Atan2(dy, dx)
}

// DistanceTo returns distance between vectors u and v
func (v *Vec2) DistanceTo(u Vec2) float64 {
	return math.Sqrt(math.Pow(u.X-v.X, 2) + math.Pow(u.Y-v.Y, 2))
}

// Sub returns the vector v - u
func (v Vec2) Sub(u Vec2) Vec2 {
	return Vec2{v.X - u.X, v.Y - u.Y}
}

// Plus returns the vector v + u
func (v Vec2) Plus(u Vec2) Vec2 {
	return Vec2{v.X + u.X, v.Y + u.Y}
}

// Scale returns v scaled by s
func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// Len returns the magnitude of v
func (v Vec2) Len() float64 {
	return math.Hypot(v.X, v.Y)
}

// Unit returns v normalized to unit length, or the zero vector if v is zero
func (v Vec2) Unit() Vec2 {
	l := v.Len()
	if l == 0 {
		return Vec2{}
	}
	return Vec2{v.X / l, v.Y / l}
}

// Mid returns the midpoint between a and b
func Mid(a, b Vec2) Vec2 {
	return Vec2{(a.X + b.X) / 2, (a.Y + b.Y) / 2}
}

// QuadBezier samples a quadratic Bézier curve with endpoints p0, p2 and control
// point p1 at parameter t in [0, 1]
func QuadBezier(p0, p1, p2 Vec2, t float64) Vec2 {
	u := 1 - t
	// u^2*p0 + 2*u*t*p1 + t^2*p2
	return p0.Scale(u * u).Plus(p1.Scale(2 * u * t)).Plus(p2.Scale(t * t))
}

// AngleOf returns the heading angle of vector v
func AngleOf(v Vec2) float64 {
	return math.Atan2(v.Y, v.X)
}

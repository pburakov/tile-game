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

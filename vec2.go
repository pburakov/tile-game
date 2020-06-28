package main

import "math"

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) Add(u Vec2) {
	v.X += u.X
	v.Y += u.Y
}

// UnitDistance returns a unit-speed vector with a given angle and velocity
func UnitDistance(angle float64, v float64) Vec2 {
	return Vec2{v * math.Cos(angle), v * math.Sin(angle)}
}

// Angle returns angle between vectors u and v
func (v *Vec2) Angle(u Vec2) float64 {
	dx := u.X - v.X
	dy := u.Y - v.Y
	return math.Atan2(dy, dx)
}

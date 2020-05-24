package main

import "math"

type Vec2 struct {
	X, Y float64
}

func (v *Vec2) Add(u Vec2) {
	v.X += u.X
	v.Y += u.Y
}

func (v *Vec2) Unit(angle float64, c float64) Vec2 {
	return Vec2{c * math.Cos(angle), c * math.Sin(angle)}
}

func (v *Vec2) Angle(u Vec2) float64 {
	dx := u.X - v.X
	dy := u.Y - v.Y
	return math.Atan2(dy, dx)
}

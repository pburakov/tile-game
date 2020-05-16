package main

import "math"

type Vector struct {
	X, Y float64
}

func (v *Vector) Add(u Vector) {
	v.X += u.X
	v.Y += u.Y
}

func (v *Vector) Unit(c float64, u Vector) Vector {
	theta := v.Theta(u)
	return Vector{c * math.Cos(theta), c * math.Sin(theta)}
}

func (v *Vector) Theta(u Vector) float64 {
	dx := u.X - v.X
	dy := u.Y - v.Y
	return math.Atan2(dy, dx)
}

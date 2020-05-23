package main

import (
	"math"
	"testing"
)

func TestRadToAngle(t *testing.T) {
	type args struct {
		r float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 deg", args{0.0}, 0},
		{"45 deg", args{math.Pi / 4}, 45},
		{"180 deg", args{math.Pi}, 180},
		{"210 deg", args{math.Pi * 7 / 6}, 210},
		{"2pi", args{math.Pi * 2}, 0},
		{"neg 45 deg", args{-math.Pi / 4}, -45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadToAngle(tt.args.r); got != tt.want {
				t.Errorf("RadToAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}

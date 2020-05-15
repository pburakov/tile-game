package main

type Train struct {
	X, Y         float64 // Position of train's geometric center
	BaseVelocity float64
}

// TrainPosition returns position of train's top left corner
func (t *Train) Position() (x float64, y float64) {
	return t.X - (trainWidth / 2), t.Y - (trainHeight / 2)
}

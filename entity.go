package main

type Train struct {
	X, Y         float64 // Position of train's geometric center
	BaseVelocity float64
	Target       Point // A destination point the train is currently moving into
}

// TrainPosition returns position of train's top left corner
func (t *Train) Position() (x float64, y float64) {
	// since there's only one sprite now. shift gometric center towards forward axis.
	// might lead to readlistic collision detection/ this will change for
	// multi-directional trains
	return -10 + t.X - (trainWidth / 2), t.Y - (trainHeight / 2)
}

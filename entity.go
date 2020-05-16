package main

type Train struct {
	Position     Vector // Position of train's geometric center
	BaseVelocity float64
	Target       Vector // A destination point the train is currently moving into
}

// TrainPosition returns position of train's top left corner
func (t *Train) TopLeft() Vector {
	// since there's only one sprite now. shift geometric center towards forward axis.
	// might lead to realistic collision detection/ this will change for
	// multi-directional trains
	return Vector{-10 + t.Position.X - (trainWidth / 2), t.Position.Y - (trainHeight / 2)}
}

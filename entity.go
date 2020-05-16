package main

type Train struct {
	BaseVelocity float64
	Cars         []*Car
}

type Car struct {
	Position Vector // Position of train's geometric center
	Target   Vector // A destination point the train is currently moving into
}

// TrainPosition returns position of train's top left corner
func (c *Car) TopLeft() Vector {
	// since there's only one sprite now. shift geometric center towards forward axis.
	// might lead to realistic collision detection/ this will change for
	// multi-directional trains
	return Vector{-10 + c.Position.X - (carWidth / 2), c.Position.Y - (carHeight / 2)}
}

package main

type Train struct {
	BaseVelocity float64
	Cars         []*Car
}

type Car struct {
	Position Vec2 // Position of train's geometric center
	Target   Vec2 // A destination point the train is currently moving into
}

// TrainPosition returns position of train's top left corner
func (c *Car) TopLeft() Vec2 {
	// since there's only one sprite now. shift geometric center towards forward axis.
	// might lead to realistic collision detection/ this will change for
	// multi-directional trains
	return Vec2{c.Position.X - (carWidth / 2), c.Position.Y - (carHeight / 2)}
}

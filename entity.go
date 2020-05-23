package main

type Train struct {
	Velocity float64 // Train's base velocity
	Cars     []*Car
}

type Car struct {
	Position Vec2 // Position of train's geometric center
	Target   Vec2 // A destination point the train is currently moving into
}

type PathNode struct {
	Position Vec2
	Prev     *PathNode
	Next     *PathNode
}

type Tile struct {
	// Entry and exit points for the train
	In     *PathNode
	Out    *PathNode
	Sprite byte
}

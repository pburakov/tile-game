package main

type Train struct {
	Velocity float64 // Train's base velocity
	Cars     []*Car
}

type Car struct {
	Position Vec2      // Position of train's geometric center
	Target   *PathNode // A destination point the train is currently moving into
}

type PathNode struct {
	Position Vec2
	Fwd      *PathNode // Next node, going west to east, south to north
	Rev      *PathNode // Prev node, going east to west, north to south
}

type Tile struct {
	Node   *PathNode
	Sprite byte
}

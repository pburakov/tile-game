package main

type Direction bool

const (
	reverse = true
	forward = false
)

type Train struct {
	Velocity float64 // Train's base velocity
	Cars     []*Car

	// Direction denotes train direction
	// Forward denotes going rightwards (W -> E) and upwards (S -> N)
	// Reverse denotes going leftwards (E -> W) and downwards (N -> S)
	Direction Direction
}

type Car struct {
	Position Vec2      // Position of train's geometric center
	Target   *PathNode // A destination point the train is currently moving into
}

type PathNode struct {
	Position Vec2
	Fwd      *PathNode // Next node in forward direction
	Rev      *PathNode // Prev node in forward direction
}

type Tile struct {
	Node   *PathNode
	Sprite byte
}

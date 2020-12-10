package main

type Heading bool

const (
	push = true
	pull = false
)

type Train struct {
	Velocity float64 // Velocity is train's base velocity
	Cars     []*Car

	// Heading denotes whether head car is pushing or pulling. Although a train
	// can have two locomotives the head is always relative to its starting
	// direction
	Heading Heading
}

type Car struct {
	Position Vec2      // Position of car's geometric center
	Angle    float64   // Angle is the last known angle the car is traveling at
	Target   *PathNode // Target is the point the car is currently moving towards
	Source   *PathNode // Source is the point the car is currently moving from
}

type PathNode struct {
	// Id is a PathNode's unique identifier, normally the same as the
	// ordinal number of a tile it sits on
	Position Vec2
	// AdjL, AdjR, AdjU, AdjD are the links to adjacent nodes in four directions
	AdjL, AdjR, AdjU, AdjD *PathNode
}

type Tile struct {
	Node *PathNode // Node is a pointer to a PathNode placed on this tile
	// Sprites contains sprite offsets to render at this tile.
	// Tile with a switch will have multiple tracks sprite layers.
	Sprites []byte
}

// HasSprite returns true if tile has a given sprite
func (t Tile) HasSprite(b byte) bool {
	for _, s := range t.Sprites {
		if s == b {
			return true
		}
	}
	return false
}

package main

// Heading selects how an angle maps to a sprite. It is kept around because
// AngleToDirection takes it; movement no longer depends on it since each car's
// sprite is chosen straight from its travel direction.
type Heading bool

const (
	push = true
	pull = false
)

type Train struct {
	Velocity float64 // Velocity is the head's speed in pixels per frame

	Cars []*Car // Cars[0] is always the leading car (gets the locomotive sprite)

	// route is the sequence of path nodes the train currently rides on, ordered
	// from the rear-most node (route[0]) to the frontier node ahead of the head
	// (route[len-1]). The smooth curve the cars follow is built from it.
	route []*PathNode

	// s is the arc length of the head measured from the start of the curve built
	// from route. Following cars sit at s - i*carSpacing behind it.
	s float64
}

type Car struct {
	Position Vec2    // Position of car's geometric center
	Angle    float64 // Angle is the direction the car is travelling at
}

type PathNode struct {
	// Position is the on-screen position of the node, normally a tile center
	Position Vec2

	// Locks counts how many trains currently occupy (or reserve) this node.
	// Single-goroutine state, no synchronization needed.
	Locks int

	// AdjL, AdjR, AdjU, AdjD are the links to adjacent nodes in four directions
	AdjL, AdjR, AdjU, AdjD *PathNode
}

func (p *PathNode) GetLocks() int {
	return p.Locks
}

type Tile struct {
	Node *PathNode // Node is a pointer to a PathNode placed on this tile

	// Sprites contains sprite byte offsets to render at this tile. Tile with a
	// switch will have multiple tracks sprite layers.
	Sprites []byte

	// PSprite is the index of a primary sprite, e.g. used for track switching
	PSprite int
}

// HasSprite returns true if the tile has a given sprite
func (t Tile) HasSprite(b byte) bool {
	for _, s := range t.Sprites {
		if s == b {
			return true
		}
	}
	return false
}

// PrimarySprite returns true if the primary sprite of a tile is a given sprite.
// If tile has a track with a switch on it, primary sprite will match the current
// switch selector setting.
func (t Tile) PrimarySprite(b byte) bool {
	return t.Sprites[t.PSprite] == b
}

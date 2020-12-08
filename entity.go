package main

type Heading bool

const (
	push = true
	pull = false
)

type Train struct {
	Velocity float64 // Train's base velocity
	Cars     []*Car

	// Heading denotes whether head car is pushing or pulling
	// Although a train can have two locomotives the head is always
	// relative to its starting direction
	Heading Heading
}

type Car struct {
	Position Vec2      // Position of car's geometric center
	Angle    float64   // Angle is the last known angle the car is traveling at
	Target   *PathNode // Target is the point the car is currently moving towards
	Source   *PathNode // Source is the point the car is currently moving from
}

type NodeId int

type PathNode struct {
	// Id is a PathNode's unique identifier, normally the same as the
	// ordinal number of a tile it sits on
	Id       NodeId
	Position Vec2
	// Adj contains links to adjacent nodes in any direction
	Adj map[NodeId]*PathNode
	// If PathNode is a switch, Selector1 and Selector2 point to currently selected links
	Selector1, Selector2 NodeId
}

type Tile struct {
	Node   *PathNode
	Sprite byte
}

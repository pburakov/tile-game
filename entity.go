package main

import "fmt"

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

func (c *Car) DebugInfo() string {
	return fmt.Sprintf("angle: %d; path: %d -> %d", RadToDegrees(c.Angle), c.Source.Id, c.Target.Id)
}

type NodeId int

type PathNode struct {
	Id       NodeId // Node's unique identifier
	Position Vec2
	Adj      map[NodeId]*PathNode // Links to adjacent nodes in any direction
	IsTurn   bool                 // Denotes whether node is a 90-degree turn in any direction
}

func NewPathNode(p Vec2, isTurn bool) *PathNode {
	return &PathNode{
		Id:       newUid(),
		Position: p,
		Adj:      make(map[NodeId]*PathNode),
		IsTurn:   isTurn,
	}
}

type Tile struct {
	Node   *PathNode
	Sprite byte
}

var uid NodeId = -1

func newUid() NodeId {
	uid++
	return uid
}

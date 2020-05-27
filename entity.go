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
	// relative to its starting direction.
	Heading Heading
}

type Car struct {
	Position Vec2      // Position of car's geometric center
	Target   *PathNode // A target point the car is currently moving into
	Source   *PathNode // A source point the car is currently moving from
}

type NodeId int

type PathNode struct {
	Id       NodeId // Node's unique identifier
	Position Vec2
	Adj      map[NodeId]*PathNode // Links to adjacent nodes in any direction
}

func NewPathNode(p Vec2) *PathNode {
	return &PathNode{
		Id:       newUid(),
		Position: p,
		Adj:      make(map[NodeId]*PathNode),
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

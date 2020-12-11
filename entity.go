package main

import (
	"sync/atomic"
	"time"
)

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
	Locks    int32 // Locks counts the number of cars locked on this node

	// AdjL, AdjR, AdjU, AdjD are the links to adjacent nodes in four directions
	AdjL, AdjR, AdjU, AdjD *PathNode
}

func (p *PathNode) AddLock() {
	atomic.AddInt32(&p.Locks, 1)
}

func (p *PathNode) ReleaseLock() {
	// TODO: Fix me. Lock is released with a lag to smooth out the gap between cars
	go func() {
		time.Sleep(650 * time.Millisecond)
		atomic.AddInt32(&p.Locks, -1)
	}()
}

func (p *PathNode) GetLocks() int32 {
	return atomic.LoadInt32(&p.Locks)
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

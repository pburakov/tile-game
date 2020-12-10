package main

import (
	"math"
)

// MoveTrains updates trains state
func MoveTrains() {
	for _, t := range trains {
		moveTrain(t)
	}
}

func moveTrain(t *Train) {
	h := getTrainHeadCar(t)
	// Check if train's head car can proceed, otherwise it's a dead end
	if atTarget(h, t.Velocity) {
		newTarget := findNextTarget(h)
		if newTarget == nil {
			reverseTrain(t)
		}
	}

	// Calculate route for each car, including head
	for _, c := range t.Cars {
		if atTarget(c, t.Velocity) {
			newTarget := findNextTarget(c)
			if newTarget != nil {
				c.Source, c.Target = c.Target, newTarget
			} // TODO: else a car probably derailed
		}

		c.Angle = c.Position.Angle(c.Target.Position)
		u := UnitDistance(c.Angle, t.Velocity)
		c.Position.Add(u)
	}
}

// findNextTarget finds the next target node for car to follow. Currently
// selects the first node that is not the source.
func findNextTarget(c *Car) *PathNode {
	u := c.Target
	for _, v := range []*PathNode{u.AdjL, u.AdjR, u.AdjU, u.AdjD} {
		if v != nil && v != c.Source {
			return v
		}
	}
	return nil
}

func reverseTrain(t *Train) {
	for _, c := range t.Cars {
		c.Target, c.Source = c.Source, c.Target
	}
	if t.Heading == push {
		t.Heading = pull
	} else {
		t.Heading = push
	}
}

func atTarget(c *Car, threshold float64) bool {
	return math.Abs(c.Position.X-c.Target.Position.X) <= threshold &&
		math.Abs(c.Position.Y-c.Target.Position.Y) <= threshold
}

func getTrainHeadCar(t *Train) *Car {
	if t.Heading == pull {
		return t.Cars[0]
	} else {
		return t.Cars[len(t.Cars)-1]
	}
}

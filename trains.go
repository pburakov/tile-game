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

	// Calculate angle and velocity for each car, including head. All cars
	// except the head will adjust own velocity to chase the previous car in
	// order stay as close to each other as possible. This loop goes in both
	// directions. The order in which cars are arranged depends on the train's
	// heading.
	var prev *Car
	for i, k := 0, len(t.Cars)-1; i < len(t.Cars) && k >= 0; i, k = i+1, k-1 {
		var c *Car
		if t.Heading == pull {
			c = t.Cars[i]
		} else {
			c = t.Cars[k]
		}

		if atTarget(c, t.Velocity) {
			newTarget := findNextTarget(c)
			if newTarget != nil {
				c.Source, c.Target = c.Target, newTarget
			} // TODO: else a car probably derailed
		}

		velocity := t.Velocity
		if prev != nil {
			dir := AngleToDirection(c.Angle, t.Heading)
			expectedDist := 1 + math.Max(
				float64(GetCarSprite(dir).Bounds().Dx()),
				float64(GetCarSprite(dir).Bounds().Dy()),
			)
			if c.Position.DistanceTo(prev.Position) > expectedDist {
				velocity = velocity * 2
			} else if c.Position.DistanceTo(prev.Position) < expectedDist {
				velocity = velocity * 0.5
			}
		}

		c.Angle = c.Position.Angle(c.Target.Position)
		u := UnitDistance(c.Angle, velocity)
		c.Position.Add(u)

		prev = c
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

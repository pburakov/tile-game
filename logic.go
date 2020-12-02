package main

import (
	"log"
	"math"
	"time"
)

var (
	Done = make(chan bool) // exit signal
)

const slowVelocityMultiplier = 0.77

func init() {
	log.SetFlags(0)

	launchClock()
	log.Print("launched logic clock")
}

func launchClock() {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-Done:
				log.Print("logic clock stopped")
				return
			case <-ticker.C:
				cycle()
			}
		}
	}()
}

// cycle updates entities state and game logic
func cycle() {
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

// findNextTarget finds the next target node for car to follow
func findNextTarget(c *Car) *PathNode {
	// For now, pick the next node that is not the source
	for k, v := range c.Target.Adj {
		if k != c.Source.Id {
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

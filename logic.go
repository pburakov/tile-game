package main

import (
	"log"
	"math"
	"time"
)

var (
	Done = make(chan bool) // exit signal
)

const slowVelocityMultiplier = 0.66

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
	if atDeadEnd(t) {
		reverseTrain(t)
	}

	for _, c := range t.Cars {
		if atTarget(c) {
			findNextTarget(t.Direction, c)
		}

		angle := c.Position.Angle(c.Target.Position)
		v := AdjustedVelocity(angle, t.Velocity)

		u := c.Position.Unit(angle, v)
		c.Position.Add(u)
	}
}

// AdjustedVelocity calculates adjusted car velocity, which must be slower at turns
func AdjustedVelocity(rad float64, baseVelocity float64) float64 {
	a := RadToDegrees(rad)
	if (30 < a && a <= 60) || // up-left
		(120 < a && a <= 150) || // up-right
		(210 < a && a <= 240) || // down-right
		(300 < a && a <= 330) { // down-left
		return slowVelocityMultiplier * baseVelocity
	} else {
		return baseVelocity
	}
}

// findNextTarget updates the car to follow the next target node
func findNextTarget(d Direction, c *Car) {
	switch d {
	case forward:
		if c.Target.Fwd != nil {
			c.Source = c.Target
			c.Target = c.Target.Fwd
		}
	case reverse:
		if c.Target.Rev != nil {
			c.Source = c.Target
			c.Target = c.Target.Rev
		} else {
			c.Target = c.Source
		}
	}
}

func reverseTrain(t *Train) {
	if t.Direction == forward {
		t.Direction = reverse
	} else {
		t.Direction = forward
	}
	for _, c := range t.Cars {
		c.Target, c.Source = c.Source, c.Target
	}
}

func atDeadEnd(t *Train) bool {
	if t.Direction == forward {
		c := t.Cars[0]
		return atTarget(c) && c.Target.Fwd == nil
	} else {
		c := t.Cars[len(t.Cars)-1]
		return atTarget(c) && c.Target.Rev == nil
	}
}

func atTarget(c *Car) bool {
	return math.Abs(c.Position.X-c.Target.Position.X) <= 1.0 &&
		math.Abs(c.Position.Y-c.Target.Position.Y) <= 1.0
}

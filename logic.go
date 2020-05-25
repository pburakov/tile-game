package main

import (
	"log"
	"math"
	"time"
)

var (
	Done = make(chan bool) // exit signal
)

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
	for _, c := range t.Cars {
		if atTarget(c) {
			findNextTarget(t, c)
		}
		angle := c.Position.Angle(c.Target.Position)
		direction := AngleToDirection(angle)

		// Calculate car velocity, must be slower at turns
		v := t.Velocity
		if direction == upLeft || direction == upRight ||
			direction == downLeft || direction == downRight {
			v = 0.66 * t.Velocity
		}

		u := c.Position.Unit(angle, v)
		c.Position.Add(u)
	}
}

func findNextTarget(t *Train, c *Car) {
	switch t.Direction {
	case forward:
		if c.Target.Fwd != nil {
			c.Target = c.Target.Fwd
		} else {
			t.Direction = reverse
		}
	case reverse:
		if c.Target.Rev != nil {
			c.Target = c.Target.Rev
		} else {
			t.Direction = forward
		}
	}
}

func atTarget(c *Car) bool {
	return math.Abs(c.Position.X-c.Target.Position.X) <= 1.0 &&
		math.Abs(c.Position.Y-c.Target.Position.Y) <= 1.0
}

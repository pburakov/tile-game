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
			// reset position to match target
			// reset a potential rounding and float errors
			c.Position.X = c.Target.X
			c.Position.Y = c.Target.Y
			findNextTarget(c)
		}
		angle := c.Position.Angle(c.Target)
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

func findNextTarget(c *Car) {
	// emulate target switching for now
	if c.Target.X == 48 {
		c.Target.X = 48 + 16*2
		c.Target.Y = 198 - 16*2
	} else if c.Target.X == 48+16*2 {
		c.Target.X = 48 + 16*5
		c.Target.Y = 198 - 16*2
	}
}

func atTarget(c *Car) bool {
	return math.Abs(c.Position.X-c.Target.X) < 1.0 &&
		math.Abs(c.Position.Y-c.Target.Y) < 1.0
}

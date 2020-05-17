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
		u := c.Position.Unit(angle, t.BaseVelocity, c.Target)
		c.Position.Add(u)
	}
}

func findNextTarget(c *Car) {
	// emulate target switching for now
	if c.Target.X == 96 {
		c.Target.X = 96 + 16
		c.Target.Y = 198 - 16
	} else if c.Target.X == 96+16 {
		c.Target.X = 96 + 16 + 64
		c.Target.Y = 198 - 16
	}
}

func atTarget(c *Car) bool {
	return math.Abs(c.Position.X-c.Target.X) < 1.0 &&
		math.Abs(c.Position.Y-c.Target.Y) < 1.0
}

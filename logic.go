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
	// calculate vector to target with x and y components
	dx := (t.Target.X - t.X)
	dy := (t.Target.Y - t.Y)
	theta := math.Atan2(dy, dx)
	t.X += t.BaseVelocity * math.Cos(theta)
	t.Y += t.BaseVelocity * math.Sin(theta)
}

package main

import (
	"log"
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
		u := c.Position.Unit(t.BaseVelocity, c.Target)
		c.Position.Add(u)
	}
}

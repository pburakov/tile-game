package logic

import (
	"github.com/pburakov/tile-game/world"
	"log"
	"time"
)

var (
	Done = make(chan bool) // exit signal
)

func init() {
	launchClock()
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
	log.Print("launched logic clock")
}

// cycle updates inner state
func cycle() {
	world.Train.X += world.Train.XVelocity
	world.Train.X += world.Train.YVelocity
}

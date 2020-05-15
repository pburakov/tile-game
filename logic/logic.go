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

// cycle updates inner state
func cycle() {
	world.Train.X += world.Train.XVelocity
	world.Train.X += world.Train.YVelocity
}

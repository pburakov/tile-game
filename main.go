package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pburakov/tile-game/draw"
	"github.com/pburakov/tile-game/logic"
	"log"
)

func main() {
	if err := ebiten.Run(draw.Update, draw.ScreenWidth, draw.ScreenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
	logic.Done <- true
}

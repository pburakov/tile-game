package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	if err := ebiten.Run(Update, ScreenWidth, ScreenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
	Done <- true
}

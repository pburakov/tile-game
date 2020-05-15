package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/pburakov/tile-game/draw"
	"github.com/pburakov/tile-game/logic"
	"github.com/pburakov/tile-game/world"
	"log"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := draw.Background(screen)
	if err != nil {
		return err
	}

	err = draw.Tiles(&world.Tiles, screen)
	if err != nil {
		return err
	}

	err = draw.Train(&world.Train, screen)
	if err != nil {
		return err
	}

	x, y := ebiten.CursorPosition()
	err = draw.Cursor(x, y, screen)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, draw.ScreenWidth, draw.ScreenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
	logic.Done <- true
}

package main

import (
	"github.com/hajimehoshi/ebiten"
)

func Update(screen *ebiten.Image) error {
	// Handling changes in user input must go before drawing
	HandleInput()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Draw static world
	err := DrawBackground(screen)
	if err != nil {
		return err
	}
	err = DrawTiles(&world, screen)
	if err != nil {
		return err
	}

	// Draw cursor
	err = DrawCursor(&selector, screen)
	if err != nil {
		return err
	}

	// Render moving assets
	err = DrawTrains(&trains, screen)
	if err != nil {
		return err
	}

	return nil
}

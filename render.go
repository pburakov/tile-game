package main

import (
	"github.com/hajimehoshi/ebiten"
)

func Update(screen *ebiten.Image) error {
	// Handle input
	x, y := ebiten.CursorPosition()
	_, wy := ebiten.Wheel()
	selector.ApplyDelta(wy)

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
	err = DrawCursor(x, y, &selector, screen)
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

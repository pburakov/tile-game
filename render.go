package main

import (
	"github.com/hajimehoshi/ebiten"
)

func Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	err := DrawBackground(screen)
	if err != nil {
		return err
	}

	err = DrawTiles(&tiles, screen)
	if err != nil {
		return err
	}

	err = DrawTrains(&trains, screen)
	if err != nil {
		return err
	}

	x, y := ebiten.CursorPosition()
	err = DrawCursor(x, y, screen)
	if err != nil {
		return err
	}

	return nil
}

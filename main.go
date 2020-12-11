package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	if err := ebiten.Run(Update, ScreenWidth, ScreenHeight, 2, "Tiles"); err != nil {
		log.Fatal(err)
	}
}

const (
	debugPaths         = false
	debugPathFollowing = false
)

func Update(screen *ebiten.Image) error {
	// Handling changes in user input must go before drawing
	HandleInput()

	MoveTrains()

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

	if debugPaths {
		DrawPaths(&world, screen)
	}

	// Draw cursor
	if !selector.ControlMode {
		err = DrawCursor(&selector, screen)
		if err != nil {
			return err
		}
	}

	// Render moving assets
	err = DrawTrains(&trains, screen, debugPathFollowing)
	if err != nil {
		return err
	}

	return nil
}

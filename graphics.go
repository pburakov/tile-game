package main

import (
	"github.com/hajimehoshi/ebiten"
	_ "image/png"
)

const (
	TileSize     = 16
	ScreenWidth  = 320
	ScreenHeight = 240 // 20 x 15 tiles
	TilesPerRow  = ScreenWidth / TileSize
)

func DrawCursor(x int, y int, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// get top-left coordinates of the nearest tile
	px, py := TileToPosition(PositionToTile(x, y))
	op.GeoM.Translate(float64(px), float64(py))

	err := screen.DrawImage(getSprite(cursor).(*ebiten.Image), op)
	if err != nil {
		return err
	}
	return nil
}

func DrawTiles(tiles *[]byte, screen *ebiten.Image) error {
	for i, t := range *tiles {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(OrdinalToPosition(i))

		err := screen.DrawImage(getSprite(t).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func DrawTrains(trains *[]*Train, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	for _, t := range *trains {
		op.GeoM.Translate(t.X, t.Y)

		err := screen.DrawImage(getSprite(train).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func DrawBackground(screen *ebiten.Image) error {
	for i := 0; i < ScreenWidth*ScreenHeight/TileSize*TileSize; i++ {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(OrdinalToPosition(i))

		err := screen.DrawImage(getSprite(grass).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

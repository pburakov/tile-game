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

	err := screen.DrawImage(getTileSprite(cursor).(*ebiten.Image), op)
	if err != nil {
		return err
	}
	return nil
}

func DrawTiles(m *Map, screen *ebiten.Image) error {
	for i, t := range m.getAll() {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(OrdinalToPosition(i))

		err := screen.DrawImage(getTileSprite(t).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func DrawTrains(trains *[]*Train, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	for _, t := range *trains {
		op.GeoM.Translate(t.Position())

		err := screen.DrawImage(trainSprite.(*ebiten.Image), op)
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

		err := screen.DrawImage(getTileSprite(grass).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

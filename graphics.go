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
	v := TileToPosition(PositionToTile(x, y))
	op.GeoM.Translate(v.X, v.Y)

	err := screen.DrawImage(getTileSprite(cursor).(*ebiten.Image), op)
	if err != nil {
		return err
	}
	return nil
}

func DrawTiles(m *Map, screen *ebiten.Image) error {
	for i, t := range m.getAll() {
		op := &ebiten.DrawImageOptions{}
		v := OrdinalToPosition(i)
		op.GeoM.Translate(v.X, v.Y)

		err := screen.DrawImage(getTileSprite(t).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func DrawTrains(trains *[]*Train, screen *ebiten.Image) error {
	for _, t := range *trains {
		for i, c := range t.Cars {
			v := c.TopLeft()
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(v.X, v.Y)
			var err error
			if i == 0 {
				err = screen.DrawImage(headSprite.(*ebiten.Image), op)
			} else {
				err = screen.DrawImage(carSprite.(*ebiten.Image), op)
			}
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DrawBackground(screen *ebiten.Image) error {
	for i := 0; i < ScreenWidth*ScreenHeight/TileSize*TileSize; i++ {
		op := &ebiten.DrawImageOptions{}
		v := OrdinalToPosition(i)
		op.GeoM.Translate(v.X, v.Y)

		err := screen.DrawImage(getTileSprite(grass).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

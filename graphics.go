package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
)

const (
	TileSize     = 16
	ScreenWidth  = 320
	ScreenHeight = 240 // 20 x 15 tiles
	TilesPerRow  = ScreenWidth / TileSize
)

var DebugColor = color.RGBA{0xff, 0xff, 0x00, 0xff}

func DrawCursor(sel *Selector, screen *ebiten.Image) error {
	op := &ebiten.DrawImageOptions{}
	// get top-left coordinates of the nearest tile
	v := TileToPosition(PositionToTile(sel.CurX, sel.CurY))
	op.GeoM.Translate(v.X, v.Y)

	// Brush
	err := screen.DrawImage(GetTileSprite(sel.GetCurrentSelection()).(*ebiten.Image), op)
	if err != nil {
		return err
	}

	// Cursor highlight frame
	err = screen.DrawImage(GetTileSprite(cursor).(*ebiten.Image), op)
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

		err := screen.DrawImage(GetTileSprite(t.Sprite).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

func DrawPaths(m *Map, screen *ebiten.Image) {
	for _, t := range m.getAll() {
		u := t.Node
		if u != nil {
			ebitenutil.DrawRect(
				screen,
				u.Position.X, u.Position.Y, 3, 3,
				DebugColor)
			for _, v := range u.Adj {
				ebitenutil.DrawLine(
					screen,
					u.Position.X, u.Position.Y,
					v.Position.X, v.Position.Y,
					DebugColor)
			}
		}
	}
}

func DrawTrains(trains *[]*Train, screen *ebiten.Image) error {
	fmt.Print("\033[0;0H")
	for tId, t := range *trains {
		fmt.Printf("Train %d:\n", tId)
		for cId, c := range t.Cars {
			fmt.Printf("Car %d: %s\n", cId, c.DebugInfo())

			angle := c.Position.Angle(c.Target.Position)
			dir := AngleToDirection(angle, t.Heading)

			var img *ebiten.Image
			if cId == 0 {
				img = GetHeadSprite(dir).(*ebiten.Image)
			} else {
				img = GetCarSprite(dir).(*ebiten.Image)
			}

			v := CarTopLeft(c, img)
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(v.X, v.Y)

			err := screen.DrawImage(img, op)
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

		err := screen.DrawImage(GetTileSprite(grass).(*ebiten.Image), op)
		if err != nil {
			return err
		}
	}
	return nil
}

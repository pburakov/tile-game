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

var DebugColor = color.RGBA{R: 0xff, B: 0xff, A: 0xff}

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

		for _, s := range t.Sprites {
			err := screen.DrawImage(GetTileSprite(s).(*ebiten.Image), op)
			if err != nil {
				return err
			}
		}

		if len(t.Sprites) > 1 { // Draw selector only for layered switch tiles
			s := SwitchOffset(t.Sprites[t.PSprite], rail, swch)
			err := screen.DrawImage(GetTileSprite(s).(*ebiten.Image), op)
			if err != nil {
				return err
			}
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
			ebitenutil.DebugPrintAt(screen, fmt.Sprint(u.GetLocks()), int(u.Position.X), int(u.Position.Y))
			for _, v := range []*PathNode{u.AdjL, u.AdjR, u.AdjU, u.AdjD} {
				if v != nil {
					ebitenutil.DrawLine(
						screen,
						u.Position.X, u.Position.Y,
						v.Position.X, v.Position.Y,
						DebugColor)
				}
			}
		}
	}
}

func DrawTrains(trains *[]*Train, screen *ebiten.Image, debug bool) error {
	for _, t := range *trains {
		for cId, c := range t.Cars {
			dir := AngleToDirection(c.Angle, t.Heading)

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

			if debug {
				ebitenutil.DrawLine(
					screen,
					c.Position.X, c.Position.Y,
					c.Target.Position.X, c.Target.Position.Y,
					DebugColor)
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

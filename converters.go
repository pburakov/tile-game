package main

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

// TileToPosition returns position of top-left corner of a tile with (tx, ty) coordinates
func TileToPosition(tx int, ty int) Vec2 {
	return Vec2{float64(tx * TileSize), float64(ty * TileSize)}
}

// PositionToTile returns matrix coordinates of a tile encompassing (x, y) position
func PositionToTile(x int, y int) (tx int, ty int) {
	return (x - x%TileSize) / TileSize, (y - y%TileSize) / TileSize
}

// OrdinalToPosition returns position of top-left corner of a tile with and ordinal number i
func OrdinalToPosition(i int) Vec2 {
	return Vec2{float64(i%TilesPerRow) * TileSize, float64(i/TilesPerRow) * TileSize}
}

//TileToOrdinal returns ordinal number of a tile with (tx, ty) matrix coordinates
func TileToOrdinal(tx int, ty int) int {
	return TilesPerRow*ty + tx
}

func AngleToDirection(rad float64) string {
	// TODO: fix me, calculates diagonals incorrectly
	a := RadToAngle(-rad)
	if 30 < a && a <= 60 {
		return upRight
	} else if 120 < a && a <= 150 {
		return upLeft
	} else if 150 < a && a <= 210 {
		return left
	} else if 210 < a && a <= 240 {
		return downLeft
	} else if 300 < a && a <= 330 {
		return downRight
	}
	return right
}

func RadToAngle(r float64) int {
	return int(math.Round(180*r/math.Pi)) % 360
}

// carTopLeft returns position of car's top left corner based on sprite
func CarTopLeft(c *Car, img *ebiten.Image) Vec2 {
	width, height := img.Size()
	return Vec2{c.Position.X - (float64(width) / 2), c.Position.Y - (float64(height) / 2)}
}

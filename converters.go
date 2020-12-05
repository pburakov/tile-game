package main

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

// TileToPosition returns position of top-left corner of a tile with (tx, ty) coordinates
func TileToPosition(tx, ty int) Vec2 {
	return Vec2{float64(tx * TileSize), float64(ty * TileSize)}
}

// PositionToTile returns matrix coordinates of a tile encompassing (x, y) position
func PositionToTile(x, y int) (tx int, ty int) {
	return (x - x%TileSize) / TileSize, (y - y%TileSize) / TileSize
}

// OrdinalToPosition returns position of top-left corner of a tile with an ordinal number i
func OrdinalToPosition(i int) Vec2 {
	return Vec2{float64(i%TilesPerRow) * TileSize, float64(i/TilesPerRow) * TileSize}
}

// TileToOrdinal returns ordinal number of a tile with (tx, ty) matrix coordinates
func TileToOrdinal(tx, ty int) int {
	return TilesPerRow*ty + tx
}

// AngleToDirection returns sprite direction token, adjusted for train's heading
func AngleToDirection(rad float64, h Heading) string {
	a := RadToDegrees(rad)
	// Circle is split into 8x45-degree segments
	switch h {
	case pull:
		if 22 < a && a <= 67 {
			return upRight
		} else if 67 < a && a <= 112 {
			return up
		} else if 112 < a && a <= 157 {
			return upLeft
		} else if 157 < a && a <= 202 {
			return left
		} else if 202 < a && a <= 247 {
			return downLeft
		} else if 247 < a && a <= 292 {
			return down
		} else if 292 < a && a <= 337 {
			return downRight
		}
		return right
	case push:
		if 22 < a && a <= 67 {
			return downLeft
		} else if 67 < a && a <= 112 {
			return down
		} else if 112 < a && a <= 157 {
			return downRight
		} else if 157 < a && a <= 202 {
			return right
		} else if 202 < a && a <= 247 {
			return upRight
		} else if 247 < a && a <= 292 {
			return up
		} else if 292 < a && a <= 337 {
			return upLeft
		}
		return left
	default:
		return right
	}
}

func RadToDegrees(r float64) int {
	return int(math.Round(180*(2*math.Pi-r)/math.Pi)) % 360
}

// CarTopLeft returns position of car's top left corner based on sprite
func CarTopLeft(c *Car, img *ebiten.Image) Vec2 {
	width, height := img.Size()
	return Vec2{c.Position.X - (float64(width) / 2), c.Position.Y - (float64(height) / 2)}
}

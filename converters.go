package main

import (
	"math"

	"github.com/hajimehoshi/ebiten"
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

func DirectionFromAngle(angle float64) string {
	if -math.Pi/8 <= angle && angle < math.Pi/8 {
		return left
	} else if -math.Pi/4 <= angle && angle < -math.Pi/8 {
		return upLeft
	}
	return left
}

// carTopLeft returns position of car's top left corner based on sprite
func CarTopLeft(c *Car, img *ebiten.Image) Vec2 {
	width, height := img.Size()
	return Vec2{c.Position.X - (float64(width) / 2), c.Position.Y - (float64(height) / 2)}
}

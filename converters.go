package main

// TileToPosition returns position of top-left corner of a tile with (tx, ty) coordinates
func TileToPosition(tx int, ty int) (x int, y int) {
	return tx * TileSize, ty * TileSize
}

// PositionToTile returns matrix coordinates of a tile encompassing (x, y) position
func PositionToTile(x int, y int) (tx int, ty int) {
	return (x - x%TileSize) / TileSize, (y - y%TileSize) / TileSize
}

func OrdinalToPosition(i int) (x float64, y float64) {
	return float64(i%TilesPerRow) * TileSize, float64(i/TilesPerRow) * TileSize
}

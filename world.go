package main

const (
	MapWidth  = 20
	MapHeight = 15
)

type Map struct {
	tiles [MapWidth * MapHeight]byte
}

var world = Map{}

func init() {
	world.setTile(0, 12, rail+hor)
	world.setTile(1, 12, rail+hor)
	world.setTile(2, 12, rail+hor)
	world.setTile(3, 12, rail+hor)
}

func (m *Map) setTile(tx int, ty int, t byte) {
	i := TileToOrdinal(tx, ty)
	m.tiles[i] = t
}

func (m *Map) getTile(tx, ty int) byte {
	i := TileToOrdinal(tx, ty)
	return m.tiles[i]
}
func (m *Map) getAll() *[MapWidth * MapHeight]byte {
	return &m.tiles
}

var trains = []*Train{{
	X:            float64(0),
	Y:            float64(198),
	BaseVelocity: 1.0,
}}

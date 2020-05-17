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
	world.setTile(4, 12, rail+hor)
	world.setTile(5, 12, rail+hor)
	world.setTile(6, 12, rail+ul)
	world.setTile(6, 11, rail+dr)
	world.setTile(7, 11, rail+hor)
	world.setTile(8, 11, rail+hor)
	world.setTile(9, 11, rail+hor)
	world.setTile(10, 11, rail+hor)
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
	Velocity: 1.0,
	Cars: []*Car{
		{
			Position: Vec2{0, 198},
			Target:   Vec2{96, 198},
		}, {
			Position: Vec2{-25, 198},
			Target:   Vec2{96, 198},
		}, {
			Position: Vec2{-25 * 2, 198},
			Target:   Vec2{96, 198},
		}, {
			Position: Vec2{-25 * 3, 198},
			Target:   Vec2{96, 198},
		},
	},
}}

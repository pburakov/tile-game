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
	world.setTile(3, 12, rail+ul)
	world.setTile(3, 11, rail+dr)
	world.setTile(4, 11, rail+ul)
	world.setTile(4, 10, rail+dr)
	world.setTile(5, 10, rail+hor)
	world.setTile(6, 10, rail+hor)
	world.setTile(7, 10, rail+hor)
	world.setTile(8, 10, rail+hor)
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

var trains = []*Train{
	{
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 + 16*4, 160 - 16*0},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 + 16*4, 160 - 16*4},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 - 16*0, 160 - 16*4},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 - 16*4, 160 - 16*4},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 - 16*4, 160 - 16*0},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 - 16*4, 160 + 16*4},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 - 16*0, 160 + 16*4},
			},
		},
	}, {
		Velocity: 1.0,
		Cars: []*Car{
			{
				Position: Vec2{160, 160},
				Target:   Vec2{160 + 16*4, 160 + 16*4},
			},
		},
	},
}

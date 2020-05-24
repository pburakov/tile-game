package main

const (
	MapWidth  = 20
	MapHeight = 15
)

type Map struct {
	tiles [MapWidth * MapHeight]Tile
}

var world = Map{}

func init() {}

func (m *Map) setTile(tx int, ty int, t Tile) {
	i := TileToOrdinal(tx, ty)
	m.tiles[i] = t
}

func (m *Map) getTile(tx, ty int) Tile {
	i := TileToOrdinal(tx, ty)
	return m.tiles[i]
}

func (m *Map) getAll() *[MapWidth * MapHeight]Tile {
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

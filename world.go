package main

const (
	MapWidth  = 20
	MapHeight = 15
)

type Map struct {
	tiles [MapWidth * MapHeight]Tile
}

var world = Map{}

func init() {
	world.setTile(0, 2, rail+hor)
	world.setTile(1, 2, rail+hor)
	world.setTile(2, 2, rail+hor)

	world.setTile(6, 12, rail+ver)
	world.setTile(6, 13, rail+ver)
	world.setTile(6, 14, rail+ver)

	trains = []*Train{
		{
			Velocity: 1.0,
			Cars: []*Car{
				{
					Position: world.getTile(0, 2).Node.Position,
					Target:   world.getTile(0, 2).Node,
				},
			},
		},
		{
			Velocity: 1.0,
			Cars: []*Car{
				{
					Position: world.getTile(6, 14).Node.Position,
					Target:   world.getTile(6, 14).Node,
				},
			},
		},
	}
}

var trains []*Train

func (m *Map) setTile(tx, ty int, b byte) {
	i := TileToOrdinal(tx, ty)
	if i < 0 || i >= len(m.tiles)-1 {
		return
	}

	// Coordinates of top-left corner
	v := TileToPosition(tx, ty)

	adjL, adjR, adjU, adjD :=
		world.getTile(tx-1, ty),
		world.getTile(tx+1, ty),
		world.getTile(tx, ty-1),
		world.getTile(tx, ty+1)

	n := &PathNode{Position: Vec2{v.X + 8, v.Y + 6}}

	switch b {
	case rail + hor:
		if adjL != nil {
			n.Rev = adjL.Node
			if adjL.Node != nil {
				adjL.Node.Fwd = n
			}
		}
		if adjR != nil {
			n.Fwd = adjR.Node
			if adjR.Node != nil {
				adjR.Node.Rev = n
			}
		}
	case rail + ver:
		if adjU != nil {
			n.Fwd = adjU.Node
			if adjU.Node != nil {
				adjU.Node.Rev = n
			}
		}
		if adjD != nil {
			n.Rev = adjD.Node
			if adjD.Node != nil {
				adjD.Node.Fwd = n
			}
		}
	}

	if i >= 0 && i < len(m.tiles) {
		m.tiles[i].Sprite = b
		m.tiles[i].Node = n
	}
}

func (m *Map) getTile(tx, ty int) *Tile {
	i := TileToOrdinal(tx, ty)
	if i >= 0 && i < len(m.tiles) {
		return &m.tiles[i]
	} else {
		return nil
	}
}

func (m *Map) removeTile(tx, ty int) {
	i := TileToOrdinal(tx, ty)
	m.tiles[i].Sprite = 0
	if m.tiles[i].Node != nil {
		// Disconnect adjacent connected nodes, if any
		if m.tiles[i].Node.Fwd != nil {
			m.tiles[i].Node.Fwd.Rev = nil
			m.tiles[i].Node.Fwd = nil
		}
		if m.tiles[i].Node.Rev != nil {
			m.tiles[i].Node.Rev.Fwd = nil
			m.tiles[i].Node.Rev = nil
		}
		m.tiles[i].Node = nil
	}
}

func (m *Map) getAll() *[MapWidth * MapHeight]Tile {
	return &m.tiles
}

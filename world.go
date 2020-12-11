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
	// TODO: test world, remove tiles and trains below
	world.setTile(0, 2, rail+hor)
	world.setTile(1, 2, rail+hor)
	world.setTile(2, 2, rail+hor)
	world.setTile(3, 2, rail+hor)
	world.setTile(4, 2, rail+hor)
	world.setTile(5, 2, rail+dl)
	world.setTile(5, 3, rail+ver)
	world.setTile(5, 4, rail+ver)
	world.setTile(5, 5, rail+ver)
	world.setTile(5, 6, rail+ver)

	world.setTile(6, 14, rail+ver)
	world.setTile(6, 13, rail+ver)
	world.setTile(6, 12, rail+ver)
	world.setTile(6, 11, rail+ver)
	world.setTile(6, 10, rail+ver)
	world.setTile(6, 9, rail+dr)
	world.setTile(7, 9, rail+hor)
	world.setTile(8, 9, rail+hor)
	world.setTile(9, 9, rail+hor)
	world.setTile(10, 9, rail+hor)

	trains = []*Train{
		{
			Velocity: 0.25,
			Cars: []*Car{
				SpawnCar(Vec2{0, 38}, world.getTile(0, 2).Node),
				SpawnCar(Vec2{-25, 38}, world.getTile(0, 2).Node),
				SpawnCar(Vec2{-50, 38}, world.getTile(0, 2).Node),
			},
		},
		{
			Velocity: 0.25,
			Cars: []*Car{
				SpawnCar(Vec2{104, 240}, world.getTile(6, 14).Node),
				SpawnCar(Vec2{104, 260}, world.getTile(6, 14).Node),
				SpawnCar(Vec2{104, 280}, world.getTile(6, 14).Node),
			},
		},
	}
}

func SpawnCar(p Vec2, target *PathNode) *Car {
	return &Car{
		Position: p,
		Target:   target,
		// For new cars that start off the screen, the source node is imaginary
		Source: &PathNode{Position: p},
	}
}

var trains []*Train

func (m *Map) setTile(tx, ty int, b byte) {
	if tx < 0 || tx >= MapWidth || ty < 0 || ty >= MapHeight {
		return
	}
	i := TileToOrdinal(tx, ty)
	// Check if not out of bounds
	if i < 0 || i > len(m.tiles)-1 {
		return
	}

	t := &m.tiles[i]
	// Check if there is same path already on this tile or it has car locked on it
	if t.HasSprite(b) || (t.Node != nil && t.Node.GetLocks() > 0) {
		return
	}

	t.Sprites = append(t.Sprites, b)
	// For a tile with switches, the last added sprite will be the primary one
	t.PSprite = len(t.Sprites) - 1

	// Coordinates of top-left corner
	v := TileToPosition(tx, ty)
	t.Node = &PathNode{Position: Vec2{v.X + 8, v.Y + 6}}

	// Adjacent nodes
	l, r, u, d :=
		world.getTile(tx-1, ty),
		world.getTile(tx+1, ty),
		world.getTile(tx, ty-1),
		world.getTile(tx, ty+1)

	// Set connections for all possible directions
	for _, s := range t.Sprites {
		connect(l, r, u, d, t, s)
	}
}

// connect makes a connection to adjacent node paths, depending on the selected sprite
// if adjacent nodes are present.
func connect(l, r, u, d, t *Tile, s byte) {
	if t.Node == nil {
		return
	}

	// First remove any current adjacent connections
	disconnect(t.Node, t.Node.AdjL)
	disconnect(t.Node, t.Node.AdjR)
	disconnect(t.Node, t.Node.AdjU)
	disconnect(t.Node, t.Node.AdjD)

	switch s {
	case rail + hor:
		if l != nil && l.Node != nil &&
			(l.PrimarySprite(rail+hor) || l.PrimarySprite(rail+dr) || l.PrimarySprite(rail+ur)) {
			l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
		}
		if r != nil && r.Node != nil &&
			(r.PrimarySprite(rail+hor) || r.PrimarySprite(rail+dl) || r.PrimarySprite(rail+ul)) {
			r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
		}
	case rail + ver:
		if u != nil && u.Node != nil &&
			(u.PrimarySprite(rail+ver) || u.PrimarySprite(rail+dr) || u.PrimarySprite(rail+dl)) {
			u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
		}
		if d != nil && d.Node != nil &&
			(d.PrimarySprite(rail+ver) || d.PrimarySprite(rail+ur) || d.PrimarySprite(rail+ul)) {
			d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
		}
	case rail + dl:
		if d != nil && d.Node != nil &&
			(d.PrimarySprite(rail+ver) || d.PrimarySprite(rail+ur) || d.PrimarySprite(rail+ul)) {
			d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
		}
		if l != nil && l.Node != nil &&
			(l.PrimarySprite(rail+hor) || l.PrimarySprite(rail+dr) || l.PrimarySprite(rail+ur)) {
			l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
		}
	case rail + dr:
		if d != nil && d.Node != nil &&
			(d.PrimarySprite(rail+ver) || d.PrimarySprite(rail+ur) || d.PrimarySprite(rail+ul)) {
			d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
		}
		if r != nil && r.Node != nil &&
			(r.PrimarySprite(rail+hor) || r.PrimarySprite(rail+dl) || r.PrimarySprite(rail+ul)) {
			r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
		}
	case rail + ul:
		if u != nil && u.Node != nil &&
			(u.PrimarySprite(rail+ver) || u.PrimarySprite(rail+dr) || u.PrimarySprite(rail+dl)) {
			u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
		}
		if l != nil && l.Node != nil &&
			(l.PrimarySprite(rail+hor) || l.PrimarySprite(rail+dr) || l.PrimarySprite(rail+ur)) {
			l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
		}
	case rail + ur:
		if u != nil && u.Node != nil &&
			(u.PrimarySprite(rail+ver) || u.PrimarySprite(rail+dr) || u.PrimarySprite(rail+dl)) {
			u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
		}
		if r != nil && r.Node != nil &&
			(r.PrimarySprite(rail+hor) || r.PrimarySprite(rail+dl) || r.PrimarySprite(rail+ul)) {
			r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
		}
	}
}

// disconnect removes any connections between the two path nodes
func disconnect(a, b *PathNode) {
	if a != nil {
		if a.AdjL == b {
			a.AdjL = nil
		}
		if a.AdjR == b {
			a.AdjR = nil
		}
		if a.AdjU == b {
			a.AdjU = nil
		}
		if a.AdjD == b {
			a.AdjD = nil
		}
	}
	if b != nil {
		if b.AdjL == a {
			b.AdjL = nil
		}
		if b.AdjR == a {
			b.AdjR = nil
		}
		if b.AdjU == a {
			b.AdjU = nil
		}
		if b.AdjD == a {
			b.AdjD = nil
		}
	}
}

func (m *Map) getTile(tx, ty int) *Tile {
	if tx < 0 || tx >= MapWidth || ty < 0 || ty >= MapHeight {
		return nil
	}
	i := TileToOrdinal(tx, ty)
	if i >= 0 && i < len(m.tiles) {
		return &m.tiles[i]
	} else {
		return nil
	}
}

func (m *Map) removeTile(tx, ty int) {
	i := TileToOrdinal(tx, ty)
	u := m.tiles[i].Node
	if u != nil && u.GetLocks() == 0 {
		disconnect(u, u.AdjL)
		disconnect(u, u.AdjR)
		disconnect(u, u.AdjU)
		disconnect(u, u.AdjD)
		m.tiles[i].Node = nil
		m.tiles[i].Sprites = []byte{}
	}
}

func (m *Map) getAll() *[MapWidth * MapHeight]Tile {
	return &m.tiles
}

func (m *Map) ToggleSwitch(tx, ty int) {
	t := m.getTile(tx, ty)
	// Check is a switch and doesn't have any car locked on it
	if t != nil && len(t.Sprites) > 1 && t.Node.GetLocks() == 0 {
		t.PSprite++
		if t.PSprite >= len(t.Sprites) {
			t.PSprite = 0
		} else if t.PSprite < 0 {
			t.PSprite = len(t.Sprites) - 1
		}

		// Adjacent nodes
		l, r, u, d :=
			world.getTile(tx-1, ty),
			world.getTile(tx+1, ty),
			world.getTile(tx, ty-1),
			world.getTile(tx, ty+1)

		connect(l, r, u, d, t, t.Sprites[t.PSprite])
	}
}

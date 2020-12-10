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
	// Check if not out of bounds or there is a same path already on this tile
	if i < 0 || i > len(m.tiles)-1 || world.getTile(tx, ty).HasSprite(b) {
		return
	}

	// Coordinates of top-left corner
	v := TileToPosition(tx, ty)

	t := &m.tiles[i]
	t.Sprites = append(t.Sprites, b)
	t.Node = &PathNode{Position: Vec2{v.X + 8, v.Y + 6}}

	connect(tx, ty)
}

// connect connects node on tile at coordinates tx, ty with path nodes on
// adjacent tiles if applicable
func connect(tx, ty int) {
	// adjacent nodes
	l, r, u, d :=
		world.getTile(tx-1, ty),
		world.getTile(tx+1, ty),
		world.getTile(tx, ty-1),
		world.getTile(tx, ty+1)

	t := world.getTile(tx, ty)
	for _, s := range t.Sprites {
		switch s {
		case rail + hor:
			if l != nil && l.Node != nil && (l.HasSprite(rail+hor) || l.HasSprite(rail+dr) || l.HasSprite(rail+ur)) {
				l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
			}
			if r != nil && r.Node != nil && (r.HasSprite(rail+hor) || r.HasSprite(rail+dl) || r.HasSprite(rail+ul)) {
				r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
			}
			disconnect(t.Node, t.Node.AdjD)
			disconnect(t.Node, t.Node.AdjU)
		case rail + ver:
			if u != nil && u.Node != nil && (u.HasSprite(rail+ver) || u.HasSprite(rail+dr) || u.HasSprite(rail+dl)) {
				u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
			}
			if d != nil && d.Node != nil && (d.HasSprite(rail+ver) || d.HasSprite(rail+ur) || d.HasSprite(rail+ul)) {
				d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
			}
			disconnect(t.Node, t.Node.AdjL)
			disconnect(t.Node, t.Node.AdjR)
		case rail + dl:
			if d != nil && d.Node != nil && (d.HasSprite(rail+ver) || d.HasSprite(rail+ur) || d.HasSprite(rail+ul)) {
				d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
			}
			if l != nil && l.Node != nil && (l.HasSprite(rail+hor) || l.HasSprite(rail+dr) || l.HasSprite(rail+ur)) {
				l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
			}
			disconnect(t.Node, t.Node.AdjU)
			disconnect(t.Node, t.Node.AdjR)
		case rail + dr:
			if d != nil && d.Node != nil && (d.HasSprite(rail+ver) || d.HasSprite(rail+ur) || d.HasSprite(rail+ul)) {
				d.Node.AdjU, t.Node.AdjD = t.Node, d.Node
			}
			if r != nil && r.Node != nil && (r.HasSprite(rail+hor) || r.HasSprite(rail+dl) || r.HasSprite(rail+ul)) {
				r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
			}
			disconnect(t.Node, t.Node.AdjU)
			disconnect(t.Node, t.Node.AdjL)
		case rail + ul:
			if u != nil && u.Node != nil && (u.HasSprite(rail+ver) || u.HasSprite(rail+dr) || u.HasSprite(rail+dl)) {
				u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
			}
			if l != nil && l.Node != nil && (l.HasSprite(rail+hor) || l.HasSprite(rail+dr) || l.HasSprite(rail+ur)) {
				l.Node.AdjR, t.Node.AdjL = t.Node, l.Node
			}
			disconnect(t.Node, t.Node.AdjD)
			disconnect(t.Node, t.Node.AdjR)
		case rail + ur:
			if u != nil && u.Node != nil && (u.HasSprite(rail+ver) || u.HasSprite(rail+dr) || u.HasSprite(rail+dl)) {
				u.Node.AdjD, t.Node.AdjU = t.Node, u.Node
			}
			if r != nil && r.Node != nil && (r.HasSprite(rail+hor) || r.HasSprite(rail+dl) || r.HasSprite(rail+ul)) {
				r.Node.AdjL, t.Node.AdjR = t.Node, r.Node
			}
			disconnect(t.Node, t.Node.AdjD)
			disconnect(t.Node, t.Node.AdjL)
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
	m.tiles[i].Sprites = []byte{}
	u := m.tiles[i].Node
	if u != nil {
		disconnect(u, u.AdjL)
		disconnect(u, u.AdjR)
		disconnect(u, u.AdjU)
		disconnect(u, u.AdjD)
	}
	m.tiles[i].Node = nil
}

func (m *Map) getAll() *[MapWidth * MapHeight]Tile {
	return &m.tiles
}

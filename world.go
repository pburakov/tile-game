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
			Velocity: 1.0,
			Cars: []*Car{
				SpawnCar(Vec2{0, 38}, world.getTile(0, 2).Node),
				SpawnCar(Vec2{-25, 38}, world.getTile(0, 2).Node),
				SpawnCar(Vec2{-50, 38}, world.getTile(0, 2).Node),
			},
		},
		{
			Velocity: 1.0,
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
		Source: NewPathNode(p, false),
	}
}

var trains []*Train

func (m *Map) setTile(tx, ty int, b byte) {
	if tx < 0 || tx >= MapWidth || ty < 0 || ty >= MapHeight {
		return
	}
	i := TileToOrdinal(tx, ty)
	// Check if not out of bounds or there is a path already on this tile
	if i < 0 || i > len(m.tiles)-1 || world.getTile(tx, ty).Sprite != none {
		return
	}

	// Coordinates of top-left corner
	v := TileToPosition(tx, ty)

	t := &m.tiles[i]
	t.Sprite = b
	t.Node = NewPathNode(
		Vec2{v.X + 8, v.Y + 6}, b != rail+hor && b != rail+ver,
	)

	connectAdj(tx, ty)
}

// connectAdj connects node on tile at coordinates tx, ty with path nodes on adjacent tiles if applicable
func connectAdj(tx, ty int) {
	// adjacent nodes
	l, r, u, d :=
		world.getTile(tx-1, ty),
		world.getTile(tx+1, ty),
		world.getTile(tx, ty-1),
		world.getTile(tx, ty+1)

	t := world.getTile(tx, ty)
	switch t.Sprite {
	case rail + hor:
		if l != nil && (l.Sprite == rail+hor || l.Sprite == rail+dr || l.Sprite == rail+ur) {
			Connect(l, t)
		}
		if r != nil && (r.Sprite == rail+hor || r.Sprite == rail+dl || r.Sprite == rail+ul) {
			Connect(r, t)
		}
	case rail + ver:
		if u != nil && (u.Sprite == rail+ver || u.Sprite == rail+dr || u.Sprite == rail+dl) {
			Connect(u, t)
		}
		if d != nil && (d.Sprite == rail+ver || d.Sprite == rail+ur || d.Sprite == rail+ul) {
			Connect(d, t)
		}
	case rail + dl:
		if d != nil && (d.Sprite == rail+ver || d.Sprite == rail+ur || d.Sprite == rail+ul) {
			Connect(d, t)
		}
		if l != nil && (l.Sprite == rail+hor || l.Sprite == rail+dr || l.Sprite == rail+ur) {
			Connect(l, t)
		}
	case rail + dr:
		if d != nil && (d.Sprite == rail+ver || d.Sprite == rail+ur || d.Sprite == rail+ul) {
			Connect(d, t)
		}
		if r != nil && (r.Sprite == rail+hor || r.Sprite == rail+dl || r.Sprite == rail+ul) {
			Connect(r, t)
		}
	case rail + ul:
		if u != nil && (u.Sprite == rail+ver || u.Sprite == rail+dr || u.Sprite == rail+dl) {
			Connect(u, t)
		}
		if l != nil && (l.Sprite == rail+hor || l.Sprite == rail+dr || l.Sprite == rail+ur) {
			Connect(l, t)
		}
	case rail + ur:
		if u != nil && (u.Sprite == rail+ver || u.Sprite == rail+dr || u.Sprite == rail+dl) {
			Connect(u, t)
		}
		if r != nil && (r.Sprite == rail+hor || r.Sprite == rail+dl || r.Sprite == rail+ul) {
			Connect(r, t)
		}
	}
}

// Connect connects nodes on tiles a and b
func Connect(a, b *Tile) {
	if a == nil || b == nil || a.Node == nil || b.Node == nil {
		return
	}
	a.Node.Adj[b.Node.Id] = b.Node
	b.Node.Adj[a.Node.Id] = a.Node
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
	m.tiles[i].Sprite = 0
	u := m.tiles[i].Node
	if u != nil {
		// Disconnect adjacent connected nodes, if any
		for _, v := range u.Adj {
			delete(v.Adj, u.Id)
		}
		m.tiles[i].Node = nil
	}
}

func (m *Map) getAll() *[MapWidth * MapHeight]Tile {
	return &m.tiles
}

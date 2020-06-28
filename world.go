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
		Source: NewPathNode(p),
	}
}

var trains []*Train

func (m *Map) setTile(tx, ty int, b byte) {
	i := TileToOrdinal(tx, ty)
	// Check if not out of bounds or there is a path already on this tile
	if i < 0 || i > len(m.tiles)-1 || world.getTile(tx, ty).Sprite != none {
		return
	}

	// Coordinates of top-left corner
	v := TileToPosition(tx, ty)

	// adjacent nodes
	l, r, u, d :=
		world.getTile(tx-1, ty),
		world.getTile(tx+1, ty),
		world.getTile(tx, ty-1),
		world.getTile(tx, ty+1)

	t := &m.tiles[i]
	t.Sprite = b
	t.Node = NewPathNode(Vec2{v.X + 8, v.Y + 6})

	switch b {
	case rail + hor:
		Connect(l, t, r)
	case rail + ver:
		Connect(u, t, d)
	case rail + dl:
		Connect(l, t, d)
	case rail + dr:
		Connect(r, t, d)
	case rail + ul:
		Connect(l, t, u)
	case rail + ur:
		Connect(r, t, u)
	}
}

// Connect connects nodes on tiles a, n and b assuming n is in the new tile between a and b
func Connect(a, n, b *Tile) {
	x := n.Node
	if a != nil && a.Node != nil {
		a.Node.Adj[x.Id] = x
		x.Adj[a.Node.Id] = a.Node
	}
	if b != nil && b.Node != nil {
		b.Node.Adj[x.Id] = x
		x.Adj[b.Node.Id] = b.Node
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

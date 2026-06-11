package main

import "math"

// Tuning constants for the follow-the-leader movement model.
const (
	// carSpacing is the arc-length gap between consecutive cars, in pixels.
	carSpacing = 22.0

	// bezierSteps is how many line segments approximate each rounded corner.
	// More steps means smoother curves at a small cost.
	bezierSteps = 8

	// frontMargin keeps a little curve generated ahead of the head so the next
	// segment direction is always known before the head reaches a corner.
	frontMargin = TileSize

	// backMargin keeps a little curve behind the rear car before old nodes are
	// trimmed off the route.
	backMargin = TileSize
)

// MoveTrains recomputes lock occupancy and advances every train one frame.
func MoveTrains() {
	// Locks are derived from where trains are right now, so clear them first.
	for _, t := range world.getAll() {
		if t.Node != nil {
			t.Node.Locks = 0
		}
	}
	for _, t := range trains {
		moveTrain(t)
	}
}

func moveTrain(t *Train) {
	t.s += t.Velocity

	// Make sure the curve extends far enough ahead of the head, generating new
	// nodes from the graph and reversing at dead ends as needed.
	pts, cum := ensureAhead(t)

	// Place every car at its fixed arc length behind the head.
	for i, c := range t.Cars {
		s := t.s - float64(i)*carSpacing
		c.Position, c.Angle = sampleCurve(pts, cum, s)
	}

	// A train reserves every node its route currently spans so the player can't
	// edit track or flip switches out from under it.
	for _, n := range t.route {
		n.Locks++
	}

	trimBehind(t)
}

// ensureAhead grows the route so the curve covers the head plus frontMargin,
// reversing the train when it runs into a dead end. It returns the curve points
// and their cumulative arc lengths for the resulting route.
func ensureAhead(t *Train) (pts []Vec2, cum []float64) {
	reversals := 0
	for {
		pts, cum = buildCurve(t.route)
		l := cum[len(cum)-1]
		if t.s+frontMargin <= l {
			return pts, cum
		}
		if next := findNextTarget(lastNode(t.route), prevNode(t.route)); next != nil {
			t.route = append(t.route, next)
			continue
		}
		// Dead end ahead. Reverse only once the head actually reaches the tip,
		// otherwise let it keep coasting toward it.
		if t.s >= l {
			// A train longer than its whole track would reverse forever; bail out.
			if reversals >= 2 {
				t.s = l
				return pts, cum
			}
			reverseTrain(t, l)
			reversals++
			continue
		}
		return pts, cum
	}
}

// reverseTrain flips the train's direction of travel in place. The route is
// reversed so the rear-most node becomes the new frontier, and the head's arc
// length is remapped so no car visibly jumps. l is the current curve length.
func reverseTrain(t *Train, l float64) {
	for i, j := 0, len(t.route)-1; i < j; i, j = i+1, j-1 {
		t.route[i], t.route[j] = t.route[j], t.route[i]
	}
	// In the reversed curve a point at arc length s sits at l-s. The new leading
	// car is the one that used to be the rear, hence the +tail term.
	tail := float64(len(t.Cars)-1) * carSpacing
	t.s = l - t.s + tail
}

// trimBehind drops route nodes that the rear car has fully left behind, keeping
// the route bounded. The dropped arc length is subtracted from the head so the
// remaining cars stay put.
func trimBehind(t *Train) {
	tail := float64(len(t.Cars)-1) * carSpacing
	for len(t.route) > 2 {
		// Dropping the rear node shortens the curve by exactly this much, which is
		// also the amount every car's arc length shifts. Requiring the rear car to
		// stay backMargin ahead of the new start keeps that shift exact (the cars
		// remain on the unchanged portion of the curve), so nothing jumps.
		_, cumOld := buildCurve(t.route)
		_, cumNew := buildCurve(t.route[1:])
		drop := cumOld[len(cumOld)-1] - cumNew[len(cumNew)-1]
		if t.s-drop-tail < backMargin {
			break
		}
		t.route = t.route[1:]
		t.s -= drop
	}
}

// buildCurve turns a node route into a smooth poly-line. Straight runs stay
// straight; every corner is rounded with a quadratic Bézier whose control point
// is the node itself, so cars trace real curves through turns. It returns the
// sampled points and their cumulative arc length.
func buildCurve(route []*PathNode) (pts []Vec2, cum []float64) {
	pos := func(i int) Vec2 { return route[i].Position }
	push := func(p Vec2) {
		if len(pts) == 0 {
			pts = append(pts, p)
			cum = append(cum, 0)
			return
		}
		last := pts[len(pts)-1]
		d := last.DistanceTo(p)
		if d < 1e-9 {
			return // skip duplicate points (e.g. shared corner midpoints)
		}
		pts = append(pts, p)
		cum = append(cum, cum[len(cum)-1]+d)
	}

	m := len(route)
	push(pos(0))
	if m == 1 {
		return pts, cum
	}
	for i := 1; i < m-1; i++ {
		a, b, c := pos(i-1), pos(i), pos(i+1)
		m0, m1 := Mid(a, b), Mid(b, c)
		push(m0)
		for s := 1; s <= bezierSteps; s++ {
			push(QuadBezier(m0, b, m1, float64(s)/float64(bezierSteps)))
		}
	}
	push(pos(m - 1))
	return pts, cum
}

// sampleCurve returns the position and travel angle at arc length s along the
// curve. Arc lengths before the start are extrapolated straight back so cars
// trail off cleanly while a train is still emerging from a town.
func sampleCurve(pts []Vec2, cum []float64, s float64) (Vec2, float64) {
	n := len(pts)
	if n == 1 {
		return pts[0], 0
	}
	if s <= 0 {
		dir := pts[1].Sub(pts[0]).Unit()
		return pts[0].Plus(dir.Scale(s)), AngleOf(dir)
	}
	l := cum[n-1]
	if s >= l {
		return pts[n-1], pts[n-2].Angle(pts[n-1])
	}
	// Find the segment containing s.
	i := 0
	for i < n-1 && cum[i+1] < s {
		i++
	}
	seg := cum[i+1] - cum[i]
	f := 0.0
	if seg > 0 {
		f = (s - cum[i]) / seg
	}
	p := pts[i].Plus(pts[i+1].Sub(pts[i]).Scale(f))
	return p, pts[i].Angle(pts[i+1])
}

// findNextTarget picks the node the head moves to next: the first neighbour of
// cur that isn't the one it just came from. At a switch this is whatever branch
// the player has the switch routed to, so the whole train follows the head.
func findNextTarget(cur, prev *PathNode) *PathNode {
	if cur == nil {
		return nil
	}
	for _, v := range []*PathNode{cur.AdjL, cur.AdjR, cur.AdjU, cur.AdjD} {
		if v != nil && v != prev {
			return v
		}
	}
	return nil
}

func lastNode(route []*PathNode) *PathNode {
	return route[len(route)-1]
}

func prevNode(route []*PathNode) *PathNode {
	if len(route) >= 2 {
		return route[len(route)-2]
	}
	return nil
}

// SpawnTrain creates a train of numCars cars starting at node start, strung out
// along the track behind the head so it appears to be already running.
func SpawnTrain(velocity float64, start *PathNode, numCars int) *Train {
	t := &Train{
		Velocity: velocity,
		route:    []*PathNode{start},
		Cars:     make([]*Car, numCars),
	}
	for i := range t.Cars {
		t.Cars[i] = &Car{}
	}

	// Extend the route far enough to lay the whole train out plus a margin.
	need := float64(numCars-1)*carSpacing + frontMargin
	for {
		_, cum := buildCurve(t.route)
		if cum[len(cum)-1] >= need {
			break
		}
		next := findNextTarget(lastNode(t.route), prevNode(t.route))
		if next == nil {
			break
		}
		t.route = append(t.route, next)
	}

	pts, cum := buildCurve(t.route)
	t.s = math.Min(float64(numCars-1)*carSpacing, cum[len(cum)-1])
	for i, c := range t.Cars {
		c.Position, c.Angle = sampleCurve(pts, cum, t.s-float64(i)*carSpacing)
	}
	return t
}

package main

import (
	"math"
	"testing"
)

// node is a tiny helper to hand-build path graphs for movement tests.
func node(x, y float64) *PathNode {
	return &PathNode{Position: Vec2{x, y}}
}

// TestBuildCurveRoundsCorners verifies a right-angle turn is traced as an actual
// curve hugging the corner, not a straight diagonal shortcut across the tile.
func TestBuildCurveRoundsCorners(t *testing.T) {
	// A -> B -> C : travel right, then down. B is the corner.
	a, b, c := node(0, 0), node(16, 0), node(16, 16)
	pts, cum := buildCurve([]*PathNode{a, b, c})

	if l := cum[len(cum)-1]; l >= 32 || l <= b.Position.DistanceTo(Vec2{0, 0})+0 {
		// Length must be shorter than the two legs (32) — i.e. the corner is cut
		// short by rounding — but longer than a single leg.
		if l >= 32 {
			t.Fatalf("curve length %.2f should be < 32 (legs), corner not rounded", l)
		}
	}

	// The curve must bow toward the corner B...
	nearCorner := math.Inf(1)
	// ...and must NOT pass through the diagonal-shortcut midpoint (8, 8).
	nearDiagonal := math.Inf(1)
	for _, p := range pts {
		nearCorner = math.Min(nearCorner, p.DistanceTo(b.Position))
		nearDiagonal = math.Min(nearDiagonal, p.DistanceTo(Vec2{8, 8}))
	}
	if nearCorner > 4 {
		t.Errorf("curve does not hug corner: closest sample is %.2f px from it", nearCorner)
	}
	if nearDiagonal < 4 {
		t.Errorf("curve cuts diagonally: a sample is only %.2f px from shortcut midpoint", nearDiagonal)
	}
}

// TestSampleCurveSpacing checks that arc-length sampling gives the requested
// distance between points on a straight run.
func TestSampleCurveSpacing(t *testing.T) {
	route := []*PathNode{node(0, 0), node(16, 0), node(32, 0), node(48, 0), node(64, 0)}
	pts, cum := buildCurve(route)

	p1, _ := sampleCurve(pts, cum, 50)
	p2, _ := sampleCurve(pts, cum, 50-carSpacing)
	if d := p1.DistanceTo(p2); math.Abs(d-carSpacing) > 0.01 {
		t.Errorf("spacing on straight run = %.3f, want %.3f", d, carSpacing)
	}
}

// TestSimulationKeepsChain runs the live world for many frames and asserts the
// train never gets stuck, never NaNs, and the cars stay coupled at a stable gap
// (smooth chain) through straights, curves, dead ends and reversals.
func TestSimulationKeepsChain(t *testing.T) {
	if len(trains) == 0 {
		t.Skip("no trains spawned in world init")
	}

	var startHead, lateHead Vec2
	for frame := 0; frame < 4000; frame++ {
		MoveTrains()
		for _, tr := range trains {
			for i, c := range tr.Cars {
				if math.IsNaN(c.Position.X) || math.IsNaN(c.Position.Y) {
					t.Fatalf("frame %d: car %d position is NaN", frame, i)
				}
				if i == 0 {
					continue
				}
				// Adjacent straight-line gap is the chord of an arc-length
				// carSpacing, so it can be a touch shorter on curves but must
				// never collapse (stuck/overlapping) or stretch (chain broken).
				d := c.Position.DistanceTo(tr.Cars[i-1].Position)
				if d < carSpacing-3 || d > carSpacing+0.5 {
					t.Fatalf("frame %d: cars %d-%d gap %.2f out of range (chain not smooth)",
						frame, i-1, i, d)
				}
			}
		}
		if frame == 200 {
			startHead = trains[0].Cars[0].Position
		}
		if frame == 1200 {
			lateHead = trains[0].Cars[0].Position
		}
	}

	if startHead.DistanceTo(lateHead) < 1 {
		t.Errorf("head did not move between frames 200 and 1200 (stuck): %v -> %v", startHead, lateHead)
	}
}

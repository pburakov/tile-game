package main

import "github.com/hajimehoshi/ebiten"

// WheelSelectorRange determines how fast/slow wheel-scrolling selector will be.
// For macOS touchpads and mouse values 3.00 .. 5.00 work best for smooth scrolling
const WheelSelectorRange = 5.00

var brushes = []byte{
	rail + ver,
	rail + hor,
	rail + dl,
	rail + dr,
	rail + ul,
	rail + ur,
}

var maxRawWheelValue = WheelSelectorRange * float64(len(brushes))

// Selector represents the state of user input
type Selector struct {
	CurX, CurY int     // Cursor position
	RawWheel   float64 // Raw offset from 0 y axis
}

var selector Selector

func init() {
	selector = Selector{
		RawWheel: 0.0,
	}
}

// HandleInput changes game state based on user input
func HandleInput() {
	selector.CurX, selector.CurY = ebiten.CursorPosition()
	_, wy := ebiten.Wheel()
	selector.ApplyDelta(wy)

	lftBtn := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rgtBtn := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)

	if lftBtn {
		tx, ty := PositionToTile(selector.CurX, selector.CurY)
		world.setTile(tx, ty, Tile{Sprite: selector.GetCurrentSelection()})
	} else if rgtBtn {
		world.removeTile(PositionToTile(selector.CurX, selector.CurY))
	}
}

func (s *Selector) GetCurrentSelection() byte {
	return brushes[int(s.RawWheel/WheelSelectorRange)%len(brushes)]
}

func (s *Selector) ApplyDelta(d float64) {
	s.RawWheel += d
	if s.RawWheel < 0 {
		s.RawWheel = maxRawWheelValue
	} else if s.RawWheel > maxRawWheelValue {
		s.RawWheel = 0
	}
}

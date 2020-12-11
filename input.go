package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// WheelSelectorRange determines how fast/slow wheel-scrolling selector will be.
// For macOS touchpads and mouse values 3.00 .. 5.00 work best for smooth scrolling
const WheelSelectorRange = 1.00

var brushes = []byte{
	rail + hor,
	rail + ver,
	rail + dl,
	rail + dr,
	rail + ul,
	rail + ur,
}

var maxRawWheelValue = WheelSelectorRange * float64(len(brushes))

// Selector represents the state of user input
type Selector struct {
	CurX, CurY  int     // CurX, CurY are the coordinates of cursor position
	RawWheel    float64 // RawWheel is a raw offset value from 0 on y-axis
	Current     int     // Current is a currently selected index
	ControlMode bool    // ControlMode is true when not in construction mode
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
	selector.ApplyWheelDelta(wy)

	selNxt := inpututil.IsKeyJustPressed(ebiten.KeyQ)
	selPrv := inpututil.IsKeyJustPressed(ebiten.KeyE)
	selector.ApplyKeyDelta(selNxt, selPrv)

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		selector.ControlMode = !selector.ControlMode
	}

	if selector.ControlMode {
		ebiten.SetCursorVisible(true)
		lftBtn := inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
		if lftBtn {
			tx, ty := PositionToTile(selector.CurX, selector.CurY)
			world.ToggleSwitch(tx, ty)
		}
	} else {
		ebiten.SetCursorVisible(false)
		lftBtn := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		rgtBtn := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
		if lftBtn {
			tx, ty := PositionToTile(selector.CurX, selector.CurY)
			world.setTile(tx, ty, selector.GetCurrentSelection())
		} else if rgtBtn {
			world.removeTile(PositionToTile(selector.CurX, selector.CurY))
		}
	}
}

func (s *Selector) GetCurrentSelection() byte {
	return brushes[s.Current]
}

func (s *Selector) ApplyKeyDelta(nxt, prv bool) {
	if nxt {
		s.Current++
	} else if prv {
		s.Current--
	}
	if s.Current >= len(brushes) {
		s.Current = 0
	} else if s.Current < 0 {
		s.Current = len(brushes) - 1
	}
}

func (s *Selector) ApplyWheelDelta(d float64) {
	if d == 0.0 {
		return
	}
	s.RawWheel += d
	if s.RawWheel < 0 {
		s.RawWheel = maxRawWheelValue
	} else if s.RawWheel > maxRawWheelValue {
		s.RawWheel = 0
	}
	s.Current = int(s.RawWheel/WheelSelectorRange) % len(brushes)
}

package main

const WheelSelectorRange = 7.25 // Determines how fast/slow wheel-scrolling selector would be

var brushes = []byte{
	rail + ver,
	rail + hor,
	rail + cro,
	rail + dl,
	rail + dr,
	rail + ul,
	rail + ur,
	rail + vl,
	rail + vr,
	rail + hu,
	rail + hd,
}

var maxRawValue = WheelSelectorRange * float64(len(brushes))

type Selector struct {
	Raw float64 // Raw offset from 0 y axis
}

var selector Selector

func init() {
	selector = Selector{
		Raw: 0.0,
	}
}

func (s *Selector) GetCurrentSelection() byte {
	return brushes[int(s.Raw/WheelSelectorRange)%len(brushes)]
}

func (s *Selector) ApplyDelta(d float64) {
	s.Raw += d
	if s.Raw < 0 {
		s.Raw = maxRawValue
	} else if s.Raw > maxRawValue {
		s.Raw = 0
	}
}

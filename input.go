package main

type Selector struct {
	Brushes []byte
	Current int
}

var selector *Selector

func init() {
	selector = &Selector{
		Brushes: []byte{
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
		},
		Current: 0,
	}
}

func (s *Selector) GetCurrentSelection() byte {
	return s.Brushes[s.Current]
}

func (s *Selector) Decrement() {
	s.Current = (s.Current - 1) % len(s.Brushes)
}

func (s *Selector) Increment() {
	s.Current = (s.Current + 1) % len(s.Brushes)
}

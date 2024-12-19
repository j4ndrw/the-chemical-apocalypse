package components

type Speed int32

func (s *Speed) AsFloat() float64 {
	return float64(*s)
}

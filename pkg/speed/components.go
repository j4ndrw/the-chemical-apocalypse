package speed

type SpeedComponent int32

func (s *SpeedComponent) AsFloat() float64 {
	return float64(*s)
}

package position

type PositionComponent struct {
	X, Y int32
}

func (p *PositionComponent) FloatX() float64 {
	return float64(p.X)
}

func (p *PositionComponent) FloatY() float64 {
	return float64(p.Y)
}

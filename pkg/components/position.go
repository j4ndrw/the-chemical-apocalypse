package components

type Position struct {
	X, Y int32
}

func (p *Position) FloatX() float64 {
	return float64(p.X)
}

func (p *Position) FloatY() float64 {
	return float64(p.Y)
}

package components

type Position struct {
	X, Y          int32
	Width, Height int32
	Bound         *Bound
}

func (p *Position) Left() int32 {
	return p.X
}

func (p *Position) Right() int32 {
	return p.X + p.Width
}

func (p *Position) Top() int32 {
	return p.Y
}

func (p *Position) Bottom() int32 {
	return p.Y + p.Height
}

package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Hitbox struct {
	rl.Color
	Position
	Width, Height int32
	Bound         *Bound
}

func (h *Hitbox) Left() int32 {
	return h.Position.X
}

func (h *Hitbox) Right() int32 {
	return h.Left() + h.Width
}

func (h *Hitbox) Top() int32 {
	return h.Position.Y
}

func (h *Hitbox) Bottom() int32 {
	return h.Top() + h.Height
}

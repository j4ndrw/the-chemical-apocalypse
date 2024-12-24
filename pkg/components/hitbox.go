package components

import rl "github.com/gen2brain/raylib-go/raylib"

type Hitbox struct {
	Color         rl.Color
	Position      Position
	Direction     Direction
	Width, Height int32
	Bound         *Bound
	Hidden        bool
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

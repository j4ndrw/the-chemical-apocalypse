package hitbox

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/bound"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
)

type HitboxComponent struct {
	Color             rl.Color
	Position          position.PositionComponent
	Direction         direction.DirectionComponent
	Width, Height     int32
	Bound             *bound.BoundComponent
	Hidden            bool
	PaddingPercentage float32
}

func (h *HitboxComponent) Left() int32 {
	return h.Position.X + int32(h.PaddingPercentage*float32(h.Width))
}

func (h *HitboxComponent) Right() int32 {
	return h.Left() + h.Width - int32(h.PaddingPercentage*float32(h.Width))
}

func (h *HitboxComponent) Top() int32 {
	return h.Position.Y + int32(h.PaddingPercentage*float32(h.Height))
}

func (h *HitboxComponent) Bottom() int32 {
	return h.Top() + h.Height - int32(h.PaddingPercentage*float32(h.Height))
}

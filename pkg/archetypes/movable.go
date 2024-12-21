package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type movable struct{}

var Movable = movable{}

func (_ *movable) MoveUp(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Position.Y < h.Bound.Top {
		return
	}
	h.Position.Y -= int32(*s)
	h.Direction.Y = components.DirectionUp
}

func (_ *movable) MoveDown(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Position.Y > h.Bound.Bottom {
		return
	}
	h.Position.Y += int32(*s)
	h.Direction.Y = components.DirectionDown
}

func (_ *movable) MoveLeft(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Position.X < h.Bound.Left {
		return
	}
	h.Position.X -= int32(*s)
	h.Direction.X = components.DirectionLeft
}

func (_ *movable) MoveRight(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Position.X > h.Bound.Right {
		return
	}
	h.Position.X += int32(*s)
	h.Direction.X = components.DirectionRight
}

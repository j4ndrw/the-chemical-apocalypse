package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type movable struct{}

var Movable = movable{}

func (_ *movable) MoveUp(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Y < h.Bound.Top {
		return
	}
	h.Y -= int32(*s)
}

func (_ *movable) MoveDown(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.Y > h.Bound.Bottom {
		return
	}
	h.Y += int32(*s)
}

func (_ *movable) MoveLeft(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.X < h.Bound.Left {
		return
	}
	h.X -= int32(*s)
}

func (_ *movable) MoveRight(w *meta.Window, h *components.Hitbox, s *components.Speed) {
	if h.X > h.Bound.Right {
		return
	}
	h.X += int32(*s)
}

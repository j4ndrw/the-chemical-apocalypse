package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type movable struct{}

var Movable = movable{}

func (_ *movable) MoveUp(m *meta.Meta, h *components.Hitbox, s *components.Speed) {
	if h.Position.Y < h.Bound.Top {
		return
	}
	h.Position.Y = int32(float32(h.Position.Y) - float32(*s)*m.DeltaTime)
	h.Direction.Y = components.DirectionUp
}

func (_ *movable) MoveDown(m *meta.Meta, h *components.Hitbox, s *components.Speed) {
	if h.Position.Y > h.Bound.Bottom {
		return
	}
	h.Position.Y = int32(float32(h.Position.Y) + float32(*s)*m.DeltaTime)
	h.Direction.Y = components.DirectionDown
}

func (_ *movable) MoveLeft(m *meta.Meta, h *components.Hitbox, s *components.Speed) {
	if h.Position.X < h.Bound.Left {
		return
	}
	h.Position.X = int32(float32(h.Position.X) - float32(*s)*m.DeltaTime)
	h.Direction.X = components.DirectionLeft
}

func (_ *movable) MoveRight(m *meta.Meta, h *components.Hitbox, s *components.Speed) {
	if h.Position.X > h.Bound.Right {
		return
	}
	h.Position.X = int32(float32(h.Position.X) + float32(*s)*m.DeltaTime)
	h.Direction.X = components.DirectionRight
}

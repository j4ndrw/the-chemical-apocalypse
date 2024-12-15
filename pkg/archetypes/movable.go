package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type movable struct{}

var Movable movable = movable{}

func (_ *movable) MoveUp(w *meta.Window, p *components.Position, s *components.Speed) {
	if p.Y < p.Bound.Top {
		return
	}
	p.Y -= int32(*s)
}

func (_ *movable) MoveDown(w *meta.Window, p *components.Position, s *components.Speed) {
	if p.Y > p.Bound.Bottom {
		return
	}
	p.Y += int32(*s)
}

func (_ *movable) MoveLeft(w *meta.Window, p *components.Position, s *components.Speed) {
	if p.X < p.Bound.Left {
		return
	}
	p.X -= int32(*s)
}

func (_ *movable) MoveRight(w *meta.Window, p *components.Position, s *components.Speed) {
	if p.X > p.Bound.Right {
		return
	}
	p.X += int32(*s)
}

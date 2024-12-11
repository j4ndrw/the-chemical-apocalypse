package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type PlayerSystem struct {
	MoveUp    system.System
	MoveDown  system.System
	MoveLeft  system.System
	MoveRight system.System
}

func Player() *PlayerSystem {
	return &PlayerSystem{
		MoveUp: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Player.Position.Y -= 1
			}),
		MoveDown: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Player.Position.Y += 1
			}),
		MoveLeft: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Player.Position.X -= 1
			}),
		MoveRight: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Player.Position.X += 1
			}),
	}
}

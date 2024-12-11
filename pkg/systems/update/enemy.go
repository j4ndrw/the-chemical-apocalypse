package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type EnemySystem struct {
	MoveUp    system.System
	MoveDown  system.System
	MoveLeft  system.System
	MoveRight system.System
}

func Enemy() *EnemySystem {
	return &EnemySystem{
		MoveUp: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Enemy.Position.Y -= 3
			}),
		MoveDown: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Enemy.Position.Y += 3
			}),
		MoveLeft: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Enemy.Position.X -= 3
			}),
		MoveRight: *system.Create(
			func(s *world.World, m *meta.Meta) {
				s.Enemy.Position.X += 3
			}),
	}
}

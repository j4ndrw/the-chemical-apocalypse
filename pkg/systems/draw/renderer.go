package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type RendererSystem struct {
	Clear  system.System
	Player system.System
	Enemy  system.System
}

func Renderer() *RendererSystem {
	return &RendererSystem{
		Clear: *system.Create(func(s *world.World, m *meta.Meta) {
			rl.ClearScreenBuffers()
		}),
		Player: *system.Create(func(s *world.World, m *meta.Meta) {
			rl.DrawRectangle(
				s.Player.Position.X,
				s.Player.Position.Y,
				40,
				40,
				rl.Color{
					0,
					0xFF,
					0,
					0xFF,
				},
			)
		}),
		Enemy: *system.Create(func(s *world.World, m *meta.Meta) {
			rl.DrawRectangle(
				s.Enemy.Position.X,
				s.Enemy.Position.Y,
				40,
				40,
				rl.Color{
					0xFF,
					0,
					0,
					0xFF,
				},
			)
		}),
	}
}

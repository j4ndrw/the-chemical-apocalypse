package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type player struct{}

var Player player = player{}

func (_ *player) HandleMovement() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
			archetypes.Movable.MoveUp(&m.Window, &w.Player.Position, &w.Player.Speed)
		}

		if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
			archetypes.Movable.MoveDown(&m.Window, &w.Player.Position, &w.Player.Speed)
		}

		if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
			archetypes.Movable.MoveLeft(&m.Window, &w.Player.Position, &w.Player.Speed)
		}

		if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
			archetypes.Movable.MoveRight(&m.Window, &w.Player.Position, &w.Player.Speed)
		}
	})
}

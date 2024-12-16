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
		handle := func(
			moveFunction archetypes.MoveFn,
			undoFunction archetypes.MoveFn,
			keys ...int32,
		) {
			for _, key := range keys {
				if rl.IsKeyDown(key) {
					moveFunction(&m.Window, &w.Player.Position, &w.Player.Speed)
					if archetypes.Collidable.IsColliding(&w.Player.Position, &w.Enemy.Position) {
						undoFunction(&m.Window, &w.Player.Position, &w.Player.Speed)
					}
					return
				}
			}
		}

		handle(archetypes.Movable.MoveUp, archetypes.Movable.MoveDown, rl.KeyW, rl.KeyUp)
		handle(archetypes.Movable.MoveDown, archetypes.Movable.MoveUp, rl.KeyS, rl.KeyDown)
		handle(archetypes.Movable.MoveLeft, archetypes.Movable.MoveRight, rl.KeyA, rl.KeyLeft)
		handle(archetypes.Movable.MoveRight, archetypes.Movable.MoveLeft, rl.KeyD, rl.KeyRight)
	})
}

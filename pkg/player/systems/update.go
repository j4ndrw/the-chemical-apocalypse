package player_systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/collidable"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/movement"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/types"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type update struct{}

var Update = update{}

func (_ *update) HandleMovement() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		oldX := w.Player.Hitbox.Position.X
		oldY := w.Player.Hitbox.Position.Y

		handle := func(
			moveFunction types.MoveFn,
			undoFunction types.MoveFn,
			keys ...int32,
		) {
			for _, key := range keys {
				if rl.IsKeyDown(key) {
					moveFunction(m, &w.Player.Hitbox, &w.Player.Speed)
					if collidable.Archetype.IsColliding(
						&w.Player.Hitbox,
						utils.SliceMap(
							utils.MapValues(w.Enemies),
							func(enemy *enemy.EnemyEntity) *hitbox.HitboxComponent {
								return &enemy.Hitbox
							},
						)...,
					) {
						undoFunction(m, &w.Player.Hitbox, &w.Player.Speed)
					}
					return
				}
			}
		}

		handle(movement.Archetype.MoveUp, movement.Archetype.MoveDown, rl.KeyW, rl.KeyUp)
		handle(movement.Archetype.MoveDown, movement.Archetype.MoveUp, rl.KeyS, rl.KeyDown)
		handle(movement.Archetype.MoveLeft, movement.Archetype.MoveRight, rl.KeyA, rl.KeyLeft)
		handle(movement.Archetype.MoveRight, movement.Archetype.MoveLeft, rl.KeyD, rl.KeyRight)

		w.Player.Moving = oldX != w.Player.Hitbox.Position.X || oldY != w.Player.Hitbox.Position.Y
		if !w.Player.Moving {
			w.Player.Hitbox.Direction.X = direction.None
			w.Player.Hitbox.Direction.Y = direction.None
		}
	})
}

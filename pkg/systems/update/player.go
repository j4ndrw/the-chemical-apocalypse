package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type player struct{}

var Player = player{}

func (_ *player) HandleMovement() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration { return }

		handle := func(
			moveFunction archetypes.MoveFn,
			undoFunction archetypes.MoveFn,
			keys ...int32,
		) {
			for _, key := range keys {
				if rl.IsKeyDown(key) {
					moveFunction(&m.Window, &w.Player.Hitbox, &w.Player.Speed)
					if archetypes.Collidable.IsColliding(
						&w.Player.Hitbox,
						utils.SliceMap(
							utils.MapValues(w.Enemies),
							func(enemy *entities.Enemy) *components.Hitbox {
								return &enemy.Hitbox
							},
						)...,
					) {
						undoFunction(&m.Window, &w.Player.Hitbox, &w.Player.Speed)
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

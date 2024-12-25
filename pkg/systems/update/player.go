package systems

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/coroutines"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type player struct{}

var Player = player{}

func (_ *player) HandleMovement() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		handle := func(
			moveFunction archetypes.MoveFn,
			undoFunction archetypes.MoveFn,
			keys ...int32,
		) {
			for _, key := range keys {
				if rl.IsKeyDown(key) {
					moveFunction(m, &w.Player.Hitbox, &w.Player.Speed)
					if archetypes.Collidable.IsColliding(
						&w.Player.Hitbox,
						utils.SliceMap(
							utils.MapValues(w.Enemies),
							func(enemy *entities.Enemy) *components.Hitbox {
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

		handle(archetypes.Movable.MoveUp, archetypes.Movable.MoveDown, rl.KeyW, rl.KeyUp)
		handle(archetypes.Movable.MoveDown, archetypes.Movable.MoveUp, rl.KeyS, rl.KeyDown)
		handle(archetypes.Movable.MoveLeft, archetypes.Movable.MoveRight, rl.KeyA, rl.KeyLeft)
		handle(archetypes.Movable.MoveRight, archetypes.Movable.MoveLeft, rl.KeyD, rl.KeyRight)
	})
}

func (_ *player) UpdateIdleSprite(sheetRow int32) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		sprite := w.Player.SpriteMap[constants.Keys.Idle]

		variants := int32(float32(m.SpriteAtlas.Width) / (float32(w.Player.Hitbox.Width) / sprite.Scale))

		if sprite.Ticker.Ticker == nil {
			sprite.Ticker.Ticker = time.NewTicker(100 * time.Millisecond)
		}

		sprite.Src.Y = float32(sheetRow)
		sprite.Dest.X = float32(w.Player.Hitbox.Position.X)
		sprite.Dest.Y = float32(w.Player.Hitbox.Position.Y)
		sprite.FlipX = w.Player.Direction.X != components.DirectionNone && w.Player.Direction.X < 0
		sprite.Src.Width = float32(w.Player.Hitbox.Width)
		sprite.Src.Width /= sprite.Scale
		if sprite.FlipX {
			sprite.Src.Width *= -1
		}
		sprite.Src.Height = float32(w.Player.Hitbox.Height)
		sprite.Src.Height /= sprite.Scale

		coroutines.Sprite.Tick(
			&w.Player.Id,
			constants.Keys.Idle,
			&w.Player.SpriteMap,
			variants,
			sprite.Ticker.Ticker,
		).CallOnce()

		w.Player.SpriteMap[constants.Keys.Idle] = sprite
	})
}

package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) UpdatePlayerSprite() system.System {
	return func(w *world.World, m *meta.Meta) {
		w.Player.SpriteKey = components.SpriteKey(constants.Keys.PlayerIdle)
		if w.Player.Direction.X == components.DirectionLeft || w.Player.Direction.X == components.DirectionRight {
			w.Player.SpriteKey = components.SpriteKey(constants.Keys.PlayerMoveForward)
		}
		if w.Player.Direction.Y == components.DirectionDown {
			w.Player.SpriteKey = components.SpriteKey(constants.Keys.PlayerMoveDown)
		}
		if w.Player.Direction.Y == components.DirectionUp {
			w.Player.SpriteKey = components.SpriteKey(constants.Keys.PlayerMoveUp)
		}
		archetypes.Sprite.Animate(
			m,
			&w.Player.Id,
			&w.Player.SpriteKey,
			&w.Player.Hitbox,
			&w.Player.Direction,
			&w.Player.SpriteMap,
		)
	}
}

func (_ *sprite) UpdateEnemySprite(enemy *entities.Enemy) system.System {
	return func(w *world.World, m *meta.Meta) {
		enemy.SpriteKey = components.SpriteKey(constants.Keys.AntiPeanutIdle)
		if enemy.Direction.X == components.DirectionLeft || enemy.Direction.X == components.DirectionRight {
			enemy.SpriteKey = components.SpriteKey(constants.Keys.AntiPeanutMoveForward)
		}
		if enemy.Direction.Y == components.DirectionDown {
			enemy.SpriteKey = components.SpriteKey(constants.Keys.AntiPeanutMoveDown)
		}
		if enemy.Direction.Y == components.DirectionUp {
			enemy.SpriteKey = components.SpriteKey(constants.Keys.AntiPeanutMoveUp)
		}
		archetypes.Sprite.Animate(
			m,
			&enemy.Id,
			&enemy.SpriteKey,
			&enemy.Hitbox,
			&enemy.Direction,
			&enemy.SpriteMap,
		)
	}
}

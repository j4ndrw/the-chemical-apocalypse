package sprite_systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type update struct{}

var Update = update{}

func (_ *update) UpdatePlayerSprite() system.System {
	return func(w *world.World, m *meta.Meta) {
		w.Player.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.PlayerIdle)
		if w.Player.Hitbox.Direction.X == direction.Left || w.Player.Hitbox.Direction.X == direction.Right {
			w.Player.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.PlayerMoveForward)
		}
		if w.Player.Hitbox.Direction.Y == direction.Down {
			w.Player.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.PlayerMoveDown)
		}
		if w.Player.Hitbox.Direction.Y == direction.Up {
			w.Player.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.PlayerMoveUp)
		}
		sprite.Archetype.Animate(
			m,
			&w.Player.Id,
			&w.Player.SpriteKey,
			&w.Player.Hitbox,
			&w.Player.Hitbox.Direction,
			&w.Player.SpriteMap,
		)
	}
}

func (_ *update) UpdateEnemySprite(enemy *enemy.EnemyEntity) system.System {
	return func(w *world.World, m *meta.Meta) {
		enemy.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.AntiPeanutIdle)
		if enemy.Hitbox.Direction.X == direction.Left || enemy.Hitbox.Direction.X == direction.Right {
			enemy.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.AntiPeanutMoveForward)
		}
		if enemy.Hitbox.Direction.Y == direction.Down {
			enemy.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.AntiPeanutMoveDown)
		}
		if enemy.Hitbox.Direction.Y == direction.Up {
			enemy.SpriteKey = sprite.SpriteKeyComponent(constants.Keys.AntiPeanutMoveUp)
		}
		sprite.Archetype.Animate(
			m,
			&enemy.Id,
			&enemy.SpriteKey,
			&enemy.Hitbox,
			&enemy.Hitbox.Direction,
			&enemy.SpriteMap,
		)
	}
}

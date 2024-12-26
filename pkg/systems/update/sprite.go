package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) UpdatePlayerSprite() system.System {
	return func(w *world.World, m *meta.Meta) {
		w.Player.SpriteKey = components.SpriteKey(func() string {
			if !w.Player.Moving {
				return constants.Keys.PlayerIdle
			}
			if w.Player.Direction.Y == components.DirectionDown {
				return constants.Keys.PlayerMoveDown
			}
			if w.Player.Direction.Y == components.DirectionUp {
				return constants.Keys.PlayerMoveUp
			}
			// DEBUG
			return constants.Keys.PlayerIdle
		}())
		archetypes.Sprite.AnimateSprite(
			m,
			&w.Player.Id,
			&w.Player.SpriteKey,
			&w.Player.Hitbox,
			&w.Player.Direction,
			&w.Player.SpriteMap,
		)
	}
}

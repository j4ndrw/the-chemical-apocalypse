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

func (_ *sprite) SetupPlayerSprites() system.System {
	return func(w *world.World, m *meta.Meta) {
		makeDefault := func() *components.Sprite {
			return archetypes.Sprite.New(
				m.SpriteAtlas,
				&w.Player.Hitbox,
				8,
			)
		}
		w.Player.SpriteMap = components.SpriteMap{
			Map: map[string]*components.Sprite{
				constants.Keys.PlayerIdle:     makeDefault(),
				constants.Keys.PlayerMoveDown: makeDefault(),
				constants.Keys.PlayerMoveUp: makeDefault(),
			},
		}
	}
}

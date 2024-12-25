package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) DrawPlayerIdleSprite() system.System {
	return func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		archetypes.Sprite.DrawSprite(w.Player.SpriteMap[constants.Keys.Idle])
	}
}

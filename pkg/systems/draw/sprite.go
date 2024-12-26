package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) DrawSprite(sm *components.SpriteMap, k *components.SpriteKey) system.System {
	return func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		sm.Lock()
		archetypes.Sprite.DrawSprite(sm.Map[string(*k)])
		sm.Unlock()
	}
}

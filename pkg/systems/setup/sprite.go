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

func (_ *sprite) Setup(sm *components.SpriteMap, hitbox *components.Hitbox, ks ...string) system.System {
	return func(w *world.World, m *meta.Meta) {
		makeDefault := func() *components.Sprite {
			return archetypes.Sprite.New(
				m.SpriteAtlas,
				hitbox,
				8,
			)
		}
		sm.Map = map[string]*components.Sprite{}
		for _, k := range ks {
			sm.Map[k] = makeDefault()
		}
	}
}

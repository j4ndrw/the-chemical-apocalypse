package sprite_systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type setup struct{}

var Setup = setup{}

func (_ *setup) Setup(sm *sprite.SpriteMapComponent, h *hitbox.HitboxComponent, ks ...string) system.System {
	return func(w *world.World, m *meta.Meta) {
		makeDefault := func() *sprite.SpriteComponent {
			return sprite.Archetype.New(
				m.SpriteAtlas,
				h,
				8,
			)
		}
		sm.Map = map[string]*sprite.SpriteComponent{}
		for _, k := range ks {
			sm.Map[k] = makeDefault()
		}
	}
}

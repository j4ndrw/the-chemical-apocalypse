package sprite_systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type draw struct{}

var Draw = draw{}

func (_ *draw) Draw(sm *sprite.SpriteMapComponent, k *sprite.SpriteKeyComponent) system.System {
	return func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		sm.Lock()
		sprite.Archetype.Draw(sm.Map[string(*k)])
		sm.Unlock()
	}
}

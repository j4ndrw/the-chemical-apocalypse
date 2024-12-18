package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type worldmode struct{}

var WorldMode = worldmode{}

func (_ *worldmode) Init() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		w.PrevMode = world.WorldModeNil
		w.CurrentMode = world.WorldModeTitleScreen
	})
}

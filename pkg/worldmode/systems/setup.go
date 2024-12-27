package worldmode_systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type setup struct{}

var Setup = setup{}

func (_ *setup) Init() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		w.PrevMode = world.WorldModeNil
		w.CurrentMode = world.WorldModeTitleScreen
	})
}

package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type input struct{}

var Input = input{}

func (_ *input) HandleJourneyStart() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeTitleScreen { return }
		if rl.IsKeyPressed(rl.KeyEnter) {
			w.PrevMode = w.CurrentMode
			w.CurrentMode = world.WorldModeExploration
		}
	})
}
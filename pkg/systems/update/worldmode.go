package systems

import (
	// rl "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type worldmode struct{}

var WorldMode = worldmode{}

func (_ *worldmode) HandleCombatMode() *system.System {
	utils.Assert(false, "HandleCombatMode not yet implemented")
	return nil
}

func (_ *worldmode) HandleExplorationMode() *system.System {
	utils.Assert(false, "HandleExplorationMode not yet implemented")
	return nil
}

func (_ *worldmode) HandlePauseMode() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		isPauseKey := rl.IsKeyPressed(rl.KeyEscape)
		if (!isPauseKey) { return }

		w.PrevMode = w.CurrentMode
		w.CurrentMode = world.WorldModePause
	})
}

func (_ *worldmode) HandleTitleScreenMode() *system.System {
	utils.Assert(false, "HandleTitleScreenMode not yet implemented")
	return nil
}

func (_ *worldmode) HandleInventoryMode() *system.System {
	utils.Assert(false, "HandleInventoryMode not yet implemented")
	return nil
}

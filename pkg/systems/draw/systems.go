package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

var Systems system.SystemSlice = *system.
	Slice().
	Register(Meta.UpdateWindowSize()).
	Register(Meta.UpdateWindowOnResize()).
	Register(Renderer.Clear()).
	Register(Renderer.DrawTitleScreen()).
	Register(Renderer.DrawHitboxesInExplorationMode()).
	Register(Meta.UpdateDeltaTime())

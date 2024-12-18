package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

var Systems system.SystemSlice = *system.
	Slice().
	Register(Renderer.Clear()).
	Register(Renderer.Hitboxes()).
	Register(Meta.UpdateDeltaTime())

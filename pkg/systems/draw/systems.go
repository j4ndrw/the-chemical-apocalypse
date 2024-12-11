package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

func Systems() []system.System {
	renderer := Renderer()
	return *system.Slice().
		Register(&renderer.Clear).
		Register(&renderer.Player).
		Register(&renderer.Enemy)
}

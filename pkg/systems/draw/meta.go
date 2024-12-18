package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type metasystem struct{}

var Meta = metasystem{}

func (_ *metasystem) UpdateWindowSize() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		width := int32(rl.GetScreenWidth())
		height := int32(rl.GetScreenHeight())
		if m.Window.Width != width {
			m.Window.Width = width
		}
		if m.Window.Height != height {
			m.Window.Height = height
		}
	})
}

func (_ *metasystem) UpdateDeltaTime() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.DeltaTime = rl.GetFrameTime()
	})
}

func (_ *metasystem) UpdateWindowOnResize() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if !rl.IsWindowResized() { return }

		m.Window.Width = int32(rl.GetScreenWidth())
		m.Window.Height = int32(rl.GetScreenHeight())
	})
}

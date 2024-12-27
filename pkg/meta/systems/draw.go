package meta_systems

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type DrawSystem struct{}

var Draw = DrawSystem{}

func (_ *DrawSystem) UpdateWindowSize() system.System {
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

func (_ *DrawSystem) UpdateDeltaTime() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.DeltaTime = rl.GetFrameTime()
	})
}

func (_ *DrawSystem) UpdateFrame() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.Frame++
		if m.Frame < 0 || m.Frame >= math.MaxInt32 {
			m.Frame = 0
		}
	})
}

func (_ *DrawSystem) UpdateWindowOnResize() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if !rl.IsWindowResized() {
			return
		}

		m.Window.Width = int32(rl.GetScreenWidth())
		m.Window.Height = int32(rl.GetScreenHeight())
	})
}

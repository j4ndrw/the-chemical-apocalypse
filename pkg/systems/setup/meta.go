package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type metasystem struct{}

var Meta = metasystem{}

func (_ *metasystem) SetConfigFlags(flags uint32) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		rl.SetConfigFlags(flags)
	})
}

func (_ *metasystem) InitWindow() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		rl.InitWindow(1, 1, m.Window.Title)

		monitor := rl.GetCurrentMonitor()
		m.Window.Width = int32(rl.GetMonitorWidth(monitor))
		m.Window.Height = int32(rl.GetMonitorHeight(monitor))

		rl.SetWindowSize(int(m.Window.Width), int(m.Window.Height))
	})
}

func (_ *metasystem) SetTargetFPS(fps int32) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.TargetFPS = fps
		rl.SetTargetFPS(m.TargetFPS)
	})
}

func (_ *metasystem) LoadFont(path string) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.Font = rl.LoadFont(path)
	})
}

package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type metasystem struct{}

var Meta metasystem = metasystem{}

func (_ *metasystem) UpdateDeltaTime() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		m.DeltaTime = rl.GetFrameTime()
	})
}

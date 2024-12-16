package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy enemy = enemy{}

func (_ *enemy) PlaceEnemyInCenter() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		w.Enemy.Position.X = m.Window.Width/2 - 1
		w.Enemy.Position.Y = m.Window.Height/2 - 1
	})
}

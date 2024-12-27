package position

import (
	"math/rand"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) RandomPointOnMap(w *meta.Window) PositionComponent {
	return PositionComponent{
		X: int32(rand.Intn(int(w.Width-1)) + 1),
		Y: int32(rand.Intn(int(w.Height-1)) + 1),
	}
}

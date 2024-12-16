package archetypes

import (
	"math/rand"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type vector2 struct{}

var Vector2 vector2 = vector2{}

func (_ *vector2) RandomPointOnMap(w *meta.Window) components.Vector2 {
	return components.Vector2{
		X: int32(rand.Intn(int(w.Width-1)) + 1),
		Y: int32(rand.Intn(int(w.Height-1)) + 1),
	}
}

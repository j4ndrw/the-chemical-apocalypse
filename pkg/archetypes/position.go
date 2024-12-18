package archetypes

import (
	"math/rand"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type position struct{}

var Position = position{}

func (_ *position) RandomPointOnMap(w *meta.Window) components.Position {
	return components.Position{
		X: int32(rand.Intn(int(w.Width-1)) + 1),
		Y: int32(rand.Intn(int(w.Height-1)) + 1),
	}
}

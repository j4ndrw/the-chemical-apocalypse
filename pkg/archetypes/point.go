package archetypes

import (
	"math/rand"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type point struct{}

var Point point = point{}

func (_ *point) RandomPointOnMap(w *meta.Window) components.Point {
	return components.Point{
		X: int32(rand.Intn(int(w.Width-1)) + 1),
		Y: int32(rand.Intn(int(w.Height-1)) + 1),
	}
}

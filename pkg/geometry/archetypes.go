package geometry

import (
	"math"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) AngleFrom(x, y float32) float32 {
	if x == 0 {
		if y > 0 {
			return 90
		} else if y < 0 {
			return -90
		}
		return 0
	}
	angle := float32(math.Atan2(float64(y), float64(x)) * (180 / math.Pi))
	return angle
}

func (_ *archetype) Center(hitbox *hitbox.HitboxComponent) (float64, float64) {
	centerX := float64(hitbox.Left()+hitbox.Right()) / 2
	centerY := float64(hitbox.Top()+hitbox.Bottom()) / 2
	utils.Assert(centerX >= 0 && centerY >= 0, "Center of hitbox cannot be defined outside bounds")
	return centerX, centerY
}

func (_ *archetype) Delta(a *position.PositionComponent, b *position.PositionComponent) (float64, float64) {
	return Archetype.DeltaEx(a.FloatX(), a.FloatY(), b.FloatX(), b.FloatY())
}

func (_ *archetype) DeltaEx(ax, ay, bx, by float64) (float64, float64) {
	dx := ax - bx
	dy := ay - by
	return dx, dy
}

func (_ *archetype) SquaredDistance(dx, dy float64) float64 {
	return dx*dx + dy*dy
}

func (_ *archetype) Distance(dx, dy float64) float64 {
	return math.Sqrt(Archetype.SquaredDistance(dx, dy))
}

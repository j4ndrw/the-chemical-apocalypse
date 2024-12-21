package archetypes

import (
	"math"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type vision struct{}

var Vision = vision{}

func (_ *vision) CenterAngle(
	this *components.Hitbox,
	target *components.Position,
	v *components.Vision,
) (float32, float32, float32) {
	centerX := float64(this.Left()+this.Right()) / 2
	centerY := float64(this.Top()+this.Bottom()) / 2
	return Vision.CenterAngleEx(centerX, centerY, target, v)
}

func (_ *vision) CenterAngleEx(
	originX float64,
	originY float64,
	target *components.Position,
	v *components.Vision,
) (float32, float32, float32) {
	dx := float64(target.X - int32(originX))
	dy := float64(target.Y - int32(originY))
	return Vision.CenterAngleEx2(originX, originY, dx, dy, target, v)
}


func (_ *vision) CenterAngleEx2(
	originX float64,
	originY float64,
	dx float64,
	dy float64,
	target *components.Position,
	v *components.Vision,
) (float32, float32, float32) {
	angleToTarget := float32(math.Atan2(dy, dx) * (180 / math.Pi))

	half := v.Angle / 2
	start := angleToTarget - half
	end := angleToTarget + half

	return start, end, angleToTarget
}

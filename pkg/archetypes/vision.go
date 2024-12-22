package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type vision struct{}

var Vision = vision{}

func (_ *vision) CenterAngle(
	this *components.Hitbox,
	target *components.Position,
	v *components.Vision,
) (float32, float32, float32) {
	centerX, centerY := Geometry.Center(this)
	return Vision.CenterAngleEx(centerX, centerY, target, v)
}

func (_ *vision) CenterAngleEx(
	originX float64,
	originY float64,
	target *components.Position,
	v *components.Vision,
) (float32, float32, float32) {
	dx, dy := Geometry.DeltaEx(
		target.FloatX(),
		target.FloatY(),
		originX,
		originY,
	)
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
	angleToTarget := Geometry.AngleFrom(float32(dx), float32(dy))

	half := v.Angle / 2
	start := angleToTarget - half
	end := angleToTarget + half

	return start, end, angleToTarget
}

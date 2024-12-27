package vision

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/geometry"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) CenterAngle(
	this *hitbox.HitboxComponent,
	target *position.PositionComponent,
	vision *VisionComponent,
) (float32, float32, float32) {
	centerX, centerY := geometry.Archetype.Center(this)
	return Archetype.CenterAngleEx(centerX, centerY, target, vision)
}

func (_ *archetype) CenterAngleEx(
	originX float64,
	originY float64,
	target *position.PositionComponent,
	vision *VisionComponent,
) (float32, float32, float32) {
	dx, dy := geometry.Archetype.DeltaEx(
		target.FloatX(),
		target.FloatY(),
		originX,
		originY,
	)
	return Archetype.CenterAngleEx2(originX, originY, dx, dy, target, vision)
}

func (_ *archetype) CenterAngleEx2(
	originX float64,
	originY float64,
	dx float64,
	dy float64,
	target *position.PositionComponent,
	vision *VisionComponent,
) (float32, float32, float32) {
	angleToTarget := geometry.Archetype.AngleFrom(float32(dx), float32(dy))

	half := vision.Angle / 2
	start := angleToTarget - half
	end := angleToTarget + half

	return start, end, angleToTarget
}

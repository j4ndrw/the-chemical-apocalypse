package archetypes

import (
	"math"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type aggro struct{}

var Aggro = aggro{}

func (_ *aggro) EnterAggro(aggro *components.Aggro) {
	aggro.Aggro = true
}
func (_ *aggro) LeaveAggro(aggro *components.Aggro) {
	aggro.Aggro = false
}

func (_ *aggro) IsWithinAggroRange(
	target *components.Hitbox,
	chaser *components.Hitbox,
	aggro *components.Aggro,
) bool {
	centerTargetX, centerTargetY := Geometry.Center(target)
	centerChaserX, centerChaserY := Geometry.Center(chaser)
	dx, dy := Geometry.DeltaEx(
		centerTargetX,
		centerTargetY,
		centerChaserX,
		centerChaserY,
	)

	sqAggroRadius := aggro.Radius * aggro.Radius

	if int32(Geometry.SquaredDistance(dx, dy)) > sqAggroRadius {
		return false
	}

	_, _, angleToTarget := Vision.CenterAngleEx2(
		chaser.Position.FloatX(),
		chaser.Position.FloatY(),
		float64(dx),
		float64(dy),
		&target.Position,
		&aggro.Vision,
	)

	chaserDirectionAngle := Geometry.AngleFrom(
		float32(chaser.Direction.X),
		float32(chaser.Direction.Y),
	)

	diff := angleToTarget - chaserDirectionAngle
	diff -= utils.BoolToNumber[float32](diff > 180) * 360
	diff += utils.BoolToNumber[float32](diff < -180) * 360

	isWithinFOV := math.Abs(float64(diff)) <= float64(aggro.Vision.Angle/2)
	return isWithinFOV
}

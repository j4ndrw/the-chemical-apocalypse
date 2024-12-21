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

func (_ *aggro) IsWithinAggroRange(target *components.Hitbox, chaser *components.Hitbox, aggro *components.Aggro) bool {
	dx := target.Position.X - chaser.Position.X
	dy := target.Position.Y - chaser.Position.Y

	sqDistance := dx*dx + dy*dy
	sqAggroRadius := aggro.Radius * aggro.Radius

	if sqDistance > sqAggroRadius {
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

	chaserDirectionAngle := float32(math.Atan2(
		float64(chaser.Direction.Y),
		float64(chaser.Direction.X),
	) * (180 / math.Pi))

	angleDifference := angleToTarget - chaserDirectionAngle
	angleDifference -= utils.BoolToNumber[float32](angleDifference > 180) * 360
	angleDifference += utils.BoolToNumber[float32](angleDifference < -180) * 360

	isWithinFOV := math.Abs(float64(angleDifference)) <= float64(aggro.Vision.Angle/2)
	return isWithinFOV
}

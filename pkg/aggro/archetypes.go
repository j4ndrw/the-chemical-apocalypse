package aggro

import (
	"math"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/geometry"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/vision"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) EnterAggro(aggro *AggroComponent) {
	aggro.Aggro = true
}
func (_ *archetype) LeaveAggro(aggro *AggroComponent) {
	aggro.Aggro = false
}

func (_ *archetype) IsWithinAggroRange(
	target *hitbox.HitboxComponent,
	chaser *hitbox.HitboxComponent,
	aggro *AggroComponent,
) bool {
	centerTargetX, centerTargetY := geometry.Archetype.Center(target)
	centerChaserX, centerChaserY := geometry.Archetype.Center(chaser)
	dx, dy := geometry.Archetype.DeltaEx(
		centerTargetX,
		centerTargetY,
		centerChaserX,
		centerChaserY,
	)

	sqAggroRadius := aggro.Radius * aggro.Radius

	if int32(geometry.Archetype.SquaredDistance(dx, dy)) > sqAggroRadius {
		return false
	}

	_, _, angleToTarget := vision.Archetype.CenterAngleEx2(
		chaser.Position.FloatX(),
		chaser.Position.FloatY(),
		float64(dx),
		float64(dy),
		&target.Position,
		&aggro.Vision,
	)

	chaserDirectionAngle := geometry.Archetype.AngleFrom(
		float32(chaser.Direction.X),
		float32(chaser.Direction.Y),
	)

	diff := angleToTarget - chaserDirectionAngle
	diff -= utils.BoolToNumber[float32](diff > 180) * 360
	diff += utils.BoolToNumber[float32](diff < -180) * 360

	isWithinFOV := math.Abs(float64(diff)) <= float64(aggro.Vision.Angle/2)
	return isWithinFOV
}

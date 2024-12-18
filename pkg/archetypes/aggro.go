package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type aggro struct{}

var Aggro = aggro{}

func (_ *aggro) EnterAggro(aggro *components.Aggro) {
	aggro.Aggro = true
}
func (_ *aggro) LeaveAggro(aggro *components.Aggro) {
	aggro.Aggro = false
}

func (_ *aggro) IsWithinAggroRange(target *components.Position, chaser *components.Position, aggro *components.Aggro) bool {
	dx := target.X - chaser.X
	dy := target.Y - chaser.Y
	sqDistance := dx*dx + dy*dy
	sqAggroRadius := aggro.Radius * aggro.Radius
	return sqDistance <= sqAggroRadius
}

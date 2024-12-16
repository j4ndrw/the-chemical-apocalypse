package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type mobmovement struct{}

var MobMovement mobmovement = mobmovement{}

func (_ *mobmovement) Chase(
	chaser *components.Position,
	target *components.Point,
	speed components.Speed,
	collisions ...bool,
) {
	neighbor := PathFinding.ClosestNeighbor(
		&chaser.Point,
		target,
		func(position, direction *components.Point) *components.Point {
			return &components.Point{X: position.X + direction.X*int32(speed), Y: position.Y + direction.Y*int32(speed)}
		},
		func(position *components.Point) bool {
			for _, collision := range collisions {
				if collision {
					return false
				}
			}
			return true
		},
	)

	if neighbor != nil {
		chaser.Point = *neighbor.Position
	}
}

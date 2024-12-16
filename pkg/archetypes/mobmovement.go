package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type mobmovement struct{}

var MobMovement mobmovement = mobmovement{}

func (_ *mobmovement) Chase(
	chaser *components.Position,
	target *components.Vector2,
	speed components.Speed,
	collisions ...bool,
) {
	neighbor := PathFinding.ClosestNeighbor(
		&chaser.Vector2,
		target,
		func(position, direction *components.Vector2) *components.Vector2 {
			return &components.Vector2{X: position.X + direction.X*int32(speed), Y: position.Y + direction.Y*int32(speed)}
		},
		func(position *components.Vector2) bool {
			for _, collision := range collisions {
				if collision {
					return false
				}
			}
			return true
		},
	)

	if neighbor != nil {
		chaser.Vector2 = *neighbor.Position
	}
}

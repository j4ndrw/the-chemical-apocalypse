package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type mobmovement struct{}

var MobMovement = mobmovement{}

func (_ *mobmovement) Chase(
	chaser *components.Hitbox,
	target *components.Position,
	speed components.Speed,
	collisions ...bool,
) {
	neighbor := PathFinding.ClosestNeighbor(
		&chaser.Position,
		target,
		func(position, direction *components.Position) *components.Position {
			return &components.Position{X: position.X + direction.X*int32(speed), Y: position.Y + direction.Y*int32(speed)}
		},
		func(position *components.Position) bool {
			for _, collision := range collisions {
				if collision {
					return false
				}
			}
			return true
		},
	)

	if neighbor != nil {
		chaser.Position = *neighbor.Position
	}
}

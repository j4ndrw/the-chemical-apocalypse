package archetypes

import (
	"math"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type mobmovement struct{}

var MobMovement = mobmovement{}

func (_ *mobmovement) NearestNeighborChase(
	chaser *components.Hitbox,
	target *components.Position,
	speed components.Speed,
	deltaTime float64,
	collisions ...bool,
) {
	neighbor := PathFinding.ClosestNeighbor(
		&chaser.Position,
		target,
		func(position, direction *components.Position) *components.Position {
			return &components.Position{
				X: int32(position.FloatX() + direction.FloatX()*speed.AsFloat()*deltaTime),
				Y: int32(position.FloatY() + direction.FloatY()*speed.AsFloat()*deltaTime),
			}
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

func (_ *mobmovement) NaiveChase(
	chaser *components.Hitbox,
	target *components.Position,
	speed components.Speed,
	deltaTime float64,
	collisions ...bool,
) {
	for _, collision := range collisions {
		if collision {
			return
		}
	}

	dx := target.FloatX() - chaser.FloatX()
	dy := target.FloatY() - chaser.FloatY()

	distance := math.Sqrt(dx*dx + dy*dy)
	if distance > 0 {
		dx /= distance
		dy /= distance
	}

	chaser.Position.X += int32(dx * speed.AsFloat())
	chaser.Position.Y += int32(dy * speed.AsFloat())
}

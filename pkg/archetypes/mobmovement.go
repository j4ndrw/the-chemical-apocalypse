package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

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

	if neighbor == nil {
		return
	}

	Direction.Update(
		&chaser.Direction,
		&chaser.Position,
		neighbor.Position.X,
		neighbor.Position.Y,
	)
	chaser.Position = *neighbor.Position
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

	dx, dy := Geometry.Delta(target, &chaser.Position)
	distance := Geometry.Distance(dx, dy)

	if distance > 0 {
		dx /= distance
		dy /= distance
	}

	newX := int32(chaser.Position.FloatX() + dx*(speed.AsFloat() * deltaTime))
	newY := int32(chaser.Position.FloatY() + dy*(speed.AsFloat() * deltaTime))

	Direction.Update(
		&chaser.Direction,
		&chaser.Position,
		newX,
		newY,
	)

	chaser.Position.X = newX
	chaser.Position.Y = newY
}

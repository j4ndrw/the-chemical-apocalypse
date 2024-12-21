package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type direction struct{}

var Direction = direction{}

func (_ *direction) Update(
	direction *components.Direction,
	position *components.Position,
	newX int32,
	newY int32,
) {
	if newX > position.X {
		direction.X = components.DirectionRight
	}
	if newX < position.X {
		direction.X = components.DirectionLeft
	}
	if newY > position.Y {
		direction.Y = components.DirectionDown
	}
	if newY < position.Y {
		direction.Y = components.DirectionUp
	}
}

package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type direction struct{}

var Direction = direction{}

func (_ *direction) Update(
	direction *components.Direction,
	position *components.Position,
	newX int32,
	newY int32,
) {
	direction.X = (utils.BoolToNumber[int](newX > position.X) * components.DirectionRight) +
		(utils.BoolToNumber[int](newX < position.X) * components.DirectionLeft)
	direction.Y = (utils.BoolToNumber[int](newY > position.Y) * components.DirectionDown) +
		(utils.BoolToNumber[int](newY < position.Y) * components.DirectionUp)
}

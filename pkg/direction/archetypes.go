package direction

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) Update(
	direction *DirectionComponent,
	position *position.PositionComponent,
	newX int32,
	newY int32,
) {
	direction.X = (utils.BoolToNumber[int](newX > position.X) * Right) +
		(utils.BoolToNumber[int](newX < position.X) * Left)
	direction.Y = (utils.BoolToNumber[int](newY > position.Y) * Down) +
		(utils.BoolToNumber[int](newY < position.Y) * Up)
}

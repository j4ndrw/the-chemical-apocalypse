package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type collidable struct{}

var Collidable collidable = collidable{}

func (_ *collidable) IsColliding(this *components.Position, other *components.Position) bool {
	return this.Left() < other.Right() && this.Right() > other.Left() && this.Top() < other.Bottom() && this.Bottom() > other.Top()
}

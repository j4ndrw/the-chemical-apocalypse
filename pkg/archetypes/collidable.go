package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type collidable struct{}

var Collidable = collidable{}

func (_ *collidable) IsColliding(this *components.Hitbox, others ...*components.Hitbox) bool {
	for _, other := range others {
		colliding := this.Left() < other.Right() && this.Right() > other.Left() && this.Top() < other.Bottom() && this.Bottom() > other.Top()
		if colliding {
			return true
		}
	}
	return false
}

package collidable

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) IsColliding(
	this *hitbox.HitboxComponent,
	others ...*hitbox.HitboxComponent,
) bool {
	for _, other := range others {
		colliding :=
			this.Left() < other.Right() &&
				this.Right() > other.Left() &&
				this.Top() < other.Bottom() &&
				this.Bottom() > other.Top()
		if colliding {
			return true
		}
	}
	return false
}

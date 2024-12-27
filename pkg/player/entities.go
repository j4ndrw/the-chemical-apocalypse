package player

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/movement"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/speed"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
)

type PlayerEntity struct {
	Id        id.IdComponent
	Hitbox    hitbox.HitboxComponent
	Speed     speed.SpeedComponent
	Moving    movement.MovingComponent
	SpriteMap sprite.SpriteMapComponent
	SpriteKey sprite.SpriteKeyComponent
}

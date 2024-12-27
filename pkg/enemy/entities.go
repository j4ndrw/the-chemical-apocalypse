package enemy

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/aggro"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/movement"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/roam"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/speed"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
)

type EnemyEntity struct {
	Id        id.IdComponent
	Hitbox    hitbox.HitboxComponent
	MinSpeed  speed.SpeedComponent
	MaxSpeed  speed.SpeedComponent
	Aggro     aggro.AggroComponent
	Roam      roam.RoamComponent
	Moving    movement.MovingComponent
	SpriteMap sprite.SpriteMapComponent
	SpriteKey sprite.SpriteKeyComponent
}

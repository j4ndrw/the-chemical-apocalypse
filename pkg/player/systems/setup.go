package player_systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/player"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type setup struct{}

var Setup = setup{}

func (_ *setup) CreatePlayer() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		id := id.IdComponent(uuid.New().String())
		w.Player = player.PlayerEntity{
			Id: id,
			Hitbox: hitbox.HitboxComponent{
				Position: position.PositionComponent{X: 0, Y: 0},
				Width:    256,
				Height:   256,
				Color:    rl.Color{0, 0xFF, 0, 0xFF},
				Hidden:   true,
				PaddingPercentage:  0.5,
			},
			SpriteKey: sprite.SpriteKeyComponent(constants.Keys.PlayerIdle),
			Speed:     750,
		}
	})
}

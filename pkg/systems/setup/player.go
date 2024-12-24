package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/google/uuid"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type player struct{}

var Player = player{}

func (_ *player) CreatePlayer() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		id := components.Id(uuid.New().String())
		w.Player = entities.Player{
			Id: id,
			Hitbox: components.Hitbox{
				Position: components.Position{X: 0, Y: 0},
				Width:    256,
				Height:   256,
				Color:    rl.Color{0, 0xFF, 0, 0xFF},
				Hidden:   true,
			},
			Speed: 750,
		}
	})
}

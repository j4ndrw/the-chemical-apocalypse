package systems

import (
	"log/slog"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy = enemy{}

func (_ *enemy) CreateEnemies(
	howMany uint,
	width int32,
	height int32,
	minSpeed components.Speed,
	maxSpeed components.Speed,
	aggroRadius int32,
	maxRoamDuration uint,
) *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for i := 0; i < int(howMany); i++ {
			roamDuration := float32(rand.Intn(int(maxRoamDuration) + 1))
			id := archetypes.Id.Create()
			w.Enemies[id] = &entities.Enemy{
				Id: id,
				Hitbox: components.Hitbox{
					Position: components.Position{X: 0, Y: 0},
					Width:    width,
					Height:   height,
					Color:    rl.Color{0xFF, 0, 0, 0xFF},
				},
				MinSpeed: minSpeed,
				MaxSpeed: maxSpeed,
				Aggro:    components.Aggro{Aggro: false, Radius: aggroRadius},
				Roam:     components.Roam{Duration: roamDuration},
			}
		}
	})
}

func (_ *enemy) PlaceEnemiesInCenter() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		slog.Info("Meta", "windowwidth", m.Window.Width, "windowheight", m.Window.Height)
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				enemy.Position.X = int32((rand.Intn(int(m.Window.Width))+int(m.Window.Width/3))/2 - 1)
				enemy.Position.Y = int32((rand.Intn(int(m.Window.Height))+int(m.Window.Height/3))/2 - 1)
			}(enemy)
		}
	})
}

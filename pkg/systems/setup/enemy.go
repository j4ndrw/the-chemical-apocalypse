package systems

import (
	"math/rand"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy enemy = enemy{}

func (_ *enemy) CreateEnemies() *system.System {
	width := int32(40)
	height := int32(40)
	minSpeed := components.Speed(1)
	maxSpeed := components.Speed(3)
	radius := 100 * int32(maxSpeed)

	return system.Create(func(w *world.World, m *meta.Meta) {
		for i := 0; i < 9; i++ {
			roamDuration := float32(rand.Intn(15) + 1)
			w.Enemies = append(w.Enemies,
				&entities.Enemy{
					Position: components.Position{
						Vector2: components.Vector2{X: 0, Y: 0},
						Width:   width,
						Height:  height,
					},
					MinSpeed: minSpeed,
					MaxSpeed: maxSpeed,
					Aggro: components.Aggro{
						Aggro:  false,
						Radius: radius,
					},
					Roam: components.Roam{Duration: roamDuration},
				})
		}
	})
}

func (_ *enemy) PlaceEnemyInCenter() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				enemy.Position.X = int32((rand.Intn(int(m.Window.Width))+int(m.Window.Width/3))/2 - 1)
				enemy.Position.Y = int32((rand.Intn(int(m.Window.Height))+int(m.Window.Height/3))/2 - 1)
			}(enemy)
		}
	})
}

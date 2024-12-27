package enemy_systems

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/aggro"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/roam"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/speed"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/vision"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type setup struct{}

var Setup = setup{}

func (_ *setup) CreateEnemies(
	howMany uint,
	width int32,
	height int32,
	minSpeed speed.SpeedComponent,
	maxSpeed speed.SpeedComponent,
	aggroRadius int32,
	aggroVisionAngle float32,
	maxRoamDurationMs time.Duration,
) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for i := 0; i < int(howMany); i++ {
			id := id.Archetype.Create()
			w.Enemies[id] = &enemy.EnemyEntity{
				Id: id,
				Hitbox: hitbox.HitboxComponent{
					Position:          position.PositionComponent{X: 0, Y: 0},
					Width:             width,
					Height:            height,
					Color:             rl.Color{0xFF, 0, 0, 0xFF},
					Hidden:            true,
					PaddingPercentage: 0.5,
				},
				MinSpeed: minSpeed,
				MaxSpeed: maxSpeed,
				Aggro: aggro.AggroComponent{
					Aggro:  false,
					Radius: aggroRadius,
					Vision: vision.VisionComponent{
						Angle: aggroVisionAngle,
					},
				},
				Roam:      roam.RoamComponent{MaxDuration: maxRoamDurationMs},
				SpriteKey: sprite.SpriteKeyComponent(constants.Keys.AntiPeanutIdle),
			}
		}
	})
}

func (_ *setup) PlaceEnemyInCenter(enemy *enemy.EnemyEntity) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		enemy.Hitbox.Position.X = int32((rand.Intn(int(m.Window.Width))+int(m.Window.Width/3))/2 - 1)
		enemy.Hitbox.Position.Y = int32((rand.Intn(int(m.Window.Height))+int(m.Window.Height/3))/2 - 1)
	})
}

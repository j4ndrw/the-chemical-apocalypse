package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type renderer struct{}

var Renderer renderer = renderer{}

func (_ *renderer) Clear() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		rl.ClearScreenBuffers()
	})
}

func (_ *renderer) Player() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		w.Player.Position.Bound = &components.Bound{
			Left:   1,
			Top:    1,
			Right:  m.Window.Width - w.Player.Position.Width - 1,
			Bottom: m.Window.Height - w.Player.Position.Height - 1,
		}
		rl.DrawRectangle(
			w.Player.Position.X,
			w.Player.Position.Y,
			w.Player.Position.Width,
			w.Player.Position.Height,
			rl.Color{
				0,
				0xFF,
				0,
				0xFF,
			},
		)
	})
}

func (_ *renderer) Enemies() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				enemy.Position.Bound = &components.Bound{
					Left:   1,
					Top:    1,
					Right:  m.Window.Width - enemy.Position.Width - 1,
					Bottom: m.Window.Height - enemy.Position.Height - 1,
				}

				// TODO(j4ndrw): When you get to adding sprites, you'll have to change this
				// so that it renders the sprite instead of the rectangle
				rl.DrawRectangle(
					enemy.Position.X,
					enemy.Position.Y,
					enemy.Position.Width,
					enemy.Position.Height,
					rl.Color{
						0xFF,
						0,
						0,
						0xFF,
					},
				)
			}(enemy)
		}
	})
}

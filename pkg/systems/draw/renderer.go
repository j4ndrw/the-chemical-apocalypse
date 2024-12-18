package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
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

func (_ *renderer) PlayerHitbox() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		w.Player.Hitbox.Bound = &components.Bound{
			Left:   1,
			Top:    1,
			Right:  m.Window.Width - w.Player.Hitbox.Width - 1,
			Bottom: m.Window.Height - w.Player.Hitbox.Height - 1,
		}
		rl.DrawRectangle(
			w.Player.Position.X,
			w.Player.Position.Y,
			w.Player.Hitbox.Width,
			w.Player.Hitbox.Height,
			rl.Color{
				0,
				0xFF,
				0,
				0xFF,
			},
		)
	})
}

func (_ *renderer) Hitboxes() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		drawHitbox := func(hitbox *components.Hitbox) {
			hitbox.Bound = &components.Bound{
				Left:   1,
				Top:    1,
				Right:  m.Window.Width - hitbox.Width - 1,
				Bottom: m.Window.Height - hitbox.Height - 1,
			}

			rl.DrawRectangle(
				hitbox.Position.X,
				hitbox.Position.Y,
				hitbox.Width,
				hitbox.Height,
				hitbox.Color,
			)
		}

		drawHitbox(&w.Player.Hitbox)
		for _, enemy := range w.Enemies {
			drawHitbox(&enemy.Hitbox)
		}
	})
}

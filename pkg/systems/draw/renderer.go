package systems

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type renderer struct{}

var Renderer = renderer{}

func (_ *renderer) Clear() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		rl.ClearScreenBuffers()
	})
}

func (_ *renderer) DrawTitleScreen() *system.System {
	frame := int32(0)
	shakeOffset := func(randomNumber float32, frame int32) float32 {
		if frame%utils.RandomBetween(100, 250) == 0 {
			return randomNumber
		}
		return 0
	}

	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeTitleScreen {
			return
		}

		if frame > math.MaxInt32 {
			frame = 0 // trying to avoid int overflows
		}

		fontSize := float32(m.Window.Width / 25)
		characterSpacing := float32(fontSize / 2.25)
		sizes := rl.MeasureTextEx(
			m.Font,
			m.Window.Title,
			fontSize,
			0,
		)

		x := (float32(m.Window.Width) - sizes.X/1.25) / 2
		y := (float32(m.Window.Height) - sizes.Y) / 2

		for idx, char := range m.Window.Title {
			baseAmplitude := float32(5)
			amplitude := baseAmplitude + shakeOffset(float32(rand.Intn(50)), frame)

			baseFrequency := float32(0.01)
			frequency := baseFrequency + shakeOffset(float32(rand.Float64()*0.75), frame)

			yOffset := amplitude * float32(math.Sin(float64(frequency*float32(frame))+float64(idx)))

			rl.DrawTextCodepoint(
				m.Font,
				char,
				rl.Vector2{
					X: x + characterSpacing*float32(idx),
					Y: y + (yOffset*2),
				},
				fontSize,
				rl.Color{0xFF, 0xFF, 0xFF, 0xFF},
			)
		}
		frame++
	})
}

func (_ *renderer) DrawHitboxesInExplorationMode() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

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

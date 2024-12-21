package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
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
	draw, next := archetypes.RenderChaoticChar.Draw(
		func(frame int32) bool {
			return frame%utils.RandomBetween(500, 300) == 0
		},
		func(frame int32) bool {
			return frame%utils.RandomBetween(2000, 5000) == 0
		},
		func(frame int32) bool {
			return frame%utils.RandomBetween(1000, 5000) == 0
		},
	)

	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeTitleScreen {
			return
		}

		titleFontSize := float32(m.Window.Width / 25)
		titleCharacterSpacing := float32(titleFontSize / 2.25)
		titleSizes := rl.MeasureTextEx(
			m.Font,
			m.Window.Title,
			titleFontSize,
			0,
		)
		titleX := (float32(m.Window.Width) - titleSizes.X/1.25) / 2
		titleY := (float32(m.Window.Height) - titleSizes.Y) / 2

		for idx, char := range m.Window.Title {
			draw(
				m,
				char, idx,
				15, 50,
				0.01, 1.5,
				titleFontSize, titleCharacterSpacing,
				titleX, titleY,
			)
		}

		startYourJourneyText := "Press ENTER / START to begin your journey..."
		startYourJourneyFontSize := float32(m.Window.Width / 60)
		startYourJourneyCharacterSpacing := float32(startYourJourneyFontSize / 2.25)
		startYourJourneySizes := rl.MeasureTextEx(
			m.Font,
			startYourJourneyText,
			startYourJourneyFontSize,
			0,
		)
		startYourJourneyX := (float32(m.Window.Width) - startYourJourneySizes.X/1.25) / 2
		startYourJourneyY := (float32(m.Window.Height) - startYourJourneySizes.Y) / 2

		for idx, char := range startYourJourneyText {
			draw(
				m,
				char, idx,
				1, 10,
				0.05, 0.5,
				startYourJourneyFontSize, startYourJourneyCharacterSpacing,
				startYourJourneyX, startYourJourneyY+150,
			)
		}

		next()
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

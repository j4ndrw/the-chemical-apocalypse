package renderer_systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/aggro"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/bound"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/chaoticchar"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/geometry"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/vision"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type draw struct{}

var Draw = draw{}

func (_ *draw) Clear() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		rl.ClearScreenBuffers()
	})
}

func (_ *draw) DrawTitleScreen() system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeTitleScreen {
			return
		}

		draw := chaoticchar.Archetype.Draw(
			m,
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
				char, idx,
				1, 10,
				0.05, 0.5,
				startYourJourneyFontSize, startYourJourneyCharacterSpacing,
				startYourJourneyX, startYourJourneyY+150,
			)
		}
	})
}

func (_ *draw) DrawHitboxesInExplorationMode(hitbox *hitbox.HitboxComponent) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		hitbox.Bound = &bound.BoundComponent{
			Left:   1,
			Top:    1,
			Right:  m.Window.Width - hitbox.Width - 1,
			Bottom: m.Window.Height - hitbox.Height - 1,
		}

		if !hitbox.Hidden {
			rl.DrawRectangleLinesEx(
				rl.Rectangle{
					X:      float32(hitbox.Position.X),
					Y:      float32(hitbox.Position.Y),
					Width:  float32(hitbox.Width),
					Height: float32(hitbox.Height),
				},
				4,
				hitbox.Color,
			)
		}
	})
}

func (_ *draw) DrawAggroInExplorationMode(
	hitbox *hitbox.HitboxComponent,
	aggro *aggro.AggroComponent,
	target *position.PositionComponent,
) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		centerX, centerY := geometry.Archetype.Center(hitbox)
		startAngle, endAngle, _ := vision.Archetype.CenterAngleEx(
			centerX,
			centerY,
			target,
			&aggro.Vision,
		)

		rl.DrawCircleSector(
			rl.Vector2{
				X: float32(centerX),
				Y: float32(centerY),
			},
			float32(aggro.Radius),
			startAngle,
			endAngle,
			25,
			rl.Color{
				0x18,
				0x18,
				0x18,
				utils.ColorOpacity(0.2),
			},
		)
	})
}

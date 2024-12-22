package archetypes

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type renderchaoticchar struct{}

var RenderChaoticChar = renderchaoticchar{}

func (_ *renderchaoticchar) Draw(
	m *meta.Meta,
	shakePredicate func(frame int32) bool,
	oscillateXPredicate func(frame int32) bool,
	oscillateYPredicate func(frame int32) bool,
) (func(
	char rune,
	idx int,
	baseAmplitude float32,
	maxAmplitude float32,
	baseFrequency float32,
	maxFrequency float32,
	fontSize float32,
	characterSpacing float32,
	x float32,
	y float32,
)) {
	shakeOffset := func(randomNumber float32, frame int32) float32 {
		if shakePredicate(m.Frame) {
			return randomNumber
		}
		return 0
	}
	displaceX := func(offset float32) float32 {
		if oscillateXPredicate(m.Frame) {
			return offset
		}
		return 0
	}
	displaceY := func(offset float32) float32 {
		if oscillateYPredicate(m.Frame) {
			return offset
		}
		return 0
	}

	draw := func(
		char rune, idx int,
		baseAmplitude, maxAmplitude float32,
		baseFrequency, maxFrequency float32,
		fontSize, characterSpacing float32,
		x, y float32,
	) {
		amplitude := baseAmplitude + shakeOffset(
			float32(rand.Intn(int(maxAmplitude))),
			m.Frame,
		)
		frequency := baseFrequency + shakeOffset(
			float32(rand.Float64()*float64(maxFrequency)),
			m.Frame,
		)

		xOffset := Physics.Oscillation(
			amplitude,
			math.Sin(
				float64(frequency*float32(m.Frame))+float64(idx),
			),
		)
		yOffset := Physics.Oscillation(
			amplitude,
			math.Sin(
				float64(frequency*float32(m.Frame))+(rand.Float64()*1.25),
			),
		)

		rl.DrawTextCodepoint(
			m.Font,
			char,
			rl.Vector2{
				X: x + (displaceX(float32(rand.Intn(5))) * xOffset) + characterSpacing*float32(idx),
				Y: y + yOffset + (displaceY(2) * yOffset),
			},
			fontSize,
			rl.Color{0xFF, 0xFF, 0xFF, 0xFF},
		)
	}

	return draw
}

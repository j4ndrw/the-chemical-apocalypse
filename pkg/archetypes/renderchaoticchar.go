package archetypes

import (
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type renderchaoticchar struct{}

var RenderChaoticChar = renderchaoticchar{}

func (_ *renderchaoticchar) Draw(shakePredicate func(frame int32) bool) (func(
	m *meta.Meta,
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
), func()) {
	frame := int32(0)
	shakeOffset := func(randomNumber float32, frame int32) float32 {
		if shakePredicate(frame) {
			return randomNumber
		}
		return 0
	}

	draw := func(
		m *meta.Meta,
		char rune, idx int,
		baseAmplitude, maxAmplitude float32,
		baseFrequency, maxFrequency float32,
		fontSize, characterSpacing float32,
		x, y float32,
	) {
		if frame > math.MaxInt32 {
			frame = 0 // trying to avoid int overflows
		}

		amplitude := baseAmplitude + shakeOffset(
			float32(rand.Intn(int(maxAmplitude))),
			frame,
		)
		frequency := baseFrequency + shakeOffset(
			float32(rand.Float64()*float64(maxFrequency)),
			frame,
		)

		yOffset := amplitude * float32(
			math.Sin(
				float64(frequency*float32(frame))+float64(idx),
			),
		)

		rl.DrawTextCodepoint(
			m.Font,
			char,
			rl.Vector2{
				X: x + characterSpacing*float32(idx),
				Y: y + (yOffset * 2),
			},
			fontSize,
			rl.Color{0xFF, 0xFF, 0xFF, 0xFF},
		)
	}

	next := func() {
		frame++
	}

	return draw, next
}

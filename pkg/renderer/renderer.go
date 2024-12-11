package renderer

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/gamestate"
)

func Clear(s *gamestate.State) {
	rl.ClearScreenBuffers()
}

func Rectangle(s *gamestate.State) {
	rl.DrawRectangle(
		s.Player.Position.X,
		s.Player.Position.Y,
		200,
		200,
		rl.Color{
			0xFF,
			0,
			0,
			0xFF,
		},
	)
}

package archetypes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) DrawSpriteFromAtlas(
	atlas *rl.Texture2D,
	x int32,
	y int32,
	hitbox *components.Hitbox,
	direction *components.Direction,
) {
	width := hitbox.Width
	if direction.X != components.DirectionNone {
		width *= int32(direction.X)
	}
	rl.DrawTextureRec(
		*atlas,
		rl.Rectangle{
			X:      float32(x),
			Y:      float32(y),
			Width:  float32(width),
			Height: float32(hitbox.Height),
		},
		rl.Vector2{
			X: float32(hitbox.Position.X),
			Y: float32(hitbox.Position.Y),
		},
		rl.White,
	)
}

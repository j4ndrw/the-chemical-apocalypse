package archetypes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) DrawSprite(sprite *components.Sprite) {
	rl.DrawTexturePro(
		*sprite.ParentAtlas,
		sprite.Src,
		sprite.Dest,
		rl.Vector2{X: 0, Y: 0},
		0.0,
		rl.White,
	)
}

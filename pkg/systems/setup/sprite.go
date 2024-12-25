package systems

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) SetupPlayerSprites() system.System {
	return func(w *world.World, m *meta.Meta) {
		w.Player.SpriteMap = components.SpriteMap{
			constants.Keys.Idle: &components.Sprite{
				ParentAtlas: m.SpriteAtlas,
				Scale:       8,
				Src:         rl.Rectangle{X: 0, Y: 0, Width: float32(w.Player.Hitbox.Width), Height: float32(w.Player.Hitbox.Height)},
				Dest:        rl.Rectangle{X: float32(w.Player.Position.X), Y: float32(w.Player.Position.Y), Width: float32(w.Player.Hitbox.Width), Height: float32(w.Player.Hitbox.Height)},
				FlipX:       false,
				FlipY:       false,
				Ticker:      components.Ticker{},
			},
		}
	}
}

package sprite

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/ticker"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) New(atlas *rl.Texture2D, h *hitbox.HitboxComponent, scale float32) *SpriteComponent {
	return &SpriteComponent{
		ParentAtlas: atlas,
		Scale:       scale,
		Src: rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(h.Width),
			Height: float32(h.Height),
		},
		Dest: rl.Rectangle{
			X:      float32(h.Position.X),
			Y:      float32(h.Position.Y),
			Width:  float32(h.Width),
			Height: float32(h.Height),
		},
		FlipX:  false,
		FlipY:  false,
		Ticker: ticker.TickerComponent{},
	}
}

func (_ *archetype) Draw(sprite *SpriteComponent) {
	rl.DrawTexturePro(
		*sprite.ParentAtlas,
		sprite.Src,
		sprite.Dest,
		rl.Vector2{X: 0, Y: 0},
		0.0,
		rl.White,
	)
}

func (_ *archetype) GetSpriteRow(key string) float32 {
	if key == constants.Keys.PlayerIdle {
		return 0
	}
	if key == constants.Keys.PlayerMoveDown {
		return 1
	}
	if key == constants.Keys.PlayerMoveUp {
		return 2
	}
	if key == constants.Keys.PlayerMoveForward {
		return 3
	}
	if key == constants.Keys.AntiPeanutIdle {
		return 4
	}
	if key == constants.Keys.AntiPeanutMoveDown {
		return 5
	}
	if key == constants.Keys.AntiPeanutMoveUp {
		return 6
	}
	if key == constants.Keys.AntiPeanutMoveForward {
		return 7
	}
	return -1
}

func (_ *archetype) Animate(
	m *meta.Meta,
	id *id.IdComponent,
	key *SpriteKeyComponent,
	h *hitbox.HitboxComponent,
	dir *direction.DirectionComponent,
	spriteMap *SpriteMapComponent,
) {
	spriteMap.Lock()
	sprite := spriteMap.Map[string(*key)]
	spriteMap.Unlock()
	sheetRow := Archetype.GetSpriteRow(string(*key))

	variants := int32(float32(m.SpriteAtlas.Width) / (float32(h.Width) / sprite.Scale))

	if sprite.Ticker.Ticker == nil {
		sprite.Ticker.Ticker = time.NewTicker(75 * time.Millisecond)
	}

	sprite.Src.Y = float32(sheetRow) * (float32(h.Height) / sprite.Scale)
	sprite.Dest.X = float32(h.Position.X)
	sprite.Dest.Y = float32(h.Position.Y)
	sprite.FlipX = dir.X != direction.None && dir.X < 0
	sprite.Src.Width = float32(h.Width)
	sprite.Src.Width /= sprite.Scale
	if sprite.FlipX {
		sprite.Src.Width *= -1
	}
	sprite.Src.Height = float32(h.Height)
	sprite.Src.Height /= sprite.Scale

	Coroutine.Animate(
		id,
		string(*key),
		spriteMap,
		variants,
		sprite.Ticker.Ticker,
	).CallOnce()

	spriteMap.Lock()
	spriteMap.Map[string(*key)] = sprite
	spriteMap.Unlock()
}

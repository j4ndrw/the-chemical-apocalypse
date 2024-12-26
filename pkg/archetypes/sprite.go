package archetypes

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/coroutines"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type sprite struct{}

var Sprite = sprite{}

func (_ *sprite) New(atlas *rl.Texture2D, hitbox *components.Hitbox, scale float32) *components.Sprite {
	return &components.Sprite{
		ParentAtlas: atlas,
		Scale:       scale,
		Src: rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(hitbox.Width),
			Height: float32(hitbox.Height),
		},
		Dest: rl.Rectangle{
			X:      float32(hitbox.Position.X),
			Y:      float32(hitbox.Position.Y),
			Width:  float32(hitbox.Width),
			Height: float32(hitbox.Height),
		},
		FlipX:  false,
		FlipY:  false,
		Ticker: components.Ticker{},
	}
}

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

func (_ *sprite) GetSpriteRow(key string) float32 {
	if key == constants.Keys.PlayerIdle {
		return 0
	}
	if key == constants.Keys.PlayerMoveDown {
		return 1
	}
	if key == constants.Keys.PlayerMoveUp {
		return 2
	}
	return -1
}

func (_ *sprite) AnimateSprite(
	m *meta.Meta,
	id *components.Id,
	key *components.SpriteKey,
	hitbox *components.Hitbox,
	direction *components.Direction,
	spriteMap *components.SpriteMap,
) {
	spriteMap.Lock()
	sprite := spriteMap.Map[string(*key)]
	spriteMap.Unlock()
	sheetRow := Sprite.GetSpriteRow(string(*key))

	variants := int32(float32(m.SpriteAtlas.Width) / (float32(hitbox.Width) / sprite.Scale))

	if sprite.Ticker.Ticker == nil {
		sprite.Ticker.Ticker = time.NewTicker(100 * time.Millisecond)
	}

	sprite.Src.Y = float32(sheetRow) * (float32(hitbox.Height) / sprite.Scale)
	sprite.Dest.X = float32(hitbox.Position.X)
	sprite.Dest.Y = float32(hitbox.Position.Y)
	sprite.FlipX = direction.X != components.DirectionNone && direction.X < 0
	sprite.Src.Width = float32(hitbox.Width)
	sprite.Src.Width /= sprite.Scale
	if sprite.FlipX {
		sprite.Src.Width *= -1
	}
	sprite.Src.Height = float32(hitbox.Height)
	sprite.Src.Height /= sprite.Scale

	coroutines.Sprite.Animate(
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

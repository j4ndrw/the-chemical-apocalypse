package sprite

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/ticker"
)

type SpriteComponent struct {
	ParentAtlas *rl.Texture2D
	Scale       float32
	Src         rl.Rectangle
	Dest        rl.Rectangle
	FlipX       bool
	FlipY       bool
	Ticker      ticker.TickerComponent
}

type SpriteMapComponent struct {
	Map map[string]*SpriteComponent
	sync.Mutex
}

type SpriteKeyComponent string

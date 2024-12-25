package components

import (
	"sync"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	ParentAtlas *rl.Texture2D
	Scale       float32
	Src         rl.Rectangle
	Dest        rl.Rectangle
	FlipX       bool
	FlipY       bool
	Ticker
}

type SpriteMap struct {
	Map map[string]*Sprite
	sync.Mutex
}

type SpriteKey string

package components

import (
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

type SpriteMap map[string]*Sprite

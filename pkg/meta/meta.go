package meta

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Window struct {
	Width  int32
	Height int32
	Title  string
	Screen struct {
		Windowed   bool
		Fullscreen bool
		Borderless bool
	}
}

type Meta struct {
	Window      Window
	TargetFPS   int32
	DeltaTime   float32
	Font        rl.Font
	Frame       int32
	SpriteAtlas map[string]*rl.Texture2D
}

func New(window Window, targetFPS int32) *Meta {
	return &Meta{
		Window:      window,
		TargetFPS:   targetFPS,
		SpriteAtlas: map[string]*rl.Texture2D{},
	}
}

func Default() *Meta {
	return New(
		Window{Title: "The Chemical Apocalypse"},
		60,
	)
}

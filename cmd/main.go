package main

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/engine"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	drawsystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/draw"
	updatesystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/update"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

func main() {
	world := world.New()
	meta := meta.New(meta.Window{
		Title:  "The Chemical Apocalypse",
		Width:  640,
		Height: 480,
	}, 60)

	exitHandler := engine.Setup().
		WithWorld(world).
		WithMeta(meta).
		Async().
		Run()
	defer exitHandler()

	draw := engine.
		Block(drawsystems.Systems()...).
		WithWorld(world).
		WithMeta(meta)
	update := engine.
		Block(updatesystems.Systems()...).
		WithWorld(world).
		WithMeta(meta)

	engine.Loop(draw, update)
}

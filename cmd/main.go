package main

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/engine"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	drawsystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/draw"
	updatesystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/update"
	assertionssystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/assertions"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

func main() {
	world := world.Default()
	meta := meta.Default()

	exitHandler := engine.Setup().
		WithWorld(world).
		WithMeta(meta).
		Async().
		Run()
	defer exitHandler()

	draw := engine.
		Block(drawsystems.Systems...).
		WithWorld(world).
		WithMeta(meta)
	update := engine.
		Block(updatesystems.Systems...).
		WithWorld(world).
		WithMeta(meta)

	engine.
		Loop(draw, update).
		WithWorld(world).
		WithMeta(meta).
		WithAssertions(assertionssystems.Systems...).
		Run()
}

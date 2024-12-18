package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/engine"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	assertionssystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/assertions"
	drawsystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/draw"
	setupsystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/setup"
	updatesystems "github.com/j4ndrw/the-chemical-apocalypse/pkg/systems/update"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

func main() {
	world := world.Default()
	meta := meta.Default()

	block := func(systems ...system.System) *engine.EngineBuilder {
		return engine.
			Block(systems...).
			WithWorld(world).
			WithMeta(meta)
	}

	setup := block(setupsystems.Systems...).
		WithDeferredHandler(rl.CloseWindow)
	draw := block(drawsystems.Systems...)
	update := block(updatesystems.Systems...)

	exit := setup.Run()
	defer exit()

	engine.
		Loop(draw, update).
		WithWorld(world).
		WithMeta(meta).
		WithAssertions(assertionssystems.Systems...).
		Run()
}

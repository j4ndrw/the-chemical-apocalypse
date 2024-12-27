package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/engine"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/systems"
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

	setup := block(systems.Setup...).
		WithDeferredHandler(rl.CloseWindow)
	draw := block(systems.Draw...)
	update := block(systems.Update...)

	exit := setup.Run()
	defer exit()

	engine.
		Loop(draw, update).
		WithWorld(world).
		WithMeta(meta).
		Run()
}

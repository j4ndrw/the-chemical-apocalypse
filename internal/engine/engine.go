package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type EngineBuilder struct {
	deferredHandler     func()
	Run                 func() func()
	world               *world.World
	meta                *meta.Meta
	WithWorld           func(s *world.World) *EngineBuilder
	WithMeta            func(m *meta.Meta) *EngineBuilder
	WithDeferredHandler func(func()) *EngineBuilder
}

type EngineLoop struct {
	assertions     []system.System
	world          *world.World
	meta           *meta.Meta
	WithWorld      func(s *world.World) *EngineLoop
	WithMeta       func(m *meta.Meta) *EngineLoop
	WithAssertions func(systems ...system.System) *EngineLoop
	Run            func() func()
}

func Block(systems ...system.System) EngineBuilder {
	options := EngineBuilder{}

	options.WithWorld = func(s *world.World) *EngineBuilder {
		options.world = s
		return &options
	}

	options.WithMeta = func(m *meta.Meta) *EngineBuilder {
		options.meta = m
		return &options
	}

	options.WithDeferredHandler = func(handler func()) *EngineBuilder {
		options.deferredHandler = handler
		return &options
	}

	options.Run = func() func() {
		for _, function := range systems {
			function(options.world, options.meta)
		}
		return options.deferredHandler
	}

	return options
}

func Loop(draw *EngineBuilder, update *EngineBuilder) EngineLoop {
	options := EngineLoop{}

	options.WithAssertions = func(systems ...system.System) *EngineLoop {
		options.assertions = systems
		return &options
	}

	options.WithWorld = func(s *world.World) *EngineLoop {
		options.world = s
		return &options
	}

	options.WithMeta = func(m *meta.Meta) *EngineLoop {
		options.meta = m
		return &options
	}

	options.Run = func() func() {
		for !rl.WindowShouldClose() {
			rl.BeginDrawing()
			draw.Run()
			rl.EndDrawing()

			update.Run()

			if len(options.assertions) > 0 {
				for _, assertion := range options.assertions {
					assertion(options.world, options.meta)
				}
			}
		}
		return func() {}
	}

	return options
}

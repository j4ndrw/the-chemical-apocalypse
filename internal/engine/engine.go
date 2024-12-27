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
	world          *world.World
	meta           *meta.Meta
	WithWorld      func(s *world.World) *EngineLoop
	WithMeta       func(m *meta.Meta) *EngineLoop
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
		}
		return func() {}
	}

	return options
}

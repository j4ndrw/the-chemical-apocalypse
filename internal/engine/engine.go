package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type EngineBuilder struct {
	parallel  bool
	Run       func() func()
	Async     func() *EngineBuilder
	Sync      func() *EngineBuilder
	world     *world.World
	meta      *meta.Meta
	WithWorld func(s *world.World) *EngineBuilder
	WithMeta  func(m *meta.Meta) *EngineBuilder
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

func Setup(systems ...system.System) EngineBuilder {
	options := EngineBuilder{parallel: false}

	options.Async = func() *EngineBuilder {
		options.parallel = true
		return &options
	}

	options.Sync = func() *EngineBuilder {
		options.parallel = false
		return &options
	}

	options.WithWorld = func(s *world.World) *EngineBuilder {
		options.world = s
		return &options
	}

	options.WithMeta = func(m *meta.Meta) *EngineBuilder {
		options.meta = m
		return &options
	}

	options.Run = func() func() {
		rl.InitWindow(options.meta.Window.Width, options.meta.Window.Height, options.meta.Window.Title)
		rl.SetTargetFPS(options.meta.TargetFPS)

		for _, function := range systems {
			if options.parallel {
				go function(options.world, options.meta)
			} else {
				function(options.world, options.meta)
			}
		}

		return rl.CloseWindow
	}

	return options
}

func Block(systems ...system.System) EngineBuilder {
	options := EngineBuilder{parallel: false}

	options.Async = func() *EngineBuilder {
		options.parallel = true
		return &options
	}

	options.Sync = func() *EngineBuilder {
		options.parallel = false
		return &options
	}

	options.WithWorld = func(s *world.World) *EngineBuilder {
		options.world = s
		return &options
	}

	options.WithMeta = func(m *meta.Meta) *EngineBuilder {
		options.meta = m
		return &options
	}

	options.Run = func() func() {
		for _, function := range systems {
			if options.parallel {
				go function(options.world, options.meta)
			} else {
				function(options.world, options.meta)
			}
		}
		return func() {}
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

package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type EngineBuilder struct {
	parallel  bool
	world     *world.World
	meta      *meta.Meta
	Run       func() func()
	Async     func() *EngineBuilder
	Sync      func() *EngineBuilder
	WithWorld func(s *world.World) *EngineBuilder
	WithMeta  func(m *meta.Meta) *EngineBuilder
}

func Setup(systems ...system.System) EngineBuilder {
	options := EngineBuilder{
		parallel:  false,
		world:     nil,
		Async:     nil,
		Sync:      nil,
		Run:       nil,
		WithWorld: nil,
	}

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
	options := EngineBuilder{
		parallel:  false,
		world:     nil,
		Async:     nil,
		Sync:      nil,
		Run:       nil,
		WithWorld: nil,
	}

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

func Loop(draw *EngineBuilder, update *EngineBuilder) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		draw.Run()
		rl.EndDrawing()

		update.Run()
	}
}

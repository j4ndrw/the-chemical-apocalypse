package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/gamestate"
)

type enginebuilder struct {
	parallel  bool
	state     *gamestate.State
	Run       func() func()
	Async     func() *enginebuilder
	Sync      func() *enginebuilder
	WithState func(s *gamestate.State) *enginebuilder
}

func Setup(functions ...func(s *gamestate.State)) enginebuilder {
	options := enginebuilder{
		parallel:  false,
		state:     nil,
		Async:     nil,
		Sync:      nil,
		Run:       nil,
		WithState: nil,
	}

	options.Async = func() *enginebuilder {
		options.parallel = true
		return &options
	}

	options.Sync = func() *enginebuilder {
		options.parallel = false
		return &options
	}

	options.WithState = func(s *gamestate.State) *enginebuilder {
		options.state = s
		return &options
	}

	options.Run = func() func() {
		rl.InitWindow(800, 450, "raylib [core] example - basic window")

		rl.SetTargetFPS(60)

		for _, function := range functions {
			if options.parallel {
				go function(options.state)
			} else {
				function(options.state)
			}
		}

		return rl.CloseWindow
	}

	return options
}

func Loop(draw *enginebuilder, update *enginebuilder) {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		draw.Run()
		rl.EndDrawing()

		update.Run()
	}

}

func Block(functions ...func(s *gamestate.State)) enginebuilder {
	options := enginebuilder{
		parallel:  false,
		state:     nil,
		Async:     nil,
		Sync:      nil,
		Run:       nil,
		WithState: nil,
	}

	options.Async = func() *enginebuilder {
		options.parallel = true
		return &options
	}

	options.Sync = func() *enginebuilder {
		options.parallel = false
		return &options
	}

	options.WithState = func(s *gamestate.State) *enginebuilder {
		options.state = s
		return &options
	}

	options.Run = func() func() {
		for _, function := range functions {
			if options.parallel {
				go function(options.state)
			} else {
				function(options.state)
			}
		}
		return func() {}
	}

	return options
}

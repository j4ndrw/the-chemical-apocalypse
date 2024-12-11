package main

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/engine"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/game"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/gamestate"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/renderer"
)

func main() {
	state := gamestate.New()

	exitHandler := engine.
		Setup().
		WithState(&state).
		Async().
		Run()
	defer exitHandler()

	draw := engine.
		Block(renderer.Clear, renderer.Rectangle).
		WithState(&state).
		Sync()
	update := engine.
		Block(game.MovePlayerDown, game.MovePlayerRight).
		WithState(&state).
		Sync()

	engine.Loop(draw, update)
}

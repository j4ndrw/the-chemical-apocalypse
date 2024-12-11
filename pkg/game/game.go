package game

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/gamestate"

func MovePlayerLeft(s *gamestate.State) {
	s.Player.MoveLeft(1)
}

func MovePlayerRight(s *gamestate.State) {
	s.Player.MoveRight(1)
}

func MovePlayerUp(s *gamestate.State) {
	s.Player.MoveUp(1)
}

func MovePlayerDown(s *gamestate.State) {
	s.Player.MoveDown(1)
}

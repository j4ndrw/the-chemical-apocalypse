package world

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
)

type WorldMode int32

const (
	WorldModeNil = iota
	WorldModeCombat
	WorldModeExploration
	WorldModePause
	WorldModeTitleScreen
	WorldModeInventory
)

type WorldAction int32

const (
	WorldActionQuit = iota
)

type World struct {
	Player      entities.Player
	Enemies     map[components.Id]*entities.Enemy
	CurrentMode WorldMode
	PrevMode    WorldMode
}

func Default() *World {
	return &World{
		Player:  entities.Player{},
		Enemies: map[components.Id]*entities.Enemy{},
	}
}

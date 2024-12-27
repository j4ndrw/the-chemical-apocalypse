package world

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/player"
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
	Player      player.PlayerEntity
	Enemies     map[id.IdComponent]*enemy.EnemyEntity
	CurrentMode WorldMode
	PrevMode    WorldMode
}

func Default() *World {
	return &World{
		Player:  player.PlayerEntity{},
		Enemies: map[id.IdComponent]*enemy.EnemyEntity{},
	}
}

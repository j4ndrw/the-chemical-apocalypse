package world

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
)

type World struct {
	Player      entities.Player
	Enemies     map[components.Id]*entities.Enemy
}

func Default() *World {
	return &World{
		Player:  entities.Player{},
		Enemies: map[components.Id]*entities.Enemy{},
	}
}

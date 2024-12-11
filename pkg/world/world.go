package world

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
)

type World struct {
	Player entities.Player
	Enemy  entities.Enemy
}

func New() *World {
	return &World{
		Player: entities.Player{
			Position: components.Position{X: 0, Y: 0},
		},
		Enemy: entities.Enemy{
			Position: components.Position{X: 0, Y: 0},
		},
	}
}

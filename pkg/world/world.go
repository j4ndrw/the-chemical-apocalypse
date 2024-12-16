package world

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
)

type World struct {
	Player entities.Player
	Enemy  entities.Enemy
}

func Default() *World {
	return &World{
		Player: entities.Player{
			Position: components.Position{
				Vector2: components.Vector2{X: 0, Y: 0},
				Width:   40,
				Height:  40,
			},
			Speed: 10,
		},
		Enemy: entities.Enemy{
			Position: components.Position{
				Vector2: components.Vector2{X: 0, Y: 0},
				Width:   40,
				Height:  40,
			},
			Speed: 5,
		},
	}
}

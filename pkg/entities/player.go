package entities

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type Player struct {
	components.Hitbox
	components.Speed
}

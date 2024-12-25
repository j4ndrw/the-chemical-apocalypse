package entities

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type Enemy struct {
	components.Id
	components.Hitbox
	MinSpeed components.Speed
	MaxSpeed components.Speed
	components.Aggro
	components.Roam
	components.Moving
}

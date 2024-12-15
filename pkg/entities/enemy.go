package entities

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type Enemy struct {
	components.Position
	components.Speed
}

package entities

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type Player struct {
	components.Id
	components.Hitbox
	components.Speed
	components.Moving
	SpriteMap components.SpriteMap
	SpriteKey components.SpriteKey
}

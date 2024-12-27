package id

import "github.com/google/uuid"

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) Create() IdComponent {
	return IdComponent(uuid.New().String())
}

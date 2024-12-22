package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type MoveFn func(m *meta.Meta, h *components.Hitbox, s *components.Speed)
type ChaseFn func(
	chaser *components.Hitbox,
	target *components.Position,
	speed components.Speed,
	deltaTime float64,
	collisions ...bool,
)

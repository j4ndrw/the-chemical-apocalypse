package types

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/speed"
)

type MoveFn func(m *meta.Meta, h *hitbox.HitboxComponent, s *speed.SpeedComponent)
type ChaseFn func(
	chaser *hitbox.HitboxComponent,
	target *position.PositionComponent,
	speed speed.SpeedComponent,
	deltaTime float64,
	collisions ...bool,
)

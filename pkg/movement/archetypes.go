package movement

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/geometry"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/hitbox"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/pathfinding"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/speed"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) NearestNeighborChase(
	chaser *hitbox.HitboxComponent,
	target *position.PositionComponent,
	speed speed.SpeedComponent,
	deltaTime float64,
	collisions ...bool,
) {
	neighbor := pathfinding.Archetype.ClosestNeighbor(
		&chaser.Position,
		target,
		func(pos, direction *position.PositionComponent) *position.PositionComponent {
			return &position.PositionComponent{
				X: int32(pos.FloatX() + direction.FloatX()*speed.AsFloat()*deltaTime),
				Y: int32(pos.FloatY() + direction.FloatY()*speed.AsFloat()*deltaTime),
			}
		},
		func(p *position.PositionComponent) bool {
			for _, collision := range collisions {
				if collision {
					return false
				}
			}
			return true
		},
	)

	if neighbor == nil {
		return
	}

	direction.Archetype.Update(
		&chaser.Direction,
		&chaser.Position,
		neighbor.Position.X,
		neighbor.Position.Y,
	)
	chaser.Position = *neighbor.Position
}

func (_ *archetype) NaiveChase(
	chaser *hitbox.HitboxComponent,
	target *position.PositionComponent,
	speed speed.SpeedComponent,
	deltaTime float64,
	collisions ...bool,
) {
	for _, collision := range collisions {
		if collision {
			return
		}
	}

	dx, dy := geometry.Archetype.Delta(target, &chaser.Position)
	distance := geometry.Archetype.Distance(dx, dy)

	if distance > 0 {
		dx /= distance
		dy /= distance
	}

	newX := int32(chaser.Position.FloatX() + dx*(speed.AsFloat() * deltaTime))
	newY := int32(chaser.Position.FloatY() + dy*(speed.AsFloat() * deltaTime))

	direction.Archetype.Update(
		&chaser.Direction,
		&chaser.Position,
		newX,
		newY,
	)

	chaser.Position.X = newX
	chaser.Position.Y = newY
}

func (_ *archetype) MoveUp(m *meta.Meta, h *hitbox.HitboxComponent, s *speed.SpeedComponent) {
	if h.Position.Y < h.Bound.Top {
		return
	}
	h.Position.Y = int32(float32(h.Position.Y) - float32(*s)*m.DeltaTime)
	h.Direction.Y = direction.Up
}

func (_ *archetype) MoveDown(m *meta.Meta, h *hitbox.HitboxComponent, s *speed.SpeedComponent) {
	if h.Position.Y > h.Bound.Bottom {
		return
	}
	h.Position.Y = int32(float32(h.Position.Y) + float32(*s)*m.DeltaTime)
	h.Direction.Y = direction.Down
}

func (_ *archetype) MoveLeft(m *meta.Meta, h *hitbox.HitboxComponent, s *speed.SpeedComponent) {
	if h.Position.X < h.Bound.Left {
		return
	}
	h.Position.X = int32(float32(h.Position.X) - float32(*s)*m.DeltaTime)
	h.Direction.X = direction.Left
}

func (_ *archetype) MoveRight(m *meta.Meta, h *hitbox.HitboxComponent, s *speed.SpeedComponent) {
	if h.Position.X > h.Bound.Right {
		return
	}
	h.Position.X = int32(float32(h.Position.X) + float32(*s)*m.DeltaTime)
	h.Direction.X = direction.Right
}

package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy enemy = enemy{}

func (_ *enemy) ApproachPlayer() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		player := w.Player.Position
		enemy := w.Enemy.Position

		neighbor := archetypes.PathFinding.ClosestNeighbor(
			&enemy.Vector2,
			&player.Vector2,
			func(position, direction *components.Vector2) *components.Vector2 {
				return &components.Vector2{X: position.X + direction.X*int32(w.Enemy.Speed), Y: position.Y + direction.Y*int32(w.Enemy.Speed)}
			},
			func(position *components.Vector2) bool {
				return !archetypes.Collidable.IsColliding(&w.Player.Position, &w.Enemy.Position)
			},
		)

		if neighbor != nil {
			w.Enemy.Position.Vector2 = *neighbor.Position
		}
	})
}

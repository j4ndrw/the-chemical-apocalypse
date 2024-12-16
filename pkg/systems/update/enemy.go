package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy enemy = enemy{}

func (_ *enemy) ChasePlayer() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				if !enemy.Aggro.Aggro {
					return
				}

				archetypes.MobMovement.Chase(
					&enemy.Position,
					&w.Player.Position.Point,
					enemy.MaxSpeed,
					archetypes.Collidable.IsColliding(
						&w.Player.Position,
						&enemy.Position,
					),
				)
			}(enemy)
		}
	})
}

func (_ *enemy) RoamMindlessly() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				if enemy.Aggro.Aggro {
					enemy.Roam.Timer = 0
					return
				}

				if enemy.Roam.Timer <= 0 {
					enemy.Roam.Direction = archetypes.Point.RandomPointOnMap(&m.Window)
					enemy.Roam.Timer = enemy.Roam.Duration
					return
				}

				archetypes.MobMovement.Chase(
					&enemy.Position,
					&enemy.Roam.Direction,
					enemy.MinSpeed,
				)
				enemy.Roam.Timer -= m.DeltaTime
			}(enemy)
		}
	})
}

func (_ *enemy) WatchAggro() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			func(enemy *entities.Enemy) {
				if archetypes.Aggro.IsWithinAggroRange(
					&w.Player.Position.Point,
					&enemy.Position.Point,
					&enemy.Aggro,
				) {
					archetypes.Aggro.EnterAggro(&enemy.Aggro)
				} else {
					archetypes.Aggro.LeaveAggro(&enemy.Aggro)
				}
			}(enemy)
		}
	})
}

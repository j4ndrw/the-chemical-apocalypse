package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type enemy struct{}

var Enemy = enemy{}

func (_ *enemy) ChasePlayer(enemy *entities.Enemy) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if !enemy.Aggro.Aggro {
			return
		}

		archetypes.MobMovement.NaiveChase(
			&enemy.Hitbox,
			&w.Player.Hitbox.Position,
			enemy.MaxSpeed,
			float64(m.DeltaTime),
			archetypes.Collidable.IsColliding(
				&w.Player.Hitbox,
				&enemy.Hitbox,
				),
			)
	})
}

func (_ *enemy) RoamMindlessly(enemy *entities.Enemy) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if enemy.Aggro.Aggro {
			enemy.Roam.Timer = 0
			return
		}

		if enemy.Roam.Timer <= 0 {
			enemy.Roam.Where = archetypes.Position.RandomPointOnMap(&m.Window)
			enemy.Roam.Timer = enemy.Roam.Duration
			return
		}

		archetypes.MobMovement.NaiveChase(
			&enemy.Hitbox,
			&enemy.Roam.Where,
			enemy.MinSpeed,
			float64(m.DeltaTime),
		)
		enemy.Roam.Timer -= m.DeltaTime
	})
}

func (_ *enemy) WatchAggro(enemy *entities.Enemy) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}



		if archetypes.Aggro.IsWithinAggroRange(
			&w.Player.Hitbox.Position,
			&enemy.Hitbox.Position,
			&enemy.Aggro,
			) {
			archetypes.Aggro.EnterAggro(&enemy.Aggro)
		} else {
			archetypes.Aggro.LeaveAggro(&enemy.Aggro)
		}
	})
}

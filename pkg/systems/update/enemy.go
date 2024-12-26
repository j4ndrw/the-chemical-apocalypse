package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/archetypes"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/coroutines"
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

		enemy.Moving = true
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
			archetypes.Roam.ResetTicker(&enemy.Roam)
			coroutines.Roam.Tick(&enemy.Id, &enemy.Roam).Remove()
			return
		}

		if enemy.Roam.Ticker.Ticker == nil {
			archetypes.Roam.StartTicker(&enemy.Roam, &m.Window)
			coroutines.Roam.Tick(&enemy.Id, &enemy.Roam).Remove()
			return
		}

		coroutines.Roam.Tick(&enemy.Id, &enemy.Roam).CallOnce()

		oldX := enemy.Hitbox.Position.X
		oldY := enemy.Hitbox.Position.Y

		archetypes.MobMovement.NaiveChase(
			&enemy.Hitbox,
			&enemy.Roam.Where,
			enemy.MinSpeed,
			float64(m.DeltaTime),
		)
		enemy.Moving = oldX != enemy.Position.X || oldY != enemy.Position.Y
		if !enemy.Moving {
			enemy.Direction.X = components.DirectionNone
			enemy.Direction.Y = components.DirectionNone
		}
	})
}

func (_ *enemy) WatchAggro(enemy *entities.Enemy) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if archetypes.Aggro.IsWithinAggroRange(
			&w.Player.Hitbox,
			&enemy.Hitbox,
			&enemy.Aggro,
		) {
			archetypes.Aggro.EnterAggro(&enemy.Aggro)
		} else {
			archetypes.Aggro.LeaveAggro(&enemy.Aggro)
		}
	})
}

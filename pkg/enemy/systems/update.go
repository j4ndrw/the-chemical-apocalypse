package enemy_systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/aggro"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/collidable"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/direction"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/movement"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/roam"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type update struct{}

var Update = update{}

func (_ *update) ChasePlayer(enemy *enemy.EnemyEntity) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if !enemy.Aggro.Aggro {
			return
		}

		enemy.Moving = true
		movement.Archetype.NaiveChase(
			&enemy.Hitbox,
			&w.Player.Hitbox.Position,
			enemy.MaxSpeed,
			float64(m.DeltaTime),
			collidable.Archetype.IsColliding(
				&w.Player.Hitbox,
				&enemy.Hitbox,
			),
		)
	})
}

func (_ *update) RoamMindlessly(enemy *enemy.EnemyEntity) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if enemy.Aggro.Aggro {
			roam.Archetype.ResetTicker(&enemy.Roam)
			roam.Coroutine.Tick(&enemy.Id, &enemy.Roam).Remove()
			return
		}

		if enemy.Roam.Ticker.Ticker == nil {
			roam.Archetype.StartTicker(&enemy.Roam, &m.Window)
			roam.Coroutine.Tick(&enemy.Id, &enemy.Roam).Remove()
			return
		}

		roam.Coroutine.Tick(&enemy.Id, &enemy.Roam).CallOnce()

		oldX := enemy.Hitbox.Position.X
		oldY := enemy.Hitbox.Position.Y

		movement.Archetype.NaiveChase(
			&enemy.Hitbox,
			&enemy.Roam.Where,
			enemy.MinSpeed,
			float64(m.DeltaTime),
		)
		enemy.Moving = oldX != enemy.Hitbox.Position.X || oldY != enemy.Hitbox.Position.Y
		if !enemy.Moving {
			enemy.Hitbox.Direction.X = direction.None
			enemy.Hitbox.Direction.Y = direction.None
		}
	})
}

func (_ *update) WatchAggro(enemy *enemy.EnemyEntity) system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {
		if w.CurrentMode != world.WorldModeExploration {
			return
		}

		if aggro.Archetype.IsWithinAggroRange(
			&w.Player.Hitbox,
			&enemy.Hitbox,
			&enemy.Aggro,
		) {
			aggro.Archetype.EnterAggro(&enemy.Aggro)
		} else {
			aggro.Archetype.LeaveAggro(&enemy.Aggro)
		}
	})
}

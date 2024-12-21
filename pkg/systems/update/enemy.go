package systems

import (
	"math/rand"
	"time"

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
			enemy.Roam.Duration = 0
			enemy.Roam.ElapsedTime = 0
			enemy.Roam.Started = false
			if enemy.Roam.Ticker != nil {
				enemy.Roam.Ticker.Stop()
				enemy.Roam.Ticker = nil
			}
			return
		}

		if enemy.Roam.Ticker == nil {
			enemy.Roam.ElapsedTime = 0
			enemy.Roam.Started = false
			enemy.Roam.Duration = time.Duration(rand.Intn(int(enemy.Roam.MaxDuration)) + 1)
			enemy.Roam.Where = archetypes.Position.RandomPointOnMap(&m.Window)
			enemy.Roam.Ticker = time.NewTicker(time.Second)
			return
		}

		go func() {
			if enemy.Roam.Ticker == nil {
				return
			}
			if enemy.Roam.Started {
				return
			}

			enemy.Roam.Started = true
			for {
				select {
				case <-enemy.Roam.Ticker.C:
					enemy.Roam.ElapsedTime += time.Second
					if enemy.Roam.ElapsedTime >= enemy.Roam.Duration {
						enemy.Roam.Ticker.Stop()
						enemy.Roam.Ticker = nil
						return
					}
				}
			}
		}()

		archetypes.MobMovement.NaiveChase(
			&enemy.Hitbox,
			&enemy.Roam.Where,
			enemy.MinSpeed,
			float64(m.DeltaTime),
		)
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

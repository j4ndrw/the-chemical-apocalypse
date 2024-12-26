package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

var Systems system.SystemSlice = *system.
	Slice().
	Register(Input.HandleJourneyStart()).
	Register(Player.HandleMovement()).
	Register(Sprite.UpdatePlayerSprite()).
	Register(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			Sprite.UpdateEnemySprite(enemy).Apply(w, m)
			Enemy.WatchAggro(enemy).Apply(w, m)
			Enemy.RoamMindlessly(enemy).Apply(w, m)
			Enemy.ChasePlayer(enemy).Apply(w, m)
		}
	})

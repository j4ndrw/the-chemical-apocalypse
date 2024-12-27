package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	enemy_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy/systems"
	input_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/input/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	player_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/player/systems"
	sprite_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

var Update = *system.
	Slice().
	Register(input_systems.Update.HandleJourneyStart()).
	Register(player_systems.Update.HandleMovement()).
	Register(sprite_systems.Update.UpdatePlayerSprite()).
	Register(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			sprite_systems.Update.UpdateEnemySprite(enemy).Apply(w, m)
			enemy_systems.Update.WatchAggro(enemy).Apply(w, m)
			enemy_systems.Update.RoamMindlessly(enemy).Apply(w, m)
			enemy_systems.Update.ChasePlayer(enemy).Apply(w, m)
		}
	})

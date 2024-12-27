package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	meta_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/meta/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	renderer_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/renderer/systems"
	sprite_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

var Draw = *system.
	Slice().
	Register(meta_systems.Draw.UpdateWindowSize()).
	Register(meta_systems.Draw.UpdateWindowOnResize()).
	Register(renderer_systems.Draw.Clear()).
	Register(renderer_systems.Draw.DrawTitleScreen()).
	Register(func(w *world.World, m *meta.Meta) {
		sprite_systems.Draw.Draw(&w.Player.SpriteMap, &w.Player.SpriteKey).Apply(w, m)
		renderer_systems.Draw.DrawHitboxesInExplorationMode(&w.Player.Hitbox).Apply(w, m)
		for _, enemy := range w.Enemies {
			renderer_systems.Draw.DrawHitboxesInExplorationMode(&enemy.Hitbox).Apply(w, m)
			renderer_systems.Draw.DrawAggroInExplorationMode(
				&enemy.Hitbox,
				&enemy.Aggro,
				func() *position.PositionComponent{
					if enemy.Aggro.Aggro {
						return &w.Player.Hitbox.Position
					}
					return &enemy.Roam.Where
				}(),
			).Apply(w, m)
			sprite_systems.Draw.Draw(&enemy.SpriteMap, &enemy.SpriteKey).Apply(w, m)
		}
	}).
	Register(meta_systems.Draw.UpdateDeltaTime()).
	Register(meta_systems.Draw.UpdateFrame())

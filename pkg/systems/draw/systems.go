package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

var Systems system.SystemSlice = *system.
	Slice().
	Register(Meta.UpdateWindowSize()).
	Register(Meta.UpdateWindowOnResize()).
	Register(Renderer.Clear()).
	Register(Renderer.DrawTitleScreen()).
	Register(func(w *world.World, m *meta.Meta) {
		Renderer.DrawHitboxesInExplorationMode(&w.Player.Hitbox).Apply(w, m)
		Renderer.DrawPlayerSprite().Apply(w, m)
		for _, enemy := range w.Enemies {
			Renderer.DrawHitboxesInExplorationMode(&enemy.Hitbox).Apply(w, m)
			Renderer.DrawAggroInExplorationMode(
				&enemy.Hitbox,
				&enemy.Aggro,
				func() *components.Position {
					if enemy.Aggro.Aggro {
						return &w.Player.Position
					}
					return &enemy.Roam.Where
				}(),
			).Apply(w, m)
		}
	}).
	Register(Meta.UpdateDeltaTime()).
	Register(Meta.UpdateFrame())

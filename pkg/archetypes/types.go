package archetypes

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type MoveFn func(w *meta.Window, p *components.Position, s *components.Speed)

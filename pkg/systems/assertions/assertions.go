package systems

import (
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type assertion struct{}

var Assertions assertion = assertion{}

func (_ *assertion) Apply() *system.System {
	return system.Create(func(w *world.World, m *meta.Meta) {})
}

package archetypes

import (
	"github.com/google/uuid"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type id struct{}

var Id = id{}

func (_ *id) Create() components.Id {
	return components.Id(uuid.New().String())
}

package system

import (
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

type System func(s *world.World, m *meta.Meta)
type SystemSlice []System

func (s *SystemSlice) Register(system *System) *SystemSlice {
	*s = append(*s, *system)
	return s
}

func Slice() *SystemSlice {
	return &SystemSlice{}
}

func Create(system System) *System {
	return &system
}

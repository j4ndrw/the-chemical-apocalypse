package aggro

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/vision"


type AggroComponent struct {
	Aggro  bool
	Radius int32
	Vision vision.VisionComponent
}

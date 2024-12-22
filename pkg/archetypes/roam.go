package archetypes

import (
	"math/rand"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
)

type roam struct{}

var Roam = roam{}

func (_ *roam) ResetTicker(roam *components.Roam) {
	roam.Duration = 0
	roam.ElapsedTime = 0
	if roam.Ticker.Ticker != nil {
		roam.Ticker.Lock()
		roam.Ticker.Stop()
		roam.Ticker.Unlock()
		roam.Ticker.Ticker = nil
	}
}

func (_ *roam) StartTicker(roam *components.Roam, window *meta.Window) {
	roam.ElapsedTime = 0
	roam.Duration = time.Duration(rand.Intn(int(roam.MaxDuration)) + 1)
	roam.Where = Position.RandomPointOnMap(window)
	roam.Ticker.Lock()
	roam.Ticker.Ticker = time.NewTicker(time.Second)
	roam.Ticker.Unlock()
}

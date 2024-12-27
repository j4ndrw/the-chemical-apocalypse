package roam

import (
	"math/rand"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
)

type archetype struct{}

var Archetype = archetype{}

func (_ *archetype) ResetTicker(roam *RoamComponent) {
	roam.Duration = 0
	roam.ElapsedTime = 0
	if roam.Ticker.Ticker != nil {
		roam.Ticker.Lock()
		roam.Ticker.Stop()
		roam.Ticker.Unlock()
		roam.Ticker.Ticker = nil
	}
}

func (_ *archetype) StartTicker(roam *RoamComponent, window *meta.Window) {
	roam.ElapsedTime = 0
	roam.Duration = time.Duration(rand.Intn(int(roam.MaxDuration)) + 1)
	roam.Where = position.Archetype.RandomPointOnMap(window)
	roam.Ticker.Lock()
	roam.Ticker.Ticker = time.NewTicker(time.Second)
	roam.Ticker.Unlock()
}

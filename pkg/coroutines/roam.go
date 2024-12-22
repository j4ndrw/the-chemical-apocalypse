package coroutines

import (
	"fmt"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/async"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/components"
)

type roam struct{}

var Roam = roam{}

func (_ *roam) Tick(id *components.Id, roam *components.Roam) async.Coroutine {
	return async.Create(
		fmt.Sprintf("%s-%s", string(*id), "Roam"),
		func(done chan bool) {
			if roam.Ticker.Ticker == nil {
				done <- true
				return
			}

			for {
				select {
				case <-roam.Ticker.C:
					roam.ElapsedTime += time.Second
					if roam.ElapsedTime >= roam.Duration {
						roam.Ticker.Lock()
						roam.Ticker.Stop()
						roam.Ticker.Unlock()
						roam.Ticker.Ticker = nil
						done <- true
						return
					}
				}
			}
		},
	)
}

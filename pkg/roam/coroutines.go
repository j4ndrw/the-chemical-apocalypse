package roam

import (
	"fmt"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/async"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
)

type coroutine struct{}

var Coroutine = coroutine{}

func (_ *coroutine) Tick(id *id.IdComponent, roam *RoamComponent) async.Coroutine {
	return async.Create(
		fmt.Sprintf("%s-%s", string(*id), constants.Keys.Roam),
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

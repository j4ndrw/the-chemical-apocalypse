package coroutines

import (
	"fmt"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/async"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/entities"
)

type enemy struct{}

var Enemy = enemy{}

func (_ *enemy) Roam(enemy *entities.Enemy) async.Coroutine {
	return async.Create(
		fmt.Sprintf("%s-%s", string(enemy.Id), "Roam"),
		func(done chan bool) {
		if enemy.Roam.Ticker == nil {
			done <- true
			return
		}

		for {
			select {
			case <-enemy.Roam.Ticker.C:
				enemy.Roam.ElapsedTime += time.Second
				if enemy.Roam.ElapsedTime >= enemy.Roam.Duration {
					enemy.Roam.Ticker.Stop()
					enemy.Roam.Ticker = nil
					done <- true
					return
				}
			}
		}
	},
	)
}

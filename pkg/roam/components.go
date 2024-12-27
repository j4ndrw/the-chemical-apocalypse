package roam

import (
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/pkg/position"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/ticker"
)

type RoamComponent struct {
	Where       position.PositionComponent
	MaxDuration time.Duration
	Duration    time.Duration
	Ticker      ticker.TickerComponent
	ElapsedTime time.Duration
}

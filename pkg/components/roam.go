package components

import "time"

type Roam struct {
	Where       Position
	MaxDuration time.Duration
	Duration    time.Duration
	Ticker      Ticker
	ElapsedTime time.Duration
}

package components

import "time"

type Roam struct {
	Where       Position
	MaxDuration time.Duration
	Duration    time.Duration
	Ticker      *time.Ticker
	ElapsedTime time.Duration
	Started     bool
}

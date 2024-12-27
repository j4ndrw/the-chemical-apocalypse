package ticker

import (
	"sync"
	"time"
)

type TickerComponent struct {
	*time.Ticker
	sync.Mutex
}

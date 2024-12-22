package components

import (
	"sync"
	"time"
)

type Ticker struct {
	*time.Ticker
	sync.Mutex
}

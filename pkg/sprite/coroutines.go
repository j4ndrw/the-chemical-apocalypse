package sprite

import (
	"fmt"
	"math"
	"time"

	"github.com/j4ndrw/the-chemical-apocalypse/internal/async"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/id"
)

type coroutine struct{}

var Coroutine = coroutine{}

func (_ *coroutine) Animate(
	id *id.IdComponent,
	key string,
	spriteMap *SpriteMapComponent,
	maxVariants int32,
	ticker *time.Ticker,
) async.Coroutine {
	return async.Create(
		fmt.Sprintf("%s-%s-%s", string(*id), "Sprite", key),
		func(done chan bool) {
			if ticker == nil {
				done <- true
				return
			}

			for {
				select {
				case <-ticker.C:
					spriteMap.Lock()
					sprite := spriteMap.Map[key]
					offset := float32(math.Abs(float64(sprite.Src.Width)))
					sprite.Src.X += offset
					if int(sprite.Src.X) >= int(maxVariants*int32(offset)) {
						sprite.Src.X = 0
					}
					spriteMap.Map[key] = sprite
					spriteMap.Unlock()
				}
			}
		},
	)
}

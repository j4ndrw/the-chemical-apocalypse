package systems

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
)

var Systems system.SystemSlice = *system.
	Slice().
	Register(Meta.SetConfigFlags(
		rl.FlagWindowResizable |
			rl.FlagWindowTopmost,
	)).
	Register(Meta.InitWindow()).
	Register(Meta.SetTargetFPS(60)).
	Register(
		Meta.LoadFont(
			fmt.Sprintf(
				"%s/assets/fonts/lunchtime-doubly-so/lunchds.ttf",
				utils.GetCwd(),
			),
		),
	).
	Register(WorldMode.Init()).
	Register(Player.CreatePlayer()).
	Register(Enemy.CreateEnemies(
		1,
		100,
		100,
		2,
		5,
		300,
		10,
	)).
	Register(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			Enemy.PlaceEnemyInCenter(enemy).Apply(w, m)
		}
	})

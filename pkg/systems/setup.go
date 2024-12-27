package systems

import (
	"fmt"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/system"
	"github.com/j4ndrw/the-chemical-apocalypse/internal/utils"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/constants"
	enemy_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/enemy/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/meta"
	meta_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/meta/systems"
	player_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/player/systems"
	sprite_systems "github.com/j4ndrw/the-chemical-apocalypse/pkg/sprite/systems"
	"github.com/j4ndrw/the-chemical-apocalypse/pkg/world"
	worldmode "github.com/j4ndrw/the-chemical-apocalypse/pkg/worldmode/systems"
)

var Setup = *system.
	Slice().
	Register(meta_systems.Setup.SetConfigFlags(
		rl.FlagWindowResizable |
			rl.FlagWindowTopmost,
	)).
	Register(meta_systems.Setup.InitWindow()).
	Register(meta_systems.Setup.SetTargetFPS(60)).
	Register(
		meta_systems.Setup.LoadFont(
			fmt.Sprintf(
				"%s/assets/fonts/lunchtime-doubly-so/lunchds.ttf",
				utils.GetCwd(),
			),
		),
	).
	Register(worldmode.Setup.Init()).
	Register(player_systems.Setup.CreatePlayer()).
	Register(enemy_systems.Setup.CreateEnemies(
		1,
		256,
		256,
		200,
		700,
		500,
		135,
		10*time.Second,
	)).
	Register(func(w *world.World, m *meta.Meta) {
		for _, enemy := range w.Enemies {
			enemy_systems.Setup.PlaceEnemyInCenter(enemy).Apply(w, m)
		}
	}).
	Register(
		meta_systems.Setup.LoadPlayerSpriteAtlas(
			fmt.Sprintf(
				"%s/assets/sprites/the-chemical-apocalypse-atlas.png",
				utils.GetCwd(),
			),
		),
	).
	Register(func(w *world.World, m *meta.Meta) {
		sprite_systems.Setup.Setup(
			&w.Player.SpriteMap,
			&w.Player.Hitbox,
			constants.Keys.PlayerIdle,
			constants.Keys.PlayerMoveUp,
			constants.Keys.PlayerMoveDown,
			constants.Keys.PlayerMoveForward,
		).Apply(w, m)
		for _, enemy := range w.Enemies {
			sprite_systems.Setup.Setup(
				&enemy.SpriteMap,
				&enemy.Hitbox,
				constants.Keys.AntiPeanutIdle,
				constants.Keys.AntiPeanutMoveUp,
				constants.Keys.AntiPeanutMoveDown,
				constants.Keys.AntiPeanutMoveForward,
			).Apply(w, m)
		}
	})

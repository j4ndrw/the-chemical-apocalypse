package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

var Systems system.SystemSlice = *system.
	Slice().
	Register(Player.CreatePlayer()).
	Register(Enemy.CreateEnemies(
		1,
		100,
		100,
		1,
		3,
		300,
		15,
	)).
	Register(Enemy.PlaceEnemiesInCenter())

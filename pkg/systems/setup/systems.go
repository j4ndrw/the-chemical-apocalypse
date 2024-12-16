package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

var Systems system.SystemSlice = *system.
	Slice().
	Register(Enemy.CreateEnemies(
		10,
		15,
		15,
		1,
		3,
		300,
		15,
	)).
	Register(Enemy.PlaceEnemiesInCenter())

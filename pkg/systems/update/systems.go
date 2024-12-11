package systems

import "github.com/j4ndrw/the-chemical-apocalypse/internal/system"

func Systems() system.SystemSlice {
	player := Player()
	enemy := Enemy()
	return *system.Slice().
		Register(&player.MoveDown).
		Register(&player.MoveRight).
		Register(&enemy.MoveDown).
		Register(&enemy.MoveRight)
}

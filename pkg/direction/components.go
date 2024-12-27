package direction

const (
	None = 0
	Up = -1
	Down = 1
	Left = -1
	Right = 1
)

type DirectionComponent struct{
	X, Y int
}

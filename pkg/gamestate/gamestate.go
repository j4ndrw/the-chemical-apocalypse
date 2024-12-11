package gamestate

type Position struct {
	X, Y int32
}

type Player struct {
	Position Position
}

type State struct {
	Player Player
}

func New() State {
	return State{
		Player: Player{
			Position: Position{X: 0, Y: 0},
		},
	}
}

func (p *Player) MoveLeft(step int32) {
	p.Position.X -= step
}

func (p *Player) MoveRight(step int32) {
	p.Position.X += step
}

func (p *Player) MoveUp(step int32) {
	p.Position.Y -= step
}

func (p *Player) MoveDown(step int32) {
	p.Position.Y += step
}

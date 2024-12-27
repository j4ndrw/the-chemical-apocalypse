package pathfinding

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/position"

type archetype struct{}

var Archetype = archetype{}

type node struct {
	Position *position.PositionComponent
	distance float64
}

func (_ *archetype) euclidianDistance(a, b *position.PositionComponent) float64 {
	return float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (p *archetype) getNeighbors(
	pos *position.PositionComponent,
	mapNeighbor func(pos *position.PositionComponent, direction *position.PositionComponent) *position.PositionComponent,
	isNeighborValid func(pos *position.PositionComponent) bool,
) []position.PositionComponent {
	directions := []position.PositionComponent{
		{X: 0, Y: 1},   // Up
		{X: 1, Y: 0},   // Right
		{X: 0, Y: -1},  // Down
		{X: -1, Y: 0},  // Left
		{X: -1, Y: 1},  // Top left
		{X: 1, Y: 1},   // Top right
		{X: -1, Y: -1}, // Bottom left
		{X: 1, Y: -1},  // Bottom right
	}

	var neighbors []position.PositionComponent
	for _, dir := range directions {
		newPos := mapNeighbor(pos, &dir)
		if isNeighborValid(newPos) {
			neighbors = append(neighbors, *newPos)
		}
	}

	return neighbors
}

func (p *archetype) findNearestNeighbor(
	pos, goal *position.PositionComponent,
	closest *node,
) *node {
	neighbor := &node{Position: pos, distance: p.euclidianDistance(pos, goal)}
	if closest == nil || neighbor.distance < closest.distance {
		return neighbor
	}
	return closest
}

func (p *archetype) ClosestNeighbor(
	current, goal *position.PositionComponent,
	mapNeighbor func(
		pos *position.PositionComponent,
		direction *position.PositionComponent,
	) *position.PositionComponent,
	isNeighborValid func(pos *position.PositionComponent) bool,
) *node {
	if current.X == goal.X && current.Y == goal.Y {
		return nil
	}

	var closest *node = nil
	for _, neighbor := range p.getNeighbors(
		current,
		mapNeighbor,
		isNeighborValid,
	) {
		closest = p.findNearestNeighbor(
			&neighbor,
			goal,
			closest,
		)
	}
	return closest
}

package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type pathfinding struct{}

var PathFinding = pathfinding{}

type node struct {
	Position *components.Position
	distance float64
}

func (_ *pathfinding) euclidianDistance(a, b *components.Position) float64 {
	return float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (p *pathfinding) getNeighbors(
	pos *components.Position,
	mapNeighbor func(position *components.Position, direction *components.Position) *components.Position,
	isNeighborValid func(position *components.Position) bool,
) []components.Position {
	directions := []components.Position{
		{X: 0, Y: 1},   // Up
		{X: 1, Y: 0},   // Right
		{X: 0, Y: -1},  // Down
		{X: -1, Y: 0},  // Left
		{X: -1, Y: 1},  // Top left
		{X: 1, Y: 1},   // Top right
		{X: -1, Y: -1}, // Bottom left
		{X: 1, Y: -1},  // Bottom right
	}

	var neighbors []components.Position
	for _, dir := range directions {
		newPos := mapNeighbor(pos, &dir)
		if isNeighborValid(newPos) {
			neighbors = append(neighbors, *newPos)
		}
	}

	return neighbors
}

func (p *pathfinding) findNearestNeighbor(
	position *components.Position,
	goal *components.Position,
	closest *node,
) *node {
	neighbor := &node{Position: position, distance: p.euclidianDistance(position, goal)}
	if closest == nil || neighbor.distance < closest.distance {
		return neighbor
	}
	return closest
}

func (p *pathfinding) ClosestNeighbor(
	current, goal *components.Position,
	mapNeighbor func(
		position *components.Position,
		direction *components.Position,
	) *components.Position,
	isNeighborValid func(position *components.Position) bool,
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

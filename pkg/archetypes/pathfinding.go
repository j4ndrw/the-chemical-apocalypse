package archetypes

import "github.com/j4ndrw/the-chemical-apocalypse/pkg/components"

type pathfinding struct{}

var PathFinding pathfinding = pathfinding{}

type node struct {
	Position      *components.Vector2
	actualcost    float64
	heuristiccost float64
	totalcost     float64
	parent        *node
}

func (_ *pathfinding) heuristic(a, b *components.Vector2) float64 {
	return float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (p *pathfinding) getNeighbors(
	pos *components.Vector2,
	mapNeighbor func(position *components.Vector2, direction *components.Vector2) *components.Vector2,
	isNeighborValid func(position *components.Vector2) bool,
) []components.Vector2 {
	directions := []components.Vector2{
		{X: 0, Y: 1},   // Up
		{X: 1, Y: 0},   // Right
		{X: 0, Y: -1},  // Down
		{X: -1, Y: 0},  // Left
		{X: -1, Y: 1},  // Top left
		{X: 1, Y: 1},   // Top right
		{X: -1, Y: -1}, // Bottom left
		{X: 1, Y: -1},  // Bottom right
	}

	var neighbors []components.Vector2
	for _, dir := range directions {
		newPos := mapNeighbor(pos, &dir)
		if isNeighborValid(newPos) {
			neighbors = append(neighbors, *newPos)
		}
	}

	return neighbors
}

func (p *pathfinding) findNearestNeighbor(
	position *components.Vector2,
	goal *components.Vector2,
	actualcost float64,
	closestNeighbor *node,
) *node {
	neighbor := &node{Position: position}
	if actualcost < neighbor.actualcost || neighbor.actualcost == 0 {
		neighbor.actualcost = actualcost
		neighbor.heuristiccost = p.heuristic(position, goal)
		neighbor.totalcost = neighbor.actualcost + neighbor.heuristiccost
	}

	if closestNeighbor == nil || neighbor.totalcost < closestNeighbor.totalcost {
		return neighbor
	}
	return closestNeighbor
}

func (p *pathfinding) ClosestNeighbor(
	start, goal *components.Vector2,
	mapNeighbor func(
		position *components.Vector2,
		direction *components.Vector2,
	) *components.Vector2,
	isNeighborValid func(position *components.Vector2) bool,
) *node {
	current := &node{Position: start, actualcost: 0, heuristiccost: p.heuristic(start, goal)}
	current.totalcost = current.actualcost + current.heuristiccost

	if start.X == goal.X && start.Y == goal.Y {
		return nil
	}

	neighbors := p.getNeighbors(current.Position, mapNeighbor, isNeighborValid)

	var closestNeighbor *node = nil
	for _, position := range neighbors {
		closestNeighbor = p.findNearestNeighbor(
			&position,
			goal,
			current.actualcost+1,
			closestNeighbor,
		)
	}

	if closestNeighbor == nil {
		return nil
	}
	return closestNeighbor
}

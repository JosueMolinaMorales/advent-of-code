package sixteen

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

type State struct {
	Cost     int
	Facing   types.Direction
	Position types.Vector
	Path     []types.Vector
}

func SolveDay16() {
	start, end, grid := parseGrid("./inputs/day_16.txt")
	fmt.Println("Day 16 Part 1: ", dijkstra(start, end, grid))
	fmt.Println("Day 16 Part 2: ", dijkstraUniqueCells(start, end, grid))
}

func parseGrid(path string) (types.Vector, types.Vector, [][]string) {
	input, err := util.LoadFileAsString(path)
	if err != nil {
		panic(err)
	}

	var start, end types.Vector
	grid := [][]string{}

	for i, line := range strings.Split(input, "\n") {
		row := []string{}
		for j, col := range strings.Split(line, "") {
			if col == "S" {
				start = *types.NewVector(i, j)
			}
			if col == "E" {
				end = *types.NewVector(i, j)
			}
			row = append(row, col)
		}
		grid = append(grid, row)
	}
	return start, end, grid
}

func dijkstra(start, end types.Vector, grid [][]string) int {
	dist := map[types.Vector]int{}
	for i, row := range grid {
		for j, col := range row {
			if col != "#" {
				dist[*types.NewVector(i, j)] = math.MaxInt
			}
		}
	}

	compareCosts := func(a, b interface{}) int {
		return -utils.IntComparator(a.(State).Cost, b.(State).Cost)
	}
	q := priorityqueue.NewWith(compareCosts)

	dist[start] = 0
	q.Enqueue(State{Cost: 0, Position: start, Facing: types.DIRECTION_EAST})
	minCost := math.MaxInt

	for !q.Empty() {
		curr, _ := q.Dequeue()
		state := curr.(State)

		if state.Position == end {
			minCost = min(minCost, state.Cost)
			continue
		}

		if state.Cost > dist[state.Position] || state.Cost >= minCost {
			continue
		}

		for _, neighbor := range adjacent(state.Position, grid) {
			newCost := calculateCost(state, neighbor)
			if newCost < dist[neighbor] {
				dist[neighbor] = newCost
				q.Enqueue(State{
					Cost:     newCost,
					Position: neighbor,
					Facing:   determineDirection(state.Position, neighbor),
				})
			}
		}
	}

	return minCost
}

func dijkstraUniqueCells(start, end types.Vector, grid [][]string) int {
	dist := map[string]int{}
	minCost := dijkstra(start, end, grid)

	for i, row := range grid {
		for j, col := range row {
			if col != "#" {
				for _, dir := range []types.Direction{types.DIRECTION_EAST, types.DIRECTION_NORTH, types.DIRECTION_SOUTH, types.DIRECTION_WEST} {
					dist[toString(State{Position: *types.NewVector(i, j), Facing: dir})] = minCost
				}
			}
		}
	}

	compareCosts := func(a, b interface{}) int {
		return -utils.IntComparator(a.(State).Cost, b.(State).Cost)
	}
	q := priorityqueue.NewWith(compareCosts)
	startState := State{Cost: 0, Position: start, Facing: types.DIRECTION_EAST}
	q.Enqueue(startState)
	dist[toString(startState)] = 0
	uniquePoints := hashset.New()

	for !q.Empty() {
		curr, _ := q.Dequeue()
		state := curr.(State)
		state.Path = append(state.Path, state.Position)

		if state.Cost > dist[toString(state)] || state.Cost > minCost {
			continue
		}

		if state.Position == end && state.Cost == minCost {
			for _, point := range state.Path {
				uniquePoints.Add(point)
			}
			continue
		}

		for _, neighbor := range adjacent(state.Position, grid) {
			newState := generateNextState(state, neighbor)
			if newState.Cost <= dist[toString(newState)] {
				dist[toString(newState)] = newState.Cost
				q.Enqueue(newState)
			}
		}
	}

	return uniquePoints.Size()
}

func calculateCost(state State, neighbor types.Vector) int {
	cost := state.Cost + 1
	if state.Facing != determineDirection(state.Position, neighbor) {
		cost += 1000
	}
	return cost
}

func generateNextState(state State, neighbor types.Vector) State {
	newDir := determineDirection(state.Position, neighbor)
	newCost := calculateCost(state, neighbor)
	return State{
		Cost:     newCost,
		Position: neighbor,
		Facing:   newDir,
		Path:     slices.Clone(state.Path),
	}
}

func determineDirection(from, to types.Vector) types.Direction {
	switch {
	case to.X == from.X && to.Y > from.Y:
		return types.DIRECTION_EAST
	case to.X == from.X && to.Y < from.Y:
		return types.DIRECTION_WEST
	case to.X < from.X:
		return types.DIRECTION_NORTH
	default:
		return types.DIRECTION_SOUTH
	}
}

func adjacent(point types.Vector, grid [][]string) []types.Vector {
	directions := []types.Vector{
		{X: 0, Y: 1}, {X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: -1},
	}

	neighbors := []types.Vector{}
	for _, dir := range directions {
		x, y := point.X+dir.X, point.Y+dir.Y
		if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) && grid[x][y] != "#" {
			neighbors = append(neighbors, *types.NewVector(x, y))
		}
	}
	return neighbors
}

func toString(s State) string {
	return fmt.Sprintf("%d,%d,%s", s.Position.X, s.Position.Y, s.Facing)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

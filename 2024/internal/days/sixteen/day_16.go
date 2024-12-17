package sixteen

import (
	"fmt"
	"math"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type State struct {
	Cost     int
	Facing   types.Direction
	Position types.Vector
}

func SolveDay16() {
	fmt.Println(solvePartOne())
}

func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_16.txt")
	if err != nil {
		panic(err)
	}
	grid := [][]string{}
	start := *types.NewVector()
	end := *types.NewVector()

	for i, line := range strings.Split(input, "\n") {
		row := []string{}
		for j, col := range strings.Split(line, "") {
			if col == "S" {
				start.X = i
				start.Y = j
			}
			if col == "E" {
				end.X = i
				end.Y = j
			}
			row = append(row, col)
		}

		grid = append(grid, row)
	}

	return dijkstra(start, end, grid)
}

func dijkstra(start types.Vector, end types.Vector, g [][]string) int {
	dist := map[types.Vector]int{}

	for i, row := range g {
		for j, col := range row {
			if col == "#" {
				continue
			}
			dist[*types.NewVector(i, j)] = math.MaxInt
		}
	}
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		aCost := a.(State).Cost
		bCost := a.(State).Cost
		return -utils.IntComparator(aCost, bCost)
	})
	minCost := math.MaxInt

	dist[start] = 0
	q.Enqueue(State{
		Cost:     0,
		Position: start,
		Facing:   types.DIRECTION_EAST,
	})

	for !q.Empty() {
		s, _ := q.Dequeue()
		p := s.(State)
		if p.Position == end && p.Cost < minCost {
			minCost = p.Cost
			continue
		}

		if p.Cost > dist[p.Position] || p.Cost >= minCost {
			continue
		}

		for _, neighbor := range adjacent(p.Position, g) {
			cost := p.Cost
			dx, dy := neighbor.X-p.Position.X, neighbor.Y-p.Position.Y
			newDir := p.Facing
			if dx == 0 && dy == 1 && p.Facing != types.DIRECTION_EAST {
				// To the right
				cost += 1000
				newDir = types.DIRECTION_EAST
			} else if dx == 0 && dy == -1 && p.Facing != types.DIRECTION_WEST {
				cost += 1000
				newDir = types.DIRECTION_WEST
			} else if dx == -1 && dy == 0 && p.Facing != types.DIRECTION_NORTH {
				cost += 1000
				newDir = types.DIRECTION_NORTH
			} else if dx == 1 && dy == 0 && p.Facing != types.DIRECTION_SOUTH {
				cost += 1000
				newDir = types.DIRECTION_SOUTH
			}
			cost += 1

			if cost < dist[neighbor] {
				dist[neighbor] = cost
				q.Enqueue(State{
					Cost:     cost,
					Facing:   newDir,
					Position: neighbor,
				})
			}
		}
	}

	return minCost
}

func printMap(state State, m [][]string) {
	for i, row := range m {
		for j, col := range row {
			if *types.NewVector(i, j) == state.Position {
				switch state.Facing {
				case types.DIRECTION_EAST:
					fmt.Print(">")
				case types.DIRECTION_NORTH:
					fmt.Print("^")
				case types.DIRECTION_SOUTH:
					fmt.Print("v")
				case types.DIRECTION_WEST:
					fmt.Print("<")
				}
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func adjacent(point types.Vector, m [][]string) []types.Vector {
	directions := [][]int{
		{0, 1}, // Right

		{1, 0},  // Up
		{-1, 0}, // Down
		{0, -1}, // Left
	}

	adj := []types.Vector{}
	for _, dir := range directions {
		dx, dy := point.X+dir[0], point.Y+dir[1]
		if dx < 0 || dx >= len(m) || dy < 0 || dy >= len(m[0]) {
			continue
		}
		if m[dx][dy] == "#" {
			continue
		}
		adj = append(adj, *types.NewVector(dx, dy))
	}

	return adj
}

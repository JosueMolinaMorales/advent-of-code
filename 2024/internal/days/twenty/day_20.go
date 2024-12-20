package twenty

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

func SolveDayTwenty() {
	fmt.Println("Day 20 Part 1: ", race(2))
	fmt.Println("Day 20 Part 2: ", race(20))
}

func setup() ([][]string, types.Vector, types.Vector) {
	input, err := util.LoadFileAsString("./inputs/day_20.txt")
	if err != nil {
		panic(err)
	}

	racetrack := [][]string{}
	start := *types.NewVector()
	end := *types.NewVector()
	for i, line := range strings.Split(input, "\n") {
		row := []string{}
		for j, ch := range strings.Split(line, "") {
			v := *types.NewVector(i, j)
			if ch == "S" {
				start = v
			} else if ch == "E" {
				end = v
			}
			row = append(row, ch)
		}
		racetrack = append(racetrack, row)
	}

	return racetrack, start, end
}

func race(cheat int) int {
	grid, start, end := setup()
	path := regPathFind(start, end, grid)

	c := 0
	save := 100
	// Iterate through the first len(path)-save paths to find the cheat start
	for i, p := range path[:len(path)-save] {
		// jump ahead save steps and see if we can reach it by cheating
		for j, q := range path[i+save:] {
			if d := q.ManhanttanDistance(p); d <= cheat && d <= j {
				c++
			}
		}
	}
	return c
}

type State struct {
	Time int
	Pos  types.Vector
}

func regPathFind(start, end types.Vector, grid [][]string) []types.Vector {
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		aTime := a.(State).Time
		bTime := b.(State).Time
		return utils.IntComparator(aTime, bTime)
	})
	q.Enqueue(State{
		Time: 0,
		Pos:  start,
	})
	visited := hashset.New(start)
	prev := map[types.Vector]types.Vector{}
	for !q.Empty() {
		p, _ := q.Dequeue()
		state := p.(State)
		if state.Pos == end {
			break
		}

		for _, neighbor := range adjacent(state.Pos, grid) {
			if visited.Contains(neighbor) {
				continue
			}
			visited.Add(neighbor)
			prev[neighbor] = state.Pos
			q.Enqueue(State{
				Time: state.Time + 1,
				Pos:  neighbor,
			})
		}

	}

	path := []types.Vector{end}
	curr := end
	for curr != start {
		curr = prev[curr]
		path = append(path, curr)
	}

	return path
}

func adjacent(point types.Vector, grid [][]string) []types.Vector {
	directions := []types.Vector{
		{X: 0, Y: 1}, {X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: -1},
	}

	neighbors := []types.Vector{}
	for _, dir := range directions {
		x, y := point.X+dir.X, point.Y+dir.Y
		n := *types.NewVector(x, y)
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			continue
		}
		if grid[x][y] != "#" {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func printMap(m [][]string, path []types.Vector) {
	// Add path
	for _, cell := range path {
		m[cell.X][cell.Y] = "O"
	}
	for _, row := range m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

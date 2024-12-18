package eighteen

import (
	"fmt"
	"math"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

const (
	ROW_LEN   = 71
	COL_LEN   = 71
	KB_FALLEN = 1024
)

var (
	START = *types.NewVector(0, 0)
	END   = *types.NewVector(ROW_LEN-1, COL_LEN-1)
)

func SolveDay18() {
	fmt.Println("Day 18 Part 1: ", solvePartOne())
	fmt.Println("Day 18 Part 2: ", solvePartTwo())
}

func setup() ([]types.Vector, [][]string) {
	input, err := util.LoadFileAsString("./inputs/day_18.txt")
	if err != nil {
		panic(err)
	}
	coords := []types.Vector{}
	for _, coord := range strings.Split(input, "\n") {
		parts := strings.Split(coord, ",")
		coords = append(coords, *types.NewVector(util.ToInt(parts[1]), util.ToInt(parts[0])))
	}

	grid := [][]string{}
	for i := 0; i < ROW_LEN; i++ {
		row := []string{}
		for j := 0; j < COL_LEN; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	return coords, grid
}

func solvePartOne() int {
	coords, grid := setup()
	for i := 0; i < KB_FALLEN; i++ {
		block := coords[i]
		// Add the block
		grid[block.X][block.Y] = "#"
	}
	return bfs(START, END, grid)
}

func solvePartTwo() string {
	coords, grid := setup()
	found := types.Vector{}

	left := float64(KB_FALLEN)
	right := float64(len(coords))

	for left < right {
		m := int(math.Floor((left + right) / 2.0))
		found = coords[m]
		for i := KB_FALLEN; i < m; i++ {
			block := coords[i]
			// Add the block
			grid[block.X][block.Y] = "#"
		}
		dist := bfs(START, END, grid)
		if dist == math.MaxInt {
			right = float64(m)
		} else {
			left = float64(m) + 1
		}
		for i := m; i > KB_FALLEN; i-- {
			block := coords[i]
			// Add the block
			grid[block.X][block.Y] = "."
		}
	}

	return fmt.Sprintf("%d,%d", found.Y, found.X)
}

type State struct {
	Dist int
	Pos  types.Vector
}

func bfs(start types.Vector, end types.Vector, m [][]string) int {
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		aDist := a.(State).Dist
		bDist := b.(State).Dist
		return utils.IntComparator(aDist, bDist)
	})
	visited := hashset.New(start)
	q.Enqueue(State{
		Dist: 0,
		Pos:  start,
	})

	for !q.Empty() {
		p, _ := q.Dequeue()
		state := p.(State)

		if state.Pos == end {
			return state.Dist
		}

		for _, neighbor := range adjacent(state.Pos, m) {
			if visited.Contains(neighbor) {
				continue
			}
			visited.Add(neighbor)

			q.Enqueue(State{
				Dist: state.Dist + 1,
				Pos:  neighbor,
			})
		}

	}
	return math.MaxInt
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

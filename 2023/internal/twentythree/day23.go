package twentythree

import (
	"math"
	"os"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`

func RunDayTwentyThree() {
	input, err := os.ReadFile("./input/day23.txt")
	if err != nil {
		panic("Failed to read day 23 input file")
	}
	println("Day 23 Part 1:", partOne(string(input)))
	println("Day 23 Part 2:", partTwo(string(input)))
}

func partTwo(input string) int {
	grid, _, _ := parse(input)
	start := Point{0, iterators.IndexOf(grid[0], ".")}
	end := Point{len(grid) - 1, iterators.IndexOf(grid[len(grid)-1], ".")}
	graph := condenseMap(grid, start, end)

	visited := make(map[Point]bool, 0)
	result := dfsGraph(graph, start, end, visited)
	return result
}

type State struct {
	n  int
	pt Point
}

func partOne(input string) int {
	m, s, e := parse(input)
	return longestSimplePath(m, s, e, false)
}

var (
	Right = Point{0, 1}
	Left  = Point{0, -1}
	Up    = Point{-1, 0}
	Down  = Point{1, 0}
)

type Graph = map[Point]map[Point]int

func condenseMap(grid [][]string, start, end Point) Graph {
	points := []Point{
		start, end,
	}
	for r, row := range grid {
		for c, col := range row {
			if col == "#" {
				continue
			}
			neighbors := 0
			for _, dir := range []Point{{r - 1, c}, {r + 1, c}, {r, c - 1}, {r, c + 1}} {
				if dir.X < len(grid) && dir.X >= 0 && dir.Y < len(grid[0]) && dir.Y >= 0 && grid[dir.X][dir.Y] != "#" {
					neighbors++
				}
			}
			if neighbors >= 3 {
				points = append(points, Point{X: r, Y: c})
			}
		}
	}

	graph := make(map[Point]map[Point]int, 0)
	for _, p := range points {
		graph[p] = make(map[Point]int, 0)
	}

	for _, pt := range points {
		stack := []State{{0, pt}}
		seen := make(map[Point]bool, 0)
		seen[pt] = true

		for len(stack) > 0 {
			state := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			n, r, c := state.n, state.pt.X, state.pt.Y

			if n != 0 && iterators.Contains(points, Point{r, c}) {
				graph[pt][Point{r, c}] = n
				continue
			}

			for _, delta := range []Point{RIGHT, LEFT, UP, DOWN} {
				nr, nc := r+delta.X, c+delta.Y
				if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) && grid[nr][nc] != "#" && !seen[Point{nr, nc}] {
					stack = append(stack, State{n + 1, Point{nr, nc}})
					seen[Point{nr, nc}] = true
				}
			}
		}
	}

	return graph
}

func dfsGraph(graph Graph, pt Point, end Point, visited map[Point]bool) int {
	if pt == end {
		return 0
	}

	m := math.MinInt

	visited[pt] = true
	for nx, weight := range graph[pt] {
		if !visited[nx] {
			m = max(m, dfsGraph(graph, nx, end, visited)+weight)
		}
	}
	visited[pt] = false

	return m
}

func dfs(matrix [][]string, visited [][]bool, current Point, end Point, currentPathLength int, ignoreSlopes bool, enableBacktrack bool, maxLength *int) {
	// Check boundaries and whether the cell is visited
	if visited[current.X][current.Y] {
		return
	}

	// Mark the current cell as visited
	visited[current.X][current.Y] = true

	// Check if the current cell is the destination
	if current.X == end.X && current.Y == end.Y {
		if currentPathLength > *maxLength {
			*maxLength = currentPathLength
		}
	}

	// Explore adjacent cells
	neighbors := getNeighbors(current, matrix, ignoreSlopes)
	for _, neighbor := range neighbors {
		dfs(matrix, visited, neighbor, end, currentPathLength+1, ignoreSlopes, enableBacktrack, maxLength)
	}
	// Backtrack - mark the current cell as not visited
	if enableBacktrack {
		visited[current.X][current.Y] = false
	}
}

func longestSimplePath(matrix [][]string, start, end Point, ignoreSlopes bool) int {
	rows, cols := len(matrix), len(matrix[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	maxPathLength := 0
	dfs(matrix, visited, start, end, 0, ignoreSlopes, true, &maxPathLength)
	return maxPathLength
}

type Point struct {
	X, Y int
}

func parse(in string) ([][]string, Point, Point) {
	m := make([][]string, 0)
	for _, l := range strings.Split(in, "\n") {
		row := make([]string, 0)
		for _, c := range l {
			row = append(row, string(c))
		}
		m = append(m, row)
	}
	start := Point{0, 0}
	end := Point{len(m[0]) - 1, 0}
	for y, row := range m[0] {
		if row == "." {
			start.Y = y
		}
	}
	for y, row := range m[len(m)-1] {
		if row == "." {
			end.Y = y
		}
	}
	return m, start, end
}

const (
	RIGHT_SLOPE = ">"
	LEFT_SLOPE  = "<"
	UP_SLOPE    = "^"
	DOWN_SLOPE  = "v"
)

var RIGHT, LEFT, UP, DOWN = Point{0, 1}, Point{0, -1}, Point{-1, 0}, Point{1, 0}

func getNeighbors(p Point, m [][]string, ignoreSlopes bool) []Point {
	directions := []Point{Right, Left, Up, Down} // Up, Down, Left, Right

	neighbors := make([]Point, 0)
	for _, dir := range directions {
		newX, newY := p.X+dir.X, p.Y+dir.Y
		if newX < 0 || newX >= len(m) || newY < 0 || newY >= len(m[0]) {
			continue
		}
		if m[newX][newY] == "#" {
			continue
		}
		if ignoreSlopes {
			neighbors = append(neighbors, Point{newX, newY})
			continue
		}
		switch m[p.X][p.Y] {
		case "<":
			if dir == Left {
				neighbors = append(neighbors, Point{newX, newY})
			}
		case ">":
			if dir == Right {
				neighbors = append(neighbors, Point{newX, newY})
			}
		case "v":
			if dir == Down {
				neighbors = append(neighbors, Point{newX, newY})
			}
		case "^":
			if dir == Up {
				neighbors = append(neighbors, Point{newX, newY})
			}
		default:
			neighbors = append(neighbors, Point{newX, newY})
		}

	}

	return neighbors
}

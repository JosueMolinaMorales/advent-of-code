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

type State struct {
	n  int
	pt Point
}

var (
	Right      = Point{0, 1}
	Left       = Point{0, -1}
	Up         = Point{-1, 0}
	Down       = Point{1, 0}
	Directions = map[string][]Point{
		">": {Right},
		"<": {Left},
		"^": {Up},
		"v": {Down},
		".": {Right, Left, Up, Down},
	}
)

type Graph = map[Point]map[Point]int

type Point struct {
	X, Y int
}

func partTwo(input string) int {
	grid, _, _ := parse(input)
	start := Point{0, iterators.IndexOf(grid[0], ".")}
	end := Point{len(grid) - 1, iterators.IndexOf(grid[len(grid)-1], ".")}
	graph := condenseMap(grid, start, end, true)

	visited := make(map[Point]bool, 0)
	result := dfs(graph, start, end, visited)
	return result
}

func partOne(input string) int {
	grid, start, end := parse(input)
	graph := condenseMap(grid, start, end, false)

	visited := make(map[Point]bool, 0)
	result := dfs(graph, start, end, visited)
	return result
}

func condenseMap(grid [][]string, start, end Point, ignoreSlopes bool) Graph {
	points := []Point{
		start, end,
	}
	for r, row := range grid {
		for c, col := range row {
			if col == "#" {
				continue
			}
			neighbors := 0
			for _, dir := range []Point{Right, Left, Up, Down} {
				dx := r + dir.X
				dy := c + dir.Y
				// Bound check & wall check
				if dx < 0 || dx >= len(grid) || dy < 0 || dy >= len(grid[0]) || grid[dx][dy] == "#" {
					continue
				}
				neighbors += 1
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

			directions := Directions[grid[r][c]]
			if ignoreSlopes {
				directions = Directions["."]
			}
			for _, delta := range directions {
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

func dfs(graph Graph, pt Point, end Point, visited map[Point]bool) int {
	if pt == end {
		return 0
	}

	m := math.MinInt

	visited[pt] = true
	for nx, weight := range graph[pt] {
		if !visited[nx] {
			m = max(m, dfs(graph, nx, end, visited)+weight)
		}
	}
	visited[pt] = false

	return m
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
	start := Point{0, iterators.IndexOf(m[0], ".")}
	end := Point{len(m[0]) - 1, iterators.IndexOf(m[len(m)-1], ".")}
	return m, start, end
}

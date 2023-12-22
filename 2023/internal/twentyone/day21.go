package twentyone

import (
	"fmt"
	"os"
	"strings"
)

const testInput = `...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........`

func RunDayTwentyOne() {
	input, err := os.ReadFile("./input/day21.txt")
	if err != nil {
		panic("Failed to read day 21 input")
	}
	fmt.Println("Day 21 Part one:", partOne(string(input)))
	fmt.Println("Day 21 Part two:", partTwo(string(input)))
}

func parse(input string) ([][]string, Plot) {
	garden := make([][]string, 0)
	start := Plot{0, 0, 0}
	for i, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for j, char := range line {
			row = append(row, string(char))
			if char == 'S' {
				start = Plot{i, j, 0}
			}
		}
		garden = append(garden, row)
	}

	return garden, start
}

func partTwo(input string) int {
	garden, _ := parse(input)
	expanded := expandGrid(garden, 5)
	start := Plot{len(expanded) / 2, len(expanded[0]) / 2, 0}
	size := len(garden)
	half := size / 2
	p := make([]int, 0)
	for _, n := range []int{half, half + size, half + 2*size} {
		p = append(p, bfs(expanded, start, n))
	}

	// https://en.wikipedia.org/wiki/Polynomial_regression
	a := (p[2] + p[0] - 2*p[1]) / 2
	b := p[1] - p[0] - a
	c := p[0]
	n := 202300
	result := a*n*n + b*n + c

	return result
}

func partOne(input string) int {
	garden, start := parse(input)
	return bfs(garden, start, 64)
}

type Plot struct {
	x, y, step int
}

func bfs(garden [][]string, start Plot, steps int) int {
	queue := make([]Plot, 0)
	queue = append(queue, start)
	visited := make(map[Plot]bool, 0)
	end := make(map[[2]int]bool, 0)
	plots := 0
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.step == steps {
			end[[2]int{n.x, n.y}] = true
			plots += 1
			continue
		}

		for _, adj := range adj(n, n.step, garden) {
			if !visited[adj] {
				visited[adj] = true
				queue = append(queue, adj)
			}
		}
	}
	// printGarden(garden, end)
	return plots
}

func expandGrid(graden [][]string, times int) [][]string {
	newGrid := make([][]string, 0)
	for k := 0; k < times; k++ {
		for _, row := range graden {
			newRow := make([]string, 0)
			for i := 0; i < times; i++ {
				newRow = append(newRow, row...)
			}
			newGrid = append(newGrid, newRow)
		}
	}
	return newGrid
}

func adj(node Plot, step int, garden [][]string) []Plot {
	directions := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	adj := make([]Plot, 0)
	for _, dir := range directions {
		dx, dy := node.x+dir[0], node.y+dir[1]
		// Bound Check
		if dx < 0 || dx >= len(garden) || dy < 0 || dy >= len(garden[0]) {
			continue
		}
		if garden[dx][dy] == "#" {
			continue
		}
		adj = append(adj, Plot{dx, dy, step + 1})
	}

	return adj
}

// For fun
func printGarden(garden [][]string, visited map[[2]int]bool) {
	for i, row := range garden {
		for j, col := range row {
			if visited[[2]int{i, j}] {
				fmt.Print("V")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()

	}
}

package ten

import (
	"os"
	"slices"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
*/

/*
Starting at point S, only 2 pipes connect to S
Every pipe in the main loop connects to its two neighbors
*/

const complexLoop = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const fullComplexLoop = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

const partTwoExample = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

type Point struct {
	Row int
	Col int
	// Distance from starting point
	Dist      int
	Neighbors int
}

func RunDayTen() {
	input, err := os.ReadFile("./input/day10input.txt")
	if err != nil {
		panic("Failed to read day 10 input")
	}
	res, _ := partOne(string(input))
	println("Part 1:", res)
	res = partTwo(partTwoExample)
	println("Part 2:", res)
}

func partTwo(input string) int {
	// Use the path found in part one
	matrix, _ := parseInput(input)
	_, path := partOne(input)
	// Find all the points within the matrix that are "."
	points := make([]Point, 0)
	for r, row := range matrix {
		for c, ch := range row {
			if ch == "." {
				points = append(points, Point{Row: r, Col: c})
			}
		}
	}
	// Find all points that are "." and are within the area of the loop
	points = iterators.Filter(points, func(p Point) bool {
		return isInArea(&p, path)
	})
	// For every point, check to see if its enclosed by the loop
	return 0
}

func isInArea(p *Point, path []Point) bool {
}

func partOne(input string) (int, []Point) {
	matrix, startingPoint := parseInput(input)

	visited := make([]Point, 0)
	stack := make([]Point, 0)
	// Add starting point to stack
	stack = append(stack, startingPoint)
	path := make([]Point, 0)
	for len(stack) > 0 {
		// Pop from stack
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// Check if it has been visited
		if iterators.Some(visited, func(point Point) bool {
			return point.Row == p.Row && point.Col == p.Col
		}) {
			continue
		}
		// Mark as visited
		visited = append(visited, p)
		// Check if it is a loop
		neighbors := findConnectingNeighbors(matrix, &p)

		// Add to path
		path = append(path, p)
		// Add neighbors to stack
		for _, n := range neighbors {
			n.Dist = p.Dist + 1
			// Add to stack if it has not been visited
			if !iterators.Some(visited, func(point Point) bool {
				return point.Row == n.Row && point.Col == n.Col
			}) {
				stack = append(stack, n)
			}
		}
	}
	// Find the middle point, this is the farthest point from the starting point
	mid := len(path) / 2
	return path[mid].Dist, path
}

func parseInput(input string) ([][]string, Point) {
	matrix := make([][]string, 0)
	startingPoint := Point{}
	for r, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for c, ch := range line {
			// Convert rune to string
			if ch == 'S' {
				startingPoint.Col = c
				startingPoint.Row = r
			}
			row = append(row, string(ch))
		}
		matrix = append(matrix, row)
	}
	return matrix, startingPoint
}

func findConnectingNeighbors(matrix [][]string, p *Point) []Point {
	directions := []Point{
		{-1, 0, 0, 0}, // UP
		{0, 1, 0, 0},  // Right
		{1, 0, 0, 0},  // Down
		{0, -1, 0, 0}, // Left
	}

	pipes := make([]Point, 0)
	for _, dir := range directions {
		dx := p.Row + dir.Row
		dy := p.Col + dir.Col
		// Check if it doesnt exceed bounds
		if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
			continue
		}
		point := Point{Row: dx, Col: dy, Neighbors: 0}
		if matrix[dx][dy] == "S" {
			p.Neighbors += 1
		}
		// Check if the pipe connects to the point
		switch matrix[p.Row][p.Col] {
		case "|":
			// Vertical pipe, only north and south
			// { "|", "7", "F", "L", "J", }
			if (dir.Row == 1 || dir.Row == -1) && slices.Contains([]string{"|", "7", "F", "L", "J"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "-":
			// Horizontal pipe, only east and west
			// { "-", "L", "J", "7", "F"}
			if dir.Row == 0 && slices.Contains([]string{"-", "L", "J", "7", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "L":
			// East, North pipe. Only South and West
			if (dir.Row == -1 || dir.Col == 1) && slices.Contains([]string{"|", "-", "J", "7", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "J":
			// North and west pipe, only south and east
			if (dir.Row == -1 || dir.Col == -1) && slices.Contains([]string{"|", "-", "L", "7", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "7":
			// South and west pipe, only north and east
			if (dir.Row == 1 || dir.Col == -1) && slices.Contains([]string{"|", "-", "L", "J", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "F":
			// South and east pipe, only north and west
			if (dir.Row == 1 || dir.Col == 1) && slices.Contains([]string{"|", "-", "L", "J", "7"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		case "S":
			// Starting point, Only Two pipes connect to it
			// Check where the dx, dy point is
			// Up, check all pipes that connect to south
			if dir.Row == -1 && slices.Contains([]string{"|", "7", "F", "L", "J"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
			// Down, check all pipes that connect to north
			if dir.Row == 1 && slices.Contains([]string{"|", "7", "F", "L", "J"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
			// Left, check all pipes that connect to east
			if dir.Col == -1 && slices.Contains([]string{"-", "L", "J", "7", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
			// Right, check all pipes that connect to west
			if dir.Col == 1 && slices.Contains([]string{"-", "L", "J", "7", "F"}, matrix[dx][dy]) {
				p.Neighbors += 1
				pipes = append(pipes, point)
			}
		}
	}

	return pipes
}

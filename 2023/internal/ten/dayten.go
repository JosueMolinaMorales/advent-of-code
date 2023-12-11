package ten

import (
	"fmt"
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

const fullComplexLoop = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const partTwoExample = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const partTwoExampleB = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

const partTwoLargerExample = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

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
	// res, _ := partOne(string(input))
	// println("Part 1:", res)
	// // 680 too high
	res := partTwo(string(input))
	println("Part 2:", res)
	// fmt.Println(p2())
}

func partTwo(input string) int {
	// Use the path found in part one
	matrix, sp := parseInput(input)
	_, path := partOne(input)

	// Convert S to a pipe based on the pipes that connect to it
	p1 := path[1]
	p2 := path[len(path)-1]

	directions := []struct {
		Row       int
		Col       int
		Direction string
	}{
		{-1, 0, "UP"},
		{0, 1, "RIGHT"},
		{1, 0, "DOWN"},
		{0, -1, "LEFT"},
	}
	p1Dir := ""
	p2Dir := ""
	for _, dir := range directions {
		dx := sp.Row + dir.Row
		dy := sp.Col + dir.Col
		if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
			continue
		}
		if dx == p1.Row && dy == p1.Col {
			p1Dir = dir.Direction
		}
		if dx == p2.Row && dy == p2.Col {
			p2Dir = dir.Direction
		}
	}

	switch fmt.Sprintf("%s,%s", p1Dir, p2Dir) {
	case "UP,RIGHT", "RIGHT,UP":
		matrix[sp.Row][sp.Col] = "L"
	case "UP,DOWN", "DOWN,UP":
		matrix[sp.Row][sp.Col] = "|"
	case "UP,LEFT", "LEFT,UP":
		matrix[sp.Row][sp.Col] = "J"
	case "RIGHT,DOWN", "DOWN,RIGHT":
		matrix[sp.Row][sp.Col] = "7"
	case "RIGHT,LEFT", "LEFT,RIGHT":
		matrix[sp.Row][sp.Col] = "-"
	case "DOWN,LEFT", "LEFT,DOWN":
		matrix[sp.Row][sp.Col] = "F"
	}

	// Print out the loop
	for r := range matrix {
		for c := range matrix[r] {
			if iterators.Some(path, func(point Point) bool {
				return point.Row == r && point.Col == c
			}) {
				fmt.Print(matrix[r][c])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	sum := 0
	for r := range matrix {
		inside := false
		for c := range matrix[r] {
			if !iterators.Some(path, func(point Point) bool {
				return point.Row == r && point.Col == c
			}) {
				if inside {
					sum++
				}
			} else if matrix[r][c] == "|" ||
				matrix[r][c] == "L" ||
				matrix[r][c] == "J" {
				inside = !inside
			}
		}
	}

	return sum
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

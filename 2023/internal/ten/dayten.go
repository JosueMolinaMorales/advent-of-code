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
	res, _ := partOne(string(input))
	println("Part 1:", res)
	res = partTwo(string(input))
	println("Part 2:", res)
}

func partTwo(input string) int {
	// Use the path found in part one
	matrix, _ := parseInput(input)
	_, path := partOne(input)
	// Find all the points within the matrix that are "."
	points := make([]Point, 0)
	for r, row := range matrix {
		for c := range row {
			// If the point is not in the main loop
			if !iterators.Some(path, func(point Point) bool {
				return point.Row == r && point.Col == c
			}) {
				points = append(points, Point{Row: r, Col: c})
			}
		}
	}

	fmt.Println("Len points:", len(points))
	// Find all points that are "." and are within the area of the loop
	points = iterators.Filter(points, func(p Point) bool {
		return isInArea(&p, path, matrix)
	})
	fmt.Println("Len Points in area:", len(points))
	// Group adjacent points
	groups := groupAdjacentPoints(points)

	fmt.Println("Len Groups:", (groups))

	// Find all groups that are enclosed
	enclosedGroups := make([][]Point, 0)
	for _, group := range groups {
		isEnclosed := true
		for _, point := range group {
			if !isEnclosedInLoop(&point, group, groups, path, matrix) {
				fmt.Println("Point not enclosed:", point)
				isEnclosed = false
				break
			}
		}
		if isEnclosed {
			enclosedGroups = append(enclosedGroups, group)
		}
	}
	fmt.Println("Len Enclosed Groups:", len(enclosedGroups))
	pointCount := 0
	for _, group := range enclosedGroups {
		pointCount += len(group)
	}

	return pointCount
}

func groupAdjacentPoints(points []Point) [][]Point {
	groups := make([][]Point, 0)
	visited := make([]Point, 0)

	for _, point := range points {
		if !slices.ContainsFunc(visited, func(p Point) bool {
			return p.Row == point.Row && p.Col == point.Col
		}) {
			group := make([]Point, 0)
			dfs(point, points, &visited, &group)
			groups = append(groups, group)
		}
	}

	return groups
}

func dfs(point Point, points []Point, visited *[]Point, group *[]Point) {
	if slices.ContainsFunc(*visited, func(p Point) bool {
		return p.Row == point.Row && p.Col == point.Col
	}) {
		return
	}

	*visited = append(*visited, point)
	*group = append(*group, point)

	// Define the possible neighbors (up, down, left, right)
	neighbors := []Point{
		{point.Row - 1, point.Col, 0, 0},
		{point.Row + 1, point.Col, 0, 0},
		{point.Row, point.Col - 1, 0, 0},
		{point.Row, point.Col + 1, 0, 0},
	}

	for _, neighbor := range neighbors {
		if slices.ContainsFunc(points, func(p Point) bool {
			return p.Row == neighbor.Row && p.Col == neighbor.Col
		}) && !slices.ContainsFunc(*visited, func(p Point) bool {
			return p.Row == neighbor.Row && p.Col == neighbor.Col
		}) {
			dfs(neighbor, points, visited, group)
		}
	}
}

func isEnclosedInLoop(p *Point, group []Point, groups [][]Point, loop []Point, matrix [][]string) bool {
	// fmt.Println("Current point:", p.Row, p.Col)
	// Check The points adj to the point, if they are not part of the main loop, then it is not enclosed
	directions := []Point{
		{-1, 0, 0, 0}, // UP
		{0, 1, 0, 0},  // Right
		{1, 0, 0, 0},  // Down
		{0, -1, 0, 0}, // Left
	}
	adjNotEnclosed := false
	for _, dir := range directions {
		dx := p.Row + dir.Row
		dy := p.Col + dir.Col
		// Check if it doesnt exceed bounds
		if dx < 0 || dx >= len(matrix) || dy < 0 || dy >= len(matrix[0]) {
			continue
		}
		if !iterators.Some(loop, func(point Point) bool {
			return point.Row == dx && point.Col == dy
		}) && !iterators.Some(group, func(point Point) bool {
			return point.Row == dx && point.Col == dy
		}) {
			adjNotEnclosed = true
			break
		}
	}

	if adjNotEnclosed {
		fmt.Println("Adjacent point not enclosed", p.Row, p.Col)
		return false
	}

	topPoint := Point{Row: p.Row - 1, Col: p.Col, Dist: 0, Neighbors: 0}
	leftVertical := []string{"|", "7", "J"}
	rightVertical := []string{"|", "F", "L"}
	canGoUp := false
	if topPoint.Row > 0 && slices.Contains(leftVertical, matrix[topPoint.Row][topPoint.Col]) {
		// Keep going up
		topRightPoint := Point{Row: p.Row - 1, Col: p.Col + 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(leftVertical, matrix[topPoint.Row][topPoint.Col]) &&
			slices.Contains(rightVertical, matrix[topRightPoint.Row][topRightPoint.Col])) ||
			(slices.Contains(leftVertical, matrix[topPoint.Row][topPoint.Col]) && !isPointInMainLoop(&topRightPoint, loop)) ||
			(slices.Contains(rightVertical, matrix[topRightPoint.Row][topRightPoint.Col]) && !isPointInMainLoop(&topPoint, loop)) {
			topPoint.Row -= 1
			topRightPoint.Row -= 1

			if !isPointInMainLoop(&topPoint, loop) && !isPointInMainLoop(&topRightPoint, loop) && isNotPartOfGroup(&topPoint, groups) && isNotPartOfGroup(&topRightPoint, groups) {
				canGoUp = true
				break
			}

			if topPoint.Row < 0 || topRightPoint.Row < 0 {
				canGoUp = true
				break
			}
		}
	} else if topPoint.Row > 0 && slices.Contains(rightVertical, matrix[topPoint.Row][topPoint.Col]) {
		// Keep going up
		topLeftPoint := Point{Row: p.Row - 1, Col: p.Col - 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(rightVertical, matrix[topPoint.Row][topPoint.Col]) &&
			slices.Contains(leftVertical, matrix[topLeftPoint.Row][topLeftPoint.Col])) ||
			(slices.Contains(rightVertical, matrix[topPoint.Row][topPoint.Col]) && !isPointInMainLoop(&topLeftPoint, loop)) ||
			(slices.Contains(leftVertical, matrix[topLeftPoint.Row][topLeftPoint.Col]) && !isPointInMainLoop(&topPoint, loop)) {
			topPoint.Row -= 1
			topLeftPoint.Row -= 1

			if !isPointInMainLoop(&topPoint, loop) && !isPointInMainLoop(&topLeftPoint, loop) && isNotPartOfGroup(&topPoint, groups) && isNotPartOfGroup(&topLeftPoint, groups) {
				fmt.Printf("Point can go up: (%d, %d)\n", p.Row, p.Col)
				canGoUp = true
				break
			}

			if topPoint.Row < 0 || topLeftPoint.Row < 0 {
				fmt.Printf("Point can go up: (%d, %d)\n", p.Row, p.Col)
				canGoUp = true
				break
			}
		}
	}

	leftPoint := Point{Row: p.Row, Col: p.Col - 1, Dist: 0, Neighbors: 0}
	topHorizontal := []string{"-", "L", "J"}
	bottomHorizontal := []string{"-", "7", "F"}
	canGoLeft := false

	if leftPoint.Col > 0 && slices.Contains(topHorizontal, matrix[leftPoint.Row][leftPoint.Col]) {
		// Keep going left
		bottomLeftPoint := Point{Row: p.Row + 1, Col: p.Col - 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(topHorizontal, matrix[leftPoint.Row][leftPoint.Col]) &&
			slices.Contains(bottomHorizontal, matrix[bottomLeftPoint.Row][bottomLeftPoint.Col])) ||
			(slices.Contains(topHorizontal, matrix[leftPoint.Row][leftPoint.Col]) && !isPointInMainLoop(&bottomLeftPoint, loop)) ||
			(slices.Contains(bottomHorizontal, matrix[bottomLeftPoint.Row][bottomLeftPoint.Col]) && !isPointInMainLoop(&leftPoint, loop)) {
			leftPoint.Col -= 1
			bottomLeftPoint.Col -= 1

			if !isPointInMainLoop(&leftPoint, loop) && !isPointInMainLoop(&bottomLeftPoint, loop) && isNotPartOfGroup(&leftPoint, groups) && isNotPartOfGroup(&bottomLeftPoint, groups) {
				canGoLeft = true
				break
			}

			if leftPoint.Col < 0 || bottomLeftPoint.Col < 0 {
				canGoLeft = true
				break
			}
		}
	} else if leftPoint.Col > 0 && slices.Contains(bottomHorizontal, matrix[leftPoint.Row][leftPoint.Col]) {
		// Keep going left
		topLeftPoint := Point{Row: p.Row - 1, Col: p.Col - 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(bottomHorizontal, matrix[leftPoint.Row][leftPoint.Col]) &&
			slices.Contains(topHorizontal, matrix[topLeftPoint.Row][topLeftPoint.Col])) ||
			(slices.Contains(bottomHorizontal, matrix[leftPoint.Row][leftPoint.Col]) && !isPointInMainLoop(&topLeftPoint, loop)) ||
			(slices.Contains(topHorizontal, matrix[topLeftPoint.Row][topLeftPoint.Col]) && !isPointInMainLoop(&leftPoint, loop)) {
			leftPoint.Col -= 1
			topLeftPoint.Col -= 1

			if !isPointInMainLoop(&leftPoint, loop) && !isPointInMainLoop(&topLeftPoint, loop) && isNotPartOfGroup(&leftPoint, groups) && isNotPartOfGroup(&topLeftPoint, groups) {
				canGoLeft = true
				break
			}

			if leftPoint.Col < 0 || topLeftPoint.Col < 0 {
				canGoLeft = true
				break
			}
		}
	}

	rightPoint := Point{Row: p.Row, Col: p.Col + 1, Dist: 0, Neighbors: 0}
	canGoRight := false
	if rightPoint.Col < len(matrix[0]) && slices.Contains(topHorizontal, matrix[rightPoint.Row][rightPoint.Col]) {
		// Keep going right
		bottomRightPoint := Point{Row: p.Row + 1, Col: p.Col + 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(topHorizontal, matrix[rightPoint.Row][rightPoint.Col]) &&
			slices.Contains(bottomHorizontal, matrix[bottomRightPoint.Row][bottomRightPoint.Col])) ||
			(slices.Contains(topHorizontal, matrix[rightPoint.Row][rightPoint.Col]) && !isPointInMainLoop(&bottomRightPoint, loop)) ||
			(slices.Contains(bottomHorizontal, matrix[bottomRightPoint.Row][bottomRightPoint.Col]) && !isPointInMainLoop(&rightPoint, loop)) {
			rightPoint.Col += 1
			bottomRightPoint.Col += 1

			if !isPointInMainLoop(&rightPoint, loop) && !isPointInMainLoop(&bottomRightPoint, loop) && isNotPartOfGroup(&rightPoint, groups) && isNotPartOfGroup(&bottomRightPoint, groups) {
				canGoRight = true
				break
			}

			if rightPoint.Col >= len(matrix[0]) || bottomRightPoint.Col >= len(matrix[0]) {
				canGoRight = true
				break
			}
		}
	} else if rightPoint.Col < len(matrix[0]) && slices.Contains(bottomHorizontal, matrix[rightPoint.Row][rightPoint.Col]) {
		// Keep going right
		topRightPoint := Point{Row: p.Row - 1, Col: p.Col + 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(bottomHorizontal, matrix[rightPoint.Row][rightPoint.Col]) &&
			slices.Contains(topHorizontal, matrix[topRightPoint.Row][topRightPoint.Col])) ||
			(slices.Contains(bottomHorizontal, matrix[rightPoint.Row][rightPoint.Col]) && !isPointInMainLoop(&topRightPoint, loop)) ||
			(slices.Contains(topHorizontal, matrix[topRightPoint.Row][topRightPoint.Col]) && !isPointInMainLoop(&rightPoint, loop)) {
			rightPoint.Col += 1
			topRightPoint.Col += 1

			if !isPointInMainLoop(&rightPoint, loop) && !isPointInMainLoop(&topRightPoint, loop) && isNotPartOfGroup(&rightPoint, groups) && isNotPartOfGroup(&topRightPoint, groups) {
				canGoRight = true
				break
			}

			if rightPoint.Col >= len(matrix[0]) || topRightPoint.Col >= len(matrix[0]) {
				canGoRight = true
				break
			}
		}
	}

	bottomPoint := Point{Row: p.Row + 1, Col: p.Col, Dist: 0, Neighbors: 0}
	canGoDown := false
	if bottomPoint.Row < len(matrix) && slices.Contains(leftVertical, matrix[bottomPoint.Row][bottomPoint.Col]) {
		// Keep going down
		bottomRightPoint := Point{Row: p.Row + 1, Col: p.Col + 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(leftVertical, matrix[bottomPoint.Row][bottomPoint.Col]) &&
			slices.Contains(rightVertical, matrix[bottomRightPoint.Row][bottomRightPoint.Col])) ||
			(slices.Contains(leftVertical, matrix[bottomPoint.Row][bottomPoint.Col]) && !isPointInMainLoop(&bottomRightPoint, loop)) ||
			(slices.Contains(rightVertical, matrix[bottomRightPoint.Row][bottomRightPoint.Col]) && !isPointInMainLoop(&bottomPoint, loop)) {
			bottomPoint.Row += 1
			bottomRightPoint.Row += 1

			// If both points are out of the loop, then it is not enclosed
			if !isPointInMainLoop(&bottomPoint, loop) && !isPointInMainLoop(&bottomRightPoint, loop) && isNotPartOfGroup(&bottomPoint, groups) && isNotPartOfGroup(&bottomRightPoint, groups) {
				canGoDown = true
				break
			}

			if bottomPoint.Row >= len(matrix) || bottomRightPoint.Row >= len(matrix) {
				canGoDown = true
				break
			}
		}
	} else if bottomPoint.Row < len(matrix) && slices.Contains(rightVertical, matrix[bottomPoint.Row][bottomPoint.Col]) {
		// Keep going down
		bottomLeftPoint := Point{Row: p.Row + 1, Col: p.Col - 1, Dist: 0, Neighbors: 0}
		for (slices.Contains(rightVertical, matrix[bottomPoint.Row][bottomPoint.Col]) &&
			slices.Contains(leftVertical, matrix[bottomLeftPoint.Row][bottomLeftPoint.Col])) ||
			(slices.Contains(rightVertical, matrix[bottomPoint.Row][bottomPoint.Col]) && !isPointInMainLoop(&bottomLeftPoint, loop)) ||
			(slices.Contains(leftVertical, matrix[bottomLeftPoint.Row][bottomLeftPoint.Col]) && !isPointInMainLoop(&bottomPoint, loop)) {
			bottomPoint.Row += 1
			bottomLeftPoint.Row += 1

			if !isPointInMainLoop(&bottomPoint, loop) && !isPointInMainLoop(&bottomLeftPoint, loop) && isNotPartOfGroup(&bottomPoint, groups) && isNotPartOfGroup(&bottomLeftPoint, groups) {
				canGoDown = true
				break
			}

			if bottomPoint.Row >= len(matrix) || bottomLeftPoint.Row >= len(matrix) {
				canGoDown = true
				break
			}

		}
	}

	return !canGoUp && !canGoLeft && !canGoRight && !canGoDown
}

func isNotPartOfGroup(p *Point, groups [][]Point) bool {
	for _, group := range groups {
		if iterators.Some(group, func(point Point) bool {
			return point.Row == p.Row && point.Col == p.Col
		}) {
			return false
		}
	}
	return true
}

func isPointInMainLoop(p *Point, loop []Point) bool {
	return iterators.Some(loop, func(point Point) bool {
		return point.Row == p.Row && point.Col == p.Col
	})
}

func isInArea(p *Point, loop []Point, matrix [][]string) bool {
	// Go above the point until either a pipe within the loop is found or the top of the matrix is reached
	top := false
	for i := p.Row; i >= 0; i-- {
		// The current point was found within the loop
		if iterators.Some(loop, func(point Point) bool {
			return point.Row == i && point.Col == p.Col
		}) {
			top = true
		}
	}

	// Go below the point until either a pipe within the loop is found or the bottom of the matrix is reached
	below := false
	for i := p.Row; i < len(matrix); i++ {
		// The current point was found within the loop
		if iterators.Some(loop, func(point Point) bool {
			return point.Row == i && point.Col == p.Col
		}) {
			below = true
		}
	}

	// Go left of the point until either a pipe within the loop is found or the left of the matrix is reached
	left := false
	for i := p.Col; i >= 0; i-- {
		// The current point was found within the loop
		if iterators.Some(loop, func(point Point) bool {
			return point.Row == p.Row && point.Col == i
		}) {
			left = true
		}
	}

	// Go right of the point until either a pipe within the loop is found or the right of the matrix is reached
	right := false
	for i := p.Col; i < len(matrix[0]); i++ {
		// The current point was found within the loop
		if iterators.Some(loop, func(point Point) bool {
			return point.Row == p.Row && point.Col == i
		}) {
			right = true
		}
	}

	return top && below && left && right
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

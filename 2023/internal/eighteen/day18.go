package eighteen

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

type Plan struct {
	Direction string
	Steps     int
	RGB       string
}

type Point = [2]int

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

var HexDirection = map[string]string{
	"0": RIGHT,
	"1": DOWN,
	"2": LEFT,
	"3": UP,
}

var Directions = map[string]Point{
	RIGHT: {0, 1},
	LEFT:  {0, -1},
	UP:    {-1, 0},
	DOWN:  {1, 0},
}

func RunDayEighteen() {
	input, err := os.ReadFile("./input/day18.txt")
	if err != nil {
		panic("Failed to read day 18 input")
	}
	fmt.Println("Day 18 Part 1", partOne(string(input)))
	fmt.Println("Day 18 Part 2", partTwo(string(input)))
}

func partTwo(input string) int {
	digPlan := make([]Plan, 0)
	for _, line := range strings.Split(input, "\n") {
		plan := Plan{}
		p := strings.Split(line, " ")
		plan.RGB = strings.ReplaceAll(p[2], "#", "")[1 : len(p[2])-2]
		plan.Direction = HexDirection[plan.RGB[len(plan.RGB)-1:]]
		hex, _ := strconv.ParseInt(plan.RGB[:len(plan.RGB)-1], 16, 64)
		plan.Steps = int(hex)
		digPlan = append(digPlan, plan)
	}
	points, b := expandPoints(digPlan)
	area := shoelaceFormula(points)
	return int(area) + b/2 + 1
}

func expandPoints(digPlan []Plan) ([]Point, int) {
	points := []Point{{0, 0}}
	b := 0
	for _, plan := range digPlan {
		dir := Directions[plan.Direction]
		lp := points[len(points)-1]
		steps := plan.Steps
		b += steps
		points = append(points, Point{lp[0] + dir[0]*steps, lp[1] + dir[1]*steps})
	}
	return points, b
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func shoelaceFormula(points []Point) float64 {
	result := 0
	for i := 0; i < len(points)-1; i++ {
		x1, y1 := points[i][0], points[i][1]
		x2, y2 := points[i+1][0], points[i+1][1]
		result += int(x1*y2 - x2*y1)
	}
	return math.Abs(float64(result)) / 2
}

func partOne(input string) int {
	digPlan := make([]Plan, 0)
	for _, line := range strings.Split(input, "\n") {
		plan := Plan{}
		p := strings.Split(line, " ")
		plan.Direction = p[0]
		plan.Steps, _ = strconv.Atoi(p[1])
		plan.RGB = p[2]
		digPlan = append(digPlan, plan)
	}
	points, b := expandPoints(digPlan)
	area := shoelaceFormula(points)
	// print(points)
	return int(area) + b/2 + 1
}

// For fun functions
func print(points []Point) {
	fmt.Println("BEFORE")
	printMatrix(points)
	fmt.Println("AFTER")
	fmt.Println()
	fmt.Println()
	minX, minY, maxX, maxY := getMinMax(points)
	p := make(map[Point]bool, 0)
	for _, point := range points {
		p[point] = true
	}
	floodFill(1, 1, minX, minY, maxX, maxY, p)
	points = make([]Point, 0)
	for point := range p {
		points = append(points, point)
	}
	printMatrix(points)
}

func floodFill(x, y, minX, minY, maxX, maxY int, points map[Point]bool) int {
	if x < minX || y < minY || x >= maxX || y >= maxY || points[Point{x, y}] {
		// Out of bounds or already visited, stop
		return 0
	}

	// Mark the current point as visited
	points[Point{x, y}] = true

	// Recursively visit the neighboring points
	count := 1
	count += floodFill(x+1, y, minX, minY, maxX, maxY, points)
	count += floodFill(x-1, y, minX, minY, maxX, maxY, points)
	count += floodFill(x, y+1, minX, minY, maxX, maxY, points)
	count += floodFill(x, y-1, minX, minY, maxX, maxY, points)

	return count
}

func printMatrix(points []Point) {
	minX, minY, maxX, maxY := getMinMax(points)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			if iterators.Contains(points, Point{i, j}) {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func getMinMax(points []Point) (int, int, int, int) {
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	for _, p := range points {
		minX = int(math.Min(float64(minX), float64(p[0])))
		minY = int(math.Min(float64(minY), float64(p[1])))
		maxX = int(math.Max(float64(maxX), float64(p[0])))
		maxY = int(math.Max(float64(maxY), float64(p[1])))
	}
	return minX, minY, maxX, maxY
}

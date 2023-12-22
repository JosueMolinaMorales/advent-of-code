package eleven

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

type Point struct {
	x, y int
}

// https://en.wikipedia.org/wiki/Taxicab_geometry
func (p Point) manhattanDistance(other Point) int {
	return int(math.Abs(float64(p.x-other.x)) + math.Abs(float64(p.y-other.y)))
}

func (p Point) mins(other Point) Point {
	return Point{x: min(p.x, other.x), y: min(p.y, other.y)}
}

func (p Point) maxs(other Point) Point {
	return Point{x: max(p.x, other.x), y: max(p.y, other.y)}
}

func RunDayEleven() {
	input, err := os.ReadFile("./input/day11.txt")
	if err != nil {
		panic("Failed to read day 11 input file")
	}
	fmt.Println("Day 11 Part 1:", partOne(string(input)))
	fmt.Println("Day 11 Part 2:", partTwo(string(input)))
}

func parseInput(input string) ([]Point, []int, []int) {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, ch := range line {
			// Convert rune to string
			row = append(row, string(ch))
		}
		matrix = append(matrix, row)
	}

	// Find all '#' in the matrix
	galaxies := make([]Point, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != "." {
				galaxies = append(galaxies, Point{x: i, y: j})
			}
		}
	}

	// Determine which rows are empty and which columns are empty
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		emptySpace := iterators.Every(matrix[i], func(ch string) bool {
			return ch == "."
		})
		if emptySpace {
			emptyRows = append(emptyRows, i)
		}
	}

	// Check columns for empty space
	i, j := 0, 0
	for j < len(matrix[0]) {
		emptySpace := true
		for i < len(matrix) {
			if matrix[i][j] != "." {
				emptySpace = false
				break
			}
			i++
		}
		if emptySpace {
			emptyCols = append(emptyCols, j)
		}
		j++
		i = 0
	}

	return galaxies, emptyRows, emptyCols
}

func partTwo(input string) int {
	galaxies, emptyRows, emptyCols := parseInput(input)

	sum := sumDistances(galaxies, emptyRows, emptyCols, 999_999)

	return sum
}

func sumDistances(galaxies []Point, emptyRows, emptyCols []int, expansion int) int {
	sum := 0
	// For each start, calc manhanttan distance to every other galaxy
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			// Calc manhattan distance
			dist := galaxies[i].manhattanDistance(galaxies[j])
			// Add expansion to the dist for each row or column it would travese
			minPoint := galaxies[i].mins(galaxies[j])
			maxPoint := galaxies[i].maxs(galaxies[j])

			// Check if the path would cross an empty row
			rowCount := 0
			for _, row := range emptyRows {
				if row >= minPoint.x && row <= maxPoint.x {
					rowCount++
				}
			}
			colCount := 0
			for _, col := range emptyCols {
				if col >= minPoint.y && col <= maxPoint.y {
					colCount++
				}
			}

			sum += dist + (rowCount * expansion) + (colCount * expansion)
		}
	}

	return sum
}

func partOne(input string) int {
	galaxies, emptyRows, emptyCols := parseInput(input)
	sum := sumDistances(galaxies, emptyRows, emptyCols, 1)
	return sum
}

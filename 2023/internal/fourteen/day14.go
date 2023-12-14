package fourteen

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const testInput = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

type Point struct {
	Row int
	Col int
	Val rune
}

func RunDayFourteen() {
	input, err := os.ReadFile("./input/day14.txt")
	if err != nil {
		panic("Failed to read file for day 14")
	}
	start := time.Now()
	fmt.Println("Part 1", partOne(string(input)))
	elasped := time.Since(start)
	fmt.Println("Part 1 took", elasped)

	start = time.Now()
	fmt.Println("Part 2", partTwo(string(input)))
	elasped = time.Since(start)
	fmt.Println("Part 2 took", elasped)
}

const (
	RoundRock = iota
	CubedRock
)

func partOne(input string) int {
	m := parseInput(input)
	tiltNorth(&m)
	return load(&m)
}

func parseInput(input string) [][]string {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, c := range line {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func load(matrix *[][]string) int {
	load := 0
	for i, row := range *matrix {
		for _, c := range row {
			if c == "O" {
				load += len(*matrix) - i
			}
		}
	}
	return load
}

func partTwo(input string) int {
	matrix := parseInput(input)
	seen := make(map[string]int)
	loads := make([]int, 1)
	currBoard := ""
	i := 1
	for i < 100_000_000 {
		// Tilt the board north
		tiltNorth(&matrix)
		// Tilt the board west
		tiltWest(&matrix)
		// Tilt the board south
		tiltSouth(&matrix)
		// Tilt the board east
		tiltEast(&matrix)

		loads = append(loads, load(&matrix))
		currBoard = buildKey(&matrix)
		if _, ok := seen[currBoard]; ok {
			break
		}
		seen[currBoard] = i
		i++
	}

	lenLoop := i - seen[currBoard] // The length of the loop
	idx := seen[currBoard]         // The index of the first element of the loop

	t := idx + (1_000_000_000-idx)%lenLoop
	return loads[t]
}

func buildKey(matrix *[][]string) string {
	board := ""
	for _, row := range *matrix {
		board += strings.Join(row, "")
	}
	return board
}

func tiltNorth(matrix *[][]string) {
	for row := 1; row < len(*matrix); row++ {
		for col := 0; col < len((*matrix)[0]); col++ {
			if (*matrix)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := row; i > 0; i-- {
					if (*matrix)[i-1][col] == "O" || (*matrix)[i-1][col] == "#" {
						// Hit rock
						(*matrix)[i][col] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*matrix)[0][col] = "O"
				}

				if rockPlaced != row {
					(*matrix)[row][col] = "."
				}
			}
		}
	}
}

func tiltWest(matrix *[][]string) {
	for col := 1; col < len((*matrix)[0]); col++ {
		for row := 0; row < len(*matrix); row++ {
			if (*matrix)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := col; i > 0; i-- {
					if (*matrix)[row][i-1] == "O" || (*matrix)[row][i-1] == "#" {
						// Hit rock
						(*matrix)[row][i] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*matrix)[row][0] = "O"
				}

				if rockPlaced != col {
					(*matrix)[row][col] = "."
				}
			}
		}
	}
}

func tiltSouth(matrix *[][]string) {
	for row := len(*matrix) - 2; row >= 0; row-- {
		for col := 0; col < len((*matrix)[0]); col++ {
			if (*matrix)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := row; i < len(*matrix)-1; i++ {
					if (*matrix)[i+1][col] == "O" || (*matrix)[i+1][col] == "#" {
						// Hit rock
						(*matrix)[i][col] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*matrix)[len(*matrix)-1][col] = "O"
				}

				if rockPlaced != row {
					(*matrix)[row][col] = "."
				}
			}
		}
	}
}

func tiltEast(matrix *[][]string) {
	for col := len((*matrix)[0]) - 2; col >= 0; col-- {
		for row := 0; row < len(*matrix); row++ {
			if (*matrix)[row][col] == "O" {
				// Check if there is already a rock below it
				rockPlaced := -1
				for i := col; i < len((*matrix)[0])-1; i++ {
					if (*matrix)[row][i+1] == "O" || (*matrix)[row][i+1] == "#" {
						// Hit rock
						(*matrix)[row][i] = "O" // Move the rock to the place
						rockPlaced = i
						break
					}
				}

				if rockPlaced == -1 {
					(*matrix)[row][len((*matrix)[0])-1] = "O"
				}

				if rockPlaced != col {
					(*matrix)[row][col] = "."
				}
			}
		}
	}
}

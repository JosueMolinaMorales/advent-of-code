package sixteen

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const testInput = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func RunDaySixteen() {
	input, err := os.ReadFile("./input/day16.txt")
	if err != nil {
		panic("Failed to read day 16 input")
	}
	fmt.Println("Part 1:", partOne(string(input)))
	fmt.Println("Part 2:", partTwo(string(input)))
}

const (
	Eastward = iota
	Northward
	Southward
	Westward
)

type (
	Beam = [3]int
	Tile = [2]int
)

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

func partTwo(input string) int {
	matrix := parseInput(input)
	initialBeams := make([]Beam, 0)
	// Top row go down
	for i := 0; i < len(matrix[0]); i++ {
		initialBeams = append(initialBeams, Beam{0, i, Southward})
		initialBeams = append(initialBeams, Beam{len(matrix) - 1, i, Northward})
	}
	// Left Column go right
	for i := 0; i < len(matrix); i++ {
		initialBeams = append(initialBeams, Beam{i, 0, Eastward})
		initialBeams = append(initialBeams, Beam{i, len(matrix[0]) - 1, Westward})
	}

	max := 0
	for _, beam := range initialBeams {
		energized := countEnergized(beam, &matrix)
		max = int(math.Max(float64(max), float64(energized)))
	}
	return max
}

func partOne(input string) int {
	matrix := parseInput(input)
	energized := countEnergized(Beam{0, 0, Southward}, &matrix)
	return energized
}

func outOfBounds(row, col int, matrix *[][]string) bool {
	return row < 0 || row >= len(*matrix) || col < 0 || col >= len((*matrix)[0])
}

func countEnergized(startingBeam Beam, matrix *[][]string) int {
	queue := make([]Beam, 0)
	queue = append(queue, startingBeam)
	energized := make(map[Tile]bool, 0)
	visitedState := make(map[Beam]bool, 0)

	for len(queue) > 0 {
		beam := queue[0]
		queue = queue[1:]

		if outOfBounds(beam[0], beam[1], matrix) {
			continue
		}
		// Check if i have already been to this state
		if visitedState[beam] {
			continue
		}
		energized[Tile{beam[0], beam[1]}] = true
		visitedState[beam] = true
		dx, dy := 0, 0
		switch beam[2] {
		case Westward:
			dy = -1
		case Northward:
			dx = -1
		case Southward:
			dx = 1
		case Eastward:
			dy = 1
		}

		// Landed on tile, check what the tile is
		tile := (*matrix)[beam[0]][beam[1]]
		// Need to figure out looping
		switch tile {
		case ".":
			// Just continue on
			beam[0], beam[1] = beam[0]+dx, beam[1]+dy
			queue = append(queue, beam)
		case "|":
			if beam[2] == Northward || beam[2] == Southward {
				beam[0], beam[1] = beam[0]+dx, beam[1]+dy
				queue = append(queue, beam)
			} else {
				// Go down first
				queue = append(queue, Beam{beam[0] + 1, beam[1], Southward})
				// Go up
				queue = append(queue, Beam{beam[0] - 1, beam[1], Northward})
			}
		case "-":
			if beam[2] == Westward || beam[2] == Eastward {
				beam[0], beam[1] = beam[0]+dx, beam[1]+dy
				queue = append(queue, beam)
			} else {
				// Go left first
				queue = append(queue, Beam{beam[0], beam[1] - 1, Westward})
				// Go right
				queue = append(queue, Beam{beam[0], beam[1] + 1, Eastward})
			}
		case "/":
			switch beam[2] {
			case Eastward:
				dx, dy = -1, 0
				beam[2] = Northward
			case Westward:
				dx, dy = 1, 0
				beam[2] = Southward
			case Northward:
				dx, dy = 0, 1
				beam[2] = Eastward
			case Southward:
				dx, dy = 0, -1
				beam[2] = Westward
			}
			beam[0], beam[1] = beam[0]+dx, beam[1]+dy
			queue = append(queue, beam)
		case "\\":
			switch beam[2] {
			case Eastward:
				dx, dy = 1, 0
				beam[2] = Southward
			case Westward:
				dx, dy = -1, 0
				beam[2] = Northward
			case Northward:
				dx, dy = 0, -1
				beam[2] = Westward
			case Southward:
				dx, dy = 0, 1
				beam[2] = Eastward
			}
			beam[0], beam[1] = beam[0]+dx, beam[1]+dy
			queue = append(queue, beam)
		}
	}

	return len(energized)
}

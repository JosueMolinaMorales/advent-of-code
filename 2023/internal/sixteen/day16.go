package sixteen

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
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
	input, _ := os.ReadFile("./input/day16.txt")
	// fmt.Println("Part 1:", partOne(testInput))
	// fmt.Println("Part 1:", partOne(string(input)))
	// Too Low 8434
	fmt.Println("Part 2:", partTwo(string(input)))
	fmt.Println("Part 2:", partTwo(testInput))
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

func partTwo(input string) int {
	matrix := make([][]string, 0)

	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, c := range line {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}
	fmt.Println(len(matrix))
	fmt.Println(len(matrix[0]))
	energies := make([]int, 0)
	// Top row go down
	for i := 0; i < len(matrix[0]); i++ {
		energized := make(map[Tile]bool, 0)
		visitedState := make([]Beam, 0)
		followBeam(Beam{0, i, Southward}, &matrix, &energized, &visitedState)
		energies = append(energies, len(energized))
	}
	fmt.Println("Current max", iterators.Max(energies))
	// Bottom row go up
	for i := 0; i < len(matrix[0]); i++ {
		energized := make(map[Tile]bool, 0)
		visitedState := make([]Beam, 0)
		followBeam(Beam{len(matrix) - 1, i, Northward}, &matrix, &energized, &visitedState)
		energies = append(energies, len(energized))
	}
	fmt.Println("Current max", iterators.Max(energies))
	// Left Column go right
	for i := 0; i < len(matrix); i++ {
		energized := make(map[Tile]bool, 0)
		visitedState := make([]Beam, 0)
		followBeam(Beam{i, 0, Eastward}, &matrix, &energized, &visitedState)
		energies = append(energies, len(energized))
	}
	fmt.Println("Current max", iterators.Max(energies))
	// Right Column go left
	for i := 0; i < len(matrix); i++ {
		energized := make(map[Tile]bool, 0)
		visitedState := make([]Beam, 0)
		followBeam(Beam{len(matrix) - 1, i, Westward}, &matrix, &energized, &visitedState)
		energies = append(energies, len(energized))
	}
	fmt.Println("Current max", iterators.Max(energies))

	max := iterators.Max(energies)

	return max
}

func partOne(input string) int {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, c := range line {
			row = append(row, string(c))
		}
		matrix = append(matrix, row)
	}

	energized := make(map[Tile]bool, 0)
	visitedState := make([]Beam, 0)
	followBeam(Beam{0, 0, Eastward}, &matrix, &energized, &visitedState)

	return len(energized)
}

func outOfBounds(row, col int, matrix *[][]string) bool {
	return row < 0 || row >= len(*matrix) || col < 0 || col >= len((*matrix)[0])
}

func followBeam(beam Beam, matrix *[][]string, energized *map[Tile]bool, visitedState *[]Beam) {
	if outOfBounds(beam[0], beam[1], matrix) {
		return
	}

	// Check if i have already been to this state
	if slices.ContainsFunc(*visitedState, func(i [3]int) bool {
		return beam[0] == i[0] && beam[1] == i[1] && beam[2] == i[2]
	}) {
		return
	}

	(*energized)[Tile{beam[0], beam[1]}] = true
	*visitedState = append(*visitedState, beam)
	// Print map
	// printMap(*matrix, *energized)

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
		followBeam(beam, matrix, energized, visitedState)
	case "|":
		if beam[2] == Northward || beam[2] == Southward {
			beam[0], beam[1] = beam[0]+dx, beam[1]+dy
			followBeam(beam, matrix, energized, visitedState)
		} else {
			// Go down first
			followBeam(Beam{beam[0] + 1, beam[1], Southward}, matrix, energized, visitedState)
			// Go up
			followBeam(Beam{beam[0] - 1, beam[1], Northward}, matrix, energized, visitedState)
		}
	case "-":
		if beam[2] == Westward || beam[2] == Eastward {
			beam[0], beam[1] = beam[0]+dx, beam[1]+dy
			followBeam(beam, matrix, energized, visitedState)
		} else {
			// Go left first
			followBeam(Beam{beam[0], beam[1] - 1, Westward}, matrix, energized, visitedState)
			// Go right
			followBeam(Beam{beam[0], beam[1] + 1, Eastward}, matrix, energized, visitedState)
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
		followBeam(beam, matrix, energized, visitedState)
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
		followBeam(beam, matrix, energized, visitedState)
	}
}

func printMap(matrix [][]string, energized map[[2]int]bool) {
	for i, line := range matrix {
		for j := range line {
			if energized[Tile{i, j}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

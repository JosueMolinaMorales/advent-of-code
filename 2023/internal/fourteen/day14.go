package fourteen

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
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
	// fmt.Println("Part 1", partOne(testInput))
	// fmt.Println("Part 1", partOne(string(input)))
	fmt.Println("Part 2", partTwo(string(input)))
	// fmt.Println("Part 2", partTwo(testInput)) // 104866 too high
	// 103436 not correct
}

const (
	RoundRock = iota
	CubedRock
)

func partOne(input string) int {
	rr, cr := parseInput(input)
	nRows := len(strings.Split(input, "\n"))

	pr, _ := tiltNorth(rr, cr)

	sum := 0
	c := 0
	for _, pr := range pr {
		if pr[2] == RoundRock {
			c += 1
			sum += (nRows - pr[0])
		}
	}

	return sum
}

func parseInput(input string) ([][3]int, [][3]int) {
	rr, cr := make([][3]int, 0), make([][3]int, 0)
	for i, line := range strings.Split(input, "\n") {
		for j, c := range line {
			if c == 'O' {
				rr = append(rr, [3]int{i, j, RoundRock})
			} else if c == '#' {
				cr = append(cr, [3]int{i, j, CubedRock})
			}
		}
	}
	return rr, cr
}

func load(rr [][3]int, nRows int) int {
	load := 0
	for _, r := range rr {
		load += (nRows - r[0])
	}
	return load
}

func partTwo(input string) int {
	rr, cr := parseInput(input)
	nRows := len(strings.Split(input, "\n"))
	nCols := len(strings.Split(input, "\n")[0])
	seen := make(map[string]int)
	loads := make([]int, 1)
	currBoard := ""
	i := 1
	for i < 100_000_000 {
		// Tilt the board north
		rr, cr = tiltNorth(rr, cr)
		// Tilt the board west
		rr, cr = tiltWest(rr, cr)
		// Tilt the board south
		rr, cr = tiltSouth(rr, cr, nRows)
		// Tilt the board east
		rr, cr = tiltEast(rr, cr, nCols)

		loads = append(loads, load(rr, nRows))
		currBoard = buildKey(input, rr, cr)
		if _, ok := seen[currBoard]; ok {
			break
		}
		seen[currBoard] = i
		i++
	}

	lam := i - seen[currBoard] // The lenght of the loop
	mu := seen[currBoard]      // The index of the first element of the loop

	times := mu + (1_000_000_000-mu)%lam
	return loads[times]
}

func buildKey(input string, rr [][3]int, cr [][3]int) string {
	board := ""
	for i, line := range strings.Split(input, "\n") {
		for j := range line {
			if iterators.Some(rr, func(p [3]int) bool {
				return p[0] == i && p[1] == j
			}) {
				board += "O"
			} else if iterators.Some(cr, func(p [3]int) bool {
				return p[0] == i && p[1] == j
			}) {
				board += "#"
			} else {
				board += "."
			}
		}
	}
	return board
}

func tiltNorth(rr [][3]int, cr [][3]int) ([][3]int, [][3]int) {
	// Tilt the board up
	placedRocks := make([][3]int, 0)
	// Add crs
	placedRocks = append(placedRocks, cr...)

	sort.Slice(rr, func(i, j int) bool {
		return rr[i][0] < rr[j][0]
	})
	for i := 0; len(rr) > 0; i++ {
		// Check if there is already a rock below it
		rockPlaced := false
		r := rr[0]
		rr = rr[1:]
		for i := r[0] - 1; i >= 0; i-- {
			if iterators.Some(rr, func(p [3]int) bool {
				return p[1] == r[1] && p[0] == i
			}) || iterators.Some(placedRocks, func(p [3]int) bool {
				return p[1] == r[1] && p[0] == i
			}) {
				// Hit rock, move up 1 and place
				r[0] = i + 1
				rockPlaced = true
				placedRocks = append(placedRocks, r)
				break
			}
		}
		if !rockPlaced {
			r[0] = 0
			placedRocks = append(placedRocks, r)
		}
	}

	return iterators.Filter(placedRocks, func(p [3]int) bool {
		return p[2] == RoundRock
	}), cr
}

func tiltWest(rr [][3]int, cr [][3]int) ([][3]int, [][3]int) {
	// Tilt the board up
	placedRocks := make([][3]int, 0)
	// Add crs
	placedRocks = append(placedRocks, cr...)
	sort.Slice(rr, func(i, j int) bool {
		return rr[i][1] < rr[j][1]
	})
	for i := 0; len(rr) > 0; i++ {
		// Check if there is already a rock below it
		rockPlaced := false
		r := rr[0]
		rr = rr[1:]
		for i := r[1] - 1; i >= 0; i-- {
			if iterators.Some(rr, func(p [3]int) bool {
				return p[1] == i && p[0] == r[0]
			}) || iterators.Some(placedRocks, func(p [3]int) bool {
				return p[1] == i && p[0] == r[0]
			}) {
				// Hit rock, move left 1 and place
				r[1] = i + 1
				rockPlaced = true
				placedRocks = append(placedRocks, r)
				break
			}
		}
		if !rockPlaced {
			r[1] = 0
			placedRocks = append(placedRocks, r)
		}
	}

	return iterators.Filter(placedRocks, func(p [3]int) bool {
		return p[2] == RoundRock
	}), cr
}

func tiltSouth(rr [][3]int, cr [][3]int, maxRows int) ([][3]int, [][3]int) {
	// Tilt the board up
	placedRocks := make([][3]int, 0)
	// Add crs
	placedRocks = append(placedRocks, cr...)

	sort.Slice(rr, func(i, j int) bool {
		return rr[i][0] > rr[j][0]
	})
	for i := 0; len(rr) > 0; i++ {
		// Check if there is already a rock below it
		rockPlaced := false
		r := rr[0]
		rr = rr[1:]
		for i := r[0] + 1; i < maxRows; i++ {
			if iterators.Some(rr, func(p [3]int) bool {
				return p[1] == r[1] && p[0] == i
			}) || iterators.Some(placedRocks, func(p [3]int) bool {
				return p[1] == r[1] && p[0] == i
			}) {
				// Hit rock, move down 1 and place
				r[0] = i - 1
				rockPlaced = true
				placedRocks = append(placedRocks, r)
				break
			}
		}
		if !rockPlaced {
			r[0] = maxRows - 1
			placedRocks = append(placedRocks, r)
		}
	}

	return iterators.Filter(placedRocks, func(p [3]int) bool {
		return p[2] == RoundRock
	}), cr
}

func tiltEast(rr [][3]int, cr [][3]int, maxCols int) ([][3]int, [][3]int) {
	// Tilt the board up
	placedRocks := make([][3]int, 0)
	// Add crs
	placedRocks = append(placedRocks, cr...)

	sort.Slice(rr, func(i, j int) bool {
		return rr[i][1] > rr[j][1]
	})
	for i := 0; len(rr) > 0; i++ {
		// Check if there is already a rock below it
		rockPlaced := false
		r := rr[0]
		rr = rr[1:]
		for i := r[1] + 1; i < maxCols; i++ {
			if iterators.Some(rr, func(p [3]int) bool {
				return p[1] == i && p[0] == r[0]
			}) || iterators.Some(placedRocks, func(p [3]int) bool {
				return p[1] == i && p[0] == r[0]
			}) {
				// Hit rock, move right 1 and place
				r[1] = i - 1
				rockPlaced = true
				placedRocks = append(placedRocks, r)
				break
			}
		}
		if !rockPlaced {
			r[1] = maxCols - 1
			placedRocks = append(placedRocks, r)
		}
	}

	return iterators.Filter(placedRocks, func(p [3]int) bool {
		return p[2] == RoundRock
	}), cr
}

package thirteen

import (
	"fmt"
	"os"
	"strings"
)

const testInput = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func RunDayThirteen() {
	input, err := os.ReadFile("./input/day13.txt")
	if err != nil {
		panic("Failed to read input for day 13")
	}

	// 26869 too small
	fmt.Println("Part 1", partOne(string(input)))
	fmt.Println(partOne(testInput))
}

func partOne(input string) int {
	maps := make([][][]string, 0)
	for _, m := range strings.Split(input, "\n\n") {
		nm := make([][]string, 0)
		for _, l := range strings.Split(m, "\n") {
			nm = append(nm, strings.Split(l, ""))
		}
		maps = append(maps, nm)
	}

	// for _, m := range maps {
	// 	for _, l := range m {
	// 		fmt.Println(l)
	// 	}
	// 	fmt.Println()
	// }

	sum := 0
	for _, m := range maps {
		ref, c := getCount(m)
		switch ref {
		case HRef:
			sum += (c * 100)
		case VRef:
			sum += c
		}
	}
	return sum
}

type Reflection struct {
	P1 int
	P2 int
}

const (
	HRef = iota
	VRef
)

func getCount(m [][]string) (int, int) {
	// Find the middle reflection point
	// This will be the point where (row or col) i == j
	// Start with Rows
	i, j := 0, 1
	rowPoints := make([]Reflection, 0)
	for i < len(m) && j < len(m) {
		found := checkRowsMatch(i, j, m)
		if found {
			rowPoints = append(rowPoints, Reflection{P1: i, P2: j})
		}
		i++
		j++
	}

	// Find col
	i, j = 0, 1
	colPoint := make([]Reflection, 0)
	for i < len(m[0]) && j < len(m[0]) {
		found := checkColumnsMatch(i, j, m)
		if found {
			colPoint = append(colPoint, Reflection{P1: i, P2: j})
		}
		i++
		j++
	}

	c := 0
	var p Reflection
	for _, rp := range rowPoints {
		c += 2
		dTop := rp.P1 - 1
		dBottom := rp.P2 + 1
		for dTop >= 0 && dBottom < len(m) {
			c += 1
			// Check if the rows are the same
			if !checkRowsMatch(dTop, dBottom, m) {
				c = 0
				break
			}
			dTop -= 1
			dBottom += 1
		}
		if c != 0 {
			p = rp
			break
		}
	}
	if c != 0 {
		return HRef, p.P2
	}
	c = 0
	for _, cp := range colPoint {
		c += 2
		dLeft := cp.P1 - 1
		dRight := cp.P2 + 1
		for dLeft >= 0 && dRight < len(m[0]) {
			if !checkColumnsMatch(dLeft, dRight, m) {
				c = 0
				break
			}
			c += 1
			dLeft -= 1
			dRight += 1
		}
		if c != 0 {
			p = cp
			break
		}
	}

	return VRef, p.P2
}

func checkRowsMatch(i, j int, m [][]string) bool {
	for k := 0; k < len(m[0]); k++ {
		if m[i][k] != m[j][k] {
			return false
		}
	}
	return true
}

func checkColumnsMatch(i, j int, m [][]string) bool {
	for k := 0; k < len(m); k++ {
		if m[k][i] != m[k][j] {
			return false
		}
	}
	return true
}

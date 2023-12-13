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

	fmt.Println("Part 1", partOne(string(input)))
	fmt.Println("Part 2", partTwo(string(input))) // 25448 too low
	// fmt.Println(partOne(testInput))
	fmt.Println(partTwo(testInput))
}

func partOne(input string) int {
	maps := parseInput(input)
	sum := 0
	for _, m := range maps {
		dir, c := getCount(m, nil)

		switch dir {
		case HRef:
			sum += (c.P2 * 100)
		case VRef:
			sum += c.P2
		}
	}

	return sum
}

func partTwo(input string) int {
	maps := parseInput(input)
	sum := 0
	for _, m := range maps {
		sum += summary(m)
	}

	return sum
}

func parseInput(input string) [][][]string {
	maps := make([][][]string, 0)
	for _, m := range strings.Split(input, "\n\n") {
		nm := make([][]string, 0)
		for _, l := range strings.Split(m, "\n") {
			nm = append(nm, strings.Split(l, ""))
		}
		maps = append(maps, nm)
	}

	return maps
}

func summary(m [][]string) int {
	_, ogRef := getCount(m, nil)
	contrbution := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			// Copy the map
			nm := make([][]string, len(m))
			for k := 0; k < len(m); k++ {
				nm[k] = make([]string, len(m[0]))
				for l := 0; l < len(m[0]); l++ {
					nm[k][l] = m[k][l]
				}
			}

			// Flip the point
			if nm[i][j] == "#" {
				nm[i][j] = "."
			} else {
				nm[i][j] = "#"
			}

			// Get the new reflection point
			dir, ref := getCount(nm, &ogRef)

			if ref.P1 != ogRef.P1 && ref.P2 != ogRef.P2 && ref.P1 != ref.P2 {
				// fmt.Println("Found a new reflection point", ref, c)
				switch dir {
				case HRef:
					contrbution += (ref.P2 * 100)
				case VRef:
					contrbution += ref.P2
				}
				return contrbution
			}
		}
	}

	return contrbution
}

type Reflection struct {
	P1 int
	P2 int
}

const (
	HRef = iota
	VRef
)

func transpose(m [][]string) [][]string {
	nm := make([][]string, len(m[0]))
	for i := 0; i < len(m[0]); i++ {
		nm[i] = make([]string, len(m))
		for j := 0; j < len(m); j++ {
			nm[i][j] = m[j][i]
		}
	}
	return nm
}

func findReflectionPoints(m [][]string, ignore Reflection) []Reflection {
	i, j := 0, 1
	rowPoints := make([]Reflection, 0)
	for i < len(m) && j < len(m) {
		found := checkRowsMatch(i, j, m)
		if found && (i != ignore.P1 && j != ignore.P2) {
			rowPoints = append(rowPoints, Reflection{P1: i, P2: j})
		}
		i++
		j++
	}

	return rowPoints
}

func getCount(m [][]string, ignore *Reflection) (int, Reflection) {
	// Find the middle reflection point
	// This will be the point where (row or col) i == j
	// Start with Rows

	if ignore == nil {
		ignore = &Reflection{P1: -1, P2: -1}
	}

	rowPoints := findReflectionPoints(m, *ignore)
	transposed := transpose(m)
	colPoints := findReflectionPoints(transposed, *ignore)

	dir, p := getCorrectReflection(rowPoints, m, HRef)
	if p.P1 != -1 && p.P2 != -1 {
		return dir, p
	}

	dir, p = getCorrectReflection(colPoints, transposed, VRef)
	if p.P1 != -1 && p.P2 != -1 {
		return dir, p
	}

	return 0, Reflection{P1: -1, P2: -1}
}

func getCorrectReflection(refPoints []Reflection, m [][]string, reflection int) (int, Reflection) {
	c := 0
	for _, rp := range refPoints {
		c += 2
		dTop := rp.P1 - 1
		dBottom := rp.P2 + 1
		for dTop >= 0 && dBottom < len(m) {
			// Check if the rows are the same
			if !checkRowsMatch(dTop, dBottom, m) {
				c = 0
				break
			}
			dTop -= 1
			dBottom += 1
			c += 1
		}
		if c != 0 {
			return reflection, rp
		}
	}
	return reflection, Reflection{P1: -1, P2: -1}
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

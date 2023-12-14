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

type Reflection struct {
	P1        int
	P2        int
	Direction int
}

const (
	HRef = iota
	VRef
)

func RunDayThirteen() {
	input, err := os.ReadFile("./input/day13.txt")
	if err != nil {
		panic("Failed to read input for day 13")
	}

	fmt.Println("Part 1", partOne(string(input)))
	fmt.Println("Part 2", partTwo(string(input)))
}

func partOne(input string) int {
	maps := parseInput(input)
	sum := 0
	for _, m := range maps {
		ref := getCount(m, nil)
		switch ref.Direction {
		case HRef:
			sum += (ref.P2 * 100)
		case VRef:
			sum += ref.P2
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
	ogRef := getCount(m, nil)
	contrbution := 0
	// Copy the map
	nm := make([][]string, len(m))
	for k := 0; k < len(m); k++ {
		nm[k] = make([]string, len(m[0]))
		for l := 0; l < len(m[0]); l++ {
			nm[k][l] = m[k][l]
		}
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			// Flip the point
			if nm[i][j] == "#" {
				nm[i][j] = "."
			} else {
				nm[i][j] = "#"
			}

			// Get the new reflection point
			ref := getCount(nm, func(p Reflection) bool {
				// Ignore the new reflection point if it is the same as the original
				return p.P1 == ogRef.P1 && p.P2 == ogRef.P2 && p.Direction == ogRef.Direction
			})

			// If a reflection point was found, add the contribution
			if ref.P1 != -1 && ref.P2 != -1 { // The reflection point is not -1
				switch ref.Direction {
				case HRef:
					contrbution += (ref.P2 * 100)
				case VRef:
					contrbution += ref.P2
				}
				return contrbution
			}
			// Flip the point back
			if nm[i][j] == "#" {
				nm[i][j] = "."
			} else {
				nm[i][j] = "#"
			}
		}
	}

	return contrbution
}

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

func findReflectionPoints(m [][]string, direction int, ignore func(p Reflection) bool) []Reflection {
	i, j := 0, 1
	rowPoints := make([]Reflection, 0)
	for i < len(m) && j < len(m) {
		found := checkRowsMatch(i, j, m)
		if found && !ignore(Reflection{P1: i, P2: j, Direction: direction}) {
			rowPoints = append(rowPoints, Reflection{P1: i, P2: j, Direction: direction})
		}
		i++
		j++
	}

	return rowPoints
}

func getCount(m [][]string, ignoreFunc func(p Reflection) bool) Reflection {
	// Find the middle reflection point
	// This will be the point where (row or col) i == j
	if ignoreFunc == nil {
		ignoreFunc = func(p Reflection) bool {
			return false
		}
	}
	rowPoints := findReflectionPoints(m, HRef, ignoreFunc)
	transposed := transpose(m)
	colPoints := findReflectionPoints(transposed, VRef, ignoreFunc)

	p := getCorrectReflection(rowPoints, m, HRef)
	if p.P1 != -1 && p.P2 != -1 {
		return p
	}

	p = getCorrectReflection(colPoints, transposed, VRef)
	if p.P1 != -1 && p.P2 != -1 {
		return p
	}
	return Reflection{P1: -1, P2: -1}
}

func getCorrectReflection(refPoints []Reflection, m [][]string, reflection int) Reflection {
	for _, rp := range refPoints {
		refFound := true
		for dt, db := rp.P1-1, rp.P2+1; dt >= 0 && db < len(m); dt, db = dt-1, db+1 {
			// Check if the rows are the same
			if !checkRowsMatch(dt, db, m) {
				refFound = false
				break
			}
		}
		if refFound {
			return rp
		}
	}
	return Reflection{P1: -1, P2: -1}
}

func checkRowsMatch(i, j int, m [][]string) bool {
	for k := 0; k < len(m[0]); k++ {
		if m[i][k] != m[j][k] {
			return false
		}
	}
	return true
}

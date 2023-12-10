package ten

import (
	"fmt"
	"strings"
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

const complexLooop = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func RunDayTen() {
	partOne(complexLooop)
}

func partOne(input string) int {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, ch := range line {
			// Convert rune to string
			row = append(row, string(ch))
		}
		matrix = append(matrix, row)
	}

	fmt.Println(matrix)
	return 9
}

func partTwo(input string) int {
	return 0
}

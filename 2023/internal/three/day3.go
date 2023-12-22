package three

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

type Number struct {
	// Value holds the value of the number
	Value int
	// Start holds the starting position of the number
	Start int
	// End holds the ending position of the number
	End int
	// Row holds what row the number is in
	Row int
}

type Point struct {
	X     int
	Y     int
	Value rune
}

var testInput string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

var adjPoints [][]int = [][]int{
	{-1, 0},  // Going up
	{1, 0},   // Going down
	{0, -1},  // Going left
	{0, 1},   // Going right
	{-1, -1}, // Top Left
	{-1, 1},  // Top Right
	{1, -1},  // Bottom Left
	{1, 1},   // Bottom Right
}

const GEAR rune = '*'

func RunDayThree() {
	input, err := os.ReadFile("./input/day3.txt")
	if err != nil {
		panic("Could not read day 3 file")
	}
	res := partOne(string(input))
	fmt.Println("Day 3 Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Day 3 Part 2:", res)
}

func partOne(input string) int {
	numbers, specialChars := parseInput(input)

	// Loop through specialChars, look at all adjacent points
	// If a number range falls into that spot, add it to sum
	sum := 0
	seen := make([]Number, 0)
	for _, char := range specialChars {
		for _, adj := range adjPoints {
			dx, dy := adj[0], adj[1]
			x, y := char.X+dx, char.Y+dy
			for _, num := range numbers {
				contains := containsPoint(seen, num)
				if !contains && x == num.Row && y >= num.Start && y <= num.End {
					// Add point to seen
					seen = append(seen, num)
					sum += num.Value
				}
			}
		}
	}
	return sum
}

func partTwo(input string) int {
	numbers, specialChars := parseInput(input)
	gears := iterators.Filter(specialChars, func(p Point) bool {
		return p.Value == GEAR
	})

	sum := 0
	seen := make([]Number, 0)
	for _, gear := range gears {
		count := 0
		for _, adj := range adjPoints {
			dx, dy := adj[0], adj[1]
			x, y := gear.X+dx, gear.Y+dy

			// The gear needs to be adj to EXACTLY two numbers
			for _, num := range numbers {
				contains := containsPoint(seen, num)
				if !contains && x == num.Row && y >= num.Start && y <= num.End {
					// Add point to seen
					seen = append(seen, num)
					count += 1
				}
			}
		}
		if count == 2 {
			// Get the last two points
			sum += (seen[len(seen)-1].Value * seen[len(seen)-2].Value)
		}

	}

	return sum
}

func containsPoint(seen []Number, num Number) bool {
	return slices.ContainsFunc(seen, func(s Number) bool { return s.Row == num.Row && s.End == num.End })
}

func parseInput(input string) ([]Number, []Point) {
	numbers := make([]Number, 0)
	specialChars := make([]Point, 0)

	for row, line := range strings.Split(input, "\n") {
		col := 0
		for col < len(line) {
			// If the character is a number, parse it
			if line[col] >= '0' && line[col] <= '9' {
				start := col
				val := ""
				for col < len(line) && line[col] >= '0' && line[col] <= '9' {
					val += string(line[col])
					col++
				}
				num, err := strconv.Atoi(val)
				if err != nil {
					fmt.Printf("Failed to convert %s to number\n", val)
				}
				numbers = append(numbers, Number{
					Value: num,
					Start: start,
					End:   col - 1,
					Row:   row,
				})
			} else if line[col] != '.' {
				specialChars = append(specialChars, Point{
					X: row,
					Y: col,
					// For part two, store the value
					Value: rune(line[col]),
				})
				col++
			} else {
				col++
			}
		}
	}

	return numbers, specialChars
}

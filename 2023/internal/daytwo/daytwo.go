package daytwo

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
	"github.com/josuemolinamorales/aoc-2023/utils/maps"
)

const (
	MAX_RED   int = 12
	MAX_GREEN int = 13
	MAX_BLUE  int = 14
)

var testInput string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func RunDayTwo() {
	input, err := os.ReadFile("./input/day2input.txt")
	if err != nil {
		panic("Failed to read input file for day 2")
	}
	res := partOne(string(input))
	fmt.Println("Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Part 2:", res)
}

func partOne(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")
	for i, l := range lines {
		// Remove "Game x" from beginning of string
		l = strings.Replace(l, fmt.Sprintf("Game %d: ", (i+1)), "", 1)
		sets := strings.Split(l, "; ")
		validGame := true
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				cs := strings.Split(color, " ")
				count, err := strconv.Atoi(strings.TrimSpace(cs[0]))
				if err != nil {
					fmt.Printf("Failed to convert %s to a number\n", cs[0])
				}

				switch cs[1] {
				case "blue":
					if count > MAX_BLUE {
						validGame = false
					}
				case "red":
					if count > MAX_RED {
						validGame = false
					}
				case "green":
					if count > MAX_GREEN {
						validGame = false
					}
				}

			}
		}
		if validGame {
			sum += (i + 1)
		}
	}

	return sum
}

func partTwo(input string) int {
	sum := 0

	lines := strings.Split(input, "\n")
	for i, l := range lines {
		// Remove "Game x" from beginning of string
		l = strings.Replace(l, fmt.Sprintf("Game %d: ", (i+1)), "", 1)
		sets := strings.Split(l, "; ")

		colorMax := map[string]int{
			"blue":  0,
			"red":   0,
			"green": 0,
		}
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				cs := strings.Split(color, " ")
				count, err := strconv.Atoi(strings.TrimSpace(cs[0]))
				if err != nil {
					fmt.Printf("Failed to convert %s to a number\n", cs[0])
				}
				c := cs[1]

				colorMax[c] = int(math.Max(float64(colorMax[c]), float64(count)))
			}
		}

		product := iterators.Product(maps.Values(colorMax))

		sum += product
	}

	return sum
}

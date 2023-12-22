package one

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/josuemolinamorales/aoc-2023/utils"
	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
	"github.com/josuemolinamorales/aoc-2023/utils/maps"
)

func RunDayOne() {
	input, err := os.ReadFile("./input/day1.txt")
	if err != nil {
		panic("Input file not found")
	}

	res := partOne(string(input))
	fmt.Println("Day 1 Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Day 1 Part 2:", res)
}

func partOne(input string) int {
	// For each line of text, get the first and last digit and combine them
	// If only one digit found, use it twice
	lines := strings.Split(input, "\n")
	sum := 0

	for _, l := range lines {
		chars := strings.Split(l, "")
		digits := iterators.Filter(chars, func(c string) bool {
			return utils.IsDigit(c)
		})
		if len(digits) == 0 {
			continue
		}
		var num string
		if len(digits) > 1 {
			num = string(digits[0]) + string(digits[len(digits)-1])
		} else {
			num = string(digits[0]) + string(digits[0])
		}

		converted, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("ERROR: Could not convert %s to num\n", num)
		}
		sum += converted
	}

	return sum
}

func partTwo(input string) int {
	// Digits as strings
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	lines := strings.Split(input, "\n")
	newInput := ""
	for _, l := range lines {
		var ids map[int]string = make(map[int]string)
		for _, k := range maps.Keys(digits) {
			if strings.Contains(l, k) {
				// Find all indices of the digit word
				indices := utils.FindAllIndex(l, k)
				for _, i := range indices {
					ids[i] = strconv.Itoa(digits[k])
				}
			}
		}
		num := ""
		for i, c := range l {
			if d, ok := ids[i]; ok {
				num += d
				continue
			}
			if unicode.IsDigit(c) {
				num += string(c)
			}
		}

		newInput += num + "\n"
	}

	sum := partOne(newInput)
	return sum
}

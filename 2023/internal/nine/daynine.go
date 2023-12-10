package nine

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func RunDayNine() {
	input, err := os.ReadFile("./input/day9.txt")
	if err != nil {
		panic("Failed to read day 9 file")
	}
	fmt.Println("Part 1: ", partOne(string(input)))
	fmt.Println("Part 2: ", partTwo(string(input)))
}

func parseInput(input string) [][]int {
	numbers := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		nums := make([]int, 0)
		for _, num := range strings.Split(line, " ") {
			n, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Failed to convert", num, "to int")
			}
			nums = append(nums, n)
		}
		numbers = append(numbers, nums)
	}
	return numbers
}

func partOne(input string) int {
	numbers := parseInput(input)

	// Find the sequence then return the next number
	sum := 0
	for _, nums := range numbers {
		sum += findSequence(nums, 1)
	}
	return sum
}

func partTwo(input string) int {
	numbers := parseInput(input)

	sum := 0
	for _, nums := range numbers {
		sum += findSequence(nums, 2)
	}

	return sum
}

func findSequence(numbers []int, part int) int {
	if iterators.Every(numbers, func(n int) bool { return n == 0 }) {
		if part == 1 {
			return numbers[len(numbers)-1]
		}
		return numbers[0]
	}
	differences := make([]int, 0)
	for i := 0; i < len(numbers)-1; i++ {
		differences = append(differences, (numbers[i+1] - numbers[i]))
	}
	diff := findSequence(differences, part)

	if part == 1 {
		return diff + numbers[len(numbers)-1]
	}
	return numbers[0] - diff
}

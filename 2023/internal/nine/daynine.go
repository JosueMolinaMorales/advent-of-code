package nine

import (
	"fmt"
	"strconv"
	"strings"
)

const testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func RunDayNine() {
	partOne(testInput)
}

func partOne(input string) int {
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

	fmt.Println(numbers)

	// Find the sequence then return the next number
	for _, nums := range numbers {
		fmt.Println(findSequence(nums))
	}
	return 0
}

func findSequence(numbers []int) int {
	return 0
}

func partTwo(input string) int {
	return 0
}

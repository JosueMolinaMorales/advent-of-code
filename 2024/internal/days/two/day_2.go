package two

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

const (
	Increasing = 1
	Decreasing = -1
)

func SolveDay2() {
	res := solvePartOne()
	fmt.Println("Day 2 Part 1: ", res)
	res = solvePartTwo()
	fmt.Println("Day 2 Part 2: ", res)
}

func setup() [][]int {
	rawReports, err := util.LoadFileAsString("./inputs/day_2.txt")
	if err != nil {
		panic(err)
	}

	var reports [][]int
	for _, line := range strings.Split(strings.TrimSpace(rawReports), "\n") {
		numbers := strings.Fields(line) // Split by spaces
		report := make([]int, len(numbers))
		for i, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			report[i] = n
		}
		reports = append(reports, report)
	}

	return reports
}

func solvePartOne() int {
	reports := setup()
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	return count
}

func solvePartTwo() int {
	reports := setup()
	count := 0
	for _, report := range reports {
		if isSafeRemove(report) {
			count++
		}
	}
	return count
}

func direction(x int) int {
	switch {
	case x < 0:
		return Increasing
	case x > 0:
		return Decreasing
	default:
		return 0
	}
}

func isSafe(report []int) bool {
	dir := direction(report[0] - report[1])
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if newDir := direction(diff); newDir != dir {
			return false
		}
		if absDiff := int(math.Abs(float64(diff))); absDiff < 1 || absDiff > 3 {
			return false
		}
	}
	return true
}

func isSafeRemove(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := range report {
		portion := slices.Delete(slices.Clone(report), i, i+1)
		if isSafe(portion) {
			return true
		}
	}
	return false
}

package two

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func Solve_day_2() {
	res := solve_part_one()
	fmt.Println("Day 2 Part 1: ", res)
	res = solve_part_two()
	fmt.Println("Day 2 Part 2: ", res)
}

func setup() [][]int {
	rawReports, err := util.LoadFileAsString("./inputs/day_2.txt")
	if err != nil {
		panic(err)
	}

	// A report is good if:
	// - all numbers are increasing or decreasing
	// - adjacent numbers are differ at least 1 & at most 3
	reports := make([][]int, 0)
	for _, rr := range strings.Split(rawReports, "\n") {
		report := make([]int, 0)
		for _, s := range strings.Split(rr, " ") {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			report = append(report, n)
		}
		reports = append(reports, report)
	}

	return reports
}

func solve_part_one() int {
	reports := setup()
	res := 0
	for _, report := range reports {
		if isSafe(report) {
			res += 1
		}
	}
	return res
}

func solve_part_two() int {
	reports := setup()
	res := 0
	for _, report := range reports {
		if isSafeRemove(report) {
			res += 1
		}
	}
	return res
}

func direction(x int) int {
	if x < 0 {
		return 1
	} else if x > 0 {
		return -1
	}
	return 0
}

func isSafe(report []int) bool {
	// check the first numbers
	dir := direction(report[0] - report[1])
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		newDir := direction(diff)
		if dir != newDir {
			return false
		}
		diff = int(math.Abs(float64(diff)))
		if diff <= 0 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeRemove(report []int) bool {
	if isSafe(report) {
		return true
	}

	for i := 0; i < len(report); i++ {
		portion := slices.Delete(slices.Clone(report), i, i+1)
		if isSafe(portion) {
			return true
		}
	}
	return false
}

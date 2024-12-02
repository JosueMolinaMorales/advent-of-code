package two

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func Solve_day_2() {
	res := solve_part_one()
	fmt.Println(res)
}

func solve_part_one() int {
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

	fmt.Println(reports)

	return 0
}

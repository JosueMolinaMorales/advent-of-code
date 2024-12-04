package three

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

const (
	PATTERN        = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
	NUMBER_PATTERN = `\d{1,3}`
)

func SolveDay3() {
	res := solvePartOne()
	fmt.Println("Day 3 Part 1: ", res)
	res = solvePartTwo()
	fmt.Println("Day 3 Part 2: ", res)
}

func solve(part2 bool) int {
	memory, err := util.LoadFileAsString("./inputs/day_3.txt")
	if err != nil {
		panic(err)
	}
	re, err := regexp.Compile(PATTERN)
	if err != nil {
		panic(err)
	}
	numRe, err := regexp.Compile(NUMBER_PATTERN)
	if err != nil {
		panic(err)
	}
	matches := re.FindAllString(memory, -1)

	res := 0
	enabled := true
	for _, match := range matches {
		if part2 && match == "don't()" {
			enabled = false
		} else if part2 && match == "do()" {
			enabled = true
		} else if strings.HasPrefix(match, "mul") && enabled {
			nums := make([]int, 2)
			numsRaw := numRe.FindAllString(match, 2)
			for i := 0; i < 2; i++ {
				nums[i], err = strconv.Atoi(numsRaw[i])
				if err != nil {
					panic(err)
				}
			}
			res += nums[0] * nums[1]
		}
	}

	return res
}

func solvePartOne() int {
	return solve(false)
}

func solvePartTwo() int {
	return solve(true)
}

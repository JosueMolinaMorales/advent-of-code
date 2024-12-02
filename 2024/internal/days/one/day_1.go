package one

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay1() {
	rawLists, err := util.LoadFileAsString("./inputs/day_1.txt")
	if err != nil {
		panic(err)
	}

	res := solvePartOne(rawLists)
	fmt.Println("Day 1 Part 1: ", res)
	res = solvePartTwo(rawLists)
	fmt.Println("Day 1 Part 2: ", res)
}

func setup(rawLists string) (leftList []int, rightList []int) {
	// Split the lists into two
	for _, locations := range strings.Split(rawLists, "\n") {
		if locations == "" {
			continue
		}

		l := strings.Split(locations, "   ")
		leftLocation, err := strconv.Atoi(l[0])
		if err != nil {
			panic(err)
		}

		rightLocation, err := strconv.Atoi(l[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftLocation)
		rightList = append(rightList, rightLocation)
	}

	return leftList, rightList
}

func solvePartOne(rawLists string) int {
	leftList, rightList := setup(rawLists)
	// Sort both lists
	slices.Sort(leftList)
	slices.Sort(rightList)

	// Compare and get the differences
	n := len(leftList)
	res := 0

	for i := 0; i < n; i++ {
		res += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return res
}

func solvePartTwo(rawLists string) int {
	leftList, rightList := setup(rawLists)

	rep := make(map[int]int)
	for _, loc := range rightList {
		rep[loc] += 1
	}

	res := 0
	for _, loc := range leftList {
		res += loc * rep[loc]
	}

	return res
}

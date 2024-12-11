package eleven

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
)

func SolveDay11() {
	fmt.Println(solve(25))
	fmt.Println(solve(75))
}

func setup() []int {
	rawStones, err := util.LoadFileAsString("./inputs/day_11.txt")
	if err != nil {
		panic(err)
	}
	stones := []int{}
	for _, s := range strings.Split(rawStones, " ") {
		stones = append(stones, util.ToInt(s))
	}
	return stones
}

func solve(blinks int) int {
	stones := setup()
	count := 0
	memo := map[types.Vector]int{}
	for _, num := range stones {
		count += getRockCount(num, blinks, memo)
	}
	return count
}

func getRockCount(num int, blinks int, memo map[types.Vector]int) int {
	if blinks == 0 {
		return 1
	}

	if cached := memo[*types.NewVector(num, blinks)]; cached != 0 {
		return cached
	}

	count := 0
	for _, stone := range blink(num) {
		count += getRockCount(stone, blinks-1, memo)
	}
	memo[*types.NewVector(num, blinks)] = count

	return count
}

func blink(num int) []int {
	if num == 0 {
		return []int{1}
	}

	s := strconv.Itoa(num)
	if len(s)%2 == 0 {
		s1 := util.ToInt(s[:len(s)/2])
		s2 := util.ToInt(s[len(s)/2:])

		return []int{s1, s2}
	}

	return []int{num * 2024}
}

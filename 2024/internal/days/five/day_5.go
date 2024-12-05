package five

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/sets"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay5() {
	p1, p2 := solve()
	fmt.Println("Day 5 Part 1: ", p1)
	fmt.Println("Day 5 Part 2: ", p2)
}

func solve() (int, int) {
	rules, updates := setup()

	comp := func(a, b int) int {
		if rules[a] == nil {
			return 0
		}
		if rules[a].Contains(b) {
			return -1
		}
		return 1
	}
	part1 := 0
	part2 := 0
	for _, update := range updates {
		if slices.IsSortedFunc(update, comp) {
			part1 += update[len(update)/2]
		} else {
			slices.SortFunc(update, comp)
			part2 += update[len(update)/2]
		}
	}
	return part1, part2
}

func setup() (map[int]sets.Set, [][]int) {
	rawManual, err := util.LoadFileAsString("./inputs/day_5.txt")
	if err != nil {
		panic(err)
	}

	manualParts := strings.Split(rawManual, "\n\n")

	rawRules := manualParts[0]
	rawUpdates := manualParts[1]

	rules := make(map[int]sets.Set, 0)
	for _, rule := range strings.Split(rawRules, "\n") {
		parts := strings.Split(rule, "|")
		before, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if rules[before] == nil {
			rules[before] = hashset.New()
		}

		rules[before].Add(after)
	}

	updates := make([][]int, 0)
	for _, update := range strings.Split(rawUpdates, "\n") {
		parts := strings.Split(update, ",")
		update := make([]int, 0)
		for _, s := range parts {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	return rules, updates
}

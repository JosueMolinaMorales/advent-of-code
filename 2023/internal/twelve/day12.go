package twelve

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const testInput = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func RunDayTwelve() {
	input, err := os.ReadFile("./input/day12.txt")
	if err != nil {
		panic("failed to ready day 12 file")
	}
	fmt.Println("Day 12 Part 1", partOne(string(input)))
	fmt.Println("Day 12 Part 2", partTwo(string(input)))
}

func partOne(input string) int {
	springs, groups := parseInput(input)

	arrangements := 0
	memo := make(map[string]int, 0)
	for i, s := range springs {
		arrangements += count(s, &memo, groups[i])
	}

	return arrangements
}

func count(input string, memo *map[string]int, group []int) int {
	if len(input) == 0 {
		// If the input is empty and the group is empty, then we have a valid arrangement
		if len(group) == 0 {
			return 1
		}
		// If the input is empty and the group is not empty, then we don't have a valid arrangement
		return 0
	}

	if len(group) == 0 {
		// If the group is empty but we still have '#' in the input, then we don't have a valid arrangement
		// Since there are still more broken springs
		if strings.Contains(input, "#") {
			return 0
		}
		// If the group is empty and there are no more '#' in the input, then we have a valid arrangement
		return 1
	}

	// Make a key for the memo
	gs := make([]string, 0)
	for _, g := range group {
		gs = append(gs, strconv.Itoa(g))
	}
	key := input + strings.Join(gs, ",")

	// Check if the key exists in the memo
	if _, exists := (*memo)[key]; exists {
		return (*memo)[key]
	}

	result := 0
	// If the first character is a '.' or a '?', then we can pretend that "?" is a "." and count the number of arrangements
	if input[0] == '.' || input[0] == '?' {
		result += count(input[1:], memo, group)
	}

	// If the first character is a '#' or a '?', then we can pretend that "?" is a "#" and count the number of arrangements
	if input[0] == '#' || input[0] == '?' {
		if group[0] <= len(input) && // If the group is not out of bounds
			!strings.Contains(input[:group[0]], ".") && // If there are no '.' from the beginning of the input to the end of the group
			(group[0] == len(input) || string(input[group[0]]) != "#") { // If the group is at the end of the input or the next character is not a '#'
			substr := ""
			if group[0]+1 < len(input) {
				substr = input[group[0]+1:]
			}
			// Recursively count the number of arrangements to the substring
			result += count(substr, memo, group[1:])
		}
	}

	(*memo)[key] = result
	return result
}

func parseInput(input string) ([]string, [][]int) {
	springs := make([]string, 0)
	groups := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		ls := strings.Split(line, " ")
		springs = append(springs, ls[0])
		groupsSplit := strings.Split(ls[1], ",")
		nums := make([]int, 0)
		for _, ch := range groupsSplit {
			num, _ := strconv.Atoi(ch)
			nums = append(nums, num)
		}
		groups = append(groups, nums)
	}

	return springs, groups
}

func partTwo(input string) int {
	springs, groups := parseInput(input)
	// Expand string
	expandedSprings := make([]string, 0)
	expandedGroups := make([][]int, 0)
	for i, s := range springs {
		expanded := make([]string, 0)
		for i := 0; i < 5; i++ {
			expanded = append(expanded, s)
		}
		es := strings.Join(expanded, "?")
		expandedSprings = append(expandedSprings, es)

		// Expand groups
		expandedGroup := make([]int, 0)
		for j := 0; j < 5; j++ {
			expandedGroup = append(expandedGroup, groups[i]...)
		}

		expandedGroups = append(expandedGroups, expandedGroup)
	}

	// Break string into substrings
	arrangements := 0
	memo := make(map[string]int, 0)
	for i, ex := range expandedSprings {
		arrangements += count(ex, &memo, expandedGroups[i])
	}

	return arrangements
}

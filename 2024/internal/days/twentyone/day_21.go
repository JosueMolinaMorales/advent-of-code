package twentyone

import (
	"fmt"
	"math"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

var NumMap = map[string]Coord{
	"A": {2, 0},
	"0": {1, 0},
	"1": {0, 1},
	"2": {1, 1},
	"3": {2, 1},
	"4": {0, 2},
	"5": {1, 2},
	"6": {2, 2},
	"7": {0, 3},
	"8": {1, 3},
	"9": {2, 3},
}

var DirMap = map[string]Coord{
	"A": {2, 1},
	"^": {1, 1},
	"<": {0, 0},
	"v": {1, 0},
	">": {2, 0},
}

func SolveDay21() {
	fmt.Println("Day 21 Part 1: ", solve(2))
	fmt.Println("Day 21 Part 2: ", solve(25))
}

func solve(robots int) int {
	input, err := util.LoadFileAsString("./inputs/day_21.txt")
	if err != nil {
		panic(err)
	}
	return getSequence(strings.Split(input, "\n"), robots)
}

type Coord struct {
	x, y int
}

func getSequence(input []string, robotCount int) int {
	total := 0
	cache := make(map[string][]int)

	for _, line := range input {
		chars := strings.Split(line, "")
		moves := getNumPadSequence(chars, "A", NumMap)
		length := countSequences(moves, robotCount, 1, cache, DirMap)
		total += util.ToInt(strings.TrimLeft(line[:3], "0")) * length
	}

	return total
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func getNumPadSequence(input []string, start string, numMap map[string]Coord) []string {
	curr := numMap[start]
	seq := []string{}

	for _, char := range input {
		dest := numMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		// Build horizontal moves
		for i := 0; i < abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		// Build vertical moves
		for i := 0; i < abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		// Order moves based on position
		if curr.y == 0 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if curr.x == 0 && dest.y == 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func countSequences(input []string, maxRobots, robot int, cache map[string][]int, dirMap map[string]Coord) int {
	key := strings.Join(input, "")
	if val, ok := cache[key]; ok && robot <= len(val) && val[robot-1] != 0 {
		return val[robot-1]
	}

	if _, ok := cache[key]; !ok {
		cache[key] = make([]int, maxRobots)
	}

	seq := getDirPadSequence(input, "A", dirMap)
	if robot == maxRobots {
		return len(seq)
	}

	steps := splitSequence(seq)
	count := 0
	for _, step := range steps {
		c := countSequences(step, maxRobots, robot+1, cache, dirMap)
		count += c
	}

	cache[key][robot-1] = count
	return count
}

func getDirPadSequence(input []string, start string, dirMap map[string]Coord) []string {
	curr := dirMap[start]
	seq := []string{}

	for _, char := range input {
		dest := dirMap[char]
		dx, dy := dest.x-curr.x, dest.y-curr.y

		horiz, vert := []string{}, []string{}

		for i := 0; i < abs(dx); i++ {
			if dx >= 0 {
				horiz = append(horiz, ">")
			} else {
				horiz = append(horiz, "<")
			}
		}

		for i := 0; i < abs(dy); i++ {
			if dy >= 0 {
				vert = append(vert, "^")
			} else {
				vert = append(vert, "v")
			}
		}

		if curr.x == 0 && dest.y == 1 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else if curr.y == 1 && dest.x == 0 {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		} else if dx < 0 {
			seq = append(seq, horiz...)
			seq = append(seq, vert...)
		} else {
			seq = append(seq, vert...)
			seq = append(seq, horiz...)
		}

		curr = dest
		seq = append(seq, "A")
	}
	return seq
}

func splitSequence(input []string) [][]string {
	var result [][]string
	var current []string

	for _, char := range input {
		current = append(current, char)
		if char == "A" {
			result = append(result, current)
			current = []string{}
		}
	}
	return result
}

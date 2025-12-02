package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day1() {
	fmt.Println("2025 Day 1 Part 1:", part1("inputs/day_1/input.txt"))
	fmt.Println("2025 Day 1 Part 2:", part2("inputs/day_1/input.txt"))
}

func part1(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 1 Part 1: %s", err)
	}

	// Circular Dial. WIll be using mod
	// Start at 50, spin the dial either Left or Right
	// By a certain number. Numbers are 0 - 99, everytime it goes over or under it loops back around
	// Count the number of times it hits 0
	pos := 50
	res := 0
	for _, act := range input {
		// Get the action and the distance
		dir := string(act[0])
		dist, err := strconv.Atoi(act[1:])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 1 Part 1: %s", err)
		}

		if dir == "L" {
			pos -= dist
		}
		if dir == "R" {
			pos += dist
		}

		// Mod
		pos = util.EuclideanMod(pos, 100)

		if pos == 0 {
			res += 1
		}

		// log.Printf("The dial is rotated %s to point at %d", act, pos)

	}
	return res
}

// 6431 TOO HIGH
// 5450 TOO LOW
func part2(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 1 Part 1: %s", err)
	}

	// Now we want to count the number of times we pass 0 + the number of times we hit 0
	// During an action we could pass 0 multiple times
	// Would just be taking the dividend of (pos + dist) / 100
	pos := 50
	res := 0
	for _, act := range input {
		// Get the action and the distance
		dir := string(act[0])
		dist, err := strconv.Atoi(act[1:])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 1 Part 1: %s", err)
		}

		var direction int
		if dir == "L" {
			direction = -1
		} else {
			direction = 1
		}

		for range dist {
			pos = pos + (direction * 1)
			pos = util.EuclideanMod(pos, 100)
			if pos == 0 {
				res += 1
			}
		}
	}

	return res
}

package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day2() {
	fmt.Println("2025 Day 2 Part 1:", day2Part1("inputs/day_2/input.txt"))
	fmt.Println("2025 Day 2 Part 2:", day2Part2("inputs/day_2/input.txt"))
}

func day2Part1(path string) int {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
	}

	ranges := strings.Split(input, ",")
	// For each section, we want to find the invalid IDs
	// An invalid ID is one where the numbers repeat: 22, 1111, 1212, 4545, etc.
	// SO they will only be even length

	// First method: convert each end range to number, loop through the range, and
	// find all numbers that repeat
	invalidIDs := make([]string, 0)
	for _, r := range ranges {
		// Convert the numbers
		endings := strings.Split(r, "-")
		start, err := strconv.Atoi(endings[0])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}
		end, err := strconv.Atoi(endings[1])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}

		for i := start; i <= end; i++ {
			numStr := strconv.Itoa(i)
			mid := len(numStr) / 2

			a := numStr[0:mid]
			b := numStr[mid:]

			if a == b {
				invalidIDs = append(invalidIDs, numStr)
			}
		}
	}

	res := 0
	for _, numStr := range invalidIDs {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}
		res += num
	}

	return res
}

func day2Part2(path string) int {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
	}
	ranges := strings.Split(input, ",")
	// Now an invalid ID is one where there is a sequence of numbers
	// that repeats at least twice: 11, 1212, 135135, etc.
	invalidIDs := make(map[string]bool)
	for _, r := range ranges {
		// Convert the numbers
		endings := strings.Split(r, "-")
		start, err := strconv.Atoi(endings[0])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}
		end, err := strconv.Atoi(endings[1])
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}
		for i := start; i <= end; i++ {
			// For a given string length, i want to find the GCF of the length
			// this GCF will have to be the length of the sequence that is being repeated
			numStr := strconv.Itoa(i)
			length := len(numStr)
			tried := map[int]bool{}
			for j := 1; j <= length/2; j++ {
				gcf := util.GCF(length, j)
				// Check to see if we have tried this already
				if _, ok := tried[gcf]; ok {
					continue
				} else {
					tried[gcf] = true
				}
				// Create a string with the substring
				sub := numStr[:gcf]
				newStrArr := []string{}
				for k := 0; k < (length / j); k++ {
					newStrArr = append(newStrArr, sub)
				}
				newStr := strings.Join(newStrArr, "")
				if newStr == numStr {
					invalidIDs[newStr] = true
				}
			}
		}
	}

	res := 0
	for k := range invalidIDs {
		num, err := strconv.Atoi(k)
		if err != nil {
			log.Fatalf("ERROR: 2025 Day 2 Part 1: %s", err)
		}
		res += num
	}

	return res
}

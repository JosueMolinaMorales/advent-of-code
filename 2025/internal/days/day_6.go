package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

func Day6() {
	fmt.Println("2025 Day 6 Part 1:", day6Part1("inputs/day_6/input.txt"))
	fmt.Println("2025 Day 6 Part 2:", day6Part2("inputs/day_6/input.txt"))
}

func day6Part1(path string) int {
	equations := parseMathProblems(path)

	res := 0
	// Process each column (problem) from left to right
	for col := 0; col < len(equations[0]); col++ {
		numbers := []int{}
		// Read numbers vertically (top to bottom)
		for row := 0; row < len(equations)-1; row++ {
			num, err := strconv.Atoi(equations[row][col])
			if err != nil {
				log.Fatalf("ERROR: 2025 Day 6 Part 1: %s", err)
			}
			numbers = append(numbers, num)
		}

		// Apply the operation
		result := numbers[0]
		op := equations[len(equations)-1][col]
		for i := 1; i < len(numbers); i++ {
			switch op {
			case "+":
				result += numbers[i]
			case "*":
				result *= numbers[i]
			}
		}
		res += result
	}
	return res
}

func day6Part2(path string) int {
	equations, ops := parseMathProblemsRightToLeft(path)

	result := 0
	for i, op := range ops {
		digits := equations[i]
		res := digits[0]
		for j := 1; j < len(digits); j++ {
			switch op {
			case "*":
				res *= digits[j]
			case "+":
				res += digits[j]
			}
		}
		result += res
	}

	return result
}

func parseMathProblemsRightToLeft(path string) ([][]int, []string) {
	input, err := io.ReadFileAs2DArray(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 6 Parsing: %s", err)
	}

	// Process columns from right to left
	col := len(input[0]) - 1
	digits := []int{}
	equations := [][]int{}
	ops := []string{}

	for col >= 0 {
		sb := strings.Builder{}

		// Read column vertically (top to bottom)
		for row := range input {
			sb.WriteString(input[row][col])
		}

		s := sb.String()

		// Extract operator from bottom row
		op := string(s[len(s)-1])
		// Remove operator row
		s = s[0 : len(s)-1]

		n, err := strconv.Atoi(strings.Trim(s, " "))
		if err != nil {
			log.Fatalf("ERROR: Could not parse string: %s", err)
		}
		digits = append(digits, n)

		// Check if we hit an operator (end of problem)
		if op == "*" || op == "+" {
			equations = append(equations, digits)
			ops = append(ops, op)
			digits = []int{}
			col-- // Skip separator column
		}
		col--
	}

	return equations, ops
}

func parseMathProblems(path string) [][]string {
	input, err := io.ReadFileAsString(path)
	if err != nil {
		log.Fatalf("ERROR: 2025 Day 6 Parsing: %s", err)
	}

	data := [][]string{}
	for _, line := range strings.Split(input, "\n") {
		row := []string{}
		for _, val := range strings.Split(line, " ") {
			if len(strings.Trim(val, " ")) == 0 {
				continue
			}
			row = append(row, val)
		}
		data = append(data, row)
	}

	return data
}

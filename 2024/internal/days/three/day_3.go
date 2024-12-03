package three

import (
	"fmt"
	"strconv"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay3() {
	res := solvePartOne()
	fmt.Println(res)
	res = solvePartTwo()
	fmt.Println(res)
}

func solvePartOne() int {
	memory, err := util.LoadFileAsString("./inputs/day_3.txt")
	if err != nil {
		panic(err)
	}
	// Scan the memory, needs to match: mul(xxx,xxx)
	// where xxx are 1-3 digits
	res := 0
	i := 0
	for i < len(memory) {
		if nextWord(memory, i, "mul(") {
			i += len("mul(")
			if isValid, ans := isValidMul(memory, i); isValid {
				res += ans
			}
		}
		i += 1
	}
	return res
}

func solvePartTwo() int {
	memory, err := util.LoadFileAsString("./inputs/day_3.txt")
	if err != nil {
		panic(err)
	}
	res := 0
	mulEnabled := true
	i := 0
	for i < len(memory) {
		if nextWord(memory, i, "mul(") && mulEnabled {
			i += len("mul(")
			if isValid, ans := isValidMul(memory, i); isValid {
				res += ans
			}
		}
		if nextWord(memory, i, "don't()") {
			mulEnabled = false
			i += len("don't()")
			continue
		}
		if nextWord(memory, i, "do()") {
			mulEnabled = true
			i += len("do()")
			continue
		}

		i++

	}
	return res
}

func nextWord(memory string, i int, word string) bool {
	for j, ch := range word {
		if memory[i+j] != byte(ch) {
			return false
		}
	}
	return true
}

func isValidMul(memory string, i int) (bool, int) {
	// Next is 1-3 digits
	a, al := nextDigits(memory, i)
	if a < 0 {
		return false, -1
	}
	// Next is ','
	if !next(memory, i+al, ',') {
		return false, -1
	}
	// Next is 1-3 digits
	b, bl := nextDigits(memory, i+al+1)
	if b < 0 {
		return false, -1
	}
	// Next is ')'
	if !next(memory, i+al+bl+1, ')') {
		return false, -1
	}

	return true, a * b
}

func next(memory string, i int, char rune) bool {
	return rune(memory[i]) == char
}

func nextDigits(memory string, i int) (int, int) {
	// Digits can be 1-3 characters long
	digitStr := ""
	for j := 0; j < 3; j++ {
		potDigit := string(memory[i+j])
		if _, err := strconv.Atoi(potDigit); err == nil {
			digitStr += potDigit
		} else {
			break
		}
	}
	digits, err := strconv.Atoi(digitStr)
	if err != nil {
		return -1, 0
	}
	return digits, len(digitStr)
}

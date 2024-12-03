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
	for i, ch := range memory {
		if ch == 'm' {
			if isValid, ans := isValidMul(memory, i); isValid {
				res += ans
			}
		}
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
	// Next is 'u'
	if !next(memory, i+1, 'u') {
		return false, -1
	}
	// Next is 'l'
	if !next(memory, i+2, 'l') {
		return false, -1
	}
	// Next is '('
	if !next(memory, i+3, '(') {
		return false, -1
	}
	// Next is 1-3 digits
	a, al := nextDigits(memory, i+4)
	if a < 0 {
		return false, -1
	}
	// Next is ','
	// TODO: figure out what to add to i
	if !next(memory, i+al+4, ',') {
		return false, -1
	}
	// Next is 1-3 digits
	b, bl := nextDigits(memory, i+al+5)
	if b < 0 {
		return false, -1
	}
	// Next is ')'
	// TODO: figure out what to add to i
	if !next(memory, i+al+bl+5, ')') {
		return false, -1
	}

	return true, a * b
}

func next(memory string, i int, char rune) bool {
	// log.Printf("[DEBUG] LOOKING AT INDEX: %d for CHAR: %c", i, char)
	return rune(memory[i]) == char
}

func nextDigits(memory string, i int) (int, int) {
	// log.Printf("[DEBUG] LOOKING FOR DIGITS START AT INDEX: %d", i)
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
	// log.Printf("Digit string found: %s", digitStr)
	digits, err := strconv.Atoi(digitStr)
	if err != nil {
		return -1, 0
	}
	return digits, len(digitStr)
}

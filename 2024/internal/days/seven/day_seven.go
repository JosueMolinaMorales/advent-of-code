package seven

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay7() {
	fmt.Println(solve([]string{"+", "*"}))
	fmt.Println(solve([]string{"+", "*", "||"}))
}

func solve(operators []string) int {
	eqs := setup()
	correct := 0
	for _, v := range eqs {
		for _, op := range operators {
			if compute(v[1:], v[0], op, operators) {
				correct += v[0]
				break
			}
		}
	}
	return correct
}

func setup() [][]int {
	rawEqs, err := util.LoadFileAsString("./inputs/day_7.txt")
	if err != nil {
		panic(err)
	}

	eqs := make([][]int, 0)
	for _, line := range strings.Split(rawEqs, "\n") {
		parts := strings.Split(line, ": ")
		row := make([]int, 0)
		row = append(row, util.ToInt(parts[0]))
		for _, n := range strings.Split(parts[1], " ") {
			row = append(row, util.ToInt(n))
		}

		eqs = append(eqs, row)
	}

	return eqs
}

func compute(arr []int, target int, operator string, operators []string) bool {
	if len(arr) == 1 {
		return arr[0] == target
	}
	newValue := performOperation(arr[0], arr[1], operator)
	// Create a new slice with the computed value replacing the first two elements.
	newArr := append([]int{newValue}, arr[2:]...)
	for _, op := range operators {
		if compute(newArr, target, op, operators) {
			return true
		}
	}
	return false
}

func performOperation(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "*":
		return a * b
	case "||":
		multiplier := 1
		for b/multiplier > 0 {
			multiplier *= 10
		}
		return a*multiplier + b
	default:
		return 0
	}
}

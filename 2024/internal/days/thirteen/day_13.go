package thirteen

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"gonum.org/v1/gonum/mat"
)

func SolveDay13() {
	fmt.Println("Day 13 Part 1: ", solve(false))
	fmt.Println("Day 13 Part 2: ", solve(true))
}

func setup() [][]float64 {
	input, err := util.LoadFileAsString("./inputs/day_13.txt")
	if err != nil {
		panic(err)
	}

	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	equations := [][]float64{}
	for _, section := range strings.Split(input, "\n\n") {
		matches := re.FindAllString(section, -1)
		equation := []float64{}
		for _, s := range matches {
			equation = append(equation, float64(util.ToInt(s)))
		}
		equations = append(equations, equation)
	}

	return equations
}

func solve(part2 bool) int {
	equations := setup()
	ans := 0.
	for _, eq := range equations {
		A := mat.NewDense(2, 2, []float64{eq[0], eq[2], eq[1], eq[3]})
		var b *mat.VecDense
		if !part2 {
			b = mat.NewVecDense(2, []float64{eq[4], eq[5]})
		} else {
			b = mat.NewVecDense(2, []float64{eq[4] + 10000000000000, eq[5] + 10000000000000})
		}
		// Solve the system
		var vd mat.VecDense
		err := vd.SolveVec(A, b)
		if err != nil {
			panic(err)
		}
		x, y := math.Round(vd.AtVec(0)*100)/100, math.Round(vd.AtVec(1)*100)/100
		if (isWholeNumber(x) && isWholeNumber(y)) &&
			(x > 0 && y > 0) &&
			((x <= 100 && y <= 100) || part2) {
			ans += 3*x + y
		}
	}

	return int(ans)
}

func isWholeNumber(num float64) bool {
	return num == math.Trunc(num)
}

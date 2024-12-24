package twentyfour

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

type Gate struct {
	Value    int
	Operator string
	O1       string
	O2       string
}

func SolveDay24() {
	fmt.Println(solvePartOne())
}

func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_24.txt")
	if err != nil {
		panic(err)
	}

	gates := map[string]*Gate{}
	parts := strings.Split(input, "\n\n")
	for _, sg := range strings.Split(parts[0], "\n") {
		p := strings.Split(sg, ": ")
		gates[p[0]] = &Gate{
			Value: util.ToInt(p[1]),
		}
	}

	for _, exp := range strings.Split(parts[1], "\n") {
		p := strings.Split(exp, " ")
		gates[p[4]] = &Gate{
			Value:    -1,
			Operator: p[1],
			O1:       p[0],
			O2:       p[2],
		}
	}
	start := time.Now()
	keys := []string{}
	for k, v := range gates {
		if k[0] == 'z' {
			keys = append(keys, k)
		}
		if v.Value != -1 {
			continue
		}
		// Eval Exp
		gates[k].Value = eval(gates[k].O1, gates[k].O2, gates[k].Operator, &gates)
	}

	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(b, a)
	})

	sb := strings.Builder{}
	for _, k := range keys {
		sb.WriteString(strconv.Itoa(gates[k].Value))
	}
	bin := sb.String()

	ans, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1 took: ", time.Since(start))
	return int(ans)
}

func eval(op1, op2, operator string, gates *map[string]*Gate) int {
	if (*gates)[op1].Value == -1 {
		op1Gate := (*gates)[op1]
		(*gates)[op1].Value = eval(op1Gate.O1, op1Gate.O2, op1Gate.Operator, gates)
	}
	if (*gates)[op2].Value == -1 {
		op2Gate := (*gates)[op2]
		(*gates)[op2].Value = eval(op2Gate.O1, op2Gate.O2, op2Gate.Operator, gates)
	}
	switch operator {
	case "XOR":
		return (*gates)[op1].Value ^ (*gates)[op2].Value
	case "AND":
		return (*gates)[op1].Value & (*gates)[op2].Value
	case "OR":
		return (*gates)[op1].Value | (*gates)[op2].Value
	}
	return 0
}

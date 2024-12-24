package twentyfour

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/sets/hashset"
)

type Gate struct {
	Value    int
	Operator string
	O1       string
	O2       string
}

func SolveDay24() {
	fmt.Println("Day 24 Part 1: ", solvePartOne())
	fmt.Println("Day 24 Part 2: ", solvePartTwo())
}

func setup() map[string]*Gate {
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

	return gates
}

func solvePartOne() int {
	gates := setup()
	keys := []string{}
	for k, v := range gates {
		if k[0] == 'z' {
			keys = append(keys, k)
		}
		if v.Value != -1 {
			continue
		}
		// Eval Exp
		gates[k].Value = eval(k, gates)
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
	return int(ans)
}

func solvePartTwo() string {
	gates := setup()

	highestZ := "z00"
	for k, v := range gates {
		if k[0] == 'z' && util.ToInt(strings.TrimLeft(k, "z")) > util.ToInt(strings.TrimLeft(highestZ, "z")) {
			highestZ = k
		}
		if v.Value != -1 {
			continue
		}
		// Eval Exp
		gates[k].Value = eval(k, gates)
	}

	wrong := hashset.New()
	mainSet := hashset.New("x", "y", "z")
	for k, v := range gates {
		if k[0] == 'z' && v.Operator != "XOR" && k != highestZ {
			wrong.Add(k)
		}
		if v.Operator == "XOR" &&
			!mainSet.Contains(string(k[0])) &&
			!mainSet.Contains(string(v.O1[0])) &&
			!mainSet.Contains(string(v.O2[0])) {
			wrong.Add(k)
		}
		if v.Operator == "AND" &&
			v.O1 != "x00" && v.O2 != "x00" {
			for _, subV := range gates {
				if (k == subV.O1 || k == subV.O2) && subV.Operator != "OR" {
					wrong.Add(k)
				}
			}
		}
		if v.Operator == "XOR" {
			for _, subV := range gates {
				if (k == subV.O1 || k == subV.O2) && subV.Operator == "OR" {
					wrong.Add(k)
				}
			}
		}
	}

	wv := []string{}
	for _, n := range wrong.Values() {
		wv = append(wv, n.(string))
	}
	slices.SortFunc(wv, func(a, b string) int {
		return cmp.Compare(a, b)
	})

	return strings.Join(wv, ",")
}

func eval(k string, gates map[string]*Gate) int {
	op1 := gates[k].O1
	op2 := gates[k].O2
	operator := gates[k].Operator
	op1Gate := gates[op1]
	op2Gate := gates[op2]
	if op1Gate.Value == -1 {
		op1Gate.Value = eval(op1, gates)
	}
	if gates[op2].Value == -1 {
		op2Gate.Value = eval(op2, gates)
	}
	switch operator {
	case "XOR":
		return op1Gate.Value ^ op2Gate.Value
	case "AND":
		return op1Gate.Value & op2Gate.Value
	case "OR":
		return op1Gate.Value | op2Gate.Value
	}
	return 0
}

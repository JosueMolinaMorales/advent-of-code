package seventeen

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

var INSTRUCTIONS = map[int]InstructionFunc{
	0: adv,
	1: bxl,
	2: bst,
	4: bxc,
	6: bdv,
	7: cdv,
}

func SolveDay17() {
	fmt.Println("Day 17 Part 1: ", solvePartOne())
	fmt.Println("Day 17 Part 2: ", solvePartTwo())
}

func setup() (map[string]int, []int) {
	input, err := util.LoadFileAsString("./inputs/day_17.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	re, err := regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}

	regs := re.FindAllString(parts[0], -1)
	registers := map[string]int{
		"A": util.ToInt(regs[0]),
		"B": util.ToInt(regs[1]),
		"C": util.ToInt(regs[2]),
	}

	program := []int{}
	for _, s := range re.FindAllString(parts[1], -1) {
		program = append(program, util.ToInt(s))
	}

	return registers, program
}

func runProgram(registers map[string]int, program []int) string {
	ptr := 0
	outList := []string{}
	for ptr < len(program) {
		ins := program[ptr]
		operand := program[ptr+1]
		f := INSTRUCTIONS[ins]
		if f == nil {
			if ins == 3 {
				// Jmp
				if registers["A"] != 0 {
					ptr = operand
					continue
				}
			}
			if ins == 5 {
				// out
				outList = append(outList, strconv.Itoa(out(&registers, operand)))
			}
		} else {
			f(&registers, operand)
		}
		ptr += 2
	}

	return strings.Join(outList, ",")
}

func solvePartOne() string {
	registers, program := setup()
	return runProgram(registers, program)
}

func intArrayToString(arr []int) []string {
	s := []string{}
	for _, n := range arr {
		s = append(s, strconv.Itoa(n))
	}
	return s
}

func solvePartTwo() int {
	registers, program := setup()

	candidates := priorityqueue.NewWith(func(a, b interface{}) int {
		return utils.IntComparator(a, b)
	})

	for i := 0; i < 8; i++ {
		candidates.Enqueue(i)
	}

	for !candidates.Empty() {
		c, _ := candidates.Dequeue()
		candidate := c.(int)
		registers["A"] = candidate
		output := runProgram(registers, program)

		if output == strings.Join(intArrayToString(program), ",") {
			return candidate
		}

		length := len(strings.Split(output, ","))
		t := program[len(program)-length:]
		s := intArrayToString(t)
		if strings.Join(strings.Split(output, ","), "") == strings.Join(s, "") {
			for i := 0; i < 8; i++ {
				candidates.Enqueue((candidate << 3) + i)
			}
		}
	}

	return -1
}

type InstructionFunc = func(registers *map[string]int, operand int)

// adv performs division -> A/B where A is the value of the A register
// B is the operand^2. Result is truncated & stored in the A register
func adv(registers *map[string]int, operand int) {
	(*registers)["A"] = (*registers)["A"] / int(math.Pow(2, float64(combo(registers, operand))))
}

// bxl calculates the bitwise XOR of register B and the LITERAL operand
// stores the result in register B
func bxl(registers *map[string]int, operand int) {
	(*registers)["B"] = (*registers)["B"] ^ operand
}

// bst calculates the value of its combo operand MOD 8
// stores results to the B Register
func bst(registers *map[string]int, operand int) {
	(*registers)["B"] = combo(registers, operand) % 8
}

// bxc calculate XOR of register B and register C
// stores the result in register B, ignores the operand
func bxc(registers *map[string]int, operand int) {
	(*registers)["B"] = (*registers)["B"] ^ (*registers)["C"]
}

// out calculates the value of its combo operand mod 8 then outputs the value
func out(registers *map[string]int, operand int) int {
	return combo(registers, operand) % 8
}

// bdv works exactly like the adv expects stores it in the B register
func bdv(registers *map[string]int, operand int) {
	(*registers)["B"] = (*registers)["A"] / int(math.Pow(2, float64(combo(registers, operand))))
}

// cdv works exactly like the adv excepts stores it in the C register
func cdv(registers *map[string]int, operand int) {
	(*registers)["C"] = (*registers)["A"] / int(math.Pow(2, float64(combo(registers, operand))))
}

func combo(registers *map[string]int, operand int) int {
	if operand <= 3 {
		return operand
	}

	if operand == 4 {
		return (*registers)["A"]
	}
	if operand == 5 {
		return (*registers)["B"]
	}
	if operand == 6 {
		return (*registers)["C"]
	}

	panic("Not valid")
}

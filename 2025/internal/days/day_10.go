package days

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/dsa"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

type Machine struct {
	Target      int
	NumLights   int
	ButtonMasks []int
	Buttons     [][]int
	Joltage     []int
}

func Day10() {
	fmt.Println("2025 Day 10 Part 1:", day10Part1("inputs/day_10/input.txt"))
	fmt.Println("2025 Day 10 Part 2:", day10Part2("inputs/day_10/input.txt"))
}

func day10Part1(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("Failed to read input file for day 10")
	}

	machines := parseMachines(input, true)

	res := 0
	for _, machine := range machines {
		min := findMinButtonPressesBFS(machine)
		res += min
	}
	return res
}

func findMinButtonPressesBFS(machine Machine) int {
	queue := dsa.NewQueue[MachineState]()
	visited := make(map[int]bool)

	initialState := MachineState{
		LightState: 0,
		NumPresses: 0,
	}
	queue.Enqueue(initialState)

	for !queue.IsEmpty() {
		current := queue.Dequeue()

		if current.LightState == machine.Target {
			return current.NumPresses
		}

		if visited[current.LightState] {
			continue
		}
		visited[current.LightState] = true

		for _, button := range machine.ButtonMasks {
			newLightState := current.LightState ^ button

			if !visited[newLightState] {
				queue.Enqueue(MachineState{
					LightState: newLightState,
					NumPresses: current.NumPresses + 1,
				})
			}
		}
	}

	return -1
}

type MachineState struct {
	LightState int
	NumPresses int
}

func day10Part2(path string) int {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("Failed to read input file for day 10")
	}

	machines := parseMachines(input, false)

	res := 0
	for _, machine := range machines {
		min := solveLinearSystem(machine)
		res += min
	}
	return res
}

func parseMachines(input []string, includeBitmasks bool) []Machine {
	machines := []Machine{}
	for _, line := range input {
		machine := Machine{}
		for _, part := range strings.Split(line, " ") {
			if part[0] == '[' {
				lightsStr := part[1 : len(part)-1]
				machine.NumLights = len(lightsStr)
				if includeBitmasks {
					machine.Target = 0
					for i, light := range lightsStr {
						if light == '#' {
							machine.Target |= (1 << i)
						}
					}
				}
			}
			if part[0] == '(' {
				buttons := []int{}
				for _, bs := range strings.Split(part[1:len(part)-1], ",") {
					b, err := strconv.Atoi(string(bs))
					if err != nil {
						log.Fatalf("Failed to parse string: %s", err)
					}
					buttons = append(buttons, b)
				}
				machine.Buttons = append(machine.Buttons, buttons)

				if includeBitmasks {
					buttonMask := 0
					for _, b := range buttons {
						buttonMask |= (1 << b)
					}
					machine.ButtonMasks = append(machine.ButtonMasks, buttonMask)
				}
			}
			if part[0] == '{' {
				joltage := []int{}
				for _, vs := range strings.Split(part[1:len(part)-1], ",") {
					v, err := strconv.Atoi(string(vs))
					if err != nil {
						log.Fatalf("Failed to parse string: %s", err)
					}
					joltage = append(joltage, v)
				}
				machine.Joltage = joltage
			}
		}
		machines = append(machines, machine)
	}
	return machines
}

func solveLinearSystem(machine Machine) int {
	numCounters := len(machine.Joltage)
	numButtons := len(machine.Buttons)

	A := make([][]int, numCounters)
	for i := range A {
		A[i] = make([]int, numButtons)
	}

	for j, button := range machine.Buttons {
		for _, counter := range button {
			if counter < numCounters {
				A[counter][j] = 1
			}
		}
	}

	augmented := make([][]int, numCounters)
	for i := range augmented {
		augmented[i] = make([]int, numButtons+1)
		copy(augmented[i], A[i])
		augmented[i][numButtons] = machine.Joltage[i]
	}

	pivotCols := []int{}
	currentRow := 0

	for col := 0; col < numButtons && currentRow < numCounters; col++ {
		pivotRow := -1
		for row := currentRow; row < numCounters; row++ {
			if augmented[row][col] != 0 {
				pivotRow = row
				break
			}
		}

		if pivotRow == -1 {
			continue
		}

		augmented[currentRow], augmented[pivotRow] = augmented[pivotRow], augmented[currentRow]
		pivotCols = append(pivotCols, col)

		for row := 0; row < numCounters; row++ {
			if row == currentRow || augmented[row][col] == 0 {
				continue
			}

			lcm := lcm(augmented[currentRow][col], augmented[row][col])
			mult1 := lcm / augmented[currentRow][col]
			mult2 := lcm / augmented[row][col]

			for c := 0; c <= numButtons; c++ {
				augmented[row][c] = augmented[row][c]*mult2 - augmented[currentRow][c]*mult1
			}
		}

		currentRow++
	}

	isPivot := make(map[int]bool)
	for _, col := range pivotCols {
		isPivot[col] = true
	}

	freeVars := []int{}
	for col := 0; col < numButtons; col++ {
		if !isPivot[col] {
			freeVars = append(freeVars, col)
		}
	}

	if len(freeVars) == 0 {
		solution := make([]int, numButtons)
		for i := len(pivotCols) - 1; i >= 0; i-- {
			col := pivotCols[i]
			val := augmented[i][numButtons]

			for j := col + 1; j < numButtons; j++ {
				val -= augmented[i][j] * solution[j]
			}

			if augmented[i][col] == 0 || val%augmented[i][col] != 0 {
				return math.MaxInt32
			}
			solution[col] = val / augmented[i][col]

			if solution[col] < 0 {
				return math.MaxInt32
			}
		}

		total := 0
		for _, v := range solution {
			total += v
		}
		return total
	}

	minSum := math.MaxInt32

	maxTarget := 0
	for _, val := range machine.Joltage {
		if val > maxTarget {
			maxTarget = val
		}
	}

	upperBound := maxTarget * 5
	if upperBound < 500 {
		upperBound = 500
	}

	var tryValues func(int, []int)
	tryValues = func(idx int, values []int) {
		if idx == len(freeVars) {
			solution := make([]int, numButtons)

			for i, freeVar := range freeVars {
				solution[freeVar] = values[i]
			}

			valid := true
			for i := len(pivotCols) - 1; i >= 0; i-- {
				col := pivotCols[i]
				val := augmented[i][numButtons]

				for j := col + 1; j < numButtons; j++ {
					val -= augmented[i][j] * solution[j]
				}

				if augmented[i][col] == 0 || val%augmented[i][col] != 0 {
					valid = false
					break
				}
				solution[col] = val / augmented[i][col]

				if solution[col] < 0 {
					valid = false
					break
				}
			}

			if valid {
				sum := 0
				for _, v := range solution {
					sum += v
				}
				if sum < minSum {
					minSum = sum
				}
			}
			return
		}

		currentSum := 0
		for i := 0; i < idx; i++ {
			currentSum += values[i]
		}
		if currentSum >= minSum {
			return
		}

		for v := 0; v <= upperBound; v++ {
			tryValues(idx+1, append(values, v))
		}
	}

	tryValues(0, []int{})

	if minSum == math.MaxInt32 {
		log.Printf("Warning: No valid solution found for machine")
		return 0
	}
	return minSum
}

func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a*b) / gcd(a, b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

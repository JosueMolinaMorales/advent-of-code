package eight

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInputA = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

const testInputB = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

type Node struct {
	Value string
	Left  string
	Right string
}

func RunDayEight() {
	input, err := os.ReadFile("./input/day8.txt")
	if err != nil {
		panic("Failed to read day 8 file")
	}
	fmt.Println("Part 1:", partOne(string(input)))
	fmt.Println("Part 2:", partTwo(string(input)))
}

func parseInput(input string) (map[string]Node, string) {
	inputSplit := strings.Split(input, "\n\n")
	path := inputSplit[0]
	routes := inputSplit[1]
	pattern := `\b(\w{3}) = \((\w{3}), (\w{3})\)`

	nodes := make(map[string]Node)
	for _, line := range strings.Split(routes, "\n") {
		// Compile the regular expression
		re := regexp.MustCompile(pattern)

		// Find the matches in the input string
		matches := re.FindStringSubmatch(line)

		node := matches[1]
		left := matches[2]
		right := matches[3]

		nodes[node] = Node{
			Value: node,
			Left:  left,
			Right: right,
		}
	}

	return nodes, path
}

func partOne(input string) int {
	nodes, path := parseInput(input)
	currNode := "AAA"
	i := 0
	count := 0
	for currNode != "ZZZ" {
		// Get the next step
		rol := string(path[i])
		i += 1
		if i >= len(path) {
			i = 0
		}

		node := nodes[currNode]
		switch rol {
		case "R":
			currNode = node.Right
		case "L":
			currNode = node.Left
		}

		count += 1

	}

	return count
}

func partTwo(input string) int {
	nodes, path := parseInput(input)

	// Find all starting nodes (i.e. nodes that end with A)
	startingNodes := make([]string, 0)
	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			startingNodes = append(startingNodes, k)
		}
	}

	// Brute force wont work, need to find a cycle for each
	// Keep track of all the cycles for all the starting nodes
	cycles := make(map[string][]string)
	// Add starting to cycles
	for _, n := range startingNodes {
		cycles[n] = make([]string, 0)
	}
	// Create a slices that stores whether a starting node has a cycle yet
	cyclesFound := make([]bool, len(startingNodes))

	i := 0
	currentNodes := make([]string, 0)
	currentNodes = append(currentNodes, startingNodes...)
	for !iterators.Every(cyclesFound, func(f bool) bool { return f }) {
		// Get the next step
		rol := string(path[i])
		i += 1
		if i >= len(path) {
			i = 0
		}

		for j, n := range currentNodes {
			// Skip if cycle already found
			if cyclesFound[j] {
				continue
			}
			startingNode := startingNodes[j]
			node := nodes[n]
			switch rol {
			case "R":
				cycles[startingNode] = append(cycles[startingNode], node.Right)
				currentNodes[j] = node.Right
			case "L":
				cycles[startingNode] = append(cycles[startingNode], node.Left)
				currentNodes[j] = node.Left
			}

			// If nodes are the same, cycle was found
			// Or if direct path was found to Z
			if startingNode == currentNodes[j] || strings.HasSuffix(currentNodes[j], "Z") {
				cyclesFound[j] = true
			}
		}
	}

	// Find the common multiple of cycle lengths
	lengths := make([]*big.Int, 0)
	for _, v := range cycles {
		lengths = append(lengths, big.NewInt(int64(len(v))))
	}
	cm := multipleLCM(lengths)

	return int(cm.Int64())
}

// Calculate the greatest common divisor (GCD) using Euclid's algorithm
func gcd(a, b *big.Int) *big.Int {
	gcd := new(big.Int)
	return gcd.GCD(nil, nil, a, b)
}

// Calculate the least common multiple (LCM) of two numbers
func lcm(a, b *big.Int) *big.Int {
	// LCM(a, b) = |a * b| / GCD(a, b)
	absA := new(big.Int).Abs(a)
	absB := new(big.Int).Abs(b)
	gcdAB := gcd(absA, absB)

	// LCM = |a * b| / GCD(a, b)
	lcm := new(big.Int).Div(new(big.Int).Mul(absA, absB), gcdAB)
	return lcm
}

// Calculate the least common multiple (LCM) of multiple numbers
func multipleLCM(numbers []*big.Int) *big.Int {
	// Initialize the LCM with 1
	result := big.NewInt(1)

	// Iterate through each number and update the LCM
	for _, num := range numbers {
		result = lcm(result, num)
	}

	return result
}

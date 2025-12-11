package days

import (
	"fmt"
	"log"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/dsa"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

type PathState struct {
	CurrentNode string
	HasSeenDAC  bool
	HasSeenFFT  bool
}

func Day11() {
	fmt.Println("2025 Day 11 Part 1:", day11Part1("inputs/day_11/input.txt"))
	fmt.Println("2025 Day 11 Part 2:", day11Part2("inputs/day_11/input.txt"))
}

func parseGraph(path string) map[string][]string {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("Failed to read input for 2025 day 11: %s", err)
	}

	graph := map[string][]string{}
	for _, line := range input {
		parts := strings.Split(line, ": ")
		graph[parts[0]] = append(graph[parts[0]], strings.Split(parts[1], " ")...)
	}
	return graph
}

func day11Part1(path string) int {
	graph := parseGraph(path)
	return findAllPathsToOut("you", "out", graph, dsa.NewSet[string]())
}

func day11Part2(path string) int {
	graph := parseGraph(path)
	return findAllPathsToOutMustContain(PathState{
		CurrentNode: "svr",
		HasSeenDAC:  false,
		HasSeenFFT:  false,
	}, "out", graph, map[PathState]int{})
}

func findAllPathsToOut(currentNode string, endNode string, graph map[string][]string, visited dsa.Set[string]) int {
	if currentNode == endNode {
		return 1
	}

	visited.Add(currentNode)

	paths := 0
	for _, v := range graph[currentNode] {
		if !visited.Contains(v) {
			paths += findAllPathsToOut(v, endNode, graph, visited)
		}
	}

	visited.Remove(currentNode)
	return paths
}

func findAllPathsToOutMustContain(state PathState, endNode string, graph map[string][]string, memo map[PathState]int) int {
	if _, ok := memo[state]; ok {
		return memo[state]
	}
	if state.CurrentNode == endNode {
		if state.HasSeenDAC && state.HasSeenFFT {
			return 1
		}
		return 0
	}

	paths := 0
	for _, v := range graph[state.CurrentNode] {
		newState := PathState{
			CurrentNode: v,
			HasSeenDAC:  state.HasSeenDAC || v == "dac",
			HasSeenFFT:  state.HasSeenFFT || v == "fft",
		}

		paths += findAllPathsToOutMustContain(newState, endNode, graph, memo)
	}

	memo[state] = paths

	return paths
}

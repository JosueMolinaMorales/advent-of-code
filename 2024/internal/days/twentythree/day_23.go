package twentythree

import (
	"fmt"
	"slices"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay23() {
	fmt.Println("Day 23 Part 1: ", solvePartOne())
	fmt.Println("Day 23 Part 2: ", solvePartTwo())
}

func setup() map[string]*hashset.Set {
	input, err := util.LoadFileAsString("./inputs/day_23.txt")
	if err != nil {
		panic(err)
	}
	netMap := strings.Split(input, "\n")

	// K: PC V: Connected PCs
	connections := map[string]*hashset.Set{}
	for _, conn := range netMap {
		pcs := strings.Split(conn, "-")

		if connections[pcs[0]] == nil {
			connections[pcs[0]] = hashset.New()
		}
		connections[pcs[0]].Add(pcs[1])
		if connections[pcs[1]] == nil {
			connections[pcs[1]] = hashset.New()
		}
		connections[pcs[1]].Add(pcs[0])
	}

	return connections
}

func solvePartOne() int {
	connections := setup()
	paths := hashset.New()
	// For every pc, see if it 3 other pcs connections
	for k := range connections {
		dfs(k, k, 0, hashset.New(), []string{}, paths, connections)
	}
	count := 0
	for _, p := range paths.Values() {
		path := p.(string)
		contains := false
		for _, c := range strings.Split(path, "-") {
			if c[0] == 't' {
				contains = true
				break
			}
		}
		if contains {
			count++
		}
	}
	return count
}

func solvePartTwo() string {
	graph := setup()

	R := hashset.New() // Current clique
	P := hashset.New() // Potential vertices
	X := hashset.New() // Excluded vertices

	// Add all vertices to P
	for vertex := range graph {
		P.Add(vertex)
	}

	// Store all maximal cliques
	cliques := [][]string{}

	// Run the Bron-Kerbosch algorithm
	BronKerbosch(graph, R, P, X, &cliques)

	// Find Max Clique
	maxClique := []string{}
	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	slices.Sort(maxClique)
	return strings.Join(maxClique, ",")
}

// BronKerbosch finds all maximal cliques in the graph.
func BronKerbosch(graph map[string]*hashset.Set, R, P, X *hashset.Set, cliques *[][]string) {
	if P.Size() == 0 && X.Size() == 0 {
		// R is a maximal clique
		clique := []string{}
		for _, v := range R.Values() {
			clique = append(clique, v.(string))
		}
		*cliques = append(*cliques, clique)
		return
	}

	// Pick a pivot
	var pivot string
	for _, v := range P.Values() {
		pivot = v.(string)
		break
	}

	neighbors := graph[pivot]
	for _, v := range P.Difference(neighbors).Values() {
		vertex := v.(string)
		neighborsOfVertex := graph[vertex]

		BronKerbosch(
			graph,
			R.Union(hashset.New(vertex)),
			P.Intersection(neighborsOfVertex),
			X.Intersection(neighborsOfVertex),
			cliques,
		)

		P.Remove(vertex)
		X.Add(vertex)
	}
}

func dfs(curr string, target string, depth int, visited *hashset.Set, path []string, paths *hashset.Set, graph map[string]*hashset.Set) {
	if curr == target && depth == 3 {
		slices.Sort(path)
		paths.Add(strings.Join(path, "-"))
		return
	}
	if depth == 3 {
		return
	}
	if visited.Contains(curr) {
		return
	}
	for _, n := range graph[curr].Values() {
		visited.Add(curr)
		dfs(n.(string), target, depth+1, visited, append(path, n.(string)), paths, graph)
		visited.Remove(curr)
	}
}

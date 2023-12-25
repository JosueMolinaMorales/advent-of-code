package twentyfive

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const testInput = `jqt: rhn xhk nvd
rsh: frs pzl lsr
xhk: hfx
cmg: qnr nvd lhk bvb
rhn: xhk bvb hfx
bvb: xhk hfx
pzl: lsr hfx nvd
qnr: nvd
ntq: jqt hfx bvb xhk
nvd: lhk
lsr: lhk
rzs: qnr cmg lsr rsh
frs: qnr lhk lsr`

func RunDayTwentyFive() {
	input, err := os.ReadFile("./input/day25.txt")
	if err != nil {
		panic("Failed to read day 25 input file")
	}
	fmt.Println("Day 25 Part 1:", partOne(string(input)))
}

type (
	Vertex string
	Edge   = [2]string // [from, to]
	Graph  struct {
		Vertices map[Vertex]struct{}
		Edges    []Edge
	}
)

func (g Graph) Copy() Graph {
	vertices := make(map[Vertex]struct{})
	edges := make([]Edge, 0)
	for v := range g.Vertices {
		vertices[v] = struct{}{}
	}
	edges = append(edges, g.Edges...)
	return Graph{
		Vertices: vertices,
		Edges:    edges,
	}
}

func partOne(input string) int {
	vertices := make(map[Vertex]struct{})
	edges := make([]Edge, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ": ")
		connections := strings.Split(parts[1], " ")

		vertices[Vertex(parts[0])] = struct{}{}
		for _, c := range connections {
			edges = append(edges, Edge{parts[0], c})
		}

		for _, c := range connections {
			vertices[Vertex(c)] = struct{}{}
			edges = append(edges, Edge{c, parts[0]})
		}
	}

	graph := Graph{
		Vertices: vertices,
		Edges:    edges,
	}

	return minCut(graph)
}

// Use Karger's algorithm to find the minimum cut
// https://en.wikipedia.org/wiki/Karger%27s_algorithm
func minCut(graph Graph) int {
	i := 0
	for {
		i += 1
		if i > 1_000_000 {
			panic("Too many iterations")
		}
		// Copy the graph
		graphCopy := graph.Copy()
		for len(graphCopy.Vertices) > 2 {
			// Pick a random edge
			idx := rand.Intn(len(graphCopy.Edges))
			u := graphCopy.Edges[idx][0]
			v := graphCopy.Edges[idx][1]
			// Merge the vertices
			merge(&graphCopy, u, v)
		}

		// If the graph has 3 edges, we found the 3 cuts we need
		if len(graphCopy.Edges)/2 == 3 {
			graph = graphCopy
			break
		}

	}
	// Count The number of unique vertices
	sum := 1
	for v := range graph.Vertices {
		vertices := make(map[Vertex]struct{})
		parts := strings.Split(string(v), "-")
		for _, p := range parts {
			vertices[Vertex(p)] = struct{}{}
		}
		sum *= len(vertices)
	}
	return sum
}

func merge(graph *Graph, u, v string) {
	// Create a new node
	newNode := fmt.Sprintf("%s-%s", u, v)

	// Add the new node
	graph.Vertices[Vertex(newNode)] = struct{}{}

	// Create a new slice for the updated edges
	updatedEdges := make([]Edge, 0, len(graph.Edges))

	// Add the edges from u and v to the new node
	for _, edge := range graph.Edges {
		// Outgoing edges
		if edge[0] == u || edge[0] == v {
			edge[0] = newNode
		}
		// Incoming edges
		if edge[1] == u || edge[1] == v {
			edge[1] = newNode
		}

		// Add the updated edge to the new slice
		updatedEdges = append(updatedEdges, edge)
	}

	// Update the graph with the new edges
	graph.Edges = updatedEdges

	// Remove the vertices u and v
	delete(graph.Vertices, Vertex(u))
	delete(graph.Vertices, Vertex(v))

	// Remove self-loops
	removeSelfLoops(graph)
}

func removeSelfLoops(graph *Graph) {
	// Create a new slice for the edges without self-loops
	edgesWithoutSelfLoops := make([]Edge, 0, len(graph.Edges))

	for _, edge := range graph.Edges {
		if edge[0] != edge[1] {
			// Add the edge to the new slice
			edgesWithoutSelfLoops = append(edgesWithoutSelfLoops, edge)
		}
	}

	// Update the graph with the edges without self-loops
	graph.Edges = edgesWithoutSelfLoops
}

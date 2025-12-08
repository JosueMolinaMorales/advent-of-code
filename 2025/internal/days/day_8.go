package days

import (
	"cmp"
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/dsa"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
	"github.com/JosueMolinaMorales/aoc/2025/internal/util/types"
)

func Day8() {
	fmt.Println("2025 Day 8 Part 1:", day8Part1("inputs/day_8/input.txt"))
	fmt.Println("2025 Day 8 Part 2:", day8Part2("inputs/day_8/input.txt"))
}

func day8Part1(path string) int {
	boxes := parseBoxes(path)

	// Initialize Union-Find: each box is its own parent (representative)
	parent := make(map[types.Point3D]types.Point3D)
	for _, box := range boxes {
		parent[box] = box
	}

	// Pre-compute all edges (pairs of boxes) with their distances
	minHeap := buildDistanceHeap(boxes)

	// Connect the 1000 closest pairs using Kruskal's algorithm
	for range 1000 {
		boxDist := minHeap.PopItem()

		rootU := findRoot(boxDist.PointA, parent)
		rootV := findRoot(boxDist.PointB, parent)

		// If they're in different circuits, union them
		if rootU != rootV {
			parent[rootV] = rootU
		}
	}

	// Count the size of each circuit
	circuitSizes := make(map[types.Point3D]int)
	for _, box := range boxes {
		root := findRoot(box, parent)
		circuitSizes[root]++
	}

	// Get all circuit sizes and sort them in descending order
	sizes := make([]int, 0, len(circuitSizes))
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}
	slices.SortFunc(sizes, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	// Multiply the three largest circuit sizes
	return sizes[0] * sizes[1] * sizes[2]
}

func day8Part2(path string) int {
	boxes := parseBoxes(path)

	// Initialize Union-Find
	parent := make(map[types.Point3D]types.Point3D)
	for _, box := range boxes {
		parent[box] = box
	}

	// Pre-compute all edges with distances
	minHeap := buildDistanceHeap(boxes)

	var u, v types.Point3D
	// Keep connecting until there's only one circuit
	for minHeap.Len() > 0 {
		boxDist := minHeap.PopItem()

		rootU := findRoot(boxDist.PointA, parent)
		rootV := findRoot(boxDist.PointB, parent)

		if rootU != rootV {
			parent[rootV] = rootU
			u = boxDist.PointA
			v = boxDist.PointB

			// Count remaining circuits
			circuits := make(map[types.Point3D]bool)
			for _, box := range boxes {
				root := findRoot(box, parent)
				circuits[root] = true
			}

			// Stop when there's only one circuit
			if len(circuits) == 1 {
				break
			}
		}
	}

	return u.X * v.X
}

// parseBoxes reads the input file and parses junction box coordinates
func parseBoxes(path string) []types.Point3D {
	input, err := io.ReadFileAsLines(path)
	if err != nil {
		log.Fatalf("ERROR reading input: %s", err)
	}

	boxes := make([]types.Point3D, 0, len(input))
	for _, line := range input {
		parts := strings.Split(line, ",")
		coords := make([]int, 0, 3)
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatalf("Failed to parse coordinate: %s", err)
			}
			coords = append(coords, n)
		}
		boxes = append(boxes, types.New3DPoint(coords[0], coords[1], coords[2]))
	}
	return boxes
}

// buildDistanceHeap creates a min-heap of all box pairs sorted by distance
func buildDistanceHeap(boxes []types.Point3D) *dsa.Heap[BoxDistance] {
	minHeap := dsa.NewMinHeap(func(a, b BoxDistance) bool {
		return a.Distance < b.Distance
	})

	for i := range boxes {
		for j := i + 1; j < len(boxes); j++ {
			minHeap.PushItem(BoxDistance{
				PointA:   boxes[i],
				PointB:   boxes[j],
				Distance: calc3DEuclideanDistance(boxes[i], boxes[j]),
			})
		}
	}

	return minHeap
}

// findRoot finds the root/representative of a box's circuit with path compression
func findRoot(box types.Point3D, parent map[types.Point3D]types.Point3D) types.Point3D {
	if parent[box] != box {
		parent[box] = findRoot(parent[box], parent) // Path compression
	}
	return parent[box]
}

// BoxDistance represents a pair of junction boxes and their distance
type BoxDistance struct {
	PointA   types.Point3D
	PointB   types.Point3D
	Distance int
}

// calc3DEuclideanDistance calculates the Euclidean distance between two 3D points
func calc3DEuclideanDistance(p1, p2 types.Point3D) int {
	dx := float64(p1.X - p2.X)
	dy := float64(p1.Y - p2.Y)
	dz := float64(p1.Z - p2.Z)
	return int(math.Sqrt(dx*dx + dy*dy + dz*dz))
}

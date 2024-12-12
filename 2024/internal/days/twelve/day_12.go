package twelve

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/arrayqueue"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay12() {
	fmt.Println("Day 12 Part 1: ", solvePartOne())
	fmt.Println("Day 12 Part 2: ", solvePartTwo())
}

func setup() ([][]string, []hashset.Set) {
	rawMap, err := util.LoadFileAsString("./inputs/day_11.txt")
	if err != nil {
		panic(err)
	}

	garden := [][]string{}
	for _, line := range strings.Split(rawMap, "\n") {
		garden = append(garden, strings.Split(line, ""))
	}

	// Find all regions uses BFS
	regions := []hashset.Set{}

	for i, row := range garden {
		for j := range row {
			//. Check if the pos is already in a region
			vec := *types.NewVector(i, j)
			contains := false
			for _, r := range regions {
				if r.Contains(vec) {
					contains = true
				}
			}
			if contains {
				continue
			}

			r := bfs(vec, garden)
			regions = append(regions, r)
		}
	}

	return garden, regions
}

func solvePartOne() int {
	garden, regions := setup()
	// For every region, calculate each plots perimeter
	totalPrice := 0
	for _, region := range regions {
		p := 0
		for _, plot := range region.Values() {
			p += calcPerimeter(plot.(types.Vector), garden)
		}
		totalPrice += p * region.Size()
	}

	return totalPrice
}

func solvePartTwo() int {
	_, regions := setup()
	p2 := 0
	for _, reg := range regions {
		corners := countCorners(reg)
		p2 += reg.Size() * corners
	}
	return p2
}

func countCorners(region hashset.Set) int {
	left := hashset.New()
	right := hashset.New()
	up := hashset.New()
	down := hashset.New()

	// Assign directions
	for _, p := range region.Values() {
		point := p.(types.Vector)

		if !region.Contains(*types.NewVector(point.X-1, point.Y)) {
			up.Add(point)
		}
		if !region.Contains(*types.NewVector(point.X+1, point.Y)) {
			down.Add(point)
		}
		if !region.Contains(*types.NewVector(point.X, point.Y+1)) {
			right.Add(point)
		}
		if !region.Contains(*types.NewVector(point.X, point.Y-1)) {
			left.Add(point)
		}
	}

	corners := 0

	// Check corners for `up`
	for _, p := range up.Values() {
		point := p.(types.Vector)
		if left.Contains(point) {
			corners++
		}
		if right.Contains(point) {
			corners++
		}
		if right.Contains(*types.NewVector(point.X-1, point.Y-1)) && !left.Contains(point) {
			corners++
		}
		if left.Contains(*types.NewVector(point.X-1, point.Y+1)) && !right.Contains(point) {
			corners++
		}
	}

	// Check corners for `down`
	for _, p := range down.Values() {
		point := p.(types.Vector)
		if left.Contains(point) {
			corners++
		}
		if right.Contains(point) {
			corners++
		}
		if right.Contains(*types.NewVector(point.X+1, point.Y-1)) && !left.Contains(point) {
			corners++
		}
		if left.Contains(*types.NewVector(point.X+1, point.Y+1)) && !right.Contains(point) {
			corners++
		}
	}

	return corners
}

func calcPerimeter(plot types.Vector, garden [][]string) int {
	directions := [][]int{
		{0, 1},  // Right
		{0, -1}, // Left
		{-1, 0}, // Down
		{1, 0},  // Up
	}

	count := 0
	// Count it if, its not the same plot
	for _, dir := range directions {
		dx, dy := plot.X+dir[0], plot.Y+dir[1]
		if dx < 0 || dx >= len(garden) || dy < 0 || dy >= len(garden[0]) {
			count += 1
			continue
		}
		if garden[plot.X][plot.Y] != garden[dx][dy] {
			count++
		}
	}

	return count
}

func bfs(start types.Vector, garden [][]string) hashset.Set {
	queue := arrayqueue.New()
	visited := hashset.New()

	visited.Add(start)

	queue.Enqueue(start)

	region := hashset.New()
	for !queue.Empty() {
		v, _ := queue.Dequeue()
		vec := v.(types.Vector)
		region.Add(vec)
		for _, adj := range adjacent(vec, garden) {
			if !visited.Contains(adj) {
				visited.Add(adj)
				queue.Enqueue(adj)
			}
		}
	}

	return *region
}

func adjacent(curr_pos types.Vector, garden [][]string) []types.Vector {
	directions := [][]int{
		{0, 1},  // Right
		{0, -1}, // Left
		{-1, 0}, // Down
		{1, 0},  // Up
	}

	adjacents := []types.Vector{}
	for _, dir := range directions {
		dx, dy := curr_pos.X+dir[0], curr_pos.Y+dir[1]
		if dx < 0 || dx >= len(garden) || dy < 0 || dy >= len(garden[0]) {
			continue
		}

		if garden[curr_pos.X][curr_pos.Y] == garden[dx][dy] {
			adjacents = append(adjacents, *types.NewVector(dx, dy))
		}
	}
	return adjacents
}

package seventeen

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const testInput = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

const test = `111
551
551
551
551
551
551
551`

func RunDaySeventeen() {
	// input, err := os.ReadFile("./input/day17.txt")
	// if err != nil {
	// 	panic("Failed to read day 17 input file")
	// }
	// fmt.Println("Part 1", partOne(string(input)))
	// fmt.Println("Part 2", partTwo(string(input)))
	fmt.Println("Part 1", partOne(testInput))
}

func parseInput(input string) [][]int {
	m := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			row = append(row, num)
		}
		m = append(m, row)
	}
	return m
}

func partTwo(input string) int {
	m := parseInput(input)
	minHeatLoss := getMinHeatLoss(m, 3, 10)
	return minHeatLoss
}

func partOne(input string) int {
	m := parseInput(input)
	minHeatLoss := getMinHeatLoss(m, 0, 3)
	return minHeatLoss
}

var (
	RIGHT, LEFT, UP, DOWN = Point{1, 0}, Point{-1, 0}, Point{0, -1}, Point{0, 1}
	TURNS                 = map[Point][]Point{
		RIGHT: {UP, DOWN},
		LEFT:  {UP, DOWN},
		UP:    {LEFT, RIGHT},
		DOWN:  {LEFT, RIGHT},
	}
)

type Point struct {
	X, Y int
}

type Item struct {
	heatLoss  int
	position  Point
	direction Point
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func key(x, y int, dir Point) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, dir.X, dir.Y)
}

func getMinHeatLoss(heatLossMap [][]int, blocksBeforeTurn, maxInDirection int) int {
	// Initialize a map to store cumulative heat losses for each position and direction.
	dist := make(map[Point]map[Point]int)
	paths := make(map[string][]Point)
	// Initialize the cumulative heat losses for the starting position and directions.
	for x := 0; x < len(heatLossMap[0]); x++ {
		for y := 0; y < len(heatLossMap); y++ {
			dist[Point{x, y}] = map[Point]int{
				RIGHT: math.MaxInt,
				DOWN:  math.MaxInt,
				UP:    math.MaxInt,
				LEFT:  math.MaxInt,
			}
		}
	}

	dist[Point{0, 0}][RIGHT] = 0
	paths[key(0, 0, RIGHT)] = make([]Point, 0)
	dist[Point{0, 0}][DOWN] = 0
	paths[key(0, 0, DOWN)] = make([]Point, 0)
	dist[Point{0, 0}][UP] = 0
	paths[key(0, 0, LEFT)] = make([]Point, 0)
	dist[Point{0, 0}][LEFT] = 0
	paths[key(0, 0, UP)] = make([]Point, 0)

	// Priority queue for Dijkstra's algorithm
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{0, Point{0, 0}, RIGHT})
	heap.Push(&pq, &Item{0, Point{0, 0}, DOWN})

	// Dijkstra's algorithm loop
	for len(pq) > 0 {
		// Pop the item with the minimum heat loss from the priority queue
		item := heap.Pop(&pq).(*Item)
		heatLoss, position, direction := item.heatLoss, item.position, item.direction

		// Skip if the current heat loss is greater than the recorded heat loss for the position and direction
		if heatLoss > dist[position][direction] {
			continue
		}

		// Move in the current direction and accumulate heat losses
		x, y := position.X, position.Y
		for block := 0; block < maxInDirection; block++ {
			x, y = x+direction.X, y+direction.Y

			// Break if out of bounds
			if x < 0 || x >= len(heatLossMap[0]) || y < 0 || y >= len(heatLossMap) {
				break
			}

			// Accumulate heat losses
			heatLoss += heatLossMap[y][x]

			// Crucible needs to move a minimum of N blocks in that direction before it can turn
			if block < blocksBeforeTurn {
				continue
			}

			// Turn the crucible and update heat losses for the new direction
			for _, newDir := range TURNS[direction] {
				if heatLoss < dist[Point{x, y}][newDir] {
					dist[Point{x, y}][newDir] = heatLoss
					heap.Push(&pq, &Item{heatLoss, Point{x, y}, newDir})
					paths[key(x, y, newDir)] = append(paths[key(x, y, newDir)], position)
				}
			}
		}
	}

	// Find the minimum heat loss among the final positions for each direction
	result := math.MaxInt
	for _, direction := range []Point{RIGHT, DOWN, UP, LEFT} {
		result = int(math.Min(float64(result), float64(dist[Point{len(heatLossMap) - 1, len(heatLossMap[0]) - 1}][direction])))
	}

	for k, v, := range dist {
		fmt.Println(k, v)
	}

	// Print the path
	for k, v := range paths {
		fmt.Println(k, v)
	}

	return result
}

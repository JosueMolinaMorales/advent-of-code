package seventeen

import (
	"container/heap"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
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
	// input, _ := os.ReadFile("./input/day17.txt")
	// fmt.Println("Part 1", partOne(string(input))) // 1139 too high
	fmt.Println("Part 1", partOne(testInput))
}

func partOne(input string) int {
	m := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			row = append(row, num)
		}
		m = append(m, row)
	}

	minHeatLoss := minHeatLossDijkstra(m)
	fmt.Println(minHeatLoss)

	return 0
}

type Point struct {
	row, col         int
	heatLoss         int
	prev             *Point
	numStraightMoves int
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int      { return len(pq) }
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Get(p *Point) int {
	for i, point := range pq {
		if point.row == p.row && point.col == p.col {
			return i
		}
	}
	return -1
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func key(p *Point) string {
	prev := p.prev
	if prev == nil {
		prev = &Point{-1, -1, -1, nil, -1}
	}
	return fmt.Sprintf("%d,%d,%d,%d", p.row, p.col, prev.row, prev.col)
}

func minHeatLossDijkstra(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])

	visited := make(map[string]bool)
	// Initialize priority queue
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Initialize distance array
	distance := make([][]int, rows)
	for i := range distance {
		distance[i] = make([]int, cols)
		for j := range distance[i] {
			distance[i][j] = -1
		}
	}

	// so instead of just using the position (x, y) as key in the lookup table / the queue, you need to use
	// the position (x, y) + the previous position (x, y) + number of straight moves that took to get to the position
	// Start from the top-left corner
	start := &Point{0, 0, matrix[0][0], nil, 0}
	heap.Push(&pq, start)
	distance[0][0] = start.heatLoss

	// Directions: right, down, left, up
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	points := make([]*Point, 0)
	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Point)
		if visited[key(current)] {
			continue
		}

		points = append(points, current)
		fmt.Printf("CP (%d, %d) heatLoss: %d, key: %s\n", current.row, current.col, current.heatLoss, key(current))
		straightMoves := straightCount(current)
		fmt.Printf("straightMoves: %d\n", straightMoves)
		for _, dir := range directions {
			newRow, newCol := current.row+dir[0], current.col+dir[1]
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}
			moves := straightMoves
			if current.row == newRow || current.col == newCol {
				moves++
			} else {
				moves = 0
			}
			newPoint := &Point{newRow, newCol, matrix[newRow][newCol], current, moves}
			// If Point has not been seen before, add it to the queue
			if _, ok := visited[key(newPoint)]; !ok {
				fmt.Printf("Pushing (%d, %d) -> %s\n", newPoint.row, newPoint.col, key(newPoint))
				heap.Push(&pq, newPoint)
				continue
			}
			point := pq[pq.Get(newPoint)]
			movedDirection := false
			if straightMoves >= 3 {
				// Cant move in the same direction, need to go left or right
				if current.row == newRow || current.col == newCol {
					continue
				}
				movedDirection = true
			}

			newHeatLoss := current.heatLoss + point.heatLoss
			// Can only update the distance if the last three moves were not in the same direction
			if distance[newRow][newCol] == -1 || newHeatLoss < distance[newRow][newCol] {
				distance[newRow][newCol] = newHeatLoss
				if movedDirection {
					point.numStraightMoves = 0
				}
				fmt.Println("HERE")
				// A new heat loss was calculated. So we need to update the point within the queue
				// with the new heat loss
				idx := pq.Get(point)
				heap.Remove(&pq, idx)
				heap.Push(&pq, point)
			}

		}

	}

	// Print out distances map
	for _, row := range distance {
		fmt.Println(row)
	}

	// Print out the path
	i, _ := iterators.Find(points, func(p *Point) bool {
		return p.row == rows-1 && p.col == cols-1
	})
	current := points[i]
	path := make([]*Point, 0)
	for current != nil {
		// fmt.Printf("(%d, %d) -> ", current.row, current.col)
		current = current.prev
		path = append(path, current)
	}

	for i, row := range matrix {
		for j, col := range row {
			if slices.ContainsFunc(path, func(p *Point) bool {
				return p != nil && p.row == i && p.col == j
			}) {
				fmt.Printf("%d ", col)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}

	// The minimum heat loss at the bottom-right corner
	return distance[rows-1][cols-1]
}

func straightCount(current *Point) int {
	if current.prev == nil {
		return 0
	}
	og := current
	count := 0
	for current.prev != nil {
		if current.row == current.prev.row {
			count++
		} else {
			break
		}
		current = current.prev
	}

	if count >= 3 {
		return count
	}

	count = 0
	current = og
	for current.prev != nil {
		if current.col == current.prev.col {
			count++
		} else {
			break
		}
		current = current.prev
	}
	return count
}

// func getAdjPoints(p Point, matrix *[][]int) []Point {
// 	directions := [][2]int{
// 		{0, 1},  // Right
// 		{0, -1}, // left
// 		{1, 0},  // Down
// 		{-1, 0}, // Up
// 	}
// 	adj := make([]Point, 0)

// 	for _, dir := range directions {
// 		dx, dy := p.row+dir[0], p.col+dir[1]
// 		if dx < 0 || dx >= len(*matrix) || dy < 0 || dy >= len((*matrix)[0]) {
// 			continue
// 		}
// 		adj = append(adj, Point{dx, dy, (*matrix)[dx][dy]})
// 	}

// 	return adj
// }

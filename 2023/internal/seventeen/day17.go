package seventeen

import (
	"container/heap"
	"fmt"
	"math"
	"os"
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
	input, _ := os.ReadFile("./input/day17.txt")
	fmt.Println("Part 1", partOne(string(input)))
	// fmt.Println("Part 1", partOne(testInput))
}

func partTwo(input string) int {
	m := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]int, 0)
		for _, c := range line {
			num, _ := strconv.Atoi(string(c))
			row = append(row, num)
		}
		m = append(m, row)
	}

	minHeatLoss := minHeatLossDijkstra(m, 3, 10)

	return minHeatLoss
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

	minHeatLoss := minHeatLossDijkstra(m, 0, 3)

	return minHeatLoss
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
		if point.row == p.row && point.col == p.col && point.prev.row == p.prev.row && point.prev.col == p.prev.col && point.numStraightMoves == p.numStraightMoves {
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
	return fmt.Sprintf("%d,%d,%d,%d,%d", p.row, p.col, prev.row, prev.col, p.numStraightMoves)
}

func minHeatLossDijkstra(matrix [][]int, maxInDir int, blockBeforeTurn int) int {
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
			distance[i][j] = math.MaxInt
		}
	}

	start := &Point{0, 0, matrix[0][0], nil, 0}
	heap.Push(&pq, start)
	heap.Push(&pq, &Point{0, 1, matrix[0][1], start, 0})

	distance[0][0] = start.heatLoss

	// Directions: right, down, left, up
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	points := make([]*Point, 0)
	for pq.Len() > 0 {
		// Get the current point
		current := heap.Pop(&pq).(*Point)
		// If this point has already been visited, skip it
		if visited[key(current)] {
			continue
		}
		// Set the point to visited
		visited[key(current)] = true
		points = append(points, current)
		straightMoves := straightCount(current)

		// Look at the adjacent points
		for _, dir := range directions {
			newRow, newCol := current.row+dir[0], current.col+dir[1]
			// Bounds check
			if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols {
				continue
			}
			moves := straightMoves
			moved := false
			if moves > 3 {
				// Cant move in the same direction, need to go left or right
				if current.row == newRow || current.col == newCol {
					continue
				}
				moved = true
			}
			if moved {
				moves = 0
			} else {
				moves++
			}

			newPoint := &Point{newRow, newCol, matrix[newRow][newCol] + current.heatLoss, current, moves}
			// if the point is already in the queue, update the heat loss
			idx := pq.Get(newPoint)
			if idx != -1 {
				if newPoint.heatLoss < pq[idx].heatLoss {
					pq[idx].heatLoss = newPoint.heatLoss
					pq[idx].prev = newPoint.prev
					pq[idx].numStraightMoves = newPoint.numStraightMoves
					heap.Fix(&pq, idx)
				}
			}

			// Add the point to the queue
			heap.Push(&pq, newPoint)
			distance[newRow][newCol] = int(math.Min(float64(distance[newRow][newCol]), float64(newPoint.heatLoss)))
		}

	}

	// For fun print the path
	printPath(points, matrix)

	// The minimum heat loss at the bottom-right corner
	return distance[rows-1][cols-1]
}

func printPath(points []*Point, m [][]int) {
	// Print out the path
	i, _ := iterators.Find(points, func(p *Point) bool {
		return p.row == len(m)-1 && p.col == len(m[0])-1
	})

	current := points[i]
	path := make([]*Point, 0)
	for current != nil {
		path = append(path, current)
		current = current.prev
	}
	for i, row := range m {
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
}

func straightCount(current *Point) int {
	if current.prev == nil {
		return 0
	}
	og := current
	hCount := 0
	for current.prev != nil {
		if current.row == current.prev.row {
			hCount++
		} else {
			break
		}
		current = current.prev
	}

	if hCount >= 3 {
		return hCount
	} else if hCount != 0 {
		return hCount
	}

	count := 0
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

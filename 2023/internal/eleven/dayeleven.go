package eleven

import (
	"container/heap"
	"fmt"
	"slices"
	"strings"
	"sync"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func RunDayEleven() {
	// input, err := os.ReadFile("./input/day11.txt")
	// if err != nil {
	// 	panic("Failed to read day 11 input file")
	// }
	// fmt.Println("Part 1:", partOne(string(input)))
	fmt.Println("Part 1:", partOne(testInput))
}

func partOne(input string) int {
	matrix := make([][]string, 0)
	for _, line := range strings.Split(input, "\n") {
		row := make([]string, 0)
		for _, ch := range line {
			// Convert rune to string
			row = append(row, string(ch))
		}
		matrix = append(matrix, row)
	}

	// fmt.Println("Before expansion:")
	// printMatrix(matrix)

	// Expand matrix
	// Check rows
	rows := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		emptySpace := iterators.Every(matrix[i], func(ch string) bool {
			return ch == "."
		})
		if emptySpace {
			rows = append(rows, i)
		}
	}

	// Check columns for empty space
	i, j := 0, 0
	cols := make([]int, 0)
	for j < len(matrix[0]) {
		emptySpace := true
		for i < len(matrix) {
			if matrix[i][j] != "." {
				emptySpace = false
				break
			}
			i++
		}
		if emptySpace {
			cols = append(cols, j)
		}
		j++
		i = 0
	}

	galaxies := make([]Point, 0)
	// matrix = matrixCopy
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != "." {
				galaxies = append(galaxies, Point{x: i, y: j})
			}
		}
	}

	// Find all pairs of galaxies
	pairs := make([][]Point, 0)
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, []Point{galaxies[i], galaxies[j]})
		}
	}
	fmt.Println("Len pairs:", len(pairs))
	// For every galaxy pair, find the shortest path between them using BFS
	// Create a channel to receive results
	resultChan := make(chan Result, len(pairs))

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch goroutines for each pair
	for _, pair := range pairs {
		wg.Add(1)
		go dijkstra(matrix, pair[0], pair[1], rows, cols, 2, resultChan, &wg)
	}

	// Create a goroutine to close the resultChan when all goroutines finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Sum the path lengths as results come in
	sum := 0
	for result := range resultChan {
		sum += (result.Distance - 1)
	}

	fmt.Println("Total sum:", sum)

	return sum
}

type Point struct {
	x, y int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

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

type Item struct {
	value    Point
	priority int
}

// Result represents the result of a Dijkstra's algorithm run.
type Result struct {
	Distance int
}

func dijkstra(matrix [][]string, start, target Point, emptyRows, emptyCols []int, factor int, resultChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	visited := make(map[Point]bool)
	distance := make(map[Point]int)
	parent := make(map[Point]Point)
	pq := make(PriorityQueue, 0)

	heap.Push(&pq, &Item{value: start, priority: 0})
	distance[start] = 0

	for pq.Len() > 0 {
		currentItem := heap.Pop(&pq).(*Item)
		current := currentItem.value

		if current == target {
			// path := reconstructPath(parent, start, target)
			fmt.Printf("Found path from (%d, %d) --> (%d, %d): %d\n", start.x, start.y, target.x, target.y, distance[current])
			resultChan <- Result{Distance: distance[current]}
			return
		}

		if !visited[current] {
			visited[current] = true

			for _, neighbor := range getNeighbors(matrix, current) {
				cost := 1
				if slices.Contains(emptyRows, neighbor.x) {
					cost = factor
				} else if slices.Contains(emptyCols, neighbor.y) {
					cost = factor
				}
				if distance[neighbor] == 0 || distance[current]+cost < distance[neighbor] {
					distance[neighbor] = distance[current] + cost
					parent[neighbor] = current
					heap.Push(&pq, &Item{value: neighbor, priority: distance[neighbor]})
				}
			}
		}
	}

	resultChan <- Result{Distance: 0} // No path found
}

func getNeighbors(matrix [][]string, p Point) []Point {
	neighbors := make([]Point, 0)

	// Assuming movements are allowed in all 8 directions (up, down, left, right, and diagonals)
	directions := []Point{
		{1, 0},  // Down
		{-1, 0}, // Up
		{0, 1},  // Right
		{0, -1}, // Left
	}

	for _, dir := range directions {
		newX, newY := p.x+dir.x, p.y+dir.y

		if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
			neighbors = append(neighbors, Point{newX, newY})
		}
	}

	return neighbors
}

func reconstructPath(parent map[Point]Point, start, target Point) []Point {
	path := make([]Point, 0)
	current := target

	for current != start {
		path = append([]Point{current}, path...)
		current = parent[current]
	}

	path = append([]Point{start}, path...)
	return path
}

func printMatrix(matrix [][]string) {
	for _, row := range matrix {
		for _, ch := range row {
			print(ch)
		}
		println()
	}
}

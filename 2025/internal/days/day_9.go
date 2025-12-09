package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

type Point struct {
	x, y int
}

func Day9() {
	fmt.Println("2025 Day 9 Part 1:", day9Part1("inputs/day_9/input.txt"))
	fmt.Println("2025 Day 9 Part 2:", day9Part2("inputs/day_9/input.txt"))
}

func parsePoints(path string) ([]Point, error) {
	lines, err := io.ReadFileAsLines(path)
	if err != nil {
		return nil, err
	}

	points := make([]Point, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		points = append(points, Point{x, y})
	}
	return points, nil
}

func day9Part1(path string) int {
	points, _ := parsePoints(path)

	maxArea := 0
	// Try all pairs of points as opposite corners
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			// Calculate area of rectangle with opposite corners at p1 and p2
			// Add 1 because we need to count the tiles at the corners too
			width := int(math.Abs(float64(p1.x-p2.x))) + 1
			height := int(math.Abs(float64(p1.y-p2.y))) + 1
			area := width * height

			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func day9Part2(path string) int {
	points, _ := parsePoints(path)

	// Coordinate compression: reduce large sparse grid to only relevant coordinates
	xCoords := make(map[int]bool)
	yCoords := make(map[int]bool)
	for _, p := range points {
		xCoords[p.x] = true
		yCoords[p.y] = true
	}

	// Convert to sorted slices
	xList := make([]int, 0, len(xCoords))
	for x := range xCoords {
		xList = append(xList, x)
	}
	yList := make([]int, 0, len(yCoords))
	for y := range yCoords {
		yList = append(yList, y)
	}
	sort.Ints(xList)
	sort.Ints(yList)

	// Create mappings
	xToCompressed := make(map[int]int)
	compressedToX := make(map[int]int)
	for i, x := range xList {
		xToCompressed[x] = i
		compressedToX[i] = x
	}

	yToCompressed := make(map[int]int)
	compressedToY := make(map[int]int)
	for i, y := range yList {
		yToCompressed[y] = i
		compressedToY[i] = y
	}

	// Convert points to compressed coordinates
	compressedPoints := make([]Point, len(points))
	for i, p := range points {
		compressedPoints[i] = Point{xToCompressed[p.x], yToCompressed[p.y]}
	}

	// Build boundary edges in compressed space (all tiles on polygon perimeter)
	greenTiles := make(map[Point]bool)

	for i := 0; i < len(compressedPoints); i++ {
		p1 := compressedPoints[i]
		p2 := compressedPoints[(i+1)%len(compressedPoints)]

		if p1.x == p2.x {
			// Vertical line in compressed space
			minY, maxY := p1.y, p2.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for cy := minY; cy <= maxY; cy++ {
				greenTiles[Point{p1.x, cy}] = true
			}
		} else {
			// Horizontal line in compressed space
			minX, maxX := p1.x, p2.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for cx := minX; cx <= maxX; cx++ {
				greenTiles[Point{cx, p1.y}] = true
			}
		}
	}

	// Pre-compute all interior points in compressed space
	for cy := 0; cy < len(yList); cy++ {
		for cx := 0; cx < len(xList); cx++ {
			cp := Point{cx, cy}
			if greenTiles[cp] {
				continue
			}
			// Check if this compressed point is inside using original coordinates
			origP := Point{compressedToX[cx], compressedToY[cy]}
			if isInside(origP, points) {
				greenTiles[cp] = true
			}
		}
	}

	// Find the largest rectangle where corners are red tiles and all interior is green
	maxArea := 0
	for i := 0; i < len(compressedPoints); i++ {
		for j := i + 1; j < len(compressedPoints); j++ {
			cp1, cp2 := compressedPoints[i], compressedPoints[j]

			// Get compressed bounds
			minCX, maxCX := cp1.x, cp2.x
			if minCX > maxCX {
				minCX, maxCX = maxCX, minCX
			}
			minCY, maxCY := cp1.y, cp2.y
			if minCY > maxCY {
				minCY, maxCY = maxCY, minCY
			}

			// Check if all boundary points in compressed space are green
			valid := true

			// Check top and bottom edges
			for cx := minCX; cx <= maxCX && valid; cx++ {
				if !greenTiles[Point{cx, minCY}] || !greenTiles[Point{cx, maxCY}] {
					valid = false
				}
			}

			// Check left and right edges (excluding corners already checked)
			for cy := minCY + 1; cy < maxCY && valid; cy++ {
				if !greenTiles[Point{minCX, cy}] || !greenTiles[Point{maxCX, cy}] {
					valid = false
				}
			}

			if valid {
				// Calculate actual area using original (uncompressed) coordinates
				width := compressedToX[maxCX] - compressedToX[minCX] + 1
				height := compressedToY[maxCY] - compressedToY[minCY] + 1
				area := width * height
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

// Ray casting algorithm to determine if a point is inside a polygon
func isInside(p Point, polygon []Point) bool {
	n := len(polygon)
	inside := false

	j := n - 1
	for i := 0; i < n; i++ {
		pi, pj := polygon[i], polygon[j]

		if ((pi.y > p.y) != (pj.y > p.y)) &&
			(p.x < (pj.x-pi.x)*(p.y-pi.y)/(pj.y-pi.y)+pi.x) {
			inside = !inside
		}
		j = i
	}

	return inside
}

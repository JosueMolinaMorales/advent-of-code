package days

import (
	"fmt"
	"math"
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

	fmt.Printf("Part 2: Processing %d red tiles...\n", len(points))

	// Build a map of green tiles (tiles on the path between consecutive red tiles)
	// and tiles inside the polygon formed by the red tiles
	greenTiles := make(map[Point]bool)

	fmt.Println("Step 1/3: Building edges between red tiles...")
	// Add all points on edges between consecutive red tiles
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]

		// Add all points on the line between p1 and p2
		if p1.x == p2.x {
			// Vertical line
			minY, maxY := p1.y, p2.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}
			for y := minY; y <= maxY; y++ {
				greenTiles[Point{p1.x, y}] = true
			}
		} else {
			// Horizontal line
			minX, maxX := p1.x, p2.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			for x := minX; x <= maxX; x++ {
				greenTiles[Point{x, p1.y}] = true
			}
		}
	}

	// Find all tiles inside the polygon using scanline algorithm
	fmt.Println("Step 2/3: Finding interior tiles using ray casting...")
	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// For each scanline, find interior points using ray casting
	totalRows := maxY - minY + 1
	for y := minY; y <= maxY; y++ {
		if (y-minY)%100 == 0 {
			fmt.Printf("  Scanning row %d/%d (%.1f%%)\n", y-minY+1, totalRows, float64(y-minY+1)*100.0/float64(totalRows))
		}
		for x := minX; x <= maxX; x++ {
			p := Point{x, y}
			if greenTiles[p] {
				continue
			}
			if isInside(p, points) {
				greenTiles[p] = true
			}
		}
	}

	fmt.Printf("Step 3/3: Checking %d pairs of red tiles for valid rectangles...\n", len(points)*(len(points)-1)/2)
	// Now find the largest rectangle where opposite corners are red tiles
	// and all tiles in between are green or red
	maxArea := 0
	pairsChecked := 0
	totalPairs := len(points) * (len(points) - 1) / 2

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairsChecked++
			if pairsChecked%10000 == 0 {
				fmt.Printf("  Checked %d/%d pairs (%.1f%%), current max area: %d\n",
					pairsChecked, totalPairs, float64(pairsChecked)*100.0/float64(totalPairs), maxArea)
			}
			p1, p2 := points[i], points[j]

			// Check if rectangle only contains green/red tiles
			minRectX, maxRectX := p1.x, p2.x
			if minRectX > maxRectX {
				minRectX, maxRectX = maxRectX, minRectX
			}
			minRectY, maxRectY := p1.y, p2.y
			if minRectY > maxRectY {
				minRectY, maxRectY = maxRectY, minRectY
			}

			valid := true
			for x := minRectX; x <= maxRectX && valid; x++ {
				for y := minRectY; y <= maxRectY && valid; y++ {
					if !greenTiles[Point{x, y}] {
						valid = false
					}
				}
			}

			if valid {
				// Add 1 to include the corner tiles
				area := (maxRectX - minRectX + 1) * (maxRectY - minRectY + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Printf("Completed! Checked %d pairs total.\n", pairsChecked)
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

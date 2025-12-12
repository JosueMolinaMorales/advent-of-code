package days

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2025/internal/util/io"
)

type PresentShape struct {
	cells  [][]bool // true = #, false = .
	width  int
	height int
}

type TreeRegion struct {
	width    int
	height   int
	presents []int // count of each shape type needed
}

func Day12() {
	lines, err := io.ReadFileAsLines("inputs/day_12/input.txt")
	if err != nil {
		log.Fatalf("Failed to read input for day 12: %s", err)
	}

	shapeMap, regions := parseDay12Input(lines)
	result := solvePart1(shapeMap, regions)
	fmt.Printf("2025 Day 12 Part 1: %d\n", result)
}

func parseDay12Input(lines []string) (map[int]PresentShape, []TreeRegion) {
	shapeMap := make(map[int]PresentShape)
	regions := []TreeRegion{}

	i := 0
	// Parse shapes
	for i < len(lines) {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			i++
			continue
		}

		// Check if this is a shape definition (starts with number:)
		if strings.Contains(line, ":") && !strings.Contains(line, "x") {
			// Extract shape ID
			shapeID, _ := strconv.Atoi(strings.TrimSuffix(line, ":"))

			// Read the shape lines
			i++
			shapeLines := []string{}
			for i < len(lines) {
				currentLine := lines[i]
				trimmed := strings.TrimSpace(currentLine)
				if trimmed == "" || strings.Contains(trimmed, ":") {
					break
				}
				shapeLines = append(shapeLines, currentLine)
				i++
			}

			if len(shapeLines) > 0 {
				shape := parsePresentShape(shapeLines)
				shapeMap[shapeID] = shape
			}
		} else {
			i++
		}
	}

	// Parse regions - reset to start and look for region lines
	for i = 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		// Parse region line: "WxH: count0 count1 count2 ..."
		if strings.Contains(line, "x") && strings.Contains(line, ":") {
			parts := strings.Split(line, ":")
			if len(parts) == 2 {
				dims := strings.Split(strings.TrimSpace(parts[0]), "x")
				if len(dims) == 2 {
					width, _ := strconv.Atoi(dims[0])
					height, _ := strconv.Atoi(dims[1])

					countsStr := strings.Fields(strings.TrimSpace(parts[1]))
					presents := []int{}
					for _, cs := range countsStr {
						count, _ := strconv.Atoi(cs)
						presents = append(presents, count)
					}

					regions = append(regions, TreeRegion{width, height, presents})
				}
			}
		}
	}

	return shapeMap, regions
}

func parsePresentShape(lines []string) PresentShape {
	height := len(lines)
	width := 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}

	cells := make([][]bool, height)
	for i := range cells {
		cells[i] = make([]bool, width)
		for j := 0; j < len(lines[i]) && j < width; j++ {
			cells[i][j] = lines[i][j] == '#'
		}
	}

	return PresentShape{cells, width, height}
}

func solvePart1(shapeMap map[int]PresentShape, regions []TreeRegion) int {
	count := 0
	for _, region := range regions {
		if canFitByArea(shapeMap, region) {
			count++
		}
	}
	return count
}

// Calculate the area (number of # cells) in a shape
func getShapeArea(shape PresentShape) int {
	area := 0
	for r := 0; r < shape.height; r++ {
		for c := 0; c < shape.width; c++ {
			if shape.cells[r][c] {
				area++
			}
		}
	}
	return area
}

// canFitByArea checks if presents fit using simple area comparison
func canFitByArea(shapeMap map[int]PresentShape, region TreeRegion) bool {
	regionArea := region.width * region.height
	totalPresentArea := 0

	for shapeIdx, count := range region.presents {
		if count > 0 {
			shape, exists := shapeMap[shapeIdx]
			if !exists {
				return false
			}
			shapeArea := getShapeArea(shape)
			totalPresentArea += shapeArea * count
		}
	}

	return totalPresentArea <= regionArea
}

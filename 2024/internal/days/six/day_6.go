package six

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/sets/hashset"
)

type Guard struct {
	Direction types.Direction
	Location  types.Vector
}

func NewGuard(dir types.Direction, loc types.Vector) Guard {
	return Guard{
		Direction: dir,
		Location:  loc,
	}
}

func SolveDay6() {
	fmt.Println("Day 6 Part 1: ", solvePartOne())
	fmt.Println("Day 6 Part 2: ", solvePartTwo())
}

func setup() (Guard, *hashset.Set, int, int) {
	rawMap, err := util.LoadFileAsString("./inputs/day_6.txt")
	if err != nil {
		panic(err)
	}

	var guard Guard
	blocks := hashset.New()
	xBound, yBound := 0, 0

	for i, row := range strings.Split(rawMap, "\n") {
		xBound = i + 1
		for j, col := range strings.Split(row, "") {
			yBound = j + 1
			vec := types.NewVector(i, j)

			switch col {
			case ">":
				guard = NewGuard(types.DIRECTION_EAST, *vec)
			case "<":
				guard = NewGuard(types.DIRECTION_WEST, *vec)
			case "^":
				guard = NewGuard(types.DIRECTION_NORTH, *vec)
			case "v":
				guard = NewGuard(types.DIRECTION_SOUTH, *vec)
			case "#":
				blocks.Add(*vec)
			}
		}
	}

	return guard, blocks, xBound, yBound
}

func findPaths(g Guard, blocks *hashset.Set, xBound, yBound int, findCycle bool) (*hashset.Set, bool) {
	visited := hashset.New()
	path := hashset.New()

	for {
		state := Guard{
			Location:  g.Location,
			Direction: g.Direction,
		}

		// Stop if guard's location is off the map
		if g.Location.X < 0 || g.Location.X >= xBound || g.Location.Y < 0 || g.Location.Y >= yBound {
			return path, false
		}

		if findCycle && visited.Contains(state) {
			// Found a cycle
			return path, true
		}

		visited.Add(state)
		path.Add(g.Location)

		// Move the guard
		g = moveGuard(g, blocks)
	}
}

func solvePartOne() int {
	guard, blocks, xBound, yBound := setup()
	path, _ := findPaths(guard, blocks, xBound, yBound, false)
	return path.Size()
}

func solvePartTwo() int {
	guard, blocks, xBound, yBound := setup()

	path, _ := findPaths(guard, blocks, xBound, yBound, false)
	path.Remove(guard.Location)

	attempted := hashset.New()
	cycleCount := 0

	for _, item := range path.Values() {
		block := item.(types.Vector)
		if attempted.Contains(block) {
			continue
		}

		attempted.Add(block)
		blocks.Add(block)

		_, found := findPaths(guard, blocks, xBound, yBound, true)
		if found {
			cycleCount++
		}

		blocks.Remove(block)
	}

	return cycleCount
}

func moveGuard(g Guard, blocks *hashset.Set) Guard {
	dx, dy := g.Location.X, g.Location.Y

	switch g.Direction {
	case types.DIRECTION_EAST:
		dy++
	case types.DIRECTION_NORTH:
		dx--
	case types.DIRECTION_SOUTH:
		dx++
	case types.DIRECTION_WEST:
		dy--
	}

	// If the guard hits a block, change direction
	if blocks.Contains(*types.NewVector(dx, dy)) {
		switch g.Direction {
		case types.DIRECTION_EAST:
			g.Direction = types.DIRECTION_SOUTH
		case types.DIRECTION_NORTH:
			g.Direction = types.DIRECTION_EAST
		case types.DIRECTION_SOUTH:
			g.Direction = types.DIRECTION_WEST
		case types.DIRECTION_WEST:
			g.Direction = types.DIRECTION_NORTH
		}
		return g
	}

	g.Location.X = dx
	g.Location.Y = dy
	return g
}

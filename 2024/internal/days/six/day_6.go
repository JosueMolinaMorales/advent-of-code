package six

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/sets/hashset"
)

type guard struct {
	direction types.Direction
	loc       *types.Vector
}

func newGuard(dir types.Direction, loc *types.Vector) *guard {
	return &guard{
		direction: dir,
		loc:       loc,
	}
}

func (g *guard) CopyLocation() *types.Vector {
	return types.NewVector(g.loc.X, g.loc.Y)
}

func SolveDay6() {
	res := solvePartOne()
	fmt.Println(res)
	res = solvePartTwo()
	fmt.Println(res)
}

func setup() (*guard, *hashset.Set, int, int) {
	rawMap, err := util.LoadFileAsString("./inputs/day_6.txt")
	if err != nil {
		panic(err)
	}

	// grid := make([][]string, 0)
	var guard *guard
	blocks := hashset.New()

	xBound := 0
	yBound := 0
	for i, row := range strings.Split(rawMap, "\n") {
		xBound = i
		// r := make([]string, 0)
		for j, col := range strings.Split(row, "") {
			yBound = j
			vec := types.NewVector(i, j)
			var direction types.Direction
			switch col {
			case ">":
				direction = types.DIRECTION_EAST
			case "<":
				direction = types.DIRECTION_WEST
			case "^":
				direction = types.DIRECTION_NORTH
			case "v":
				direction = types.DIRECTION_SOUTH
			case "#":
				blocks.Add(*vec)
				continue
			default:
				continue
			}
			guard = newGuard(direction, vec)
		}
	}
	xBound++
	yBound++

	return guard, blocks, xBound, yBound
}

func solvePartOne() int {
	guard, blocks, xBound, yBound := setup()
	path := hashset.New()
	// Stop when the guard's location is off the map
	for guard.loc.X < xBound && guard.loc.X >= 0 && guard.loc.Y < yBound && guard.loc.Y >= 0 {
		// Count the guards location
		path.Add(*guard.CopyLocation())

		// Move the guard
		moveGuard(guard, blocks)
	}

	return path.Size()
}

func solvePartTwo() int {
	guard, blocks, xBound, yBound := setup()
	gSp := guard.CopyLocation()
	gsd := guard.direction
	cycles := 0
	// Loop through the entire grid and try to place an obstacle
	// Dont place an obstacle where the is already one and dont place one on the starting position of the guard
	for i := 0; i < xBound; i++ {
		for j := 0; j < yBound; j++ {
			// Place a block, if there is already a block or if its the start pos of guard, skip
			vec := *types.NewVector(i, j)
			if blocks.Contains(vec) || (vec.X == gSp.X && vec.Y == gSp.Y) {
				continue
			}
			// Add block
			blocks.Add(vec)
			// Check to see if there is a cycle
			keepGoing := true
			cycleDetected := true
			path := hashset.New()
			for keepGoing {
				// If the g is out of the map, no cycle
				if guard.loc.X >= xBound || guard.loc.X < 0 || guard.loc.Y >= yBound || guard.loc.Y < 0 {
					// fmt.Println("out of bounds")
					cycleDetected = false
					break
				}
				// check to see if the guard has already been to this location with the current direction
				state := struct {
					pos types.Vector
					dir types.Direction
				}{
					pos: *types.NewVector(guard.loc.X, guard.loc.Y),
					dir: guard.direction,
				}
				if path.Contains(state) {
					// fmt.Println("We have been to this position before: ", state, vec)
					cycleDetected = true
					break
				}
				path.Add(state)

				// Move guard
				moveGuard(guard, blocks)
			}
			if cycleDetected {
				cycles++
			}
			// Remove it
			blocks.Remove(vec)
			// Reset guards position
			guard.direction = gsd
			guard.loc.X = gSp.X
			guard.loc.Y = gSp.Y
		}
	}

	return cycles
}

func moveGuard(guard *guard, blocks *hashset.Set) {
	dx, dy := guard.loc.X, guard.loc.Y
	switch guard.direction {
	case types.DIRECTION_EAST:
		dy += 1
	case types.DIRECTION_NORTH:
		dx -= 1
	case types.DIRECTION_SOUTH:
		dx += 1
	case types.DIRECTION_WEST:
		dy -= 1
	}

	// if guard is on a block, move back and change direction
	if blocks.Contains(*types.NewVector(dx, dy)) {
		switch guard.direction {
		case types.DIRECTION_EAST:
			guard.direction = types.DIRECTION_SOUTH
		case types.DIRECTION_NORTH:
			guard.direction = types.DIRECTION_EAST
		case types.DIRECTION_SOUTH:
			guard.direction = types.DIRECTION_WEST
		case types.DIRECTION_WEST:
			guard.direction = types.DIRECTION_NORTH
		}
		return
	}

	guard.loc.X = dx
	guard.loc.Y = dy
}

func printMap(gLoc types.Vector, blocks hashset.Set, xBound, yBound int) {
	for i := 0; i < xBound; i++ {
		for j := 0; j < yBound; j++ {
			v := *types.NewVector(i, j)
			if gLoc.X == i && gLoc.Y == j {
				fmt.Print("G")
			} else if blocks.Contains(v) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("==============================")
}

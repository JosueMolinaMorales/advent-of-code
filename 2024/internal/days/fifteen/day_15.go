package fifteen

import (
	"fmt"
	"slices"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
)

func SolveDay15() {
	fmt.Println("Day 15 Part 1: ", solvePartOne())
	fmt.Println("Day 15 Part 2: ", solvePartTwo())
}

func setup() ([][]string, []string, types.Vector) {
	input, err := util.LoadFileAsString("./inputs/day_15.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(input, "\n\n")
	moves := strings.Split(strings.Join(parts[1:], ""), "")

	robot := *types.NewVector()
	m := [][]string{}
	for i, line := range strings.Split(parts[0], "\n") {
		row := []string{}
		for j, col := range strings.Split(line, "") {
			v := *types.NewVector(i, j)
			if col == "@" {
				robot = v
				row = append(row, ".")
			} else {
				row = append(row, col)
			}
		}
		m = append(m, row)
	}

	return m, moves, robot
}

func solvePartOne() int {
	m, moves, robot := setup()
	for _, move := range moves {
		switch move {
		case "<":
			moveRobot(&robot, &m, types.DIRECTION_WEST)
		case ">":
			moveRobot(&robot, &m, types.DIRECTION_EAST)
		case "v":
			moveRobot(&robot, &m, types.DIRECTION_SOUTH)
		case "^":
			moveRobot(&robot, &m, types.DIRECTION_NORTH)
		}
	}

	return calcGPS("O", m)
}

func solvePartTwo() int {
	// Enlarge map
	m, moves, robot := setup()
	newMap := [][]string{}
	for _, row := range m {
		str := strings.Join(row, "")
		str = strings.ReplaceAll(str, "#", "##")
		str = strings.ReplaceAll(str, "O", "[]")
		str = strings.ReplaceAll(str, ".", "..")
		newMap = append(newMap, strings.Split(str, ""))
	}

	robot.Y *= 2
	for _, move := range moves {
		switch move {
		case "<":
			moveRobotHorizontally(&robot, &newMap, types.DIRECTION_WEST)
		case ">":
			moveRobotHorizontally(&robot, &newMap, types.DIRECTION_EAST)
		case "v":
			moveRobotVertically(&robot, &newMap, types.DIRECTION_SOUTH)
		case "^":
			moveRobotVertically(&robot, &newMap, types.DIRECTION_NORTH)
		}
	}

	return calcGPS("[", newMap)
}

func calcGPS(cell string, m [][]string) int {
	res := 0
	for i, row := range m {
		for j, col := range row {
			if col == cell {
				res += 100*i + j
			}
		}
	}
	return res
}

func printMap(m [][]string, robot types.Vector) {
	fmt.Println()
	for i, row := range m {
		for j, col := range row {
			if *types.NewVector(i, j) == robot {
				fmt.Print("@")
			} else {
				fmt.Print(col)
			}
		}
		fmt.Println()
	}
}

func moveRobot(robot *types.Vector, m *[][]string, direction types.Direction) {
	move := movement(direction)
	dx, dy := robot.X+move.X, robot.Y+move.Y
	toMove := 0
	for {
		// if its a wall, stop
		if (*m)[dx][dy] == "#" {
			return
		}

		// if its a box, check next to the box
		if (*m)[dx][dy] == "O" {
			dx += move.X
			dy += move.Y
			toMove++
			continue
		}

		// If its a empty spot, move everything
		dv := *types.NewVector(dx, dy)
		for i := 0; i < toMove; i++ {
			// Swap
			(*m)[dv.X][dv.Y] = (*m)[dv.X-move.X][dv.Y-move.Y]
			(*m)[dv.X-move.X][dv.Y-move.Y] = "."
			dv.X -= move.X
			dv.Y -= move.Y
		}
		// Move the robot
		robot.X += move.X
		robot.Y += move.Y
		return
	}
}

func moveBoxesVertically(currBox []int, m [][]string, yDir int) ([][]int, bool) {
	// Check to the left and right for a box
	dx_1, dy_1, dx_2, dy_2 := currBox[0]+yDir, currBox[1], currBox[0]+yDir, currBox[2]
	// Check if there is a wall
	cell_1 := m[dx_1][dy_1]
	cell_2 := m[dx_2][dy_2]
	if cell_1 == "#" {
		return nil, false
	}
	if cell_2 == "#" {
		return nil, false
	}

	boxes := [][]int{}
	if cell_1 == "[" {
		boxes = append(boxes, []int{dx_1, dy_1, dy_1 + 1})
	}
	if cell_1 == "]" {
		boxes = append(boxes, []int{dx_1, dy_1 - 1, dy_1})
	}
	if cell_2 == "[" {
		boxes = append(boxes, []int{dx_2, dy_2, dy_2 + 1})
	}
	if cell_2 == "]" && cell_1 != "[" {
		boxes = append(boxes, []int{dx_2, dy_2 - 1, dy_2})
	}

	return boxes, true
}

func moveRobotVertically(robot *types.Vector, m *[][]string, direction types.Direction) {
	move := movement(direction)
	dx, dy := robot.X+move.X, robot.Y
	boxesToMove := [][]int{}
	if (*m)[dx][dy] == "#" {
		return
	}
	// Check if there are any boxes on this level
	if (*m)[dx][dy] == "]" {
		boxesToMove = append(boxesToMove, []int{dx, dy - 1, dy})
	}
	if (*m)[dx][dy] == "[" {
		boxesToMove = append(boxesToMove, []int{dx, dy, dy + 1})
	}
	// Need to check above each of the boxes
	n := len(boxesToMove)
	for i := 0; i < n; i++ {
		box := boxesToMove[i]
		boxes, canMove := moveBoxesVertically(box, *m, move.X)
		if !canMove {
			return
		}
		boxesToMove = append(boxesToMove, boxes...)
		n += len(boxes)
	}

	// Move all the boxes
	if direction == types.DIRECTION_NORTH {
		slices.SortFunc(boxesToMove, func(a, b []int) int {
			return a[0] - b[0]
		})
	}
	if direction == types.DIRECTION_SOUTH {
		slices.SortFunc(boxesToMove, func(a, b []int) int {
			return b[0] - a[0]
		})
	}

	for _, box := range boxesToMove {
		dx, dy := box[0]+move.X, box[1] // The '['
		(*m)[dx][dy] = "["
		(*m)[dx][dy+1] = "]"
		(*m)[box[0]][box[1]] = "."
		(*m)[box[0]][box[2]] = "."
	}
	// Move the robot
	robot.X += move.X
}

func moveRobotHorizontally(robot *types.Vector, m *[][]string, direction types.Direction) {
	move := movement(direction)
	dx, dy := robot.X+move.X, robot.Y+move.Y
	toMove := 0
	for {
		// Check for wall
		if (*m)[dx][dy] == "#" {
			return
		}
		cell := (*m)[dx][dy]
		// check if its part of a box
		if cell == "[" || cell == "]" {
			toMove += 2
			dy += move.Y * 2
			continue
		}
		// If its a empty spot, move everything
		dv := *types.NewVector(dx, dy)
		for i := 0; i < toMove; i++ {
			// Swap
			(*m)[dv.X][dv.Y] = (*m)[dv.X-move.X][dv.Y-move.Y]
			(*m)[dv.X-move.X][dv.Y-move.Y] = "."
			dv.X -= move.X
			dv.Y -= move.Y
		}
		// Move the robot
		robot.X += move.X
		robot.Y += move.Y
		return
	}
}

func movement(dir types.Direction) types.Vector {
	switch dir {
	case types.DIRECTION_EAST:
		return *types.NewVector(0, 1)
	case types.DIRECTION_NORTH:
		return *types.NewVector(-1, 0)
	case types.DIRECTION_SOUTH:
		return *types.NewVector(1, 0)
	case types.DIRECTION_WEST:
		return *types.NewVector(0, -1)
	default:
		panic("Not a valid direction")
	}
}

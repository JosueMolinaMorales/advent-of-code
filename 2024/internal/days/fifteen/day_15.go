package fifteen

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
)

func SolveDay15() {
	fmt.Println(solvePartOne())
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

	res := 0
	for i, row := range m {
		for j, col := range row {
			if col == "O" {
				res += calcGPS(i, j)
			}
		}
	}
	return res
}

func calcGPS(x, y int) int {
	return 100*x + y
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

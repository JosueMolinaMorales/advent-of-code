package fourteen

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
)

const X_BOUND, Y_BOUND = 103, 101

type Robot struct {
	Position types.Vector
	Velocity types.Vector
}

func SolveDay14() {
	fmt.Println(solvePartOne())
	fmt.Println(solvePartTwo())
}

func setup() []Robot {
	input, err := util.LoadFileAsString("./inputs/day_14.txt")
	if err != nil {
		panic(err)
	}

	re, err := regexp.Compile(`-?\d+`)
	if err != nil {
		panic(err)
	}

	robots := []Robot{}
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}
		matches := re.FindAllString(line, -1)
		robots = append(robots, Robot{
			Position: *types.NewVector(util.ToInt(matches[1]), util.ToInt(matches[0])),
			Velocity: *types.NewVector(util.ToInt(matches[3]), util.ToInt(matches[2])),
		})
	}

	return robots
}

func moveRobots(robots *[]Robot) {
	for i, r := range *robots {
		// Move each robot
		dx, dy := mod((r.Position.X+r.Velocity.X), X_BOUND), mod((r.Position.Y+r.Velocity.Y), Y_BOUND)
		(*robots)[i].Position.X = dx
		(*robots)[i].Position.Y = dy
	}
}

func calcSafetyScore(robots []Robot) int {
	// Split into quads
	// Find middle going horizontal
	xMid := X_BOUND / 2
	// Find middle going vertical
	yMid := Y_BOUND / 2

	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, r := range robots {
		x, y := r.Position.X, r.Position.Y
		// top left check
		if x >= 0 && x < xMid && y >= 0 && y < yMid {
			q1++
		} else if x >= 0 && x < xMid && y > yMid && y < Y_BOUND {
			// top right
			q2++
		} else if x > xMid && x < X_BOUND && y >= 0 && y < yMid {
			// bottom left
			q3++
		} else if x > xMid && x < X_BOUND && y > yMid && y < Y_BOUND {
			q4++
		}
	}
	return q1 * q2 * q3 * q4
}

func solvePartOne() int {
	robots := setup()
	seconds := 0

	for seconds < 100 {
		moveRobots(&robots)
		seconds++
	}

	return calcSafetyScore(robots)
}

func solvePartTwo() int {
	// Find where the christmas tree is
	robots := setup()
	seconds := 0

	minEntropy := math.MaxInt
	t := 0
	for seconds < 10_000 {
		moveRobots(&robots)

		entropy := calcSafetyScore(robots)
		if entropy < minEntropy {
			minEntropy = entropy
			t = seconds + 1
		}
		seconds++
	}
	return t
}

func mod(a, b int) int {
	return (a%b + b) % b
}

package twentyfour

import (
	"fmt"
	"os"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils"
	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

type Hailstone struct {
	x, y, z    float64
	vx, vy, vz float64
}

func (h Hailstone) getCoefficient() (float64, float64, float64) {
	a := h.vy
	b := -h.vx
	c := h.vy*h.x - h.vx*h.y
	return a, b, c
}

func RunDayTwentyFour() {
	input, err := os.ReadFile("./input/day24.txt")
	if err != nil {
		panic("Failed to read day 24 input file")
	}
	fmt.Println("Day 24 Part 1:", partOne(string(input), Limits{200000000000000, 400000000000000}))
	// Part 2 done in python to use sympy
}

type Limits = [2]float64

func parse(input string) []Hailstone {
	hails := make([]Hailstone, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " @ ")
		positions := strings.Split(parts[0], ",")
		velocities := strings.Split(parts[1], ",")
		hails = append(hails, Hailstone{
			x:  utils.ToFloat(strings.TrimSpace(positions[0])),
			y:  utils.ToFloat(strings.TrimSpace(positions[1])),
			z:  utils.ToFloat(strings.TrimSpace(positions[2])),
			vx: utils.ToFloat(strings.TrimSpace(velocities[0])),
			vy: utils.ToFloat(strings.TrimSpace(velocities[1])),
			vz: utils.ToFloat(strings.TrimSpace(velocities[2])),
		})
	}
	return hails
}

func partOne(input string, limits Limits) int {
	hails := parse(input)
	// Looking forward in time, how many of the hailstones' paths will intersect within a test area?
	total := 0
	for i, p1 := range hails {
		for _, p2 := range hails[:i] {
			a1, b1, c1 := p1.getCoefficient()
			a2, b2, c2 := p2.getCoefficient()

			if a1*b2 == a2*b1 {
				// Parallel lines
				continue
			}

			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)

			if x >= limits[0] && x <= limits[1] && y >= limits[0] && y <= limits[1] {
				if iterators.Every([]Hailstone{p1, p2}, func(h Hailstone) bool {
					return (x-h.x)*h.vx >= 0 && (y-h.y)*h.vy >= 0
				}) {
					total++
				}
			}
		}
	}
	return total
}

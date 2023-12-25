package twentyfour

import (
	"fmt"
	"os"
	"sort"
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
	fmt.Println("Day 24 Part 2:", partTwo(string(input)))
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

func getRockVelocity(velocities map[int][]int) int {
	possibleV := make([]int, 0)
	for x := -1000; x <= 1000; x++ {
		possibleV = append(possibleV, x)
	}

	for vel, values := range velocities {
		if len(values) < 2 {
			continue
		}

		npV := make([]int, 0)
		for _, poss := range possibleV {
			// Add a check to ensure that the denominator is not zero
			if poss-vel != 0 && (values[0]-values[1])%(poss-vel) == 0 {
				npV = append(npV, poss)
			}
		}

		possibleV = npV
	}

	return possibleV[0]
}

func partTwo(input string) int {
	// https://github.com/ayoubzulfiqar/advent-of-code/blob/main/Go/Day24/part_2.go
	pvx := make([]int, 2001)
	for x := -1000; x <= 1000; x++ {
		pvx[x+1000] = x
	}

	velX := make(map[int][]int)
	velY := make(map[int][]int)
	velZ := make(map[int][]int)

	hails := parse(input)
	for _, h := range hails {
		velX[int(h.vx)] = append(velX[int(h.vx)], int(h.x))
		velY[int(h.vy)] = append(velY[int(h.vy)], int(h.y))
		velZ[int(h.vz)] = append(velZ[int(h.vz)], int(h.z))
	}

	rvx := getRockVelocity(velX)
	rvy := getRockVelocity(velY)
	rvz := getRockVelocity(velZ)

	results := make(map[int]int, 0)

	for i, p1 := range hails {
		for _, p2 := range hails[:i] {

			ma := (p1.vy - float64(rvy)) / (p1.vx - float64(rvx))
			mb := (p2.vy - float64(rvy)) / (p2.vx - float64(rvx))

			ca := p1.y - ma*p1.x
			cb := p2.y - mb*p2.x

			rpx := (cb - ca) / (ma - mb)
			rpy := ma*float64(rpx) + ca

			time := (rpx - p1.x) / (p1.vx - float64(rvx))
			rpz := p1.z + (p1.vz-float64(rvz))*time

			result := int(rpx) + int(rpy) + int(rpz)
			if _, ok := results[result]; !ok {
				results[result] = 1
			} else {
				results[result] += 1
			}
		}
	}

	var keys []int
	for k := range results {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		return results[keys[i]] > results[keys[j]]
	})
	return keys[0]
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

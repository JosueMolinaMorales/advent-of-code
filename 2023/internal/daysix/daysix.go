package daysix

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const testInput string = `Time:      7  15   30
Distance:  9  40  200`

type TimeDistance struct {
	Time     int
	Distance int
}

func RunDaySix() {
	input, err := os.ReadFile("./input/day6input.txt")
	if err != nil {
		panic("Failed to read file for day 6")
	}

	res := partOne(string(input))
	fmt.Println("Part 1:", res)
	res = partTwo(string(input))
	fmt.Println("Part 2:", res)
}

func partOne(input string) int {
	times := make([]int, 0)
	dists := make([]int, 0)
	for i, line := range strings.Split(input, "\n") {
		numsSplit := strings.Split(line, " ")
		for _, numStr := range numsSplit {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			if i == 0 {
				times = append(times, num)
			} else {
				dists = append(dists, num)
			}
		}
	}
	timeDists := make([]TimeDistance, 0)
	for i := 0; i < len(times); i++ {
		timeDists = append(timeDists, TimeDistance{
			Time:     times[i],
			Distance: dists[i],
		})
	}

	ans := 1
	for _, td := range timeDists {
		half := int(math.Floor(float64(td.Time) / 2.0))
		count := 0

		time := half * (td.Time - half)
		for time > td.Distance {
			count += 1
			half -= 1
			time = half * (td.Time - half)
		}

		half = int(math.Floor(float64(td.Time)/2.0)) + 1
		time = half * (td.Time - half)
		for time > td.Distance {
			count += 1
			half += 1
			time = half * (td.Time - half)
		}

		ans = ans * count
	}

	return ans
}

func partTwo(input string) int {
	lines := strings.Split(input, "\n")
	timeLine := strings.ReplaceAll(strings.ReplaceAll(lines[0], "Time: ", ""), " ", "")

	time, err := strconv.Atoi(timeLine)
	if err != nil {
		fmt.Println("Failed to convert", timeLine, "to int")
	}

	distStr := strings.ReplaceAll(strings.ReplaceAll(lines[1], "Distance: ", ""), " ", "")
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		fmt.Println("Failed to convert", distStr, "to int")
	}

	count := 0
	half := int(math.Floor(float64(time) / 2.0))
	t := half * (time - half)
	for t > dist {
		count += 1
		half -= 1
		t = half * (time - half)
	}

	half = int(math.Floor(float64(time)/2.0)) + 1
	t = half * (time - half)
	for t > dist {
		count += 1
		half += 1
		t = half * (time - half)
	}

	return count
}

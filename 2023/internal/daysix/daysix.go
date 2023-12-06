package daysix

import (
	"fmt"
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
	partOne(testInput)
}

func partOne(input string) {
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
		half := td.Time / 2
		count := 0

		time := half * (td.Time - half)
		for time > td.Distance {
			count += 1
			half -= 1
			time = half * (td.Time - half)
		}

		ans *= count
	}
	fmt.Println(ans)
}

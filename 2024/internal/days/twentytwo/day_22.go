package twentytwo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay22() {
	fmt.Println("Day 22 Part 1: ", solvePartOne())
	fmt.Println("Day 22 Part 2: ", solvePartTwo())
}

func setup() []int {
	input, err := util.LoadFileAsString("./inputs/day_22.txt")
	if err != nil {
		panic(err)
	}

	secrets := []int{}
	for _, s := range strings.Split(input, "\n") {
		secrets = append(secrets, util.ToInt(s))
	}

	return secrets
}

func solvePartOne() int {
	secrets, _ := genSecrets(setup())
	sum := 0
	for _, secret := range secrets {
		sum += secret
	}
	return sum
}

func solvePartTwo() int {
	_, prices := genSecrets(setup())
	windows := map[string]int{}
	for k := 0; k < len(prices); k++ {
		v := prices[k]
		inPrices := map[string]int{}
		for i := 0; i < len(v)-4; i++ {
			window := []string{}
			for w := 0; w < 4; w++ {
				d := v[i+w+1] - v[i+w]
				window = append(window, strconv.Itoa(d))
			}
			key := strings.Join(window, "")
			// Store the price for the window if
			// - we haven't seen it before
			// - if its greater than the previously stored one but only if
			// 		its the first secret number
			if _, ok := inPrices[key]; !ok || v[i+4] > inPrices[key] && k == 0 {
				inPrices[key] = v[i+4]
			}
		}
		for k, v := range inPrices {
			windows[k] += v
		}
	}
	maxVal := 0
	for _, v := range windows {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func genSecrets(secrets []int) ([]int, map[int][]int) {
	iterations := 2000
	prices := map[int][]int{}
	for i := 0; i < iterations; i++ {
		for j := 0; j < len(secrets); j++ {
			n := secrets[j]
			if prices[j] == nil {
				prices[j] = []int{n % 10}
			}
			n = prune(mix(n*64, n))
			n = prune(mix(n/32, n))
			n = prune(mix(n*2048, n))
			secrets[j] = n
			// Store the price
			prices[j] = append(prices[j], n%10)
		}
	}

	return secrets, prices
}

func mix(n, secret int) int {
	return n ^ secret
}

func prune(n int) int {
	return n % 16777216
}

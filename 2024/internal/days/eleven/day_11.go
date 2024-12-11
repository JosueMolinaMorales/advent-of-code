package eleven

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
)

func SolveDay11() {
	fmt.Println(solvePartOne())
}

func rockMovement(stones []int, movements int, cache_0 map[int]int, cache_1 map[int]int) ([]int, int) {
	count := 0
	// fmt.Printf("stones ---> %v\n", stones)
	for i := 0; i < movements; i++ {
		n := len(stones)
		splits := 0
		for j := 0; j < n+splits; j++ {
			stone := stones[j]
			// fmt.Printf("stone is: %d at %d\n", stone, i+1)
			s := strconv.Itoa(stone)
			// if stone is == 0 then -> 1
			if (stone == 0 || stone == 1) && (cache_0 != nil && cache_1 != nil) {
				// Remove
				stones = append(stones[:j], stones[j+1:]...)
				// fmt.Printf("after removal: %v\n", stones)
				splits--
				j--
				if stone == 0 {
					// fmt.Println("DIGGING THOUGH CACHE 0 FOUND: ", cache_0[movements-i-1])
					count += cache_0[movements-i-1]
				}
				if stone == 1 {
					// fmt.Println("DIGGING THOUGH CACHE 1 FOUND: ", cache_1[movements-i-1])
					count += cache_1[movements-i-1]
				}

				continue
			}
			if stone == 0 {
				stones[j] = 1
			} else if stone >= 10 && len(s)%2 == 0 {
				// fmt.Printf("Stone is to be split: %d\n", stone)
				// if stone is >= 10 and even split into two
				s1 := util.ToInt(s[:len(s)/2])
				// fmt.Printf("	first have of stone is %d\n", s1)
				s2 := util.ToInt(s[len(s)/2:])
				// fmt.Printf("	second have of stone is %d\n", s2)
				stones = append(stones[:j], append([]int{s1, s2}, stones[j+1:]...)...)
				j += 1
				splits++
				continue
			} else {
				// otherwise stone * 2024
				// fmt.Println("Stone to be multiplied by 2024")
				stones[j] = stone * 2024
			}

		}

		// fmt.Printf("After blink %d: %v\n", i+1, stones)
		// fmt.Printf("After blink %d: %d\n", i+1, len(stones))
		// fmt.Println(len(stones))
	}
	// fmt.Printf("AFTER EVERYTHING: COUNT: %d - LEN(): %d\n", count, len(stones))
	return stones, len(stones) + count
}

func solvePartOne() int {
	// rawStones, err := util.LoadFileAsString("./inputs/day_11.txt")
	// if err != nil {
	// 	panic(err)
	// }
	rawStones := "773 79858 0 71 213357 2937 1 3998391"
	stones := []int{}
	for _, s := range strings.Split(rawStones, " ") {
		stones = append(stones, util.ToInt(s))
	}

	count := 0
	memo := map[types.Vector]int{}
	for _, num := range stones {
		count += solve(num, 75, memo)
	}

	// cache_0 := map[int]int{}
	// s_0 := []int{0}
	// for i := 0; i < 75; i++ {
	// 	s_0, cache_0[i] = rockMovement(s_0, 1, nil, nil)
	// }
	// // fmt.Println(cache_0)
	// fmt.Println("Part 1 done")
	// cache_1 := map[int]int{}
	// s_1 := []int{1}
	// for i := 0; i < 75; i++ {
	// 	s_1, cache_1[i] = rockMovement(s_1, 1, nil, nil)
	// }
	// // fmt.Println(cache_1)
	// fmt.Println("part 2 done")
	// _, res := rockMovement(stones, 75, cache_0, cache_1)

	return count
}

func solve(num int, blinks int, memo map[types.Vector]int) int {
	if blinks == 0 {
		return 1
	}

	if cached := memo[*types.NewVector(num, blinks)]; cached != 0 {
		return cached
	}

	count := 0
	for _, stone := range blink(num) {
		count += solve(stone, blinks-1, memo)
	}
	memo[*types.NewVector(num, blinks)] = count

	return count
}

func blink(num int) []int {
	if num == 0 {
		return []int{1}
	}

	s := strconv.Itoa(num)
	if len(s)%2 == 0 {
		// fmt.Printf("Stone is to be split: %d\n", stone)
		// if stone is >= 10 and even split into two
		s1 := util.ToInt(s[:len(s)/2])
		// fmt.Printf("	first have of stone is %d\n", s1)
		s2 := util.ToInt(s[len(s)/2:])

		return []int{s1, s2}
	}

	return []int{num * 2024}
}

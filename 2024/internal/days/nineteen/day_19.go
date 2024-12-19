package nineteen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay19() {
	fmt.Println(solvePartOne())
}

// TOO HIGH: 327
func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_19.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")
	towels := strings.Split(parts[0], ", ")
	designs := strings.Split(parts[1], "\n")

	// fmt.Println(towels)
	// fmt.Println(designs)

	count := 0
	for _, design := range designs {
		indices := hashset.New()
		fmt.Printf("Looking at design: %s\n", design)
		// c := map[string]int{}
		// for _, ch := range design {
		// 	c[string(ch)]++
		// }
		// d := map[string]int{}
		for _, towel := range towels {
			fmt.Printf(" looking at towel: %s\n", towel)
			re, err := regexp.Compile(fmt.Sprintf("(%s)+", towel))
			if err != nil {
				panic(err)
			}

			idxs := re.FindAllIndex([]byte(design), -1)
			fmt.Println("  idxs: ", idxs)
			for _, ra := range idxs {
				for i := ra[0]; i < ra[1]; i++ {
					fmt.Println("  Adding: ", i)
					indices.Add(i)
				}
			}
		}
		if indices.Size() == len(design) {
			fmt.Println("MATCH!")
			count++
		}
		fmt.Println()
	}
	return count
}

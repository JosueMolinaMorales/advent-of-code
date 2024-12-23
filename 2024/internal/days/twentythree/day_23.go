package twentythree

import (
	"fmt"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/emirpasic/gods/sets/hashset"
)

func SolveDay23() {
	fmt.Println(solvePartOne())
}

func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_23.txt")
	if err != nil {
		panic(err)
	}
	netMap := strings.Split(input, "\n")

	// K: PC V: Connected PCs
	connections := map[string]*hashset.Set{}
	for _, conn := range netMap {
		pcs := strings.Split(conn, "-")

		if connections[pcs[0]] == nil {
			connections[pcs[0]] = hashset.New()
		}
		connections[pcs[0]].Add(pcs[1])
		if connections[pcs[1]] == nil {
			connections[pcs[1]] = hashset.New()
		}
		connections[pcs[1]].Add(pcs[0])
	}

	// For every pc, see if it 3 other pcs connections
	for k, v := range connections {
		fmt.Println(k, v.Values())
		// c := []string{}
		// for o, v := range connections {
		// 	if v.Contains(k) {
		// 		c = append(c, o)
		// 	}
		// }
		// if len(c) >= 3 {
		// 	fmt.Printf("%s: %v\n", k, c)
		// }
	}
	return 0
}

package nine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay9() {
	fmt.Println(solvePartOne())
	fmt.Println(solvePartTwo())
}

func setup() []string {
	rawDiskMap, err := util.LoadFileAsString("./inputs/day_9.txt")
	if err != nil {
		panic(err)
	}

	diskmap := make([]int, 0)
	for _, chr := range strings.Split(rawDiskMap, "") {
		diskmap = append(diskmap, util.ToInt(chr))
	}

	space := make([]string, 0)
	id := 0
	for i, ch := range diskmap {
		if (i % 2) == 0 {
			// Even is files
			for j := 0; j < ch; j++ {
				space = append(space, strconv.Itoa(id))
			}
			id++
		} else {
			// Odd is free space
			for j := 0; j < ch; j++ {
				space = append(space, ".")
			}
		}
	}

	return space
}

func solvePartOne() int {
	space := setup()
	// Two pointers, one pointing to free space, the next pointing to the right most
	// file to be moved
	freeSpacePtr, filePtr := 0, len(space)-1

	for freeSpacePtr < filePtr {
		// Check if the freePtr is on a free space
		if space[freeSpacePtr] != "." {
			freeSpacePtr++
			continue
		}
		if space[filePtr] == "." {
			filePtr--
			continue
		}
		// free ptr on a free space
		// now swap
		space[freeSpacePtr] = space[filePtr]
		space[filePtr] = "."

		// Increment/decrement
		freeSpacePtr++
		filePtr--
	}

	return calcChecksum(space)
}

func solvePartTwo() int {
	space := setup()
	freeSpacePtr, filePtr := 0, len(space)-1
	for filePtr > 0 {
		// Check if the freePtr is on a free space
		if space[freeSpacePtr] != "." {
			freeSpacePtr++
			continue
		}
		if space[filePtr] == "." {
			filePtr--
			continue
		}

		// Get window size of file ptr
		size := 0
		j := filePtr
		for j > 0 && space[j] == space[filePtr] {
			size += 1
			j--
		}

		// Find starting position of free space window
		canFit := false
		for freeSpacePtr < filePtr {
			windowSize := 0
			j := freeSpacePtr
			for space[j] == "." {
				windowSize++
				j++
			}
			if windowSize >= size {
				canFit = true
				break
			}
			freeSpacePtr++
		}

		if !canFit {
			freeSpacePtr = 0
			filePtr = filePtr - size
			continue
		}
		// Swap
		for size > 0 {
			space[freeSpacePtr] = space[filePtr]
			space[filePtr] = "."
			freeSpacePtr++
			size--
			filePtr--
		}
		freeSpacePtr = 0
	}
	return calcChecksum(space)
}

func calcChecksum(space []string) int {
	checksum := 0
	for i, n := range space {
		if n == "." {
			continue
		}
		checksum += i * util.ToInt(n)
	}

	return checksum
}

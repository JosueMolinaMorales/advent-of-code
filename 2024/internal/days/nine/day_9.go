package nine

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
)

func SolveDay9() {
	fmt.Println("Day 9 Part 1: ", solvePartOne())
	fmt.Println("Day 9 Part 1: ", solvePartTwo())
}

// setup initializes the disk space representation based on the input file
func setup() []string {
	rawDiskMap, err := util.LoadFileAsString("./inputs/day_9.txt")
	if err != nil {
		panic(err)
	}

	space := []string{}
	id := 0

	for i, chr := range strings.Split(rawDiskMap, "") {
		count := util.ToInt(chr)

		if i%2 == 0 { // Even indices represent files
			for j := 0; j < count; j++ {
				space = append(space, strconv.Itoa(id))
			}
			id++
		} else { // Odd indices represent free space
			for j := 0; j < count; j++ {
				space = append(space, ".")
			}
		}
	}

	return space
}

// solvePartOne solves the first part of the puzzle
func solvePartOne() int {
	space := setup()
	freeSpacePtr, filePtr := 0, len(space)-1

	for freeSpacePtr < filePtr {
		// Move pointers to valid positions
		for freeSpacePtr < filePtr && space[freeSpacePtr] != "." {
			freeSpacePtr++
		}
		for freeSpacePtr < filePtr && space[filePtr] == "." {
			filePtr--
		}

		// Swap if valid positions are found
		if freeSpacePtr < filePtr {
			space[freeSpacePtr], space[filePtr] = space[filePtr], "."
			freeSpacePtr++
			filePtr--
		}
	}

	return calcChecksum(space)
}

// solvePartTwo solves the second part of the puzzle
func solvePartTwo() int {
	space := setup()
	freeSpacePtr, filePtr := 0, len(space)-1

	for filePtr > 0 {
		// Move pointers to valid positions
		for freeSpacePtr < filePtr && space[freeSpacePtr] != "." {
			freeSpacePtr++
		}
		for filePtr > 0 && space[filePtr] == "." {
			filePtr--
		}

		// Determine the size of the file block
		size := 0
		currFilePtr := filePtr
		for currFilePtr > 0 && space[currFilePtr] == space[filePtr] {
			size++
			currFilePtr--
		}

		// Check if a free space window can fit the file block
		canFit := false
		for freeSpacePtr < filePtr {
			windowSize := 0
			currFreePtr := freeSpacePtr
			for currFreePtr < len(space) && space[currFreePtr] == "." {
				windowSize++
				currFreePtr++
			}
			if windowSize >= size {
				canFit = true
				break
			}
			freeSpacePtr++
		}

		// If no fitting space, move the file pointer back
		if !canFit {
			freeSpacePtr = 0
			filePtr -= size
			continue
		}

		// Move the file block into the free space window
		for size > 0 {
			space[freeSpacePtr] = space[filePtr]
			space[filePtr] = "."
			freeSpacePtr++
			filePtr--
			size--
		}

		// Reset the free space pointer for the next iteration
		freeSpacePtr = 0
	}

	return calcChecksum(space)
}

// calcChecksum calculates the checksum for the final state of the space
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

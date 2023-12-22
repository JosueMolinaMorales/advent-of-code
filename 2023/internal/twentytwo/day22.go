package twentytwo

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils"
	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `1,0,1~1,2,1
0,0,2~2,0,2
0,2,3~2,2,3
0,0,4~0,2,4
2,0,5~2,2,5
0,1,6~2,1,6
1,1,8~1,1,9`

func RunDayTwentyTwo() {
	input, err := os.ReadFile("./input/day22.txt")
	if err != nil {
		panic("Failed to read day 22 input")
	}
	fmt.Println("Part one:", partOne(string(input)))
	fmt.Println("Part two:", partTwo(string(input)))
}

type (
	Position = [3]int // x,y,z
	Brick    struct {
		posA        Position
		posB        Position
		supports    []int
		supportedBy []int
	}
)

func parse(input string) []Brick {
	bricks := make([]Brick, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, "~")
		s := Brick{}
		pa := strings.Split(parts[0], ",")
		pos, _ := strconv.Atoi(pa[0])
		s.posA[0] = pos
		pos, _ = strconv.Atoi(pa[1])
		s.posA[1] = pos
		pos, _ = strconv.Atoi(pa[2])
		s.posA[2] = pos

		pb := strings.Split(parts[1], ",")
		pos, _ = strconv.Atoi(pb[0])
		s.posB[0] = pos
		pos, _ = strconv.Atoi(pb[1])
		s.posB[1] = pos
		pos, _ = strconv.Atoi(pb[2])
		s.posB[2] = pos

		bricks = append(bricks, s)
	}
	// Sort the Bricks by their z value
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].posA[2] < bricks[j].posA[2]
	})

	return bricks
}

func partOne(input string) int {
	bricks := parse(input)
	placedBlocks := placeBlocks(bricks)
	// Filter out the bricks that are supported by more than one brick
	filtered := iterators.Filter(placedBlocks, func(b Brick) bool {
		return iterators.Every(b.supports, func(idx int) bool {
			return len(placedBlocks[idx].supportedBy) > 1
		})
	})
	return len(filtered)
}

func partTwo(input string) int {
	bricks := parse(input)
	placedBlocks := placeBlocks(bricks)

	// For every brick, if the brick is removed, how many bricks will fall?
	// If the brick is removed, the bricks it supports will fall, and the bricks those bricks support will fall, and so on
	fallen := 0
	for i := range iterators.Reverse(placedBlocks) {
		if len(placedBlocks[i].supports) == 0 {
			continue
		}
		fallen += fallenBricks(i, placedBlocks)
	}

	return fallen
}

func fallenBricks(brickIdx int, bricks []Brick) int {
	fallen := make([]int, 0)
	type Item struct {
		idx int
		z   int
	}

	pq := utils.NewHeap(make([]interface{}, 0), func(a, b interface{}) bool {
		return a.(Item).z < b.(Item).z
	})

	heap.Init(&pq)
	heap.Push(&pq, Item{brickIdx, bricks[brickIdx].posB[2]})

	for pq.Len() > 0 {
		b := heap.Pop(&pq).(Item)
		fallen = append(fallen, b.idx)

		// Enqueue the bricks that this brick supports
		for _, idx := range bricks[b.idx].supports {
			sb := bricks[idx]
			if iterators.Every(sb.supportedBy, func(idx int) bool {
				return iterators.Contains(fallen, idx)
			}) {
				pq.Push(Item{idx, sb.posB[2]})
			}
		}
	}
	return len(fallen) - 1
}

func placeBlocks(bricks []Brick) []Brick {
	// Find the highest z value
	highestZ := 1
	placedBlocks := make([]Brick, 0)
	for _, brick := range bricks {
		// Position the brick just above the highest z
		diff := brick.posB[2] - brick.posA[2]
		brick.posA[2] = highestZ + 1
		brick.posB[2] = brick.posA[2] + diff

		for {
			canMoveDown := true
			brick = moveDown(brick)
			for i := len(placedBlocks) - 1; i >= 0; i-- {
				if placedBlocks[i].posB[2] < brick.posA[2] {
					continue
				}

				if hasCollision(&brick, &placedBlocks[i]) {
					canMoveDown = false
					lenPlacedBlocks := len(placedBlocks)
					placedBlocks[i].supports = append(placedBlocks[i].supports, lenPlacedBlocks)
					brick.supportedBy = append(brick.supportedBy, i)
				}
			}

			if !canMoveDown {
				brick = moveUp(brick)
				break
			}

			// reached the bottom
			if brick.posA[2] == 1 {
				break
			}
		}
		highestZ = int(math.Max(float64(highestZ), float64(brick.posB[2])))

		placedBlocks = append(placedBlocks, brick)
	}

	return placedBlocks
}

func hasCollision(a, b *Brick) bool {
	return (pointInRange(a.posA[2], [2]int{b.posA[2], b.posB[2]}) ||
		pointInRange(a.posB[2], [2]int{b.posA[2], b.posB[2]}) ||
		pointInRange(b.posA[2], [2]int{a.posA[2], a.posB[2]}) ||
		pointInRange(b.posB[2], [2]int{a.posA[2], a.posB[2]})) &&
		(pointInRange(a.posA[0], [2]int{b.posA[0], b.posB[0]}) ||
			pointInRange(a.posB[0], [2]int{b.posA[0], b.posB[0]}) ||
			pointInRange(b.posA[0], [2]int{a.posA[0], a.posB[0]}) ||
			pointInRange(b.posB[0], [2]int{a.posA[0], a.posB[0]})) &&
		(pointInRange(a.posA[1], [2]int{b.posA[1], b.posB[1]}) ||
			pointInRange(a.posB[1], [2]int{b.posA[1], b.posB[1]}) ||
			pointInRange(b.posA[1], [2]int{a.posA[1], a.posB[1]}) ||
			pointInRange(b.posB[1], [2]int{a.posA[1], a.posB[1]}))
}

func moveUp(b Brick) Brick {
	b.posA[2] += 1
	b.posB[2] += 1
	return b
}

func moveDown(b Brick) Brick {
	if b.posA[2] > 1 {
		b.posA[2] -= 1
		b.posB[2] -= 1
	}
	return b
}

func pointInRange(p int, r [2]int) bool {
	return p >= r[0] && p <= r[1]
}

package twenty

import (
	"fmt"
	"strings"
	"time"

	"github.com/JosueMolinaMorales/aoc/2024/internal/util"
	"github.com/JosueMolinaMorales/aoc/2024/internal/util/types"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/utils"
)

func SolveDayTwenty() {
	fmt.Println(solvePartOne())
}

func solvePartOne() int {
	input, err := util.LoadFileAsString("./inputs/day_20.txt")
	if err != nil {
		panic(err)
	}

	racetrack := [][]string{}
	start := *types.NewVector()
	end := *types.NewVector()
	for i, line := range strings.Split(input, "\n") {
		row := []string{}
		for j, ch := range strings.Split(line, "") {
			v := *types.NewVector(i, j)
			if ch == "S" {
				start = v
			} else if ch == "E" {
				end = v
			}
			row = append(row, ch)
		}
		racetrack = append(racetrack, row)
	}

	fmt.Println(start)
	fmt.Println(end)

	startTime := time.Now()
	ogTime := regPathFind(start, end, racetrack)
	fmt.Println("Time to find end path: ", time.Since(startTime))
	fmt.Println("WITH og time: ", ogTime)

	time := raceToEnd(start, end, racetrack)
	fmt.Println(time)
	return time
	// return 0
}

type State struct {
	Time       int
	Pos        types.Vector
	HasCheated bool
	Cheat      types.Vector
	Visited    hashset.Set
}

func regPathFind(start, end types.Vector, grid [][]string) int {
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		aTime := a.(State).Time
		bTime := b.(State).Time
		return utils.IntComparator(aTime, bTime)
	})
	q.Enqueue(State{
		Time:       0,
		Pos:        start,
		HasCheated: false,
		Cheat:      types.Vector{},
	})
	visited := hashset.New(start)
	for !q.Empty() {
		p, _ := q.Dequeue()
		state := p.(State)
		if state.Pos == end {
			return state.Time
		}

		neighbors, _ := adjacent(state, grid, &hashset.Set{})
		for _, neighbor := range neighbors {
			if visited.Contains(neighbor) {
				continue
			}
			visited.Add(neighbor)
			q.Enqueue(State{
				Time:       state.Time + 1,
				Pos:        neighbor,
				Cheat:      state.Cheat,
				HasCheated: state.HasCheated,
			})
		}

		// fmt.Println()
	}
	return -1
}

func minPath(start, end types.Vector, enableCheats bool, grid [][]string, seenCheats *hashset.Set) State {
	q := priorityqueue.NewWith(func(a, b interface{}) int {
		aTime := a.(State).Time
		bTime := b.(State).Time
		return utils.IntComparator(aTime, bTime)
	})
	q.Enqueue(State{
		Time:       0,
		Pos:        start,
		HasCheated: false,
		Cheat:      types.Vector{},
		Visited:    *hashset.New(start),
	})
	endState := State{}
	visited := hashset.New(start)
	for !q.Empty() {
		p, _ := q.Dequeue()
		state := p.(State)
		if state.Pos == end {
			return state
		}

		neighbors, cheats := adjacent(state, grid, seenCheats)
		for _, neighbor := range neighbors {
			if visited.Contains(neighbor) {
				continue
			}
			visited.Add(neighbor)
			q.Enqueue(State{
				Time:       state.Time + 1,
				Pos:        neighbor,
				Cheat:      state.Cheat,
				HasCheated: state.HasCheated,
			})
		}
		if enableCheats {
			for _, cheat := range cheats {
				if visited.Contains(cheat) {
					continue
				}
				visited.Add(cheat)

				q.Enqueue(State{
					Time:       state.Time + 1,
					Pos:        cheat,
					HasCheated: true,
					Cheat:      cheat,
				})
			}
		}
		// fmt.Println()
	}
	return endState
}

func raceToEnd(start, end types.Vector, grid [][]string) int {
	seenCheats := hashset.New()
	saves := map[int]int{}
	ogTime := minPath(start, end, false, grid, seenCheats).Time
	fmt.Println("OG TIME: ", ogTime)
	for i := 0; i < 10_000; i++ {
		endState := minPath(start, end, true, grid, seenCheats)

		saves[ogTime-endState.Time]++
		seenCheats.Add(endState.Cheat)

		// fmt.Println("Cheat was: ", endState.Cheat)
	}

	// How many cheats save at least 100
	fmt.Println(saves)
	// Print path

	// printMap(grid, []types.Vector{})
	count := 0
	for k, v := range saves {
		if k >= 100 {
			count += v
		}
	}
	return count
}

func adjacent(state State, grid [][]string, seenCheats *hashset.Set) ([]types.Vector, []types.Vector) {
	directions := []types.Vector{
		{X: 0, Y: 1}, {X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: -1},
	}

	neighbors := []types.Vector{}
	cheats := []types.Vector{}
	point := state.Pos
	for _, dir := range directions {
		x, y := point.X+dir.X, point.Y+dir.Y
		n := *types.NewVector(x, y)
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			continue
		}
		if grid[x][y] != "#" {
			neighbors = append(neighbors, n)
		}
		if grid[x][y] == "#" && !state.HasCheated && !seenCheats.Contains(n) {
			cheats = append(cheats, n)
		}
	}
	return neighbors, cheats
}

func printMap(m [][]string, path []types.Vector) {
	// Add path
	for _, cell := range path {
		m[cell.X][cell.Y] = "O"
	}
	for _, row := range m {
		for _, col := range row {
			fmt.Print(col)
		}
		fmt.Println()
	}
}

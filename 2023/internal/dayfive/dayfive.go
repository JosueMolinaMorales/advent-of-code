package dayfive

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

const testInput string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

/**
50 98 2 --> 50 is source start, 98 is destination start, 2 is range length
50, 51 maps to 98, 99
*/

type MappingRange struct {
	SourceStart      int
	SourceEnd        int
	DestinationStart int
	DestinationEnd   int
	Step             int
}

type FoodMap struct {
	Maps []MappingRange
}

type SeedRange struct {
	Start int
	End   int
}

func RunDayFive() {
	// input, err := os.ReadFile("./input/day5input.txt")
	// if err != nil {
	// 	panic("Could not read day 5 input file")
	// }
	// partOne(string(input))
	partTwo(testInput)
}

func partOne(input string) {
	foodMaps, seeds := parseInput(input)
	locations := make([]int, 0)

	fmt.Println(seeds)
	for _, seed := range seeds {
		for _, fm := range foodMaps {
			// fmt.Println("Map", i+1)
			for _, m := range fm.Maps {
				if seed >= m.SourceStart && seed <= m.SourceEnd {
					// New mapping --> seed + (ds - ss)
					diff := -(m.SourceStart - m.DestinationStart)

					seed = seed + diff
					break
				}
			}
			// fmt.Println("Seed: ", seed)
		}
		locations = append(locations, seed)
		// fmt.Println(locations)
	}

	min := slices.Min(locations)
	fmt.Println(min)
}

func partTwo(input string) {
	foodMaps, seeds := parseInput(input)
	// locations := make([]int, 0)
	var seedRanges []SeedRange
	i := 0
	for i < len(seeds)-1 {
		seedRanges = append(seedRanges, SeedRange{
			Start: seeds[i],
			End:   seeds[i] + seeds[i+1] - 1,
		})
		i += 2
	}

	// minSeedRanges := []SeedRange
	i = 0
	// minLocation := math.MaxInt64

	r := seedRanges[i]
	seedCount := len(seedRanges)
	fmIdx := 0
	for fmIdx < len(foodMaps) {
		// fmt.Println("Map", fmIdx+1)
		mappedSeeds := make([]int, 2)

		ranges := make([]SeedRange, 0)
		for _, m := range foodMaps[fmIdx].Maps {
			diff := -(m.SourceStart - m.DestinationStart)
			// Case 1: Zero overlap
			if (r.End < m.SourceStart) || (r.Start > m.SourceEnd) {
				// fmt.Println("Case 1: Zero overlap")
				ranges = append(ranges, SeedRange{
					Start: r.Start,
					End:   r.End,
				})
			} else {
				overlapStart := int(math.Max(float64(r.Start), float64(m.SourceStart)))
				overlapEnd := int(math.Min(float64(r.End), float64(m.SourceEnd)))

				// Case 2: Complete overlap + handles overlap for cases 3-4
				// Add the mapped overlapping range to mappings
				ranges = append(ranges, SeedRange{
					Start: overlapStart + diff,
					End:   overlapEnd + diff,
				})

				// Case 3: Seed range is larger than source range
				if r.Start < m.SourceStart {
					ranges = append(ranges, SeedRange{
						Start: r.Start,
						End:   overlapStart - 1,
					})
				}
				// Case 4
				if r.End > m.SourceEnd {
					ranges = append(ranges, SeedRange{
						Start: overlapEnd + 1,
						End:   r.End,
					})
				}
			}

			// if r.Start < m.SourceStart && r.End <= m.SourceEnd {
			// 	// Case 2: Start is less than source start but end is within source range
			// 	// fmt.Println("r.Start < m.SourceStart && r.End <= m.SourceEnd")
			// 	ranges = append(ranges, SeedRange{
			// 		Start: m.SourceStart,
			// 		End:   r.End,
			// 	})
			// } else if r.Start >= m.SourceStart && r.End > m.SourceEnd {
			// 	// Case 3: Start is in source range but end is greater than source end
			// 	// fmt.Println("r.Start >= m.SourceStart && r.End > m.SourceEnd")
			// 	ranges = append(ranges, SeedRange{
			// 		Start: r.Start,
			// 		End:   m.SourceEnd,
			// 	})
			// } else if r.Start < m.SourceStart && r.End > m.SourceEnd {
			// 	// Case 4: Start and End are both inside source range
			// 	// fmt.Println("r.Start < m.SourceStart && r.End > m.SourceEnd")
			// 	ranges = append(ranges, SeedRange{
			// 		Start: m.SourceStart,
			// 		End:   m.SourceEnd,
			// 	})
			// }
		}
		fmt.Println("Ranges", ranges)

		// mapped := false
		for rIdx, seed := range ranges {
			for j, s := range []int{seed.Start, seed.End} {
				for _, m := range foodMaps[fmIdx].Maps {
					if s >= m.SourceStart && s <= m.SourceEnd {
						// New mapping --> seed + (ds - ss)
						diff := -(m.SourceStart - m.DestinationStart)
						s = s + diff
						// mapped = true
						break
					}
				}
				mappedSeeds[j] = s
			}
			ranges[rIdx] = SeedRange{
				Start: mappedSeeds[0],
				End:   mappedSeeds[1],
			}
		}
		seedCount -= 1
		// fmt.Printf("%d --> %d\n", r.Start, mappedSeeds[0])
		// fmt.Printf("%d --> %d\n", r.End, mappedSeeds[1])

		// Add all the ranges to the seedRanges
		for _, sr := range ranges {
			seedRanges = append(seedRanges, sr)
		}

		// Remove the current range from the seedRanges
		seedRanges = append(seedRanges[:i], seedRanges[i+1:]...)

		r = seedRanges[0]
		if seedCount == 0 {
			seedCount = len(seedRanges)
			fmIdx += 1
			fmt.Println("Onto Map", fmIdx+1)
			continue
		}
	}
	fmt.Println("OUt of loop")

	// Get the minimum seed range
	min := seedRanges[0].Start
	for _, sr := range seedRanges {
		if sr.Start < min {
			min = sr.Start
		}
	}

	fmt.Println(min)
}

// func getSeedMappings(seed int, foodMaps FoodMap) int {
// 	for _, m := range fm.Maps {
// 		if seed >= m.SourceStart && seed <= m.SourceEnd {
// 			// New mapping --> seed + (ds - ss)
// 			diff := -(m.SourceStart - m.DestinationStart)
// 			seed = seed + diff
// 			break
// 		}
// 	}
// 	return seed
// }

func parseInput(input string) ([]FoodMap, []int) {
	foodMap := make([]FoodMap, 0)
	splitInput := strings.Split(input, "\n\n")

	seedsSlice := strings.Split(strings.ReplaceAll(splitInput[0], "seeds: ", ""), " ")
	seeds := make([]int, 0)
	for _, seed := range seedsSlice {
		seedInt, err := strconv.Atoi(seed)
		if err != nil {
			fmt.Println("Could not convert ", seed, "to int")
		}
		seeds = append(seeds, seedInt)
	}

	splitInput = splitInput[1:]

	for _, line := range splitInput {
		fm := FoodMap{}
		for _, mapLine := range strings.Split(line, "\n")[1:] {
			mappings := strings.Split(mapLine, " ")

			ds, err := strconv.Atoi(mappings[0])
			if err != nil {
				fmt.Println("COuld not convert", mappings[0], "to int")
				continue
			}

			ss, err := strconv.Atoi(mappings[1])
			if err != nil {
				fmt.Println("Could not convert", mappings[1], "to int")
				continue
			}

			r, err := strconv.Atoi(mappings[2])
			if err != nil {
				fmt.Println("Could not convert", mappings[2], "to int")
				continue
			}

			fm.Maps = append(fm.Maps, MappingRange{
				SourceStart:      ss,
				SourceEnd:        ss + r - 1,
				DestinationStart: ds,
				DestinationEnd:   ds + r - 1,
				Step:             r,
			})
		}

		foodMap = append(foodMap, fm)
	}

	return foodMap, seeds
}

package dayfive

import (
	"fmt"
	"os"
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

func RunDayFive() {
	input, err := os.ReadFile("./input/day5input.txt")
	if err != nil {
		panic("Could not read day 5 input file")
	}
	res := partOne(string(input))
	fmt.Println("Part one result: ", res)
	res = partTwo(string(input))
	fmt.Println("Part two result: ", res)
}

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

func partOne(input string) int {
	foodMaps, seeds := parseInput(input)
	locations := make([]int, 0)

	for _, seed := range seeds {
		for _, fm := range foodMaps {
			for _, m := range fm.Maps {
				if seed >= m.SourceStart && seed <= m.SourceEnd {
					// New mapping --> seed + (ds - ss)
					diff := -(m.SourceStart - m.DestinationStart)

					seed = seed + diff
					break
				}
			}
		}
		locations = append(locations, seed)
	}

	min := slices.Min(locations)

	return min
}

func partTwo(input string) int {
	foodMaps, seeds := parseInput(input)
	var seedRanges []SeedRange
	i := 0
	for i < len(seeds)-1 {
		seedRanges = append(seedRanges, SeedRange{
			Start: seeds[i],
			End:   seeds[i] + seeds[i+1],
		})
		i += 2
	}

	for _, fm := range foodMaps {
		new_ranges := make([]SeedRange, 0)
		for _, r := range seedRanges {
			for _, m := range fm.Maps {
				offset := m.DestinationStart - m.SourceStart
				rule_applies := r.Start <= r.End && r.Start <= m.SourceEnd && r.End >= m.SourceStart

				if rule_applies {
					if r.Start < m.SourceStart {
						new_ranges = append(new_ranges, SeedRange{
							Start: r.Start,
							End:   m.SourceStart - 1,
						})
						r.Start = m.SourceStart
						if r.Start < m.SourceEnd {
							new_ranges = append(new_ranges, SeedRange{
								Start: r.Start + offset,
								End:   r.End + offset,
							})
							r.Start = r.End + 1
						} else {
							new_ranges = append(new_ranges, SeedRange{
								Start: r.Start + offset,
								End:   m.SourceEnd + offset,
							})
							r.Start = m.SourceEnd
						}
					} else if r.Start < m.SourceEnd {
						new_ranges = append(new_ranges, SeedRange{
							Start: r.Start + offset,
							End:   r.End + offset,
						})
						r.Start = r.End + 1
					} else {
						new_ranges = append(new_ranges, SeedRange{
							Start: r.Start + offset,
							End:   m.SourceEnd + offset,
						})
						r.Start = m.SourceEnd
					}
				}
			}
			if r.Start <= r.End {
				new_ranges = append(new_ranges, r)
			}
		}

		seedRanges = new_ranges
	}

	// GEt the minimum seed range
	min := seedRanges[0].Start
	for _, sr := range seedRanges {
		if sr.Start < min {
			min = sr.Start
		}
	}

	return min
}

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

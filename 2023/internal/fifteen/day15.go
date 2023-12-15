package fifteen

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
)

const testInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func RunDayFifteen() {
	input, err := os.ReadFile("./input/day15.txt")
	if err != nil {
		panic("Failed to read day 15 file")
	}
	// fmt.Println("Part one", partOne(testInput))
	// fmt.Println("Part one", partOne(string(input)))
	// fmt.Println("Part two", partTwo(testInput))
	fmt.Println("Part two", partTwo(string(input)))
}

func partOne(input string) int {
	sum := 0
	for _, line := range strings.Split(input, ",") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		sum += hash(line)
	}

	return sum
}

/*
Determine the ASCII code for the current character of the string.
Increase the current value by the ASCII code you just determined.
Set the current value to itself multiplied by 17.
Set the current value to the remainder of dividing itself by 256.
*/
func hash(input string) int {
	current := 0
	for _, char := range input {
		current += int(char)
		current *= 17
		current %= 256
	}

	return current
}

type Value struct {
	Key   string
	Value int
}

type Box struct {
	Values []Value
}

const (
	Add = iota
	Remove
)

func partTwo(input string) int {
	hm := make([]Box, 256)
	for _, line := range strings.Split(input, ",") {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		key := ""
		value := 0
		operation := Add
		for i, char := range line {
			if char == '-' {
				operation = Remove
				break
			}
			if char == '=' {
				n, err := strconv.Atoi(line[i+1:])
				if err != nil {
					panic("Failed to convert string to int")
				}
				value = n
				break
			}
			key += string(char)
		}
		idx := hash(key)
		switch operation {
		case Add:
			// fmt.Println(idx, key, value, "Adding")
			// Loop through the values of the box
			bi, item := iterators.Find(hm[idx].Values, func(v Value) bool {
				return v.Key == key
			})
			if bi != -1 {
				// Key found in box, replace it
				item.Value = value
				hm[idx].Values[bi] = *item
			} else {
				// If not found, add it to the end of the box
				hm[idx].Values = append(hm[idx].Values, Value{Key: key, Value: value})
			}

		case Remove:
			// Find the item to remove
			i, _ := iterators.Find(hm[idx].Values, func(v Value) bool {
				return v.Key == key
			})
			if i != -1 {
				// Found item
				hm[idx].Values = append(hm[idx].Values[:i], hm[idx].Values[i+1:]...)
			}
		}
	}
	// Print boxes
	// for i, box := range hm {
	// 	if len(box.Values) == 0 {
	// 		continue
	// 	}
	// 	fmt.Printf("Box %d: [", (i + 1))
	// 	for _, value := range box.Values {
	// 		fmt.Printf("(%s, %d), ", value.Key, value.Value)
	// 	}
	// 	fmt.Println("]")
	// }
	sum := 0
	for i, box := range hm {
		for j, value := range box.Values {
			sum += (i + 1) * (j + 1) * value.Value
		}
	}

	return sum
}

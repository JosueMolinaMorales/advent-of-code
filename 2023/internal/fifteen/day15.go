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
	fmt.Println("Part one", partOne(string(input)))
	fmt.Println("Part two", partTwo(string(input)))
}

func partOne(input string) int {
	sum := 0
	for _, line := range strings.Split(input, ",") {
		sum += hash(line)
	}

	return sum
}

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

func partTwo(input string) int {
	hm := make([]Box, 256)
	for _, line := range strings.Split(input, ",") {
		key, op, value := "", "", 0
		for i, char := range line {
			if char == '-' {
				op = "remove"
				break
			}
			if char == '=' {
				n, _ := strconv.Atoi(line[i+1:])
				value = n
				op = "add"
				break
			}
			key += string(char)
		}
		idx := hash(key)
		matchKey := func(v Value) bool {
			return v.Key == key
		}
		switch op {
		case "add":
			// Find the value to change
			if bi, item := iterators.Find(hm[idx].Values, matchKey); bi != -1 {
				// Key found in box, replace it
				item.Value = value
				hm[idx].Values[bi] = *item
			} else {
				// If not found, add it to the end of the box
				hm[idx].Values = append(hm[idx].Values, Value{Key: key, Value: value})
			}
		case "remove":
			// Find the item to remove
			if i, _ := iterators.Find(hm[idx].Values, matchKey); i != -1 {
				hm[idx].Values = append(hm[idx].Values[:i], hm[idx].Values[i+1:]...)
			}
		}
	}
	sum := 0
	for i, box := range hm {
		for j, value := range box.Values {
			sum += (i + 1) * (j + 1) * value.Value
		}
	}

	return sum
}

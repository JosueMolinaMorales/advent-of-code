package twenty

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/josuemolinamorales/aoc-2023/utils"
	"github.com/josuemolinamorales/aoc-2023/utils/iterators"
	"github.com/josuemolinamorales/aoc-2023/utils/maps"
)

const testInput = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

func RunDayTwenty() {
	input, err := os.ReadFile("./input/day20.txt")
	if err != nil {
		panic("Failed to read day 20 input")
	}
	fmt.Println("Part one:", partOne(string(input)))
	fmt.Println("Part two:", partTwo(string(input)))
}

type Pulse struct {
	from  string
	to    string
	pulse int
}

type Module struct {
	key          string
	moduleType   string
	on           bool
	destinations []string
	inputs       []InputState
}

type InputState struct {
	key   string
	state int
}

const (
	HIGH_PULSE  = 0
	LOW_PULSE   = 1
	FLIP_FLOP   = "FlipFlop"
	CONJUNCTION = "Conjunction"
	BROADCASTER = "Broadcaster"
)

func parse(input string) map[string]Module {
	modules := make(map[string]Module, 0)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " -> ")
		dests := strings.Split(parts[1], ", ")
		mt := ""
		key := parts[0]
		if strings.HasPrefix(parts[0], "%") {
			// Flip flop
			mt = FLIP_FLOP
			key = key[1:]
		} else if strings.HasPrefix(parts[0], "&") {
			mt = CONJUNCTION
			key = key[1:]
		} else {
			mt = BROADCASTER
		}
		modules[key] = Module{
			key:          key,
			moduleType:   mt,
			destinations: dests,
			on:           false,
		}
	}

	// Setup inputs
	for key, module := range modules {
		for _, dest := range module.destinations {
			if modules[dest].moduleType != CONJUNCTION {
				continue
			}
			c := modules[dest]
			c.inputs = append(modules[dest].inputs, InputState{key, LOW_PULSE})
			modules[dest] = c
		}
	}
	return modules
}

func partTwo(input string) int {
	modules := parse(input)
	// Find the inputs that output to rx
	rxIn := ""
	for key, module := range modules {
		for _, dest := range module.destinations {
			if dest == "rx" {
				rxIn = key
				break
			}
		}
	}
	watch := make([]string, 0)
	for _, dest := range modules[rxIn].inputs {
		watch = append(watch, dest.key)
	}

	firstHighPulse := make(map[string]int, 0)
	for _, key := range watch {
		firstHighPulse[key] = 0
	}

	i := 0
	for !iterators.Every(maps.Values(firstHighPulse), func(value int) bool {
		return value > 0
	}) {
		i++
		pules := press_button(&modules)
		for _, pulse := range pules {
			if pulse.pulse == HIGH_PULSE && iterators.Contains(watch, pulse.from) && firstHighPulse[pulse.from] == 0 {
				firstHighPulse[pulse.from] = i
			}
		}
	}
	values := maps.Values(firstHighPulse)
	numbers := make([]*big.Int, 0)
	for _, value := range values {
		numbers = append(numbers, big.NewInt(int64(value)))
	}

	lcm := utils.MultipleLCM(numbers)
	return int(lcm.Int64())
}

func partOne(input string) int {
	modules := parse(input)
	low, high := 0, 0
	for i := 0; i < 1000; i++ {
		pulses := press_button(&modules)
		for _, pulse := range pulses {
			if pulse.pulse == HIGH_PULSE {
				high++
			} else {
				low++
			}
		}
	}
	return low * high
}

func press_button(modules *map[string]Module) []Pulse {
	firstPulse := Pulse{"button", "broadcaster", LOW_PULSE}
	queue := []Pulse{firstPulse}
	pulses := []Pulse{}
	for len(queue) > 0 {
		p := queue[0]
		pulses = append(pulses, p)
		queue = queue[1:]
		to := (*modules)[p.to]

		if to.moduleType == FLIP_FLOP {
			if p.pulse == HIGH_PULSE {
				continue
			}
			to.on = !to.on
			pulse := LOW_PULSE
			if to.on {
				pulse = HIGH_PULSE
			} else {
				pulse = LOW_PULSE
			}
			(*modules)[p.to] = to
			for _, out := range to.destinations {
				queue = append(queue, Pulse{p.to, out, pulse})
			}
		} else if to.moduleType == CONJUNCTION {
			// Update the state
			m := (*modules)[p.to]
			m.inputs = iterators.Map(m.inputs, func(input InputState) InputState {
				if input.key == p.from {
					input.state = p.pulse
				}
				return input
			})
			(*modules)[p.to] = m
			// Check if all inputs are high
			allHigh := iterators.Every(m.inputs, func(input InputState) bool {
				return input.state == HIGH_PULSE
			})
			pulse := HIGH_PULSE
			if allHigh {
				pulse = LOW_PULSE
			}
			for _, out := range to.destinations {
				queue = append(queue, Pulse{p.to, out, pulse})
			}
		} else { // Broadcaster
			for _, out := range to.destinations {
				queue = append(queue, Pulse{p.to, out, LOW_PULSE})
			}
		}
	}

	return pulses
}

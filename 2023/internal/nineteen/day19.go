package nineteen

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const testInput = `px{a<2006:qkq,m>2090:A,rfg}
pv{a>1716:R,A}
lnx{m>1548:A,A}
rfg{s<537:gd,x>2440:R,A}
qs{s>3448:A,lnx}
qkq{x<1416:A,crn}
crn{x>2662:A,R}
in{s<1351:px,qqz}
qqz{s>2770:qs,m<1801:hdj,R}
gd{a>3333:R,R}
hdj{m>838:A,pv}

{x=787,m=2655,a=1222,s=2876}
{x=1679,m=44,a=2067,s=496}
{x=2036,m=264,a=79,s=2244}
{x=2461,m=1339,a=466,s=291}
{x=2127,m=1623,a=2188,s=1013}`

func RunDayNineteen() {
	input, err := os.ReadFile("./input/day19.txt")
	if err != nil {
		panic("Failed to read day 19 input")
	}
	fmt.Println("Day 19 Part 1", partOne(string(input)))
	fmt.Println("Day 19 Part 2", partTwo(string(input)))
}

type Part = map[string]int

const (
	LESS    = "<"
	GREATER = ">"
	ACCEPT  = "A"
	REJECT  = "R"
)

type Rule struct {
	category    string
	destination string
	rule        *RuleOperation
}

type RuleOperation struct {
	operation string
	amount    int
}

type Range struct {
	low, high int
}

type PartRanges = map[string]Range

func parseInput(input string) ([]Part, map[string][]Rule) {
	parts := make([]Part, 0)
	workflows := make(map[string][]Rule)
	ws := strings.Split(input, "\n\n")[0]
	for _, l := range strings.Split(ws, "\n") {
		pat := regexp.MustCompile(`(\w+)\s*{([^}]+)}`)
		matches := pat.FindAllStringSubmatch(l, -1)

		for _, match := range matches {
			key := match[1]
			r := strings.Split(match[2], ",")
			rules := make([]Rule, 0)

			for _, rule := range r {
				newRule := Rule{}
				if strings.Contains(rule, ":") {
					// There is a mapping
					rs := strings.Split(rule, ":")
					newRule.destination = rs[1]

					ruleOp := &RuleOperation{}
					newRule.category = string(rs[0][0])
					amount, _ := strconv.Atoi(rs[0][2:])
					ruleOp.amount = amount
					ruleOp.operation = string(rs[0][1])
					newRule.rule = ruleOp
				} else {
					// Direct
					newRule.destination = rule
				}
				rules = append(rules, newRule)
			}

			workflows[key] = rules
		}
	}

	ps := strings.Split(input, "\n\n")[1]
	for _, l := range strings.Split(ps, "\n") {
		pattern := regexp.MustCompile(`(\w+)\s*=\s*([^,}]+)`)

		// Find all matches
		matches := pattern.FindAllStringSubmatch(l, -1)

		// Iterate over matches
		x, _ := strconv.Atoi(matches[0][2])
		m, _ := strconv.Atoi(matches[1][2])
		a, _ := strconv.Atoi(matches[2][2])
		s, _ := strconv.Atoi(matches[3][2])

		parts = append(parts, Part{"x": x, "m": m, "a": a, "s": s})
	}

	return parts, workflows
}

func partTwo(input string) int {
	// Setup up ranges for x,m,a,s
	ranges := PartRanges{
		"x": Range{1, 4000},
		"m": Range{1, 4000},
		"a": Range{1, 4000},
		"s": Range{1, 4000},
	}
	_, workflows := parseInput(input)
	acceptCount := solve(ranges, workflows, "in")
	return acceptCount
}

func countRanges(ranges PartRanges) int {
	count := 1
	for _, v := range ranges {
		count *= (v.high - v.low) + 1
	}

	return count
}

func solve(ranges PartRanges, workflows map[string][]Rule, cw string) int {
	// If the current state is ACCEPT, return the count of ranges
	if cw == ACCEPT {
		return countRanges(ranges)
	} else if cw == REJECT {
		return 0
	}

	// Get the rules for the current state
	rules := workflows[cw]
	// Initialize count to 0
	count := 0
	// Iterate over the rules for the current state
	for _, r := range rules {
		// Copy ranges to avoid modifying the original
		nr := PartRanges{
			"x": ranges["x"],
			"m": ranges["m"],
			"a": ranges["a"],
			"s": ranges["s"],
		}

		// Check if the rule is nil, indicating an immediate transition
		if r.rule == nil {
			cw = r.destination
			// Go to the next rule and accumulate the count
			count += solve(ranges, workflows, cw)
			continue
		}

		cat := r.category
		switch r.rule.operation {
		case LESS:
			// Adjust the ranges for the LESS operation
			nr[cat] = Range{nr[cat].low, r.rule.amount - 1}
			// Update the ranges for the existing range for next iteration
			ranges[cat] = Range{r.rule.amount, ranges[cat].high}
			// Recursively call solve with updated ranges and destination state
			count += solve(nr, workflows, r.destination)
		case GREATER:
			// Adjust the ranges for the GREATER operation
			nr[cat] = Range{r.rule.amount + 1, nr[cat].high}
			// Update the ranges for the existing range for next iteration
			ranges[cat] = Range{ranges[cat].low, r.rule.amount}
			// Recursively call solve with updated ranges and destination state
			count += solve(nr, workflows, r.destination)
		}
	}

	// Return the accumulated count
	return count
}

func partOne(input string) int {
	parts, workflows := parseInput(input)
	sum := 0

	for _, p := range parts {
		cw := "in"
		for {
			rules := workflows[cw]
			for _, r := range rules {
				if r.rule == nil {
					cw = r.destination
					break
				}
				match := false
				switch r.rule.operation {
				case LESS:
					match = p[r.category] < r.rule.amount
				case GREATER:
					match = p[r.category] > r.rule.amount
				}
				if match {
					cw = r.destination
					break
				}
			}
			if cw == ACCEPT || cw == REJECT {
				break
			}
		}

		if cw == ACCEPT {
			for _, v := range p {
				sum += v
			}
		}
	}

	return sum
}

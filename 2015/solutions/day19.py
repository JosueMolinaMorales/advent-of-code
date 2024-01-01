def parse(input):
    parts = input.split("\n\n")
    rules = set()
    for rule in parts[0].split("\n"):
        (key, value) = rule.split(" => ")
        rules.add((key, value))

    return (rules, parts[1])


def part_one(input):
    (rules, molecule) = parse(input)
    new_molecules = set()

    for (match, replace) in rules:
        for i in range(len(molecule)):
            if molecule[i:i+len(match)] == match:
                new_molecules.add(
                    molecule[:i]+replace+molecule[i+len(match):])

    return len(new_molecules)


def get_min_steps(rules, molecule):
    # Sort rules by descending length of replacement & match
    rules = sorted(rules, key=lambda r: r[1]+r[0])

    stack = [(0, molecule)]
    seen = set()
    while len(stack) > 0:
        (steps, curr) = stack.pop()
        if curr == "e":
            return steps
        seen.add(curr)

        for (match, replace) in rules:
            for i in range(len(curr)):
                if curr[i:i+len(replace)] == replace:
                    replacement = curr[:i]+match+curr[i+len(replace):]
                    stack.append((steps+1, replacement))
    return -1


def part_two(input):
    (rules, molecule) = parse(input)
    return get_min_steps(rules, molecule)


def run_day_nineteen():
    with open("./input/day19.txt", "r") as f:
        input = f.read()
    print("Day 19 part one:", part_one(input))
    print("Day 19 part two:", part_two(input))

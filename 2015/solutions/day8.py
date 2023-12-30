
def part_one(lines):
    characters = []
    for line in lines:
        characters.append(len(line) - len(eval(line)))

    return sum(characters)


def part_two(lines):
    characters = []
    for line in lines:
        characters.append(
            (2+line.count("\\")+line.count("\"")+len(line))-len(line))

    return sum(characters)


def run_day_eight():
    with open("input/day8.txt", "r") as f:
        lines = f.read().splitlines()

    print("Day 8 part one:", part_one(lines))
    print("Day 8 part two:", part_two(lines))

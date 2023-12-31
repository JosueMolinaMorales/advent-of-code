
def parse(input):
    susans = []
    for i, line in enumerate(input.split("\n")):
        line = line.replace(f"Sue {i+1}: ", "").replace(":", "").split(", ")
        sue = {"id": i+1}
        for item in line:
            item = item.split(" ")
            sue[item[0]] = int(item[1])
        susans.append(sue)

    return susans


correct_sue = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
}


def part_one(input):
    susans = parse(input)
    for sue in susans:
        correct = True
        keys = []
        for key in correct_sue.keys():
            if key in sue.keys():
                keys.append(key)
        for key in keys:
            if sue[key] != correct_sue[key]:
                correct = False
                break
        if correct:
            return sue["id"]


def part_two(input):
    susans = parse(input)
    for sue in susans:
        correct = True
        keys = []
        for key in correct_sue.keys():
            if key in sue.keys():
                keys.append(key)
        for key in keys:
            if key in ["cats", "trees"]:
                if sue[key] <= correct_sue[key]:
                    correct = False
                    break
            elif key in ["pomeranians", "goldfish"]:
                if sue[key] >= correct_sue[key]:
                    correct = False
                    break
            elif sue[key] != correct_sue[key]:
                correct = False
                break
        if correct:
            return sue["id"]


def run_day_sixteen():
    with open("input/day16.txt", "r") as f:
        input = f.read()

    print("Day 16 part one:", part_one(input))
    print("Day 16 part two:", part_two(input))

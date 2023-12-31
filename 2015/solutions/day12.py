import json


def sum_numbers(obj, ignore_red=False):
    if ignore_red and obj == "red":
        return (True, 0)
    # Check if obj is a number
    if isinstance(obj, int):
        return (False, int(obj))

    sum = 0
    # Check if obj is an array
    if isinstance(obj, list):
        for item in obj:
            _, s = sum_numbers(item, ignore_red)
            sum += s

    # Check if obj is an object
    if isinstance(obj, dict):
        for k, v in obj.items():
            (red_found, s) = sum_numbers(k, ignore_red)
            if red_found:
                return (False, 0)
            sum += s
            (red_found, s) = sum_numbers(v, ignore_red)
            if red_found:
                return (False, 0)
            sum += s
    return (False, sum)


def part_one(input):
    obj = json.loads(input)
    return sum_numbers(obj)[1]


def part_two(input):
    obj = json.loads(input)
    return sum_numbers(obj, True)[1]


def run_day_twelve():
    with open("input/day12.txt", "r") as f:
        input = f.read()

    print("Day 12 part one:", part_one(input))
    print("Day 12 part two:", part_two(input))

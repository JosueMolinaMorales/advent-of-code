import math


def parse(lines):
    wires = {}
    for (action, end) in lines:
        wires[end] = action.split(" ")
    return wires


def part_one(lines):
    return get_value(parse(lines), {}, "a")


def part_two(lines):
    wires = parse(lines)
    wires["b"] = [str(get_value(wires, {}, "a"))]
    return get_value(wires, {}, "a")


def get_value(wires, cache, wire):
    if wire in cache:
        return cache[wire]

    if str(wire).isdigit():
        return int(wire)

    action = wires[wire]
    if len(action) == 1:
        # direct assignment
        cache[wire] = get_value(wires, cache, action[0])

    if len(action) == 2:
        # bitwise not
        cache[wire] = int(math.pow(2, 16) + ~
                          get_value(wires, cache, action[1]))

    if len(action) == 3:
        if action[1] == "AND":
            cache[wire] = get_value(
                wires, cache, action[0]) & get_value(wires, cache, action[2])
        if action[1] == "OR":
            cache[wire] = get_value(
                wires, cache, action[0]) | get_value(wires, cache, action[2])
        if action[1] == "LSHIFT":
            cache[wire] = get_value(
                wires, cache, action[0]) << get_value(wires, cache, action[2])
        if action[1] == "RSHIFT":
            cache[wire] = get_value(
                wires, cache, action[0]) >> get_value(wires, cache, action[2])

    return cache[wire]


def run_day_seven():
    with open("input/day7.txt", "r") as f:
        lines = [l.split(" -> ") for l in f.read().splitlines()]

    print("Day 7 part one:", part_one(lines))
    print("Day 7 part two:", part_two(lines))

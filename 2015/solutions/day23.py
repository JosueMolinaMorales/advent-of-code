def solve(registers, input):
    instructions = [ins.replace(",", "") for ins in input.splitlines()]
    curr = 0

    while curr < len(instructions):
        parts = instructions[curr].split(" ")
        inst = parts[0]
        r = parts[1]

        if inst == "hlf":
            registers[r] //= 2
        elif inst == "tpl":
            registers[r] *= 3
        elif inst == "inc":
            registers[r] += 1
        elif inst == "jmp":
            offset = int(r)
            curr += offset
            continue
        elif inst == "jie" and registers[r] % 2 == 0:
            offset = int(parts[2])
            curr += offset
            continue
        elif inst == "jio" and registers[r] == 1:
            offset = int(parts[2])
            curr += offset
            continue

        curr += 1


def part_one(input):
    registers = {"a": 0, "b": 0}
    solve(registers, input)
    return registers["b"]


def part_two(input):
    # Register a starts at 1
    registers = {"a": 1, "b": 0}
    solve(registers, input)
    return registers["b"]


def run_day_twenty_three():
    with open("./input/day23.txt", "r") as f:
        input = f.read()

    print("Day 23 part one:", part_one(input))
    print("Day 23 part two:", part_two(input))

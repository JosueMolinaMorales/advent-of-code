

def part_one(input):
    containers = list(map(int, input.splitlines()))
    count = 0
    # Generate all possible combinations of containers
    for i in range(2**len(containers)):
        total = 0
        # Check if the combination is valid
        for j in range(len(containers)):
            # Check if the jth bit is set
            if i & (1 << j):
                total += containers[j]
        if total == 150:
            count += 1

    return count


def part_two(input):
    containers = list(map(int, input.splitlines()))
    count = 0
    min_containers = len(containers)
    # Generate all possible combinations of containers
    for i in range(2**len(containers)):
        total = 0
        num_containers = 0
        # Check if the combination is valid
        for j in range(len(containers)):
            # Check if the jth bit is set
            if i & (1 << j):
                total += containers[j]
                num_containers += 1
        if total == 150:
            if num_containers < min_containers:
                min_containers = num_containers
                count = 1
            elif num_containers == min_containers:
                count += 1

    return count


def run_day_seventeen():
    with open("input/day17.txt") as f:
        input = f.read()

    print("Day 17 part one:", part_one(input))
    print("Day 17 part two:", part_two(input))

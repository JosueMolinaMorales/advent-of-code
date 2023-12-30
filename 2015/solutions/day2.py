
def partOne(dimensions):
    sum = 0
    for l, w, h in dimensions:
        l, w, h = int(l), int(w), int(h)
        sum += 2*l*w + 2*w*h + 2*h*l + min(l*w, w*h, h*l)
    return sum


def partTwo(dimensions):
    sum = 0
    for l, w, h in dimensions:
        l, w, h = int(l), int(w), int(h)
        x = [l, w, h]
        x.sort()
        sum += (2*x[0]+2*x[1]) + (l*w*h)
    return sum


def run_day_two():
    with open("./input/day2.txt") as f:
        dimensions = f.readlines()
        dimensions = [x.rstrip().split('x') for x in dimensions]
    print("Day 2 part one:", partOne(dimensions))
    print("Day 2 part two:", partTwo(dimensions))


if __name__ == "__main__":
    run_day_two()

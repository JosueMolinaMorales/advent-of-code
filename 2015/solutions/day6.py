def part1(directions):
    points = {}
    for x in range(1000):
        for y in range(1000):
            points[(x, y)] = False

    for direction in directions:
        p1, p2 = direction[1][0:2], direction[1][2:4]
        if direction[0] == "toggle":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] = not points[(x, y)]
        elif direction[0] == "on":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] = True
        elif direction[0] == "off":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] = False

    return sum(map(lambda p: 1 if p else 0, points.values()))


def part2(directions):
    points = {}
    for x in range(1000):
        for y in range(1000):
            points[(x, y)] = 0

    for direction in directions:
        p1, p2 = direction[1][0:2], direction[1][2:4]
        if direction[0] == "toggle":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] += 2
        elif direction[0] == "on":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] += 1
        elif direction[0] == "off":
            for x in range(p1[0], p2[0] + 1):
                for y in range(p1[1], p2[1] + 1):
                    points[(x, y)] = max(0, points[(x, y)] - 1)

    return sum(points.values())


def run_day_six():
    with open("./input/day6.txt", "r") as f:
        lines = list(map(lambda l: (l[0], list(map(lambda p: int(p), l[1].split(",")))), map(lambda l: l.replace(
            " through ", ",").replace("turn ", "").split(" "), f.read().splitlines())))

    print("Day 6 part one:", part1(lines))
    print("Day 6 part two:", part2(lines))


if __name__ == "__main__":
    run_day_six()

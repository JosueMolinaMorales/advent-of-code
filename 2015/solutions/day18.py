
DOWN = (1, 0)
DOWN_LEFT = (1, -1)
DOWN_RIGHT = (1, 1)
UP = (-1, 0)
UP_LEFT = (-1, -1)
UP_RIGHT = (-1, 1)
LEFT = (0, -1)
RIGHT = (0, 1)

DIRECTIONS = [DOWN, DOWN_LEFT, DOWN_RIGHT, UP, UP_LEFT, UP_RIGHT, LEFT, RIGHT]


def parse(input):
    points = {}
    for i, row in enumerate(input.split("\n")):
        for j, c in enumerate(row):
            points[(i, j)] = c == "#"

    return points


def neighbors(point, points):
    n = []
    for dir in DIRECTIONS:
        (dx, dy) = point[0]+dir[0], point[1]+dir[1]
        if (dx, dy) not in points:
            n.append(False)
        else:
            n.append(points[(dx, dy)])

    return n


def switch_lights(lights):
    new_lights = {}
    for key in lights.keys():
        n = neighbors(key, lights)
        on_n = n.count(True)

        if lights[key] and (on_n != 2 and on_n != 3):
            # Turn off if it was on and it doesnt have 2 or three lights on
            new_lights[key] = False
        elif not lights[key] and on_n == 3:
            # Light turns on when it is off and it has three on neighbors
            new_lights[key] = True
        else:
            new_lights[key] = lights[key]

    return new_lights


def part_one(input):
    lights = parse(input)

    for _ in range(100):
        lights = switch_lights(lights)
    # Count how many lights are on
    return sum(map(lambda l: 1 if l else 0, lights.values()))


def part_two(input):
    lights = parse(input)
    # Turn on the four corners
    lights[(0, 0)] = True
    lights[(0, 99)] = True
    lights[(99, 0)] = True
    lights[(99, 99)] = True

    for _ in range(100):
        lights = switch_lights(lights)
        # Turn on the four corners
        lights[(0, 0)] = True
        lights[(0, 99)] = True
        lights[(99, 0)] = True
        lights[(99, 99)] = True

    # Count how many lights are on
    return sum(map(lambda l: 1 if l else 0, lights.values()))


def run_day_eighteen():
    with open("input/day18.txt") as f:
        input = f.read()

    print("Day 18 part one:", part_one(input))
    print("Day 18 part two:", part_two(input))

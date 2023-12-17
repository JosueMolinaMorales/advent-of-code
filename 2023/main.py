import sys
from collections import defaultdict
from queue import PriorityQueue

RIGHT, LEFT, UP, DOWN = (1, 0), (-1, 0), (0, -1), (0, 1)  # (x, y)
CRUCIBLE_TURNS = {RIGHT: [UP, DOWN], LEFT: [
    UP, DOWN], UP: [LEFT, RIGHT], DOWN: [LEFT, RIGHT]}

with open('./input/day17.txt') as f:
    heat_loss_map = [list(map(int, line.strip())) for line in f.readlines()]

    # distance is in this case cumulative heat loss for each direction
    dist = defaultdict(lambda: defaultdict(lambda: sys.maxsize))
    dist[(0, 0)][RIGHT] = 0
    dist[(0, 0)][DOWN] = 0
    dist[(0, 0)][UP] = 0
    dist[(0, 0)][LEFT] = 0

    # Dijkstra's algorithm
    pq = PriorityQueue()
    pq.put((0, (0, 0), RIGHT))
    pq.put((0, (0, 0), DOWN))

    while not pq.empty():
        heat_loss, position, direction = pq.get()

        if heat_loss > dist[position][direction]:
            continue

        # find positions where the crucible can turn
        x, y = position
        for _ in range(3):
            # move in the direction
            x, y = x + direction[0], y + direction[1]

            # out of bounds check
            if x < 0 or x >= len(heat_loss_map[0]) or y < 0 or y >= len(heat_loss_map):
                break

            # cumulate heat losses
            heat_loss += heat_loss_map[y][x]

            # turn the crucible
            for new_dir in CRUCIBLE_TURNS[direction]:
                if heat_loss < dist[(x, y)][new_dir]:
                    dist[(x, y)][new_dir] = heat_loss
                    pq.put((heat_loss, (x, y), new_dir))

    print(
        min(dist[(len(heat_loss_map) - 1, len(heat_loss_map[0]) - 1)].values()))

# 2. part - Directing the ultra crucible from the lava pool to the machine parts factory,
#           what is the least heat loss it can incur?

RIGHT, LEFT, UP, DOWN = (1, 0), (-1, 0), (0, -1), (0, 1)
CRUCIBLE_TURNS = {RIGHT: [UP, DOWN], LEFT: [
    UP, DOWN], UP: [LEFT, RIGHT], DOWN: [LEFT, RIGHT]}

with open('./input/day17.txt') as f:
    heat_loss_map = [list(map(int, line.strip())) for line in f.readlines()]

    # distance is in this case cumulative heat loss
    dist = defaultdict(lambda: defaultdict(lambda: sys.maxsize))
    dist[(0, 0)][RIGHT] = 0
    dist[(0, 0)][DOWN] = 0
    dist[(0, 0)][UP] = 0
    dist[(0, 0)][LEFT] = 0

    # Dijkstra's algorithm
    pq = PriorityQueue()
    pq.put((0, (0, 0), RIGHT))
    pq.put((0, (0, 0), DOWN))

    while not pq.empty():
        heat_loss, position, direction = pq.get()

        if heat_loss > dist[position][direction]:
            continue

        # find positions where the crucible can turn
        x, y = position
        for block in range(10):  # <-- CHANGE HERE
            # move in the direction
            x, y = x + direction[0], y + direction[1]

            # out of bounds check
            if x < 0 or x >= len(heat_loss_map[0]) or y < 0 or y >= len(heat_loss_map):
                break

            # cumulate heat losses
            heat_loss += heat_loss_map[y][x]

            # ultra crucible needs to move a minimum of four blocks in that direction before it can turn
            if block < 3:  # <-- CHANGE HERE
                continue

            # turn the crucible
            for new_dir in CRUCIBLE_TURNS[direction]:
                if heat_loss < dist[(x, y)][new_dir]:
                    dist[(x, y)][new_dir] = heat_loss
                    pq.put((heat_loss, (x, y), new_dir))

    print(
        min(dist[(len(heat_loss_map) - 1, len(heat_loss_map[0]) - 1)].values()))


# 1. and 2. part combined

RIGHT, LEFT, UP, DOWN = (1, 0), (-1, 0), (0, -1), (0, 1)
CRUCIBLE_TURNS = {RIGHT: [UP, DOWN], LEFT: [
    UP, DOWN], UP: [LEFT, RIGHT], DOWN: [LEFT, RIGHT]}


def get_min_heat_loss(heat_loss_map, blocks_before_turn, max_in_direction):
    # distance is in this case cumulative heat loss
    dist = defaultdict(lambda: defaultdict(lambda: sys.maxsize))
    dist[(0, 0)][RIGHT] = 0
    dist[(0, 0)][DOWN] = 0
    dist[(0, 0)][UP] = 0
    dist[(0, 0)][LEFT] = 0

    # Dijkstra's algorithm
    pq = PriorityQueue()
    pq.put((0, (0, 0), RIGHT))
    pq.put((0, (0, 0), DOWN))

    while not pq.empty():
        heat_loss, position, direction = pq.get()

        if heat_loss > dist[position][direction]:
            continue

        # find positions where the crucible can turn
        x, y = position
        for block in range(max_in_direction):
            # move in the direction
            x, y = x + direction[0], y + direction[1]

            # out of bounds check
            if x < 0 or x >= len(heat_loss_map[0]) or y < 0 or y >= len(heat_loss_map):
                break

            # cumulate heat losses
            heat_loss += heat_loss_map[y][x]

            # crucible needs to move a minimum of N blocks in that direction before it can turn
            if block < blocks_before_turn:
                continue

            # turn the crucible
            for new_dir in CRUCIBLE_TURNS[direction]:
                if heat_loss < dist[(x, y)][new_dir]:
                    dist[(x, y)][new_dir] = heat_loss
                    pq.put((heat_loss, (x, y), new_dir))

    return min(dist[(len(heat_loss_map) - 1, len(heat_loss_map[0]) - 1)].values())


with open('./input/day17.txt') as f:
    heat_loss_map = [list(map(int, line.strip())) for line in f.readlines()]

    print(get_min_heat_loss(heat_loss_map, 0, 3))
    print(get_min_heat_loss(heat_loss_map, 3, 10))


MOVES = {
    ">": (0, 1),
    "<": (0, -1),
    "^": (1, 0),
    "v": (-1, 0)
}


def partOne(moves):
    houses = [(0, 0)]
    for m in moves:
        move = MOVES[m]
        lh = houses[len(houses)-1]
        houses.append((lh[0]+move[0], lh[1]+move[1]))
    houses = set(houses)

    return len(houses)


def partTwo(moves):
    sh = [(0, 0)]
    rh = [(0, 0)]
    for i, m in enumerate(moves):
        move = MOVES[m]
        if i % 2 == 0:
            # Santa moves
            lh = sh[len(sh)-1]
            sh.append((lh[0]+move[0], lh[1]+move[1]))
        else:
            # Robo Santa moves
            lh = rh[len(rh)-1]
            rh.append((lh[0]+move[0], lh[1]+move[1]))

    return len(set(sh+rh))


def run_day_three():
    with open("./input/day3.txt", "r") as f:
        input = f.readline()
        print("Day 3 part one:", partOne(input))
        print("Day 3 part two:", partTwo(input))


if __name__ == "__main__":
    run_day_three()


def parse(input):
    deers = list(map(lambda r: (int(r[3]), int(r[6]), int(
        r[13])), [r.split(" ") for r in input.replace(".", "").split("\n")]))
    return deers


def part_one(input):
    r = parse(input)

    time = 2503
    max_dist = 0
    for (speed, s_dur, r_dur) in r:
        cycles = time // (s_dur+r_dur)
        dist = (speed*s_dur) * cycles
        dist += min(s_dur, time % (s_dur+r_dur)) * speed
        max_dist = max(max_dist, dist)

    return max_dist


def part_two(input):
    r = parse(input)

    time = 2503
    # Keep track of points for each deer
    points = [0 for _ in range(len(r))]
    for t in range(1, time+1):
        max_dist = 0  # Keep track of max distance
        max_dist_deers = []  # Keep track of deers that have taken the max distance
        for (i, (speed, s_dur, r_dur)) in enumerate(r):
            cycles = t // (s_dur+r_dur)
            dist = (speed*s_dur) * cycles
            dist += min(s_dur, t % (s_dur+r_dur)) * speed
            if dist > max_dist:
                max_dist = dist
                max_dist_deers = [i]
            elif dist == max_dist:
                max_dist_deers.append(i)
        for i in max_dist_deers:
            points[i] += 1

    return max(points)


def run_day_fourteen():
    with open("./input/day14.txt", "r") as f:
        input = f.read()
    print("Day 14 part one:", part_one(input))
    print("Day 14 part two:", part_two(input))

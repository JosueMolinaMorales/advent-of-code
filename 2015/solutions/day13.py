from itertools import permutations


def parse(input):
    seating_happines = list(map(lambda p: (p[0], p[2], int(
        p[3]), p[-1]), [p.replace(".", "").split(" ") for p in input.split(".\n")]))
    people = set()
    sh = {}
    for (fr, gl, units, to) in seating_happines:
        people.add(fr)
        people.add(to)
        if (fr, to) not in sh:
            sh[(fr, to)] = units if gl == "gain" else -units

    return (people, sh)


def get_optimal_happines(people, sh):
    max_h = 0
    for sa in permutations(people):
        h = 0
        for i in range(len(sa)):
            h += sh[(sa[i], sa[(i+1) % len(sa)])]
            h += sh[(sa[(i+1) % len(sa)], sa[i])]
        max_h = max(max_h, h)
    return max_h


def part_one(input):
    (people, sh) = parse(input)
    return get_optimal_happines(people, sh)


def part_two(input):
    (people, sh) = parse(input)
    # Add myself to table
    for person in people:
        sh[("me", person)] = 0
        sh[(person, "me")] = 0
    people.add("me")
    return get_optimal_happines(people, sh)


def run_day_thirteen():
    with open("./input/day13.txt", "r") as f:
        input = f.read()
    print("Day 13 part one:", part_one(input))
    print("Day 13 part two:", part_two(input))

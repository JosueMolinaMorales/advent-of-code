from itertools import combinations
import math


def solve(num_groups, weights):
    # Find total weights
    total_weight = sum(weights)
    # Find weight for each group
    weight_for_group = total_weight // num_groups
    # Find all combinations of weights that sum to weight_for_group
    groups = []
    for i in range(1, len(weights)):
        for c in combinations(weights, i):
            if sum(c) == weight_for_group:
                groups.append(c)

    # Find the smallest group
    min_len = len(groups[0])
    for g in groups:
        if len(g) < min_len:
            min_len = len(g)

    # Find all groups with the smallest length
    min_groups = []
    for g in groups:
        if len(g) == min_len:
            min_groups.append(g)

    # Find the smallest quantum entanglement
    min_qe = math.inf
    for g in min_groups:
        qe = 1
        for n in g:
            qe *= n
        if qe < min_qe:
            min_qe = qe

    return min_qe


def part_one(weights):
    return solve(3, weights)


def part_two(weights):
    return solve(4, weights)


def run_day_twenty_four():
    with open("./input/day24.txt", "r") as f:
        weights = [int(n) for n in f.read().splitlines()]

    print("Day 24 part one:", part_one(weights))
    print("Day 24 part two:", part_two(weights))

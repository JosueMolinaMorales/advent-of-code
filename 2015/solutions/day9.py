from collections import defaultdict
import math


def parse(input) -> dict:
    graph = defaultdict(list)

    for (from_node, to_node, distance) in input:
        distance = int(distance)
        graph[from_node].append((to_node, distance))
        graph[to_node].append((from_node, distance))

    return graph


def dfs(graph, start):
    # Find all paths from start to other nodes
    paths = []
    stack = [(start, [start])]
    while stack:
        (vertex, path) = stack.pop()
        for (n, w) in graph[vertex]:
            if n not in path:
                stack.append((n, path+[n]))
            else:
                paths.append(path+[n])

    paths = list(filter(lambda p: len(p) == len(
        graph.keys()), map(lambda p: p[0:-1], paths)))

    # Find shortest path
    shortest = math.inf
    longest = 0
    for p in paths:
        dist = 0
        for i in range(len(p)-1):
            for (n, w) in graph[p[i]]:
                if n == p[i+1]:
                    dist += w
        if dist < shortest:
            shortest = dist
        if dist > longest:
            longest = dist

    return (shortest, longest)


def part_one(input_data):
    graph = parse(input_data)
    min_distance = min(dfs(graph, node)[0] for node in graph.keys())
    return min_distance


def part_two(input_data):
    graph = parse(input_data)
    max_distance = max(dfs(graph, node)[1] for node in graph.keys())
    return max_distance


def run_day_nine():
    with open("input/day9.txt", "r") as f:
        input = [l.replace(" to ", " ").replace(" = ", " ").split(" ")
                 for l in f.read().splitlines()]

    print("Day 9 part one:", part_one(input))
    print("Day 9 part two:", part_two(input))


import math


def part_one(input):
    min_house = 1

    houses = [0 for _ in range(input//10)]
    for i in range(1, input//10):
        for j in range(i, input//10, i):
            houses[j] += i*10

    for i in range(len(houses)):
        if houses[i] >= input:
            min_house = i
            break
    return min_house


def part_two(input):
    size = input // 10
    houses = [0 for _ in range(size)]

    for i in range(1, size+1):
        counter = 0
        for j in range(i, size, i):
            houses[j] += i*11
            counter += 1
            if counter == 50:
                break

    for i in range(len(houses)):
        if houses[i] >= input:
            return i
    return -1


def run_day_twenty():
    input = 29000000
    print("Day 20 part one:", part_one(input))
    print("Day 20 part two:", part_two(input))

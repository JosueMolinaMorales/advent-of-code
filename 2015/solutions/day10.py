def solve(inp, repeat):
    for _ in range(repeat):
        ans = ""
        i = 0
        while i < len(inp):
            count = 1
            while i + count < len(inp) and inp[i] == inp[i + count]:
                count += 1
            ans += str(count) + inp[i]
            i += count
        inp = ans
    return len(ans)


def part_one(input):
    return solve(input, 40)


def part_two(input):
    return solve(input, 50)


def run_day_ten():
    input = "1113122113"
    print("Day 10 part one:", part_one(input))
    print("Day 10 part two:", part_two(input))


if __name__ == '__main__':
    run_day_ten()

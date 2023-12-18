
def part_one(input):
    count = 0
    for ch in input:
        if ch == '(':
            count += 1
        else:
            count -= 1
    print(count)


def part_two(input):
    position = 0
    count = 0
    for i, ch in enumerate(input):
        if ch == '(':
            count += 1
        else:
            count -= 1
            if count == -1:
                position = i + 1
                break
    print(position)


def run_day_one():
    with open("./input/dayone.txt", "r") as f:
        input = f.read()
        # part_one(input)
        part_two(input)


if __name__ == "__main__":
    run_day_one()

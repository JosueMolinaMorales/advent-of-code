
def part_one(row, column):
    # Find the position of the number based on the row and column
    pos = 1
    for i in range(1, row):
        pos += i
    for i in range(1, column):
        pos += row+i

    # Find the number of the code
    code = 20151125
    for i in range(1, pos):
        code = (code * 252533) % 33554393

    return code


def run_day_twenty_five():
    print("Day 25 part one:", part_one(2947, 3029))

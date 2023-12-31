
def parse(input):
    ingredients = []
    for line in input.splitlines():
        line = line.split()
        ingredients.append({
            "capacity": int(line[2][:-1]),
            "durability": int(line[4][:-1]),
            "flavor": int(line[6][:-1]),
            "texture": int(line[8][:-1]),
            "calories": int(line[10]),
        })
    return ingredients


def part_one(input):
    ingredients = parse(input)

    max_score = 0
    for a in range(101):
        for b in range(101-a):
            for c in range(101-a-b):
                d = 100-a-b-c
                capacity = max(0, a*ingredients[0]["capacity"] + b*ingredients[1]
                               ["capacity"] + c*ingredients[2]["capacity"] + d*ingredients[3]["capacity"])
                durability = max(0, a*ingredients[0]["durability"] + b*ingredients[1]
                                 ["durability"] + c*ingredients[2]["durability"] + d*ingredients[3]["durability"])
                flavor = max(0, a*ingredients[0]["flavor"] + b*ingredients[1]
                             ["flavor"] + c*ingredients[2]["flavor"] + d*ingredients[3]["flavor"])
                texture = max(0, a*ingredients[0]["texture"] + b*ingredients[1]
                              ["texture"] + c*ingredients[2]["texture"] + d*ingredients[3]["texture"])
                score = capacity * durability * flavor * texture
                max_score = max(max_score, score)

    return max_score


def part_two(input):
    ingredients = parse(input)

    max_score = 0
    for a in range(101):
        for b in range(101-a):
            for c in range(101-a-b):
                d = 100-a-b-c
                calories = a*ingredients[0]["calories"] + b*ingredients[1]["calories"] + \
                    c*ingredients[2]["calories"] + d*ingredients[3]["calories"]
                if calories == 500:
                    capacity = max(0, a*ingredients[0]["capacity"] + b*ingredients[1]
                                   ["capacity"] + c*ingredients[2]["capacity"] + d*ingredients[3]["capacity"])
                    durability = max(0, a*ingredients[0]["durability"] + b*ingredients[1]
                                     ["durability"] + c*ingredients[2]["durability"] + d*ingredients[3]["durability"])
                    flavor = max(0, a*ingredients[0]["flavor"] + b*ingredients[1]
                                 ["flavor"] + c*ingredients[2]["flavor"] + d*ingredients[3]["flavor"])
                    texture = max(0, a*ingredients[0]["texture"] + b*ingredients[1]
                                  ["texture"] + c*ingredients[2]["texture"] + d*ingredients[3]["texture"])
                    score = capacity * durability * flavor * texture
                    max_score = max(max_score, score)

    return max_score


def run_day_fifteen():
    with open("./input/day15.txt", "r") as f:
        input = f.read()
    print("Day 15 part one:", part_one(input))
    print("Day 15 part two:", part_two(input))

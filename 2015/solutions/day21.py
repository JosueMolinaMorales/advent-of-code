weapons = [
    {"cost": 8, "dmg": 4},
    {"cost": 10, "dmg": 5},
    {"cost": 25, "dmg": 6},
    {"cost": 40, "dmg": 7},
    {"cost": 74, "dmg": 8}
]

armors = [
    {"cost": 0, "armor": 0},
    {"cost": 13, "armor": 1},
    {"cost": 31, "armor": 2},
    {"cost": 53, "armor": 3},
    {"cost": 75, "armor": 4},
    {"cost": 102, "armor": 5}
]

rings = [
    {"cost": 0, "dmg": 0, "armor": 0},
    {"cost": 0, "dmg": 0, "armor": 0},
    {"cost": 25, "dmg": 1, "armor": 0},
    {"cost": 50, "dmg": 2, "armor": 0},
    {"cost": 100, "dmg": 3, "armor": 0},
    {"cost": 20, "dmg": 0, "armor": 1},
    {"cost": 40, "dmg": 0, "armor": 2},
    {"cost": 80, "dmg": 0, "armor": 3}
]


def fight(boss, player):
    boss_hp = boss["hp"]
    player_hp = player["hp"]

    while boss_hp > 0 and player_hp > 0:
        boss_hp -= max(player["dmg"] - boss["armor"], 1)
        if boss_hp <= 0:
            return True
        player_hp -= max(boss["dmg"] - player["armor"], 1)
        if player_hp <= 0:
            return False


def purchase_items():
    items = []
    for weapon in weapons:
        for armor in armors:
            for ring1 in rings:
                for ring2 in rings:
                    if ring1 == ring2 and ring1["cost"] != 0:
                        continue
                    cost = weapon["cost"] + armor["cost"] + \
                        ring1["cost"] + ring2["cost"]
                    dmg = weapon["dmg"] + ring1["dmg"] + ring2["dmg"]
                    armor_stat = armor["armor"] + \
                        ring1["armor"] + ring2["armor"]
                    items.append((cost, dmg, armor_stat))

    return items


def part_one(boss):
    # Need to buy one weapon, armour is optional, and can buy 0-2 rings
    min_cost = 100000
    items = purchase_items()
    for (cost, dmg, armor) in items:
        player = {
            "hp": 100,
            "dmg": dmg,
            "armor": armor
        }
        if fight(boss, player):
            min_cost = min(min_cost, cost)

    return min_cost


def part_two(boss):
    # What is the most amount of gold you can spend and still lose the fight?
    max_cost = 0
    items = purchase_items()
    for (cost, dmg, armor) in items:
        player = {
            "hp": 100,
            "dmg": dmg,
            "armor": armor
        }
        if not fight(boss, player):
            max_cost = max(max_cost, cost)

    return max_cost


def run_day_twentyone():
    boss = {
        "hp": 109,
        "dmg": 8,
        "armor": 2
    }

    print("Day 21 part one:", part_one(boss))
    print("Day 21 part two:", part_two(boss))

import math
import heapq

spells = [
    {
        "cost": 53,
        "dmg": 4,
        "heal": 0,
        "armor": 0,
        "mana": 0,
        "duration": 0
    },
    {
        "cost": 73,
        "dmg": 2,
        "heal": 2,
        "armor": 0,
        "mana": 0,
        "duration": 0
    },
    {
        "name": "Shield",
        "cost": 113,
        "dmg": 0,
        "heal": 0,
        "armor": 7,
        "mana": 0,
        "duration": 6
    },
    {
        "name": "Poison",
        "cost": 173,
        "dmg": 3,
        "heal": 0,
        "armor": 0,
        "mana": 0,
        "duration": 6
    },
    {
        "name": "Recharge",
        "cost": 229,
        "dmg": 0,
        "heal": 0,
        "armor": 0,
        "mana": 101,
        "duration": 5
    }
]


def part_one(boss):
    return fight_boss(boss["hp"], boss["dmg"], False)


def fight_boss(boss_hp, boss_dmg, hard_mode):
    # Use dijkstra's algorithm to find the shortest path
    starting_mana = 500
    player_hp = 50

    # Create a queue of states to check
    queue = []
    # Add the starting state to the queue
    heapq.heappush(queue, (0, starting_mana, player_hp, boss_hp, 0, 0, 0, 0))

    while len(queue) > 0:
        # Get the first state in the queue
        used_mana, mana, player_hp, boss_hp, turn, shield, poison, recharge = heapq.heappop(
            queue)
        # Check if the boss is dead
        if boss_hp <= 0:
            return used_mana

        if hard_mode and turn == 0:
            player_hp -= 1

        if player_hp <= 0:
            continue

        # Apply all active affects
        if shield > 0:
            shield -= 1

        if poison > 0:
            boss_hp -= 3
            poison -= 1
            if boss_hp <= 0:
                return used_mana

        if recharge > 0:
            mana += 101
            recharge -= 1

        # Player turn
        if turn == 0:
            # Apply magic missile
            if mana >= 53:
                heapq.heappush(queue, (used_mana+53, mana-53, player_hp, boss_hp-4, 1,
                                       shield, poison, recharge))
            # Apply drain
            if mana >= 73:
                heapq.heappush(queue, (used_mana + 73, mana-73, player_hp+2, boss_hp-2, 1,
                                       shield, poison, recharge))
            # Apply shield
            if mana >= 113 and shield == 0:
                heapq.heappush(queue, (used_mana + 113, mana-113, player_hp, boss_hp, 1,
                                       6, poison, recharge))
            # Apply poison
            if mana >= 173 and poison == 0:
                heapq.heappush(queue, (used_mana + 173, mana-173, player_hp, boss_hp, 1,
                                       shield, 6, recharge))
            # Apply recharge
            if mana >= 229 and recharge == 0:
                heapq.heappush(queue, (used_mana + 229, mana-229, player_hp, boss_hp, 1,
                                       shield, poison, 5))
        else:
            # Boss turn
            player_hp -= max(1, boss_dmg - (7 if shield > 0 else 0))

            # Add all possible player turns to the queue
            heapq.heappush(queue, (used_mana, mana, player_hp, boss_hp, 0,
                                   shield, poison, recharge))

    return math.inf


def part_two(boss):
    return fight_boss(boss["hp"], boss["dmg"], True)


def run_day_twenty_two():
    boss = {
        "hp": 71,
        "dmg": 10,
    }
    print("Day 22 part one:", part_one(boss))
    print("Day 22 part two:", part_two(boss))

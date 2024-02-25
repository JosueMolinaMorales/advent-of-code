use std::{
    cmp::Reverse,
    collections::{BinaryHeap, HashMap, HashSet},
    hash::Hash,
};

use crate::utils::{file_loader::FileLoader, point::Point};

pub fn solve_day_fifteen() {
    let lines = FileLoader::new("./inputs/day15.txt".into()).read_lines();
    println!("Day 15 Part 1: {}", part_one(&lines));
    println!("Day 15 Part 2: {}", part_two(&lines))
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
enum Team {
    Elf,
    Goblin,
}

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Unit {
    pos: Point,
    hp: i64,
    team: Team,
}

fn parse(lines: &Vec<String>) -> (HashSet<Point>, Vec<Unit>) {
    let mut walls = HashSet::new();
    let mut units = Vec::new();
    for (x, line) in lines.iter().enumerate() {
        for (y, c) in line.chars().enumerate() {
            let pos = Point::new(x as i32, y as i32);
            match c {
                '#' => {
                    walls.insert(pos);
                }
                'E' => {
                    units.push(Unit {
                        pos,
                        hp: 200,
                        team: Team::Elf,
                    });
                }
                'G' => {
                    units.push(Unit {
                        pos,
                        hp: 200,
                        team: Team::Goblin,
                    });
                }
                _ => {}
            }
        }
    }

    (walls, units)
}

fn solve(
    lines: &Vec<String>,
    elf_deaths_allowed: bool,
    damage: &HashMap<Team, i32>,
) -> Option<i64> {
    let (walls, mut units) = parse(lines);
    let mut rounds = 0;
    loop {
        match step(&mut units, &walls, damage, elf_deaths_allowed) {
            Ok(true) => {}
            Ok(false) => {
                return Some(total_hp(&units) * rounds);
            }
            Err(_) => {
                return None;
            }
        }
        rounds += 1;
    }
}

fn part_two(lines: &Vec<String>) -> i64 {
    let mut damage_map = HashMap::new();
    let mut damage = 4;
    damage_map.insert(Team::Goblin, 3);
    loop {
        damage_map.insert(Team::Elf, damage);
        match solve(lines, false, &damage_map) {
            Some(result) => {
                return result;
            }
            None => {
                damage += 1;
            }
        }
    }
}

fn part_one(lines: &Vec<String>) -> i64 {
    solve(
        lines,
        true,
        &HashMap::from([(Team::Elf, 3), (Team::Goblin, 3)]),
    )
    .unwrap()
}

fn total_hp(units: &Vec<Unit>) -> i64 {
    units
        .iter()
        .filter_map(|u| if u.hp > 0 { Some(u.hp) } else { None })
        .sum()
}

fn step(
    units: &mut Vec<Unit>,
    walls: &HashSet<Point>,
    damage: &HashMap<Team, i32>,
    elf_deaths_allowed: bool,
) -> Result<bool, String> {
    // Sort the units by position in reading order
    units.sort_by(|a, b| a.pos.cmp(&b.pos));
    for i in 0..units.len() {
        if units[i].hp <= 0 {
            continue;
        }
        match get_move(&units[i], &units, &walls) {
            Some(new_pos) => units[i].pos = new_pos,
            None => return Ok(false),
        }

        let attack = get_attack(&units[i], &units);
        if let Some(t) = attack {
            let unit = units[i].clone();
            let target = units
                .iter_mut()
                .find(|u| u.pos == t.pos && u.hp == t.hp)
                .unwrap();

            target.hp -= damage[&unit.team] as i64;
            if target.hp <= 0 {
                if target.team == Team::Elf && !elf_deaths_allowed {
                    return Err("Elf died".into());
                }
            }
        }
    }
    Ok(true)
}

fn get_attack(unit: &Unit, units: &Vec<Unit>) -> Option<Unit> {
    let units = units
        .iter()
        .filter(|u| u.team != unit.team && u.hp > 0 && unit.pos.distance(&u.pos) == 1)
        .map(|u| u.clone())
        .collect::<Vec<_>>();

    let mut min_hp = i64::MAX;
    units.iter().for_each(|u| {
        if u.hp < min_hp {
            min_hp = u.hp;
        }
    });
    units
        .iter()
        .filter(|u| u.hp == min_hp)
        .cloned()
        .min_by(|a, b| a.pos.cmp(&b.pos))
}

fn get_move(unit: &Unit, units: &Vec<Unit>, walls: &HashSet<Point>) -> Option<Point> {
    let occupied = get_occupied(&unit, &units, &walls);
    let targets = units
        .iter()
        .filter(|u| u.team != unit.team && u.hp > 0)
        .map(|u| u.pos.clone())
        .collect::<HashSet<_>>();
    if targets.is_empty() {
        return None;
    }

    let in_range = adjacent(&targets)
        .difference(&occupied)
        .cloned()
        .collect::<HashSet<_>>();
    let target = choose_target(&unit.pos, &in_range, &occupied);
    if target.is_none() {
        return Some(unit.pos.clone());
    }
    choose_move(&unit.pos, &target.unwrap(), &occupied)
}

fn choose_move(position: &Point, target: &Point, occupied: &HashSet<Point>) -> Option<Point> {
    if position == target {
        return Some(position.clone());
    }
    let paths = shortest_paths(position, &HashSet::from([target.clone()]), occupied);
    let starts = paths
        .iter()
        .map(|p| p.iter().nth(1).unwrap().clone())
        .collect::<HashSet<_>>();

    starts.iter().min().cloned()
}

fn choose_target(
    position: &Point,
    targets: &HashSet<Point>,
    occupied: &HashSet<Point>,
) -> Option<Point> {
    if targets.is_empty() {
        return None;
    }

    if targets.contains(&position) {
        return Some(position.clone());
    }

    let paths = shortest_paths(position, targets, occupied);
    if paths.is_empty() {
        return None;
    }

    let ends = paths
        .iter()
        .map(|p| p.last().unwrap().clone())
        .collect::<HashSet<_>>();

    ends.iter().min().cloned()
}

fn shortest_paths(
    source: &Point,
    targets: &HashSet<Point>,
    occupied: &HashSet<Point>,
) -> Vec<Vec<Point>> {
    let mut results = vec![];
    let mut best: Option<usize> = None;
    let mut visited = occupied.clone();
    let mut queue = BinaryHeap::new();
    queue.push(Reverse((0, vec![source.clone()])));

    while let Some(Reverse((dist, path))) = queue.pop() {
        if best.is_some() && path.len() > best.unwrap() {
            return results;
        }
        let node = path.last().unwrap();

        if targets.contains(node) {
            results.push(path.clone());
            best = Some(path.len());
            continue;
        }

        if visited.contains(node) {
            continue;
        }
        visited.insert(node.clone());

        for neighbor in adjacent(&HashSet::from([node.clone()])) {
            if visited.contains(&neighbor) {
                continue;
            }

            let mut new_path = path.clone();
            new_path.push(neighbor.clone());
            queue.push(Reverse((dist + 1, new_path)));
        }
    }

    results
}

fn adjacent(targets: &HashSet<Point>) -> HashSet<Point> {
    let mut in_range = HashSet::new();
    for t in targets.iter() {
        in_range.insert(Point::new(t.x, t.y - 1)); // Left
        in_range.insert(Point::new(t.x, t.y + 1)); // Right
        in_range.insert(Point::new(t.x - 1, t.y)); // Up
        in_range.insert(Point::new(t.x + 1, t.y)); // Down
    }
    in_range
}

fn get_occupied(unit: &Unit, units: &Vec<Unit>, walls: &HashSet<Point>) -> HashSet<Point> {
    let mut occupied = HashSet::new();
    for u in units.iter() {
        if u.hp > 0 && u != unit {
            occupied.insert(u.pos.clone());
        }
    }

    occupied.union(walls).cloned().collect()
}

fn print_map(units: &Vec<Unit>, walls: &HashSet<Point>) {
    let mut map = vec![];
    for u in units.iter() {
        if u.hp > 0 {
            map.push((u.pos.clone(), u.team));
        }
    }

    let max_x = map.iter().map(|(p, _)| p.x).max().unwrap();
    let max_y = map.iter().map(|(p, _)| p.y).max().unwrap();
    for x in 0..=max_x {
        for y in 0..=max_y {
            let pos = Point::new(x, y);
            let unit = map.iter().find(|(p, _)| p == &pos);
            if let Some((_, team)) = unit {
                match team {
                    Team::Elf => print!("E"),
                    Team::Goblin => print!("G"),
                }
            } else if walls.contains(&pos) {
                print!("#");
            } else {
                print!(".");
            }
        }
        println!();
    }
}

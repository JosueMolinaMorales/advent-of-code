use std::{
    cmp::Reverse,
    collections::{BinaryHeap, HashMap, HashSet, VecDeque},
    path,
};

use crate::utils::{direction::Direction, file_loader::FileLoader, point::Point};

pub fn solve_day_fifteen() {
    let input = FileLoader::new("./inputs/day15.txt".into()).read_lines();
    println!("Day 15 part 1: {}", part_one(&input));
    println!("Day 15 part 2: {}", part_two(&input))
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Clone, Hash)]
enum UnitType {
    Goblin,
    Elf,
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Debug, Clone)]
struct Unit {
    unit_type: UnitType,
    health: i32,
    position: Point,
}

fn part_two(input: &Vec<String>) -> i32 {
    let (grid, units) = parse(input);
    let mut attack_dmg = HashMap::new();
    attack_dmg.insert(UnitType::Goblin, 3);

    let mut elf_died = true;
    let mut elf_attack = 4;
    let mut rounds = 0;
    while elf_died {
        attack_dmg.insert(UnitType::Elf, elf_attack);
        rounds = solve(&units, &grid, &attack_dmg, true);
        elf_attack += 1;
        println!("Rounds: {} -> {}", elf_attack, rounds);

        elf_died = rounds == -1;
    }

    rounds
}

fn parse(input: &Vec<String>) -> (Vec<Vec<char>>, Vec<Unit>) {
    let mut grid = vec![];
    let mut units = vec![];
    input.iter().enumerate().for_each(|(x, line)| {
        let mut row = vec![];
        line.chars().enumerate().for_each(|(y, c)| match c {
            'G' | 'E' => {
                units.push(Unit {
                    unit_type: if c == 'G' {
                        UnitType::Goblin
                    } else {
                        UnitType::Elf
                    },
                    health: 200,
                    position: Point::new(x as i32, y as i32),
                });
                row.push('.');
            }
            '#' => row.push('#'),
            _ => row.push('.'),
        });
        grid.push(row);
    });

    (grid, units)
}

fn solve(
    units: &Vec<Unit>,
    grid: &Vec<Vec<char>>,
    attack_dmg: &HashMap<UnitType, i32>,
    no_elf_death: bool,
) -> i32 {
    let mut units = units.clone();
    let grid = grid.clone();
    let mut round = 0;
    loop {
        // Sort the units by reading order
        units.sort_by(|a, b| {
            if a.position.x == b.position.x {
                a.position.y.cmp(&b.position.y)
            } else {
                a.position.x.cmp(&b.position.x)
            }
        });
        // For every unit, Move then attack
        for i in 0..units.len() {
            let unit = units.get(i).unwrap();
            if unit.health <= 0 {
                // Unit is dead
                if no_elf_death && unit.unit_type == UnitType::Elf {
                    return -1;
                }
                continue;
            }
            // Identify all possible targets.
            let start = std::time::Instant::now();
            let poss_targets = get_possible_targets(&units, unit, &grid);
            // println!(
            //     "Getting Possible Targets Time: {:?}",
            //     start.elapsed().as_millis()
            // );
            // println!("{:?} --pt--> {:?}", unit.0, poss_targets);
            // Check to see if units are already in range
            let adj = units_in_range(unit, &units);
            if adj.is_empty() {
                let start = std::time::Instant::now();
                // Move to the closest target
                match select_move(unit, &units, &grid, &poss_targets) {
                    Some(step) => units.get_mut(i).unwrap().position = step,
                    None => {}
                }

                // println!("Select Move Time: {:?}", start.elapsed().as_millis());
            }
            // Attack
            let unit = units.get(i).unwrap();
            let adj = {
                let adj = units_in_range(unit, &units);
                units
                    .iter()
                    .filter(|p| adj.contains(&p.position) && p.health > 0)
                    .cloned()
                    .collect::<Vec<_>>()
            };
            if adj.is_empty() {
                continue;
            }
            // Attack the unit with the lowest health
            let min = adj.iter().min_by(|a, b| a.health.cmp(&b.health)).unwrap();
            let mut min = units
                .iter()
                .filter(|p| p.position == min.position && p.position != unit.position)
                .collect::<Vec<_>>();
            // Sort by reading order
            min.sort_by(|a, b| {
                if a.position.x == b.position.x {
                    a.position.y.cmp(&b.position.y)
                } else {
                    a.position.x.cmp(&b.position.x)
                }
            });

            let min = min.get(0).unwrap();
            let idx = units
                .iter()
                .position(|p| p.position == min.position)
                .unwrap();
            units.get_mut(idx).unwrap().health -= attack_dmg[&unit.unit_type];
        }
        // If any units died remove them
        units = units.iter().filter(|p| p.health > 0).cloned().collect();

        // Check if the game is over
        let goblins = units
            .iter()
            .filter(|unit| unit.unit_type == UnitType::Goblin)
            .count();
        let elves = units
            .iter()
            .filter(|unit| unit.unit_type == UnitType::Elf)
            .count();
        if goblins == 0 || elves == 0 {
            if no_elf_death && elves == 0 {
                return -1;
            }
            break;
        }

        round += 1;

        // println!("============ ROUND {round} ==============");
        // print_all(&grid, &units);
    }

    // println!(
    //     "Rounds: {round} -> {:?}",
    //     units.iter().map(|unit| unit.health).collect::<Vec<_>>()
    // );

    round * units.iter().map(|unit| unit.health).sum::<i32>()
}

fn part_one(input: &Vec<String>) -> i32 {
    let (grid, units) = parse(input);
    // print_grid(&grid);
    print_all(&grid, &units);

    let mut attack_dmg = HashMap::new();
    attack_dmg.insert(UnitType::Goblin, 3);
    attack_dmg.insert(UnitType::Elf, 3);

    solve(&units, &grid, &attack_dmg, false)
}

fn select_move(
    unit: &Unit,
    units: &Vec<Unit>,
    grid: &Vec<Vec<char>>,
    poss_targets: &Vec<Point>,
) -> Option<Point> {
    let mut dist: HashMap<i32, Vec<Point>> = HashMap::new();
    let mut paths = HashMap::new();
    for target in poss_targets {
        let d = find_min_path(&unit.position, target, grid, units);
        if d.is_empty() {
            continue;
        }
        paths.insert(target.clone(), d.clone());
        let len = d[0].len() as i32;
        dist.entry(len)
            .and_modify(|e| e.push(target.clone()))
            .or_insert(vec![target.clone()]);
    }

    if dist.is_empty() {
        return None;
    }

    // Find the closest target
    let min = dist.keys().min().unwrap();
    let mut min = dist.get(min).unwrap().clone();
    // Sort by reading order
    min.sort_by(|a, b| {
        if a.x == b.x {
            a.y.cmp(&b.y)
        } else {
            a.x.cmp(&b.x)
        }
    });

    // Find the closest step
    let min_steps = paths.get(min.get(0).unwrap()).unwrap();

    Some(min_steps[0][1].clone())
}

fn units_in_range(unit: &Unit, units: &Vec<Unit>) -> Vec<Point> {
    let mut in_range = vec![];
    for dir in [
        Direction::East,
        Direction::West,
        Direction::North,
        Direction::South,
    ] {
        let mut dp: Point = dir.into();
        dp.x += unit.position.x;
        dp.y += unit.position.y;
        if units
            .iter()
            .find(|point| {
                point.position == dp
                    && ((unit.unit_type == UnitType::Goblin && point.unit_type == UnitType::Elf)
                        || (unit.unit_type == UnitType::Elf && point.unit_type == UnitType::Goblin))
                    && unit.health > 0
            })
            .is_some()
        {
            in_range.push(dp);
        }
    }

    in_range
}

fn get_possible_targets(units: &Vec<Unit>, unit: &Unit, grid: &Vec<Vec<char>>) -> Vec<Point> {
    let mut poss_targets = vec![];
    for other in units.iter() {
        if ((unit.unit_type == UnitType::Goblin && other.unit_type == UnitType::Elf)
            || (unit.unit_type == UnitType::Elf && other.unit_type == UnitType::Goblin))
            && unit.health > 0
            && bfs(&unit.position, &other.position, grid, units)
        {
            poss_targets.push(other.position.clone());
        }
    }
    poss_targets
}

fn find_min_path(
    start: &Point,
    end: &Point,
    grid: &Vec<Vec<char>>,
    other_units: &Vec<Unit>,
) -> Vec<Vec<Point>> {
    let mut queue = BinaryHeap::new();
    let mut visited = HashSet::new();
    queue.push(Reverse((0, vec![start.clone()]))); // (dist, path)
    let mut res = vec![];
    let mut best = i32::MAX;
    while !queue.is_empty() {
        let (dist, path) = queue.pop().unwrap().0;
        let curr = path.last().unwrap();

        if curr == end {
            if path.len() as i32 <= best {
                res.push(path.clone());
                best = path.len() as i32;
            }
            continue;
        }

        for neigh in get_adj(&curr, end, grid, other_units) {
            let mut path = path.clone();
            if path.contains(&neigh) {
                continue;
            }
            path.push(neigh.clone());
            if visited.contains(&neigh) {
                continue;
            }
            visited.insert(neigh.clone());
            queue.push(Reverse((dist + 1, path.clone())));
        }
    }

    res
}

fn bfs(start: &Point, find: &Point, grid: &Vec<Vec<char>>, other_units: &Vec<Unit>) -> bool {
    let mut queue = VecDeque::new();
    let mut visited = HashSet::new();

    queue.push_back(start.clone());
    visited.insert(start.clone());

    while !queue.is_empty() {
        let curr = queue.pop_front().unwrap();

        if curr != *start && curr == *find {
            return true;
        }

        for neigh in get_adj(&curr, find, grid, other_units) {
            if visited.contains(&neigh) {
                continue;
            }
            visited.insert(neigh.clone());
            queue.push_back(neigh.clone());
        }
    }

    false
}

fn get_adj(point: &Point, find: &Point, grid: &Vec<Vec<char>>, units: &Vec<Unit>) -> Vec<Point> {
    let directions = vec![
        Direction::South,
        Direction::North,
        Direction::East,
        Direction::West,
    ];

    let mut adj = vec![];
    for dir in directions {
        let change: Point = dir.into();
        let (dx, dy) = (point.x + change.x, point.y + change.y);
        let np = Point::new(dx, dy);
        if dx < 0
            || dy < 0
            || dx >= grid.len() as i32
            || dy >= grid[dx as usize].len() as i32
            || grid[dx as usize][dy as usize] == '#'
            || units
                .iter()
                .find(|point| point.position == np && point.position != *find && point.health > 0)
                .is_some()
        {
            continue;
        }

        adj.push(np)
    }

    adj
}

fn print_grid(grid: &Vec<Vec<char>>) {
    for row in grid.iter() {
        for c in row.iter() {
            print!("{c}")
        }
        println!();
    }
}

fn print_all(grid: &Vec<Vec<char>>, units: &Vec<Unit>) {
    for (x, row) in grid.iter().enumerate() {
        for (y, c) in row.iter().enumerate() {
            match units
                .iter()
                .find(|p| p.position == Point::new(x as i32, y as i32))
            {
                Some(point) => print!(
                    "{}",
                    if point.unit_type == UnitType::Goblin {
                        "G"
                    } else {
                        "E"
                    }
                ),
                None => {
                    print!("{c}")
                }
            }
        }
        println!()
    }
}

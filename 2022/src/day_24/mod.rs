use std::{
    cmp::Reverse,
    collections::{self, BinaryHeap, HashMap, HashSet},
};

pub fn solve_day_24() {
    let input = include_str!("test_input.txt");
    println!("Day 24 Part one: {}", part_one(input));
    // println!("Day 24 Part two: {}", part_two(input));
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy)]
enum Move {
    North,
    South,
    East,
    West,
}

impl Move {
    fn to_point(&self) -> Point {
        match self {
            Move::North => (0, -1),
            Move::South => (0, 1),
            Move::East => (1, 0),
            Move::West => (-1, 0),
        }
    }
}

const DIRECTIONS: [Move; 4] = [Move::North, Move::South, Move::East, Move::West];

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct Blizzard {
    start_position: Point,
    position: Point,
    direction: Move,
}

impl Blizzard {
    fn position_at_time(&self, time: i32, grid: &Vec<Vec<char>>) -> Point {
        let mut position = self.start_position;
        match self.direction {
            Move::North => {
                position.0 = (self.start_position.0 - time) % grid.len() as i32;
            }
            Move::South => {
                position.0 = (self.start_position.0 + time) % grid.len() as i32;
            }
            Move::East => {
                position.1 = (self.start_position.1 + time) % grid[0].len() as i32;
            }
            Move::West => {
                position.1 = (self.start_position.1 - time) % grid[0].len() as i32;
            }
        }
        position
    }
}

type Point = (i32, i32);
struct Map {
    grid: Vec<Vec<char>>,
    blizzards: HashSet<Blizzard>,
    start: Point,
    end: Point,
}

impl Map {
    fn neighbors(&self, point: Point) -> Vec<Point> {
        let mut neighbors = Vec::new();
        for direction in DIRECTIONS.iter() {
            let new_point = (
                point.0 + direction.to_point().0,
                point.1 + direction.to_point().1,
            );
            // Bound Check
            if new_point.0 < 0
                || new_point.0 >= self.grid.len() as i32
                || new_point.1 < 0
                || new_point.1 >= self.grid[0].len() as i32
            {
                continue;
            }

            neighbors.push(new_point);
        }
        neighbors
    }

    fn is_wall(&self, point: Point) -> bool {
        self.grid[point.0 as usize][point.1 as usize] == '#'
    }

    fn is_blizzard_at_point_at_time(&self, point: Point, time: i32) -> bool {
        // We can find the blizzard's point at a time t by doing the following:
        // 1. Find the blizzard's starting point
        // 2. Find the blizzard's direction
        // 3. Find the blizzard's position at time t

        for blizzard in self.blizzards.iter() {
            let blizzard_position = blizzard.position_at_time(time, &self.grid);
            if blizzard_position == point {
                return true;
            }
        }

        false
    }
}

fn part_one(input: &str) -> i32 {
    let map = parse(input);

    0
}

#[derive(PartialEq, Eq, PartialOrd, Hash, Clone)]
struct State {
    position: Point,
    cost: i32,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.cost.cmp(&other.cost)
    }
}

fn dijkstras(map: &Map) {
    let mut heap = BinaryHeap::new();
    heap.push(Reverse(State {
        position: map.start,
        cost: 0,
    }));
    let mut visited = HashSet::new();
    while !heap.is_empty() {
        let curr = heap.pop().unwrap().0;

        // If already visited this cell
        if visited.contains(&curr) {
            continue;
        }

        visited.insert(curr.clone());
        dist.entry(curr.clone()).and_modify(|n| *n = *n + 1);

        let neighbors = map
            .neighbors(curr.position)
            .iter()
            .filter(|n| !map.is_wall(**n))
            .filter(|n| !map.is_blizzard_at_point_at_time(**n, curr.cost))
            .map(|n| *n)
            .collect::<Vec<(i32, i32)>>();

        if neighbors.len() == 0 {
            // Cannot move, stay
        }

        for neighbor in neighbors {
            dist.entry(State {
                position: neighbor,
                cost: curr.cost,
            })
            .and_modify(|cost| {
                if (*cost + 1) < curr.cost {
                    *cost = *cost + 1
                }
            })
            .or_insert(curr.cost + 1);
        }
    }
}

fn parse(input: &str) -> Map {
    let mut start = (0, 0);
    let mut end = (0, 0);
    let grid = input
        .lines()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    start.1 = grid.get(0).unwrap().iter().position(|c| *c == '.').unwrap() as i32;
    end.1 = grid
        .get(grid.len() - 1)
        .unwrap()
        .iter()
        .position(|c| *c == '.')
        .unwrap() as i32;

    // Find all blizzards
    let mut blizzards = HashSet::new();
    for (i, line) in grid.iter().enumerate() {
        for (j, c) in line.iter().enumerate() {
            let mut blizzard = Blizzard {
                start_position: (i as i32, j as i32),
                position: (i as i32, j as i32),
                direction: Move::North,
            };
            match c {
                '>' => {
                    blizzard.direction = Move::East;
                    blizzards.insert(blizzard);
                }
                '<' => {
                    blizzard.direction = Move::West;
                    blizzards.insert(blizzard);
                }
                '^' => {
                    blizzard.direction = Move::North;
                    blizzards.insert(blizzard);
                }
                'v' => {
                    blizzard.direction = Move::South;
                    blizzards.insert(blizzard);
                }
                _ => {}
            }
        }
    }

    Map {
        grid,
        blizzards,
        start,
        end,
    }
}

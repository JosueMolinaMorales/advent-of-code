use std::{
    collections::{BinaryHeap, HashMap, HashSet},
    iter,
};

mod help;
pub fn solve_day_24() {
    let input = include_str!("input.txt");
    println!("Day 24 Part one: {}", help::part_1(input));
    println!("Day 24 Part two: {}", help::part_2(input));
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
            Move::North => Point { row: -1, col: 0 },
            Move::South => Point { row: 1, col: 0 },
            Move::East => Point { row: 0, col: 1 },
            Move::West => Point { row: 0, col: -1 },
        }
    }
}

const DIRECTIONS: [Move; 4] = [Move::North, Move::South, Move::East, Move::West];

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct Blizzard {
    start_position: Point,
    direction: Move,
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy, PartialOrd, Ord)]
struct Point {
    row: i32,
    col: i32,
}

#[derive(Debug)]
struct Map {
    walls: HashSet<Point>,
    blizzards: HashSet<Blizzard>,
    row_len: i32,
    col_len: i32,
    start: Point,
    end: Point,
}

impl Map {
    fn neighbors(&self, point: Point) -> Vec<Point> {
        let mut neighbors = Vec::new();
        for direction in DIRECTIONS.iter() {
            let new_point = (
                point.row + direction.to_point().row,
                point.col + direction.to_point().col,
            );
            // Bound Check
            if new_point.0 < 0
                || new_point.0 >= self.row_len
                || new_point.1 < 0
                || new_point.1 >= self.col_len
            {
                continue;
            }

            neighbors.push(Point {
                row: new_point.0,
                col: new_point.1,
            });
        }
        neighbors
    }

    fn is_wall(&self, point: Point) -> bool {
        self.walls.contains(&point)
    }
}

fn part_one(input: &str) -> i32 {
    let map = parse(input);
    let lcm = lcm(map.row_len as usize - 2, map.col_len as usize - 2);
    println!("lcm: {}", lcm);
    let bliz_maps = bliz_maps(&map, map.row_len, map.col_len, lcm);
    // println!("bliz_maps: {:#?}", bliz_maps);
    let min_cost = dijkstras(&map, lcm, bliz_maps);
    min_cost
}

#[derive(PartialEq, Eq, PartialOrd, Hash, Clone, Debug)]
struct State {
    position: Point,
    cost: i32,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.cost.cmp(&self.cost)
    }
}

fn dijkstras(map: &Map, lcm: usize, bliz_maps: HashMap<i32, HashSet<Point>>) -> i32 {
    let mut heap = BinaryHeap::new();
    heap.push(State {
        position: map.start,
        cost: 0,
    });
    let mut visited = HashSet::new();
    visited.insert((map.start, 0));
    while let Some(state) = heap.pop() {
        if state.position == map.end {
            return state.cost;
        }

        // Each step is a cost of 1
        let new_cost = state.cost + 1;
        if new_cost % 10 == 0 {
            println!("new_cost: {}", new_cost);
        }
        let blizzards = &bliz_maps[&((new_cost as usize % lcm) as i32)];
        let neighbors = map.neighbors(state.position);

        let neighbors = neighbors
            .iter()
            .chain(iter::once(&state.position))
            .filter(|point| !map.is_wall(**point))
            .filter(|point| !blizzards.contains(point))
            .collect::<Vec<_>>();

        for new_pos in neighbors {
            if visited.insert((*new_pos, new_cost)) {
                heap.push(State {
                    position: *new_pos,
                    cost: new_cost,
                });
            }
        }
    }

    0
}

fn bliz_maps(map: &Map, rows: i32, cols: i32, max_time: usize) -> HashMap<i32, HashSet<Point>> {
    // key: turn, val: set of a bliz locations
    let mut cache = HashMap::new();

    let mut blizzards: Vec<(Point, Move)> = map
        .blizzards
        .iter()
        .map(|blizzard| (blizzard.start_position, blizzard.direction))
        .collect();

    let coords = blizzards.iter().map(|(coord, _)| *coord).collect();
    cache.insert(0, coords);

    // precompute every blizzard coord at every time before the coords repeat
    for time in 1..max_time {
        for (coord, dir) in blizzards.iter_mut() {
            let d = dir.to_point();
            *coord = Point {
                row: coord.row + d.row,
                col: coord.col + d.col,
            };
            // if next coord went to an edge, wrap
            match dir {
                Move::West => {
                    if coord.col == 0 {
                        coord.col = cols - 2;
                    }
                }
                Move::East => {
                    if coord.col == cols - 1 {
                        coord.col = 1;
                    }
                }
                Move::North => {
                    if coord.row == 0 {
                        coord.row = rows - 2;
                    }
                }
                Move::South => {
                    if coord.row == rows - 1 {
                        coord.row = 1;
                    }
                }
            }
        }
        let coords = blizzards.iter().map(|(coord, _)| *coord).collect();
        cache.insert(time as i32, coords);
    }

    cache
}

fn parse(input: &str) -> Map {
    let mut start = Point { row: 0, col: 0 };
    let mut end = Point { row: 0, col: 0 };
    let mut walls = HashSet::new();
    let mut blizzards = HashSet::new();

    for (row, line) in input.lines().enumerate() {
        for (col, c) in line.chars().enumerate() {
            if row == 0 && c == '.' {
                start = Point {
                    row: row as i32,
                    col: col as i32,
                };
            }
            if row == input.lines().count() - 1 && c == '.' {
                end = Point {
                    row: row as i32,
                    col: col as i32,
                };
            }
            match c {
                '#' => {
                    walls.insert(Point {
                        row: row as i32,
                        col: col as i32,
                    });
                }
                '>' => {
                    blizzards.insert(Blizzard {
                        start_position: Point {
                            row: row as i32,
                            col: col as i32,
                        },
                        direction: Move::East,
                    });
                }
                '<' => {
                    blizzards.insert(Blizzard {
                        start_position: Point {
                            row: row as i32,
                            col: col as i32,
                        },
                        direction: Move::West,
                    });
                }
                '^' => {
                    blizzards.insert(Blizzard {
                        start_position: Point {
                            row: row as i32,
                            col: col as i32,
                        },
                        direction: Move::North,
                    });
                }
                'v' => {
                    blizzards.insert(Blizzard {
                        start_position: Point {
                            row: row as i32,
                            col: col as i32,
                        },
                        direction: Move::South,
                    });
                }
                _ => {}
            }
        }
    }

    let row_len = input.lines().count() as i32;
    let col_len = input.lines().next().unwrap().chars().count() as i32;

    Map {
        walls,
        blizzards,
        row_len,
        col_len,
        start,
        end,
    }
}

fn lcm(first: usize, second: usize) -> usize {
    first * second / gcd(first, second)
}

fn gcd(first: usize, second: usize) -> usize {
    let mut max = first;
    let mut min = second;
    if min > max {
        std::mem::swap(&mut max, &mut min);
    }

    loop {
        let res = max % min;
        if res == 0 {
            return min;
        }

        max = min;
        min = res;
    }
}

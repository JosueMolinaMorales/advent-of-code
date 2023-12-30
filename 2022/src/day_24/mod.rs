use std::{
    collections::{BinaryHeap, HashMap, HashSet},
    iter,
};

pub fn solve_day_24() {
    let input = include_str!("input.txt");
    println!("Day 24 Part one: {}", part_one(input));
    println!("Day 24 Part two: {}", part_two(input));
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy)]
enum Direction {
    North,
    South,
    East,
    West,
}

enum Tile {
    Wall,
    Blizzard(Direction),
}

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct Blizzard {
    start_position: Point,
    direction: Direction,
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy, PartialOrd, Ord)]
struct Point {
    row: usize,
    col: usize,
}

impl Point {
    fn neighbors(&self, rows: usize, cols: usize) -> Vec<Point> {
        let mut neighbors = Vec::new();
        if self.row > 0 {
            neighbors.push(self.add_dir(Direction::North));
        }
        if self.row < rows - 1 {
            neighbors.push(self.add_dir(Direction::South));
        }
        if self.col > 0 {
            neighbors.push(self.add_dir(Direction::West));
        }
        if self.col < cols - 1 {
            neighbors.push(self.add_dir(Direction::East));
        }
        neighbors
    }

    fn add_dir(&self, dir: Direction) -> Point {
        match dir {
            Direction::North => Point {
                row: self.row - 1,
                col: self.col,
            },
            Direction::South => Point {
                row: self.row + 1,
                col: self.col,
            },
            Direction::East => Point {
                row: self.row,
                col: self.col + 1,
            },
            Direction::West => Point {
                row: self.row,
                col: self.col - 1,
            },
        }
    }
}

struct MapInfo {
    rows: usize,
    cols: usize,
    walls: HashSet<Point>,
    repeats_at: usize,
    blizzard_maps: HashMap<usize, HashSet<Point>>,
}

#[derive(PartialEq, Eq, Hash, Clone, Debug)]
struct State {
    position: Point,
    cost: usize,
}

impl Ord for State {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        other.cost.cmp(&self.cost)
    }
}

impl PartialOrd for State {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        Some(self.cmp(other))
    }
}

fn part_one(input: &str) -> i32 {
    let (map, walls, row, col) = parse(input);
    let lcm = lcm(row - 2, col - 2);

    let bliz_maps = bliz_maps(&map, row, col, lcm);

    let map_info = MapInfo {
        rows: row,
        cols: col,
        walls,
        repeats_at: lcm,
        blizzard_maps: bliz_maps,
    };

    let start = Point { row: 0, col: 1 };
    let end = Point {
        row: row - 1,
        col: col - 2,
    };

    let min_cost = dijkstras(start, end, 0, &map_info);
    min_cost as i32
}

fn part_two(input: &str) -> i32 {
    let (map, walls, row, col) = parse(input);

    let lcm = lcm(row - 2, col - 2);
    let blizzard_maps = bliz_maps(&map, row, col, lcm);
    let start = Point { row: 0, col: 1 };
    let end = Point {
        row: row - 1,
        col: col - 2,
    };

    let map_info = MapInfo {
        rows: row,
        cols: col,
        walls,
        repeats_at: lcm,
        blizzard_maps,
    };

    let there = dijkstras(start, end, 0, &map_info);
    let back = dijkstras(end, start, there, &map_info);
    dijkstras(start, end, back, &map_info) as i32
}

fn dijkstras(from: Point, to: Point, start_time: usize, map_info: &MapInfo) -> usize {
    let mut heap = BinaryHeap::new();
    let mut visited = HashSet::new();

    heap.push(State {
        position: from,
        cost: start_time,
    });
    visited.insert((from, start_time));

    while let Some(State { position, cost }) = heap.pop() {
        if position == to {
            return cost;
        }

        // Each step is a cost of 1
        let new_cost = cost + 1;

        let blizzards = &map_info.blizzard_maps[&(new_cost % map_info.repeats_at)];

        let neighbors = position
            .neighbors(map_info.rows, map_info.cols)
            .into_iter()
            .chain(iter::once(position))
            .filter(|point| !map_info.walls.contains(point))
            .filter(|point| !blizzards.contains(point));

        for new_pos in neighbors {
            if visited.insert((new_pos, new_cost)) {
                heap.push(State {
                    position: new_pos,
                    cost: new_cost,
                });
            }
        }
    }

    0
}

fn bliz_maps(
    map: &HashMap<Point, Tile>,
    rows: usize,
    cols: usize,
    max_time: usize,
) -> HashMap<usize, HashSet<Point>> {
    // key: turn, val: set of a bliz locations
    let mut cache = HashMap::new();

    let mut blizzards: Vec<(Point, Direction)> = map
        .iter()
        .filter_map(|(pos, tile)| match tile {
            Tile::Blizzard(dir) => Some((*pos, *dir)),
            _ => None,
        })
        .collect();

    let coords = blizzards.iter().map(|(coord, _)| *coord).collect();
    cache.insert(0, coords);

    // precompute every blizzard coord at every time before the coords repeat
    for time in 1..max_time {
        for (coord, dir) in blizzards.iter_mut() {
            *coord = coord.add_dir(*dir);
            // if next coord went to an edge, wrap
            match dir {
                Direction::West => {
                    if coord.col == 0 {
                        coord.col = cols - 2;
                    }
                }
                Direction::East => {
                    if coord.col == cols - 1 {
                        coord.col = 1;
                    }
                }
                Direction::North => {
                    if coord.row == 0 {
                        coord.row = rows - 2
                    }
                }
                Direction::South => {
                    if coord.row == rows - 1 {
                        coord.row = 1;
                    }
                }
            }
        }
        let coords = blizzards.iter().map(|(coord, _)| *coord).collect();
        cache.insert(time, coords);
    }

    cache
}

fn parse(input: &str) -> (HashMap<Point, Tile>, HashSet<Point>, usize, usize) {
    let mut map = HashMap::new();
    for (row, line) in input.lines().enumerate() {
        for (col, c) in line.chars().enumerate() {
            if c == '.' {
                continue;
            }
            let point = Point { row, col };
            let tile = match c {
                '#' => Tile::Wall,
                '^' => Tile::Blizzard(Direction::North),
                'v' => Tile::Blizzard(Direction::South),
                '>' => Tile::Blizzard(Direction::East),
                '<' => Tile::Blizzard(Direction::West),
                _ => panic!("Invalid tile"),
            };
            map.insert(point, tile);
        }
    }

    let row_len = input.lines().count();
    let col_len = input.lines().next().unwrap().chars().count();

    let walls = map
        .iter()
        .filter_map(|(pos, tile)| match tile {
            Tile::Wall => Some(*pos),
            _ => None,
        })
        .collect();

    (map, walls, row_len, col_len)
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

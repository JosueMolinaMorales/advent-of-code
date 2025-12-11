use std::collections::HashMap;

use crate::utils::{file_loader::FileLoader, point::Point};

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
enum Tile {
    Open,
    Trees,
    Lumber,
}

pub fn solve_day_eighteen() {
    let lines = FileLoader::new("inputs/day_18.txt".into()).read_lines();
    let grid = parse_grid(&lines);

    println!("Part 1: {}", part_one(&grid));
    println!("Part 2: {}", part_two(&grid));
}

fn parse_grid(lines: &[String]) -> Vec<Vec<Tile>> {
    lines
        .iter()
        .map(|line| {
            line.chars()
                .map(|c| match c {
                    '.' => Tile::Open,
                    '|' => Tile::Trees,
                    '#' => Tile::Lumber,
                    _ => panic!("Invalid character: {}", c),
                })
                .collect()
        })
        .collect()
}

fn part_one(grid: &[Vec<Tile>]) -> usize {
    let mut grid = grid.to_vec();
    run_simulation(&mut grid, 10);
    calculate_resource_value(&grid)
}

fn part_two(grid: &[Vec<Tile>]) -> usize {
    let mut grid = grid.to_vec();
    run_simulation_with_cycle(&mut grid, 1_000_000_000);
    calculate_resource_value(&grid)
}

fn run_simulation(grid: &mut Vec<Vec<Tile>>, minutes: usize) {
    let mut updated_grid = grid.clone();

    for _ in 0..minutes {
        for (i, row) in grid.iter().enumerate() {
            for (j, cell) in row.iter().enumerate() {
                let point = Point::new(i as i32, j as i32);
                updated_grid[i][j] = apply_rules(&point, cell, grid);
            }
        }
        std::mem::swap(grid, &mut updated_grid);
    }
}

fn run_simulation_with_cycle(grid: &mut Vec<Vec<Tile>>, target_minute: usize) {
    let mut updated_grid = grid.clone();
    let mut seen: HashMap<String, usize> = HashMap::new();

    for minute in 0..target_minute {
        for (i, row) in grid.iter().enumerate() {
            for (j, cell) in row.iter().enumerate() {
                let point = Point::new(i as i32, j as i32);
                updated_grid[i][j] = apply_rules(&point, cell, grid);
            }
        }
        std::mem::swap(grid, &mut updated_grid);

        let key = grid_to_string(grid);

        if let Some(&previous_minute) = seen.get(&key) {
            let cycle_length = minute - previous_minute;
            let remaining = target_minute - minute - 1;
            let position_in_cycle = remaining % cycle_length;
            run_simulation(grid, position_in_cycle);
            return;
        }

        seen.insert(key, minute);
    }
}

fn apply_rules(point: &Point, tile: &Tile, grid: &[Vec<Tile>]) -> Tile {
    match tile {
        Tile::Lumber => apply_lumber_condition(point, grid),
        Tile::Open => apply_open_area_condition(point, grid),
        Tile::Trees => apply_trees_condition(point, grid),
    }
}

fn calculate_resource_value(grid: &[Vec<Tile>]) -> usize {
    let (trees, lumber) = grid
        .iter()
        .flatten()
        .fold((0, 0), |(trees, lumber), tile| match tile {
            Tile::Trees => (trees + 1, lumber),
            Tile::Lumber => (trees, lumber + 1),
            _ => (trees, lumber),
        });
    trees * lumber
}

fn grid_to_string(grid: &[Vec<Tile>]) -> String {
    grid.iter()
        .flatten()
        .map(|tile| match tile {
            Tile::Lumber => '#',
            Tile::Open => '.',
            Tile::Trees => '|',
        })
        .collect()
}

fn apply_lumber_condition(point: &Point, grid: &[Vec<Tile>]) -> Tile {
    let adjacent = count_adjacent_tiles(point, grid);
    if adjacent[&Tile::Lumber] >= 1 && adjacent[&Tile::Trees] >= 1 {
        Tile::Lumber
    } else {
        Tile::Open
    }
}

fn apply_open_area_condition(point: &Point, grid: &[Vec<Tile>]) -> Tile {
    let adjacent = count_adjacent_tiles(point, grid);
    if adjacent[&Tile::Trees] >= 3 {
        Tile::Trees
    } else {
        Tile::Open
    }
}

fn apply_trees_condition(point: &Point, grid: &[Vec<Tile>]) -> Tile {
    let adjacent = count_adjacent_tiles(point, grid);
    if adjacent[&Tile::Lumber] >= 3 {
        Tile::Lumber
    } else {
        Tile::Trees
    }
}

fn count_adjacent_tiles(point: &Point, grid: &[Vec<Tile>]) -> HashMap<Tile, usize> {
    const DIRECTIONS: [(i32, i32); 8] = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];

    let mut counts = HashMap::from([(Tile::Trees, 0), (Tile::Lumber, 0), (Tile::Open, 0)]);
    let (rows, cols) = (grid.len() as i32, grid[0].len() as i32);

    for (dx, dy) in DIRECTIONS {
        let (x, y) = (point.x + dx, point.y + dy);
        if x >= 0 && x < rows && y >= 0 && y < cols {
            *counts.entry(grid[x as usize][y as usize]).or_insert(0) += 1;
        }
    }

    counts
}

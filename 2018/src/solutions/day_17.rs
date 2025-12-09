use regex::Regex;
use std::collections::{HashMap, HashSet};

use crate::utils::file_loader::FileLoader;

const SAMPLE_INPUT: &str = r#"x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504"#;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
enum Tile {
    Clay,
    Sand,
    FlowingWater, // |
    SettledWater, // ~
}

pub fn solve_day_seventeen() {
    let lines = FileLoader::new("./inputs/day_17.txt".into()).read_lines();

    println!("Part 1: {}", part_one(&lines));
    println!("Part 2: {}", part_two(&lines));
}

fn part_one(lines: &[String]) -> usize {
    let (grid, _clay_positions, min_y, max_y) = parse_input_and_simulate(lines);

    // Count all water tiles (both flowing and settled) within y bounds
    grid.iter()
        .filter(|&(&(_, y), tile)| {
            y >= min_y && y <= max_y && (*tile == Tile::FlowingWater || *tile == Tile::SettledWater)
        })
        .count()
}

fn part_two(lines: &[String]) -> usize {
    let (grid, _, min_y, max_y) = parse_input_and_simulate(lines);

    // Count only settled water within y bounds
    grid.iter()
        .filter(|&(&(_, y), tile)| y >= min_y && y <= max_y && *tile == Tile::SettledWater)
        .count()
}

fn parse_input_and_simulate(
    lines: &[String],
) -> (HashMap<(i32, i32), Tile>, HashSet<(i32, i32)>, i32, i32) {
    let re = Regex::new(r"(\w)=(\d+), (\w)=(\d+)..(\d+)").unwrap();
    let mut clay_positions = HashSet::new();

    for line in lines {
        if let Some(caps) = re.captures(line) {
            let first_coord = caps[1].chars().next().unwrap();
            let first_val: i32 = caps[2].parse().unwrap();
            let second_start: i32 = caps[4].parse().unwrap();
            let second_end: i32 = caps[5].parse().unwrap();

            if first_coord == 'x' {
                for y in second_start..=second_end {
                    clay_positions.insert((first_val, y));
                }
            } else {
                for x in second_start..=second_end {
                    clay_positions.insert((x, first_val));
                }
            }
        }
    }

    // Find bounds (only y bounds matter for counting)
    let min_y = *clay_positions.iter().map(|(_, y)| y).min().unwrap();
    let max_y = *clay_positions.iter().map(|(_, y)| y).max().unwrap();

    // Create grid with clay positions
    let mut grid: HashMap<(i32, i32), Tile> = HashMap::new();
    for &pos in &clay_positions {
        grid.insert(pos, Tile::Clay);
    }

    // Start water flow simulation from spring at (500, 0)
    flow(500, 0, &mut grid, max_y);

    (grid, clay_positions, min_y, max_y)
}

fn flow(x: i32, y: i32, grid: &mut HashMap<(i32, i32), Tile>, max_y: i32) -> bool {
    // If we're beyond max depth, stop
    if y > max_y {
        return false;
    }

    let pos = (x, y);

    // If this position already has water or clay, we're done
    if let Some(tile) = grid.get(&pos) {
        match tile {
            Tile::Clay => return true,          // Hit a solid surface
            Tile::FlowingWater => return false, // Already visited, prevent infinite loops
            Tile::SettledWater => return true,  // Hit settled water, treat as solid
            _ => {}
        }
    }

    // Mark this position as flowing water
    grid.insert(pos, Tile::FlowingWater);

    // Try to flow down first
    let below_is_solid = flow(x, y + 1, grid, max_y);

    if !below_is_solid {
        // Water flows down freely, can't settle here
        return false;
    }

    // Water can't flow down, try to spread horizontally
    let left_is_solid = flow_horizontal(x, y, -1, grid, max_y);
    let right_is_solid = flow_horizontal(x, y, 1, grid, max_y);

    if left_is_solid && right_is_solid {
        // Water is contained on both sides, settle it
        settle_row(x, y, grid);
        return true;
    }

    // Water flows off at least one side, stays as flowing water
    false
}

fn flow_horizontal(
    x: i32,
    y: i32,
    dx: i32,
    grid: &mut HashMap<(i32, i32), Tile>,
    max_y: i32,
) -> bool {
    let mut current_x = x + dx;

    loop {
        let pos = (current_x, y);

        // Check if we hit clay
        if let Some(Tile::Clay) = grid.get(&pos) {
            return true; // Hit a wall
        }

        // Check if this position already has settled water (acts as a wall)
        if let Some(Tile::SettledWater) = grid.get(&pos) {
            return true;
        }

        // Mark as flowing water
        grid.insert(pos, Tile::FlowingWater);

        // Check what's below this position
        let below = (current_x, y + 1);

        // Check if there's something solid below
        match grid.get(&below) {
            Some(Tile::Clay) | Some(Tile::SettledWater) => {
                // Solid below, continue spreading horizontally
                current_x += dx;
            }
            _ => {
                // Nothing solid below, water falls down
                flow(current_x, y + 1, grid, max_y);

                // After flowing down, check if it settled and now provides support
                match grid.get(&below) {
                    Some(Tile::SettledWater) => {
                        // Water below settled, continue spreading
                        current_x += dx;
                    }
                    _ => {
                        // Water flows off the edge
                        return false;
                    }
                }
            }
        }
    }
}

fn settle_row(x: i32, y: i32, grid: &mut HashMap<(i32, i32), Tile>) {
    // Settle water from x to the left until we hit clay
    let mut current_x = x;
    loop {
        let pos = (current_x, y);
        if let Some(Tile::Clay) = grid.get(&pos) {
            break;
        }
        grid.insert(pos, Tile::SettledWater);
        current_x -= 1;
    }

    // Settle water from x to the right until we hit clay
    current_x = x + 1;
    loop {
        let pos = (current_x, y);
        if let Some(Tile::Clay) = grid.get(&pos) {
            break;
        }
        grid.insert(pos, Tile::SettledWater);
        current_x += 1;
    }
}

#[allow(dead_code)]
fn print_grid(grid: &HashMap<(i32, i32), Tile>, clay_positions: &HashSet<(i32, i32)>) {
    let min_x = clay_positions.iter().map(|(x, _)| x).min().unwrap() - 2;
    let max_x = clay_positions.iter().map(|(x, _)| x).max().unwrap() + 2;
    let min_y = clay_positions.iter().map(|(_, y)| y).min().unwrap() - 1;
    let max_y = clay_positions.iter().map(|(_, y)| y).max().unwrap() + 1;

    for y in min_y..=max_y {
        for x in min_x..=max_x {
            if x == 500 && y == 0 {
                print!("+");
            } else {
                match grid.get(&(x, y)) {
                    Some(Tile::Clay) => print!("#"),
                    Some(Tile::FlowingWater) => print!("|"),
                    Some(Tile::SettledWater) => print!("~"),
                    _ => print!("."),
                }
            }
        }
        println!();
    }
}

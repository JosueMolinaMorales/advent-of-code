use core::panic;
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
    let mut grid: Vec<Vec<Tile>> = lines
        .iter()
        .map(|l| {
            l.split("")
                .filter(|c| c.len() > 0)
                .map(|c| match c {
                    "." => Tile::Open,
                    "|" => Tile::Trees,
                    "#" => Tile::Lumber,
                    _ => panic!(""),
                })
                .collect::<Vec<Tile>>()
        })
        .collect();

    print_grid(&grid);

    let mut updated_grid = grid.clone();
    for i in 0..10 {
        // Update the state of each tile
        for (i, row) in grid.iter().enumerate() {
            for (j, cell) in row.iter().enumerate() {
                let p = &Point::new(i as i32, j as i32);
                let new_cell = match cell {
                    &Tile::Lumber => apply_lumber_condition(p, &grid),
                    &Tile::Open => apply_open_area_condition(p, &grid),
                    &Tile::Trees => apply_trees_condition(p, &grid),
                };
                updated_grid[i][j] = new_cell;
            }
        }
        grid = updated_grid.clone();

        println!("=========== MINUTE {} ===============", i + 1);
        print_grid(&grid);
        println!()
    }

    // Count
    let mut wooded = 0;
    let mut lumber = 0;

    for row in &grid {
        for col in row {
            if col == &Tile::Trees {
                wooded += 1;
            } else if col == &Tile::Lumber {
                lumber += 1;
            }
        }
    }

    println!("Part 1: {}", wooded * lumber)
}

fn print_grid(grid: &Vec<Vec<Tile>>) {
    for row in grid {
        for col in row {
            match col {
                &Tile::Lumber => print!("#"),
                &Tile::Open => print!("."),
                &Tile::Trees => print!("|"),
            }
        }
        println!();
    }
}

fn apply_lumber_condition(coords: &Point, grid: &Vec<Vec<Tile>>) -> Tile {
    // An acre containing a lumberyard will remain a lumberyard if
    // it was adjacent to at least one other lumberyard and at least one acre containing trees.
    // Otherwise, it becomes open.
    let adj = get_adj_points(coords, grid);

    if *adj.get(&Tile::Lumber).unwrap() >= 1 && *adj.get(&Tile::Trees).unwrap() >= 1 {
        return Tile::Lumber;
    }

    Tile::Open
}
fn apply_open_area_condition(coords: &Point, grid: &Vec<Vec<Tile>>) -> Tile {
    let adj = get_adj_points(coords, grid);

    if *adj.get(&Tile::Trees).unwrap() >= 3 {
        return Tile::Trees;
    }

    Tile::Open
}
fn apply_trees_condition(coords: &Point, grid: &Vec<Vec<Tile>>) -> Tile {
    let adj = get_adj_points(coords, grid);
    if *adj.get(&Tile::Lumber).unwrap() >= 3 {
        return Tile::Lumber;
    }

    Tile::Trees
}

fn get_adj_points(curr: &Point, grid: &Vec<Vec<Tile>>) -> HashMap<Tile, usize> {
    let directions = vec![
        (-1, 0),  // Up
        (-1, -1), // Up-Left
        (1, 0),   //  Down
        (1, -1),  // Down-Left
        (-1, 1),  // Up Right
        (1, 1),   // Down Right
        (0, 1),   // Right
        (0, -1),  // Left
    ];

    let mut adj = HashMap::from([(Tile::Trees, 0), (Tile::Lumber, 0), (Tile::Open, 0)]);
    for (dx, dy) in directions {
        let (x, y) = (curr.x + dx, curr.y + dy);
        //bounds
        if x < 0 || x >= grid.len() as i32 || y < 0 || y >= grid[0].len() as i32 {
            continue;
        }

        adj.entry(grid[x as usize][y as usize])
            .and_modify(|c| *c += 1);
    }

    return adj;
}

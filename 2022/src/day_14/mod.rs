use std::{collections::HashMap, vec};

const INPUT: &str = include_str!("day_14_input.txt");

enum Part {
    One,
    Two
}

fn parse_input_and_create_grid<'a>(part: Part) -> (Vec<Vec<&'a str>>, isize) {
    let mut points: HashMap<(isize, isize), &str> = HashMap::new();
    let mut max_col = 0;
    let mut max_row = 0;
    let mut min_col = isize::MAX;
    // Draw grid
    INPUT
        .lines()
        .for_each(|lines| {
            let pairs = lines
                .split(" -> ")
                .map(|p| {
                    let mut pair = p.split(",");
                    (pair.next().unwrap().parse::<isize>().unwrap(), pair.next().unwrap().parse::<isize>().unwrap())
                }).collect::<Vec<(isize, isize)>>();

            let mut prev_pair = pairs[0];
            for (i, pair) in pairs.iter().enumerate() {
                if pair.0 > max_col {
                    max_col = pair.0;
                }
                if pair.0 < min_col {
                    min_col = pair.0;
                }
                if pair.1 > max_row {
                    max_row = pair.1;
                }
                if i == 0 {
                    points.insert(*pair, "#");
                    continue;
                }
                let dx = pair.0 - prev_pair.0;
                let dy = pair.1 - prev_pair.1;
                for j in 0..=dx.abs() {
                    points.insert(
                        (prev_pair.0 + (j * dx.signum()) , prev_pair.1),
                        "#"
                    );
                }
                for j in 0..=dy.abs() {
                    points.insert(
                        (prev_pair.0, prev_pair.1 + (j * dy.signum())), 
                        "#"
                    );
                }
                prev_pair = *pair;
            }
        });
    match part {
        Part::One => (create_grid(&points, min_col, max_row, max_col, false), min_col),
        Part::Two => {
            min_col = min_col - min_col;
            max_col = max_col + max_col;
            (create_grid(&points, min_col, max_row + 2, max_col, true), min_col)
        }
    }
}

pub fn solve_day_14() {
    solve_part_one();
    solve_part_two()
}

fn solve_part_two() {
    let res = parse_input_and_create_grid(Part::Two);
    let mut grid = res.0;
    // drop sand
    let sand_drop_point: (usize, usize) = (500, 0);
    let origin = (res.1, 0);
    let mut sand_count = 0;
    // Drop in the y direction
    'outer: loop {
        let mut x = sand_drop_point.1 + 1;
        let mut y = sand_drop_point.0 - origin.0 as usize;
        // Drop Sand till it hits something
        loop {
            let current_drop = grid[x][y];
            if grid[sand_drop_point.1][sand_drop_point.0] == "o" {
                break 'outer;
            }
            if current_drop == "o" || current_drop == "#" {
                x -= 1;
                // Obstacle hit, attempt to go diagonally left
                let diag = grid[x + 1][y - 1 as usize];
                if diag == "."{
                    x = x + 1;
                    y = y - 1 as usize;
                    continue
                }
                // Obstacle hit, attempt to go diagonally right
                let diag = grid[x + 1][y + 1];
                if diag == "." {
                    x = x + 1;
                    y = y + 1 as usize;
                    continue
                }
                break;
            }
            
            x += 1
        }
        grid[x][y] = "o";
        sand_count += 1;
    }
    print_grid(&grid);
    println!("Sand count: {}", sand_count);
}

fn create_grid<'a>(
    points: &HashMap<(isize, isize), &'a str>,
    lowest_x: isize,
    lowest_y: isize,
    greatest_x: isize,
    add_floor: bool
) -> Vec<Vec<&'a str>> {
    let mut grid: Vec<Vec<&str>> = vec![];
    let mut column_count = 0;
    for j in 0..=lowest_y {
        let mut row = vec![];
        for i in lowest_x..=greatest_x {
            if i == 500 && column_count == 0 {
                row.push("+");
                continue;
            }
            if j == lowest_y && add_floor {
                row.push("#");
                continue;
            }
            row.push(*points.get(&(i, column_count)).unwrap_or(&"."))
        }
        grid.push(row);
        column_count += 1
    }
    grid
}

fn print_grid(grid: &Vec<Vec<&str>>) {
    for i in grid.iter().enumerate() {
        for j in i.1.iter().enumerate() {
            if j.0 == 0 {
                print!("{}  ", i.0);
            }
            print!("{}", j.1);
        }
        println!();
    }
}


fn solve_part_one() {
    let res = parse_input_and_create_grid(Part::One);
    let mut grid = res.0;
    // drop sand
    let sand_drop_point: (usize, usize) = (500, 0);
    let origin = (res.1, 0);
    let mut sand_count = 0;
    // Drop in the y direction
    'outer: loop {
        let mut x = sand_drop_point.1;
        let mut y = sand_drop_point.0 - origin.0 as usize;
        // Drop Sand till it hits something
        loop {
            if x >= grid.len() || y == 0 {
                break 'outer;
            }
            let current_drop = grid[x][y];
            if current_drop == "o" || current_drop == "#" {
                x -= 1;
                // Obstacle hit, attempt to go diagonally left
                let diag = grid[x + 1][y - 1 as usize];
                if diag == "."{
                    x = x + 1;
                    y = y - 1 as usize;
                    continue
                }
                // Obstacle hit, attempt to go diagonally right
                let diag = grid[x + 1][y + 1];
                if diag == "." {
                    x = x + 1;
                    y = y + 1 as usize;
                    continue
                }
                // Obstacle hit, stay in current position
                break;
            }
            x += 1
        }
        grid[x][y] = "o";
        sand_count += 1;
    }
    print_grid(&grid);
    println!("Sand count: {}", sand_count);
}
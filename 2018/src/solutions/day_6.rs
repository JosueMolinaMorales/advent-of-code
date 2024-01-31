use std::collections::{HashMap, HashSet};

use crate::utils::{file_loader::FileLoader, point::Point};

pub fn solve_day_six() {
    let input = FileLoader::new("./inputs/day6.txt".into()).read_lines();
    println!("Day 6 part 1: {}", part_one(input.clone()));
    println!("Day 6 part 2: {}", part_two(input));
}

fn part_one(coords: Vec<String>) -> i32 {
    let (coords, mut grid) = parse(coords);

    // For each point in the grid, find the closest coordinate
    for (y, row) in grid.iter_mut().enumerate() {
        for (x, point) in row.iter_mut().enumerate() {
            let mut min_distance = std::i32::MAX;
            let mut min_coord = 0;
            for (i, point) in coords.iter().enumerate() {
                let distance = point.distance(&Point::new(x as i32, y as i32));
                if distance < min_distance {
                    min_distance = distance;
                    min_coord = i as i32;
                } else if distance == min_distance {
                    min_coord = -1;
                }
            }
            *point = min_coord as i32;
        }
    }

    // Find the largest area
    let mut areas = HashMap::new();
    for row in grid.iter() {
        for point in row.iter() {
            if *point != -1 {
                areas.entry(point).and_modify(|e| *e += 1).or_insert(1);
            }
        }
    }

    let mut infinite_areas = HashMap::new();
    for (i, row) in grid.iter().enumerate() {
        for (j, point) in row.iter().enumerate() {
            if i == 0 || j == 0 || i == grid.len() - 1 || j == row.len() - 1 {
                infinite_areas.entry(point).or_insert(true);
            }
        }
    }

    areas
        .iter()
        .filter(|(k, _)| !infinite_areas.contains_key(*k))
        .map(|(_, v)| v)
        .max()
        .unwrap()
        .clone()
}

fn part_two(input: Vec<String>) -> i32 {
    let (coords, grid) = parse(input);
    let mut safe_points = HashSet::new();
    for (y, row) in grid.iter().enumerate() {
        for (x, _) in row.iter().enumerate() {
            // For every point, sum the dist to all points
            let point = Point::new(x as i32, y as i32);
            let mut sum = 0;
            coords.iter().for_each(|coord| sum += point.distance(coord));
            if sum < 10_000 {
                safe_points.insert(point);
            }
        }
    }

    for (y, row) in grid.iter().enumerate() {
        for (x, _) in row.iter().enumerate() {
            if safe_points.contains(&Point::new(x as i32, y as i32)) {
                print!("#")
            } else {
                print!(".")
            }
        }
        println!()
    }

    safe_points.len() as i32
}

fn parse(coords: Vec<String>) -> (Vec<Point>, Vec<Vec<i32>>) {
    let coords = coords
        .iter()
        .map(|line| {
            let mut split = line.split(", ");
            let x = split.next().unwrap().parse::<i32>().unwrap();
            let y = split.next().unwrap().parse::<i32>().unwrap();
            Point::new(x, y)
        })
        .collect::<Vec<Point>>();

    let max_x = coords.iter().max_by_key(|p| p.x).unwrap().x;
    let max_y = coords.iter().max_by_key(|p| p.y).unwrap().y;
    let mut grid: Vec<Vec<i32>> = vec![];
    for _ in 0..=max_y {
        let mut row = vec![];
        for _ in 0..=max_x {
            row.push(-1);
        }
        grid.push(row);
    }

    for (i, point) in coords.iter().enumerate() {
        grid[point.y as usize][point.x as usize] = i as i32;
    }

    (coords, grid)
}

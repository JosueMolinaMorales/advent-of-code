use regex::Regex;

use crate::utils::{file_loader::FileLoader, point::Point};

pub fn solve_day_ten() {
    let input = FileLoader::new("./inputs/day10.txt".into()).read_lines();
    println!("Day 10 part 1:");
    let seconds = solve(input.clone());
    println!("Day 10 part 2: {}", seconds);
}

fn solve(input: Vec<String>) -> i32 {
    let regex = Regex::new(r"(-?[\d]+)").unwrap();
    let mut points = input
        .iter()
        .map(|line| {
            let mut matches = regex
                .find_iter(line)
                .map(|n| n.as_str().parse::<i32>().unwrap());

            Light {
                pos: Point::new(matches.next().unwrap(), matches.next().unwrap()),
                vel: Point::new(matches.next().unwrap(), matches.next().unwrap()),
            }
        })
        .collect::<Vec<Light>>();

    let mut seconds = 0;
    let mut min_area = find_grid_area(&points);
    loop {
        seconds += 1;
        // Move every point by their velocity
        points.iter_mut().for_each(|point| {
            point.pos.x += point.vel.x;
            point.pos.y += point.vel.y;
        });

        // If the area of the grid is increasing, we have passed the minimum area
        let area = find_grid_area(&points);
        if area > min_area {
            break;
        }

        min_area = area;
    }

    // Move back one second
    points.iter_mut().for_each(|point| {
        point.pos.x -= point.vel.x;
        point.pos.y -= point.vel.y;
    });

    print_grid(&points);

    seconds - 1
}

fn find_grid_area(points: &Vec<Light>) -> i64 {
    let min_x = points.iter().map(|p| p.pos.x).min().unwrap();
    let max_x = points.iter().map(|p| p.pos.x).max().unwrap();
    let min_y = points.iter().map(|p| p.pos.y).min().unwrap();
    let max_y = points.iter().map(|p| p.pos.y).max().unwrap();

    ((max_x - min_x) as i64 * (max_y - min_y) as i64).abs()
}

fn print_grid(points: &Vec<Light>) {
    let min_x = points.iter().map(|p| p.pos.x).min().unwrap();
    let max_x = points.iter().map(|p| p.pos.x).max().unwrap();
    let min_y = points.iter().map(|p| p.pos.y).min().unwrap();
    let max_y = points.iter().map(|p| p.pos.y).max().unwrap();

    for y in min_y..=max_y {
        for x in min_x..=max_x {
            if points.iter().any(|p| p.pos.x == x && p.pos.y == y) {
                print!("#");
            } else {
                print!(".");
            }
        }
        println!();
    }
}

#[derive(Debug, Clone)]
struct Light {
    pub pos: Point,
    pub vel: Point,
}

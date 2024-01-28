use std::collections::HashMap;

use regex::Regex;

use crate::utils::file_loader::FileLoader;

const TEST_INPUT: &str = r#"#1 @ 1,3: 4x4
#2 @ 3,1: 4x4
#3 @ 5,5: 2x2"#;

pub fn solve_day_three() {
    let input = FileLoader::new("./inputs/day3.txt".into()).read_lines();
    // let input: Vec<String> = TEST_INPUT.split('\n').map(String::from).collect();
    println!("Day 3 part 1: {}", part_one(input.clone()));
    println!("Day 3 part 2: {}", part_two(input))
}

fn part_one(input: Vec<String>) -> i32 {
    let re = Regex::new(r"(\d+)").unwrap();
    let claims = input
        .iter()
        .map(|line| {
            let matches: Vec<_> = re
                .find_iter(line)
                .map(|m| m.as_str().parse().unwrap())
                .collect();
            Claim {
                id: matches[0],
                position: (matches[1], matches[2]),
                size: (matches[3], matches[4]),
            }
        })
        .collect::<Vec<Claim>>();

    let mut points: HashMap<(u16, u16), u16> = HashMap::new();

    claims.iter().for_each(|claim| {
        for x in claim.position.0..claim.position.0 + claim.size.0 {
            for y in claim.position.1..claim.position.1 + claim.size.1 {
                *points.entry((x, y)).or_insert(0) += 1;
            }
        }
    });

    // print_map(points.clone());

    points.values().filter(|v| **v > 1).count() as i32
}

fn print_map(points: HashMap<(u16, u16), u16>) {
    let mut map = vec![vec!['.'; 1000]; 1000];
    points.iter().for_each(|(point, count)| {
        map[point.0 as usize][point.1 as usize] = match count {
            1 => '.',
            _ => 'X',
        }
    });

    map.iter().for_each(|row| {
        row.iter().for_each(|c| print!("{}", c));
        println!()
    });
}

fn part_two(input: Vec<String>) -> i32 {
    0
}

#[derive(Debug)]
struct Claim {
    pub id: u16,
    pub position: (u16, u16),
    pub size: (u16, u16),
}

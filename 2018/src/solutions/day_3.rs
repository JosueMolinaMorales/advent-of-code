use std::collections::HashMap;

use regex::Regex;

use crate::utils::file_loader::FileLoader;

pub fn solve_day_three() {
    let input = FileLoader::new("./inputs/day3.txt".into()).read_lines();
    println!("Day 3 part 1: {}", part_one(input.clone()));
    println!("Day 3 part 2: {}", part_two(input))
}

fn part_one(input: Vec<String>) -> i32 {
    let (_, points) = parse(input);
    points.values().filter(|v| **v > 1).count() as i32
}

fn part_two(input: Vec<String>) -> i32 {
    // What is the ID of the only claim that doesn't overlap?
    let (claims, points) = parse(input);
    claims
        .iter()
        .find(|claim| {
            for x in claim.position.0..claim.position.0 + claim.size.0 {
                for y in claim.position.1..claim.position.1 + claim.size.1 {
                    if points.get(&(x, y)).unwrap() > &1 {
                        return false;
                    }
                }
            }
            true
        })
        .unwrap()
        .id as i32
}

fn parse(input: Vec<String>) -> (Vec<Claim>, HashMap<(u16, u16), u16>) {
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

    (claims, points)
}

#[derive(Debug)]
struct Claim {
    pub id: u16,
    pub position: (u16, u16),
    pub size: (u16, u16),
}

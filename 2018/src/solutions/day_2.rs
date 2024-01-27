use std::collections::{HashMap, HashSet};

use crate::utils::file_loader::FileLoader;
pub fn solve_day_two() {
    let input = FileLoader::new("./inputs/day2.txt".into()).read_lines();
    println!("Day 2 part 1: {}", part_one(input.clone()));
    println!("Day 2 part 2: {}", part_two(input))
}

fn part_one(input: Vec<String>) -> i32 {
    let mut two_count = 0;
    let mut three_count = 0;

    input.iter().for_each(|box_id| {
        let mut counts = HashMap::new();
        box_id.chars().for_each(|c| {
            *counts.entry(c).or_insert(0) += 1;
        });
        if counts.values().any(|v| *v == 2) {
            two_count += 1
        }
        if counts.values().any(|v| *v == 3) {
            three_count += 1
        }
    });

    two_count * three_count
}

fn part_two(input: Vec<String>) -> String {
    // What letters are common between the two correct box IDs?
    let mut seen = HashSet::new();
    for box_id in input.iter() {
        for i in 0..box_id.len() {
            let mut new_id = box_id.clone();
            new_id.remove(i);
            let id = format!("{}{}", i, new_id);
            if seen.contains(&id) {
                return new_id;
            }
            seen.insert(id);
        }
    }
    "".into()
}

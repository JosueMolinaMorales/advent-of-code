use std::collections::{HashMap, HashSet};

use crate::utils::file_loader::FileLoader;
pub fn solve_day_two() {
    let input = FileLoader::new("./inputs/day2.txt".into()).read_lines();
    println!("Day 2 part 1: {}", part_one(input))
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
    let ids: Vec<HashSet<char>> = input.iter().fold(Vec::new(), |mut acc, id| {
        acc.extend(id.chars());
        acc
    });
}

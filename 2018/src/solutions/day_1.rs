use std::collections::HashSet;

use crate::utils::file_loader::FileLoader;

pub fn solve_day_one() {
    let input = FileLoader::new("./inputs/day1.txt".into()).read_lines();

    println!("Day 1 part 1: {}", part_one(input.clone()));
    println!("Day 1 part 2: {}", part_two(input))
}

fn part_one(input: Vec<String>) -> i32 {
    input
        .iter()
        .map(|s| {
            s.parse::<i32>()
                .unwrap_or_else(|_| panic!("Could not convert {} to int", s))
        })
        .sum()
}

fn part_two(input: Vec<String>) -> i32 {
    let mut seen = HashSet::new();
    let mut curr_freq = 0;
    for num in input
        .iter()
        .map(|s| {
            s.parse::<i32>()
                .unwrap_or_else(|_| panic!("Could not convert {} to int", s))
        })
        .cycle()
    {
        curr_freq += num;
        if seen.contains(&curr_freq) {
            return curr_freq;
        }

        seen.insert(curr_freq);
    }

    -1
}

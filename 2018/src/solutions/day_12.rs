use std::collections::HashSet;

use crate::utils::file_loader::FileLoader;

pub fn solve_day_twelve() {
    let input = FileLoader::new("./inputs/day12.txt".into()).read_lines();

    println!("Day 12 part 1: {}", part_one(&input));
    println!("Day 12 part 2: {}", part_two(&input));
}

fn part_one(input: &Vec<String>) -> i64 {
    let mut state = input[0].split(": ").collect::<Vec<&str>>()[1].to_string();
    println!("{:?}", state);
    let rules = input[2..]
        .iter()
        .map(|s| s.split(" => ").collect::<Vec<&str>>())
        .collect::<Vec<Vec<&str>>>();
    let mut offset = 0;
    for _ in 0..20 {
        state = format!("....{}", state);
        offset += 4;
        state = format!("{}....", state);
        state = state
            .chars()
            .collect::<Vec<char>>()
            .windows(5)
            .map(|w| {
                let s = w.iter().collect::<String>();
                match rules.iter().find(|r| r[0] == s) {
                    Some(r) => r[1].chars().next().unwrap(),
                    None => '.',
                }
            })
            .collect::<String>();
    }
    offset /= 2;
    state
        .chars()
        .enumerate()
        .map(|(i, c)| {
            if c == '#' {
                println!("{} - {} = {}", i, offset, i as i64 - offset);
                i as i64 - offset
            } else {
                0
            }
        })
        .sum()
}

fn part_two(input: &Vec<String>) -> i64 {
    let mut state = input[0].split(": ").collect::<Vec<&str>>()[1].to_string();
    let rules = input[2..]
        .iter()
        .map(|s| s.split(" => ").collect::<Vec<&str>>())
        .collect::<Vec<Vec<&str>>>();
    let mut offset = 0;
    let mut last_sum = 0;
    let mut seen = HashSet::new();
    let mut gen = 0;
    loop {
        gen += 1;
        state = format!("....{}", state);
        offset += 4;
        state = format!("{}....", state);
        state = state
            .chars()
            .collect::<Vec<char>>()
            .windows(5)
            .map(|w| {
                let s = w.iter().collect::<String>();
                match rules.iter().find(|r| r[0] == s) {
                    Some(r) => r[1].chars().next().unwrap(),
                    None => '.',
                }
            })
            .collect::<String>();

        let sum: i64 = state
            .chars()
            .enumerate()
            .map(|(i, c)| if c == '#' { i as i64 - (offset / 2) } else { 0 })
            .sum();

        let diff = sum - last_sum;
        if seen.contains(&key(&state)) {
            println!("{}", gen);
            return sum + (50000000000 - gen) * diff;
        }
        seen.insert(key(&state));

        last_sum = sum;
    }
}

fn key(s: &str) -> String {
    // remove extra '.' from the beginning and end
    s.trim_matches('.').to_string()
}

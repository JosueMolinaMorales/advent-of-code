use crate::utils::file_loader::FileLoader;
use chrono::{NaiveDateTime, Timelike};
use std::collections::HashMap;

pub fn solve_day_four() {
    let input = FileLoader::new("./inputs/day4.txt".into()).read_lines();
    println!("Day 4 part 1: {}", part_one(&input));
    println!("Day 4 part 2: {}", part_two(&input));
}

fn part_one(input: &Vec<String>) -> i32 {
    let guards = parse(input);

    let (max_guard, max_minute) = guards
        .iter()
        .max_by_key(|(_, sleep)| sleep.iter().sum::<i32>())
        .map(|(guard, sleep)| {
            let binding = sleep.iter().fold(HashMap::new(), |mut acc, &minute| {
                *acc.entry(minute).or_insert(0) += 1;
                acc
            });
            let max_minute = *binding.iter().max_by_key(|&(_, count)| count).unwrap().0;
            (*guard, max_minute)
        })
        .unwrap();

    max_guard * max_minute
}

fn part_two(input: &Vec<String>) -> i32 {
    let guards = parse(input);

    let (guard, (max_minute, _)) = guards
        .iter()
        .flat_map(|(guard, sleeps)| {
            let binding = sleeps.iter().fold(HashMap::new(), |mut acc, &minute| {
                *acc.entry(minute).or_insert(0) += 1;
                acc
            });
            let max_minute = *binding.iter().max_by_key(|&(_, count)| count).unwrap().0;
            Some((
                *guard,
                (max_minute, binding.get(&max_minute).cloned().unwrap()),
            ))
        })
        .max_by_key(|&(_, (_, count))| count)
        .unwrap();

    guard * max_minute
}

fn parse(input: &Vec<String>) -> HashMap<i32, Vec<i32>> {
    let mut schedule = input
        .iter()
        .map(|line| {
            let line = line.replace("[", "").replace("]", "");
            let mut split = line.split_whitespace();
            let date = split.next().unwrap();
            let time = split.next().unwrap();
            (
                NaiveDateTime::parse_from_str(&format!("{} {}", date, time), "%Y-%m-%d %H:%M")
                    .unwrap(),
                split.collect::<Vec<&str>>().join(" "),
            )
        })
        .collect::<Vec<(NaiveDateTime, String)>>();

    schedule.sort_by(|a, b| a.0.cmp(&b.0));

    let mut guards: HashMap<i32, Vec<i32>> = HashMap::new();
    let mut current_guard = 0;
    let mut sleep_start = 0;

    for (time, action) in schedule.iter() {
        match action {
            a if a.contains("Guard") => {
                current_guard = a
                    .split_whitespace()
                    .nth(1)
                    .unwrap()
                    .replace("#", "")
                    .parse()
                    .unwrap();
            }
            a if a.contains("falls") => {
                sleep_start = time.minute();
            }
            a if a.contains("wakes") => {
                let sleep_end = time.minute();
                let guard = guards.entry(current_guard).or_insert_with(Vec::new);
                guard.extend((sleep_start..sleep_end).map(|x| x as i32));
            }
            _ => {}
        }
    }

    guards
}

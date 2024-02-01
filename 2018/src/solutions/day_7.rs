use std::{
    cmp::Reverse,
    collections::{HashMap, HashSet},
};

use crate::utils::file_loader::FileLoader;

pub fn solve_day_seven() {
    let input = FileLoader::new("./inputs/day7.txt".into()).read_lines();

    println!("Day 7 part 1: {}", part_one(input.clone()));
    println!("Day 7 part 2: {}", part_two(input));
}

fn part_two(input: Vec<String>) -> i32 {
    let mut instructions = parse(input);
    let mut workers = vec![(String::new(), 0); 5];
    let mut seconds = 0;

    while instructions.len() > 0 || workers.iter().any(|worker| !worker.0.is_empty()) {
        // Find available workers
        let available_workers: Vec<&mut (String, u32)> =
            workers.iter_mut().filter(|v| v.0.is_empty()).collect();

        // Find available steps
        let mut available_steps = find_empty_steps(&instructions, true);
        // For every available worker, assign a step
        for available_worker in available_workers {
            if available_steps.is_empty() {
                break;
            }
            available_worker.0 = available_steps.pop().unwrap();
            // Remove
            instructions.remove(&available_worker.0);
        }

        // For every worker, working on a step, inc their time
        for (curr_step, time) in workers.iter_mut() {
            if curr_step.is_empty() {
                // Worker is not doing anything
                continue;
            }
            *time += 1;
            // Check if work is done
            if *time == (60 + (curr_step.chars().nth(0).unwrap() as u32 - 'A' as u32) + 1) {
                // Remove the curr from dep lists
                instructions.iter_mut().for_each(|(_, v)| {
                    v.remove(&curr_step.to_string());
                });
                *curr_step = String::new();
                *time = 0;
            }
        }
        // Add time
        seconds += 1;
    }

    seconds
}

fn part_one(input: Vec<String>) -> String {
    let mut instructions = parse(input);

    let mut ans = String::new();
    while instructions.len() > 0 {
        // Find my starting point
        let potential = find_empty_steps(&instructions, false);
        let curr = potential.get(0).unwrap();
        ans.push_str(curr);
        // Remove the curr from dep lists
        instructions.iter_mut().for_each(|(_, v)| {
            v.remove(curr);
        });

        // Remove curr
        instructions.remove(curr);
    }

    ans
}

fn find_empty_steps(instructions: &HashMap<String, HashSet<String>>, reverse: bool) -> Vec<String> {
    let mut steps = instructions
        .iter()
        .filter_map(|(k, v)| if v.len() == 0 { Some(k.clone()) } else { None })
        .collect::<Vec<String>>();
    if reverse {
        steps.sort_by(|a, b| Reverse(a).cmp(&Reverse(b)))
    } else {
        steps.sort();
    }
    steps
}

fn parse(input: Vec<String>) -> HashMap<String, HashSet<String>> {
    let mut steps = HashSet::new();
    let mut instructions = input.iter().fold(HashMap::new(), |mut acc, inst| {
        let mut split = inst.split_ascii_whitespace();
        let step = split.nth(1).unwrap().to_string();
        let before_step = split.nth(5).unwrap().to_string();
        steps.insert(step.clone());
        steps.insert(before_step.clone());
        acc.entry(before_step)
            .and_modify(|dep: &mut HashSet<String>| {
                dep.insert(step.clone());
            })
            .or_insert(HashSet::from([step.clone()]));
        acc
    });

    // For every step not in instructions, add an empty []
    steps.iter().for_each(|step| {
        instructions.entry(step.clone()).or_default();
    });

    instructions
}

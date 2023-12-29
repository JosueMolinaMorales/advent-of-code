use std::{collections::HashSet, fs, str::Chars};

const INPUT_FILE: &str = "./inputs/day_6_input.txt";

pub fn solve_day_six() {
    let binding = fs::read_to_string(INPUT_FILE).unwrap();
    let input_string = binding.chars();
    println!("Day 6 part one: {}", part_one(input_string.clone()));
    println!("Day 6 part two: {}", part_two(input_string));
}

fn part_two(input_string: Chars) -> i32 {
    let mut past_fourteen = Vec::new();

    let mut res = 0;
    for (i, ch) in input_string.enumerate() {
        if past_fourteen.len() >= 14 {
            // Pop stack
            past_fourteen.pop();
        }
        // Push new character
        past_fourteen.insert(0, ch);

        // Make the vector and set and assert its size is 4
        let past_fourteen_set: HashSet<char> = HashSet::from_iter(past_fourteen.iter().cloned());
        if past_fourteen_set.len() == 14 {
            res = i + 1;
            break;
        }
    }

    res as i32
}

fn part_one(input_string: Chars) -> i32 {
    let mut past_four = Vec::new();

    let mut res = 0;
    for (i, ch) in input_string.enumerate() {
        if past_four.len() >= 4 {
            // Pop stack
            past_four.pop();
        }
        // Push new character
        past_four.insert(0, ch);

        // Make the vector and set and assert its size is 4
        let past_four_set: HashSet<char> = HashSet::from_iter(past_four.iter().cloned());
        if past_four_set.len() == 4 {
            res = i + 1;
            break;
        }
    }

    res as i32
}

#[test]
fn test_part_one() {
    let res = part_one("mjqjpqmgbljsphdztnvjfqwrcgsmlb".chars());
    assert_eq!(res, 7);
}

#[test]
fn test_part_two() {
    let res = part_two("mjqjpqmgbljsphdztnvjfqwrcgsmlb".chars());
    assert_eq!(res, 19);
}

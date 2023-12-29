use std::{collections::HashMap, fs};

const INPUT_FILE: &str = "./inputs/day_3_input.txt";

trait CharacterValue {
    fn get_value(&self) -> u32;
}

impl CharacterValue for char {
    fn get_value(&self) -> u32 {
        // println!("char: {}", self);
        let mut value = *self as u32;
        if (97..=122).contains(&value) {
            // Lowercase letter
            value -= 96;
            value
        } else if (65..=90).contains(&value) {
            // Uppercase letter
            value -= 38;
            value
        } else {
            0
        }
    }
}

pub fn solve_day_3() {
    println!("Day 3 part one: {}", part_one());
    println!("Day 3 part two: {}", part_two());
}

fn part_two() -> i32 {
    let mut letter_count_first: HashMap<char, i32> = HashMap::new();
    let mut letter_count_second: HashMap<char, i32> = HashMap::new();

    let mut matching_characters: Vec<u32> = Vec::new();

    fs::read_to_string(INPUT_FILE)
        .unwrap()
        .split('\n')
        .enumerate()
        .for_each(|(index, line)| {
            let curr_index = (index + 1) % 3;
            if curr_index == 0 {
                // 3 Have passed, calculate matching character for all three
                for c in line.chars() {
                    if letter_count_first.contains_key(&c) && letter_count_second.contains_key(&c) {
                        matching_characters.push(c.get_value());
                        break;
                    }
                }
                letter_count_first = HashMap::new();
                letter_count_second = HashMap::new();
            } else {
                match curr_index {
                    1 => {
                        line.chars().for_each(|c| {
                            letter_count_first
                                .entry(c)
                                .and_modify(|v| *v += 1)
                                .or_insert(1);
                        });
                    }
                    2 => {
                        line.chars().for_each(|c| {
                            letter_count_second
                                .entry(c)
                                .and_modify(|v| *v += 1)
                                .or_insert(1);
                        });
                    }
                    _ => unreachable!(),
                }
            }
        });

    matching_characters.iter().sum::<u32>() as i32
}

fn part_one() -> i32 {
    let res: u32 = fs::read_to_string(INPUT_FILE)
        .unwrap()
        .split('\n')
        .map(|sack| {
            let len_str = sack.len();
            let mut letter_count: HashMap<char, i32> = HashMap::new();

            let mut matching_character = ' ';
            for (index, letter) in sack.chars().enumerate() {
                if (index + 1) <= (len_str / 2) {
                    letter_count
                        .entry(letter)
                        .and_modify(|c| *c += 1)
                        .or_insert(1);
                } else {
                    matching_character = if letter_count.contains_key(&letter) {
                        letter
                    } else {
                        ' '
                    };
                    if matching_character != ' ' {
                        break;
                    }
                }
            }

            // Calculate
            matching_character.get_value()
        })
        .sum();

    res as i32
}

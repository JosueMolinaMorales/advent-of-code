use std::fs;

const INPUT_FILE: &str = "./inputs/day_4_input.txt";

pub fn solve_day_four() {
    let input: Vec<Vec<u32>> = fs::read_to_string(INPUT_FILE)
        .unwrap()
        .split('\n')
        .map(|line| {
            line.replace('-', ",")
                .split(',')
                .map(|string| string.parse::<u32>().unwrap())
                .collect::<Vec<_>>()
        })
        .collect();
    println!("Day 4 part one: {}", part_one(input.clone()));
    println!("Day 4 part two: {}", part_two(input));
}

fn part_two(input: Vec<Vec<u32>>) -> i32 {
    let res: u32 = input
        .iter()
        .map(|pairs| {
            let left = (pairs[0], pairs[1]);
            let right = (pairs[2], pairs[3]);

            if right.0 >= left.0 && right.0 <= left.1 || left.0 >= right.0 && left.0 <= right.1 {
                return 1;
            }

            0
        })
        .sum();

    res as i32
}

fn part_one(input: Vec<Vec<u32>>) -> i32 {
    let res: u32 = input
        .iter()
        .map(|pairs| {
            let left = (pairs[0], pairs[1]);
            let right = (pairs[2], pairs[3]);

            // Check to see if right endpoint one is between left endpoints and
            // Check to see if right endpoint two is also in between left endpoints
            if (right.0 >= left.0 && right.0 <= left.1 && right.1 >= left.0 && right.1 <= left.1)
                || (left.0 >= right.0
                    && left.0 <= right.1
                    && left.1 >= right.0
                    && left.1 <= right.1)
            {
                // Right pair is in between left pair or left pair is in between right pair
                return 1;
            }
            0
        })
        .sum();

    res as i32
}

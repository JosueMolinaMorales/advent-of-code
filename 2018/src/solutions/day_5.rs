use crate::utils::file_loader::FileLoader;

pub fn solve_day_five() {
    let input = FileLoader::new("./inputs/day5.txt".into()).read();
    println!("Day 5 part 1: {}", part_one(input.clone()));
    println!("Day 5 part 2: {}", part_two(input))
}

fn part_one(input: String) -> i32 {
    react_polymer(input).len() as i32
}

fn part_two(input: String) -> i32 {
    let mut min = input.len();
    for c in "abcdefghijklmnopqrstuvwxyz".chars() {
        let polymer = input
            .clone()
            .replace(c, "")
            .replace(c.to_ascii_uppercase(), "");
        let polymer = react_polymer(polymer);
        if polymer.len() < min {
            min = polymer.len();
        }
    }
    min as i32
}

fn react_polymer(polymer: String) -> String {
    let mut buf = Vec::new();
    for c in polymer.chars() {
        if buf.len() > 0 && are_opp(c, buf[buf.len() - 1]) {
            buf.pop();
        } else {
            buf.push(c);
        }
    }
    buf.into_iter().collect()
}

fn are_opp(a: char, b: char) -> bool {
    (a.is_ascii_lowercase() && b.is_ascii_uppercase() && a.to_ascii_uppercase() == b)
        || (a.is_ascii_uppercase() && b.is_ascii_lowercase() && a == b.to_ascii_uppercase())
}

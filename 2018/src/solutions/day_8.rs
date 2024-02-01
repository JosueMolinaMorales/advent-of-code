use crate::utils::file_loader::FileLoader;

const TEST_INPUT: &str = r#"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"#;
pub fn solve_day_eight() {
    let input = FileLoader::new("./inputs/day8.txt".into()).read_lines();
    println!("Day 8 part 1: {}", part_one(input.clone()));
    println!("Day 8 part 2: {}", part_two(input));
}

fn part_one(input: Vec<String>) -> i32 {
    0
}

fn part_two(input: Vec<String>) -> i32 {
    0
}

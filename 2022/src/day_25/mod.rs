pub fn solve_day_25() {
    let input = include_str!("input.txt");
    println!("Day 25 Part one: {}", part_one(input));
}

fn part_one(input: &str) -> String {
    let sum = input.lines().map(convert_to_decimal).sum::<i64>();
    convert_to_snafu(sum)
}

fn convert_to_decimal(snafu: &str) -> i64 {
    snafu.chars().rev().enumerate().fold(0, |acc, (idx, c)| {
        let val = match c {
            '-' => -1,
            '=' => -2,
            _ => c.to_digit(10).unwrap() as i64,
        };
        acc + (val * 5i64.pow(idx as u32))
    })
}

fn convert_to_snafu(decimal: i64) -> String {
    if decimal == 0 {
        return String::new();
    }

    let decimal_remainder = decimal % 5;
    let snafu_digit = ['0', '1', '2', '=', '-'][decimal_remainder as usize];

    let new_decimal = (decimal + 2) / 5;
    let mut snafu = convert_to_snafu(new_decimal);
    snafu.push(snafu_digit);

    snafu
}

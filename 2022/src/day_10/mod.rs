const INPUT: &str = include_str!("./day_10_input.txt");

fn solve() -> (i32, Vec<&'static str>) {
    let mut cycle_counter = 1;
    let mut register = 1;
    let mut cycle_check = 20;
    let mut sum = 0;
    let mut screen = vec![" "; 240];

    for line in INPUT.lines() {
        let exec = line.split(' ').collect::<Vec<&str>>();
        // Beginning of cycle
        cycle(
            &cycle_counter,
            &mut cycle_check,
            &register,
            &mut sum,
            &mut screen,
        );

        if exec[0] == "noop" {
            cycle_counter += 1;
            continue;
        }

        // Operation is addx
        cycle_counter += 1; // Increase cycle by 1
        cycle(
            &cycle_counter,
            &mut cycle_check,
            &register,
            &mut sum,
            &mut screen,
        );
        cycle_counter += 1; // Increase cycle by 1
                            // 2 cycles done, Add to register
        register += exec[1].parse::<i32>().unwrap();
        // End of cycle
    }

    (sum, screen)
}

fn part_one() -> i32 {
    let (sum, _) = solve();
    sum
}

fn part_two() -> String {
    let (_, screen) = solve();
    let mut letters = String::new();
    for (i, pixel) in screen.iter().enumerate() {
        if i % 40 == 0 {
            letters.push('\n')
        }
        letters += pixel;
    }
    letters
}

fn cycle(
    cycle_counter: &i32,
    cycle_check: &mut i32,
    register: &i32,
    sum: &mut i32,
    screen: &mut Vec<&str>,
) {
    if cycle_counter == cycle_check {
        *sum += cycle_counter * register;
        *cycle_check += 40;
    }
    let screen_pixel = (cycle_counter % 40) as usize;
    if (screen_pixel == (*register as usize))
        || (screen_pixel == (*register + 1) as usize)
        || (screen_pixel == (*register + 2) as usize)
    {
        screen[(cycle_counter - 1) as usize] = "#";
    }
}

pub fn solve_day_ten() {
    println!("Day 10 Part one: {}", part_one());
    println!("Day 10 Part two: {}", part_two());
}

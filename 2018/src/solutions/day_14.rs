pub fn solve_day_fourteen() {
    let input = 652601;
    println!("Day 14 part 1: {}", part_one(input));
    println!("Day 14 part 2: {}", part_two(input));
}

fn part_one(input: i64) -> String {
    let mut elves = vec![0, 1];
    let mut recipes = vec![3, 7];

    for _ in 0..(input + 10) {
        add_recipe(&mut recipes, &mut elves)
    }

    recipes[input as usize..(input as usize + 10)]
        .iter()
        .map(|&num| num.to_string())
        .collect::<Vec<String>>()
        .join("")
}

fn part_two(input: i64) -> usize {
    let mut elves = vec![0, 1];
    let mut recipes = vec![3, 7];
    let input = input.to_string();
    let input = input
        .chars()
        .map(|ch| ch.to_digit(10).unwrap() as usize)
        .collect::<Vec<usize>>();

    loop {
        add_recipe(&mut recipes, &mut elves);

        if recipes.len() <= input.len() {
            continue;
        }

        for (idx, window) in recipes[recipes.len() - input.len() - 1..]
            .windows(input.len())
            .enumerate()
        {
            if window == input.as_slice() {
                return idx + recipes.len() - input.len() - 1;
            }
        }
    }
}

fn add_recipe(recipes: &mut Vec<usize>, elves: &mut Vec<usize>) {
    let sum = recipes[elves[0]] + recipes[elves[1]];
    let sum = sum.to_string();
    sum.chars()
        .for_each(|ch| recipes.push(ch.to_digit(10).unwrap() as usize));

    elves
        .iter_mut()
        .for_each(|idx| *idx = (*idx + 1 + recipes[*idx]) % recipes.len());
}

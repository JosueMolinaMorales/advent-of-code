use std::fs;

pub fn solve_day_five() {
    let mut stacks = vec![
        vec!["N", "R", "G", "P"],
        vec!["J", "T", "B", "L", "F", "G", "D", "C"],
        vec!["M", "S", "V"],
        vec!["L", "S", "R", "C", "Z", "P"],
        vec!["P", "S", "L", "V", "C", "W", "D", "Q"],
        vec!["C", "T", "N", "W", "D", "M", "S"],
        vec!["H", "D", "G", "W", "P"],
        vec!["Z", "L", "P", "H", "S", "C", "M", "V"],
        vec!["R", "P", "F", "L", "W", "G", "Z"]
    ];
    part_one(&mut stacks.clone());
    part_two(&mut stacks);

}

fn part_one(stacks: &mut Vec<Vec<&str>>) {
    fs::read_to_string("./inputs/day_5_input.txt").unwrap()
        .split('\n')
        .for_each(|line| {
            let replaced_line = line
                .replace("move ", "")
                .replace(" from ", ",")
                .replace(" to ", ",");
            let numbers = replaced_line.split(',').map(|num| num.parse::<u32>().unwrap()).collect::<Vec<u32>>();

            let num_of_moves = numbers[0];
            let from = numbers[1]-1;
            let to = numbers[2]-1;

            let mut mut_vec = stacks.get_mut(from as usize).unwrap();

            for _ in 1..=num_of_moves {
                let val = mut_vec.pop().unwrap();
                mut_vec = stacks.get_mut(to as usize).unwrap();
                mut_vec.push(val);
                mut_vec = stacks.get_mut(from as usize).unwrap();
            }
        });

    let res = stacks.iter().map(|stack| stack.clone().pop().unwrap()).collect::<Vec<&str>>().join("");

    println!("Part 1 answer: {}", res);

}

fn part_two(stacks: &mut Vec<Vec<&str>>) {
    fs::read_to_string("./inputs/day_5_input.txt").unwrap()
        .split('\n')
        .for_each(|line| {
            let replaced_line = line
                .replace("move ", "")
                .replace(" from ", ",")
                .replace(" to ", ",");
            let numbers = replaced_line.split(',').map(|num| num.parse::<u32>().unwrap()).collect::<Vec<u32>>();

            let num_of_moves = numbers[0];
            let from = numbers[1]-1;
            let to = numbers[2]-1;


            let mut mut_vec = stacks.get_mut(from as usize).unwrap();

            let mut moving_stack = Vec::new();
            for _ in 1..=num_of_moves {
                moving_stack.push(mut_vec.pop().unwrap());
            }
            moving_stack.reverse();
            mut_vec = stacks.get_mut(to as usize).unwrap();
            mut_vec.append(&mut moving_stack)
        });
    let res = stacks.iter().map(|stack| stack.clone().pop().unwrap()).collect::<Vec<&str>>().join("");

    println!("Part 2 answer: {}", res);
}
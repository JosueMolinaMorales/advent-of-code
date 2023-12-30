const INPUT: &str = include_str!("input.txt");

pub fn solve_day_20() {
    println!("Day 20 Part one: {}", part_one());
    println!("Day 20 Part two: {}", part_two());
}

fn part_one() -> isize {
    let (mut nums, mut mix_pos) = parse();

    mix_numbers(&mut nums, &mut mix_pos);

    let res = get_sum(&mut nums, &mut mix_pos);

    res
}

fn part_two() -> isize {
    let (mut nums, mut mix_pos) = parse();
    mix_numbers(&mut nums, &mut mix_pos);

    let decrypt_key = 811589153;
    nums = nums.iter().map(|num| *num * decrypt_key).collect();
    mix_pos = (0..nums.len()).collect();
    (1..=10).for_each(|_| {
        mix_numbers(&mut nums, &mut mix_pos);
    });

    let res = get_sum(&mut nums, &mut mix_pos);

    res
}

fn parse() -> (Vec<isize>, Vec<usize>) {
    // From the input, make two lists
    let mut nums = vec![]; // Keep track of the numbers
    let mut mix_pos = vec![]; // Keep track of the mixed positions of the numbers

    INPUT.lines().enumerate().for_each(|(i, num)| {
        let num = num.parse::<isize>().unwrap();
        nums.push(num);
        mix_pos.push(i);
    });

    (nums, mix_pos)
}

fn mix_numbers(nums: &mut Vec<isize>, mix_pos: &mut Vec<usize>) {
    for (i, num) in nums.iter().enumerate() {
        // Get the current position of the number
        let mixed_index = mix_pos.iter().position(|pos| *pos == i).unwrap();
        // Remove this position from the mixed list
        mix_pos.remove(mixed_index);
        // Get the new index of the list and insert it back
        let new_mixed_index =
            (mixed_index as isize + num).rem_euclid(mix_pos.len() as isize) as usize;
        mix_pos.insert(new_mixed_index, i);
    }
}

fn get_sum(nums: &mut Vec<isize>, mix_pos: &mut Vec<usize>) -> isize {
    let zero_index = nums.iter().position(|num| *num == 0).unwrap();
    let zero_mixed_index = mix_pos
        .iter()
        .position(|mix_num| *mix_num == zero_index)
        .unwrap();
    let res: isize = [1000, 2000, 3000]
        .iter()
        .map(|offset| {
            let mixed_idx = (zero_mixed_index + offset) % mix_pos.len();
            let nums_idx = mix_pos[mixed_idx];
            nums[nums_idx]
        })
        .sum();
    res
}

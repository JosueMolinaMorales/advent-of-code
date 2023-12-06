use std::fs;

const FILE_NAME_DAY_ONE: &str = "./inputs/day_1_input.txt";

fn read_file() -> Vec<String> {
    let file = fs::read_to_string(FILE_NAME_DAY_ONE).unwrap();

    let file_vec: Vec<String> = file.split('\n').map(|str| str.trim().to_string()).collect();

    file_vec
}

pub fn solve_day_one() {
    // Read file
    let cal_str_vec = read_file();
    part_one(cal_str_vec.clone());
    part_two(cal_str_vec)
}

fn part_one(cal_str_vec: Vec<String>) {
        // At every empty index, take the sum of the previous values
        let mut sum_vec = vec![];
        let mut local_sum = 0;
        cal_str_vec.iter().for_each(|cal| {
            if cal.is_empty() {
                sum_vec.push(local_sum);
                local_sum = 0;
            } else {
                local_sum += cal.parse::<i32>().unwrap();
            }
        });
        // Push the last local_sum
        sum_vec.push(local_sum);
    
        println!("The max number of calories is: {} ", sum_vec.iter().max().unwrap())
}

fn part_two(cal_str_vec: Vec<String>) {
    let mut sum_vec = vec![];
        let mut local_sum = 0;
        cal_str_vec.iter().for_each(|cal| {
            if cal.is_empty() {
                sum_vec.push(local_sum);
                local_sum = 0;
            } else {
                local_sum += cal.parse::<i32>().unwrap();
            }
        });
        // Push the last local_sum
        sum_vec.push(local_sum);

        sum_vec.sort();

        let top_3: i32 = sum_vec.iter().rev().take(3).sum();
    
        println!("The top three calories summed is: {} ", top_3)
}

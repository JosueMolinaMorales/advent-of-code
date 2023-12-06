use std::fs;
use std::io::{self, BufRead};
pub fn solve_day_eight(){
    
    let file = fs::File::open("./inputs/day_8_input.txt").unwrap();
    let reader = io::BufReader::new(file);

    let mut grid = Vec::new();

    for line in reader.lines() {
        let line = line.unwrap();
        for num in line.split_whitespace() {
            let num = num.split("").filter_map(|n| n.trim().parse::<u32>().ok()).collect::<Vec<u32>>();
            grid.push(num);
        }
    }

    let mut count = 0;
    let mut scenic_score: Vec<u32> = vec![];
    for (i, row) in grid.iter().enumerate() {
        for (j, &num) in row.iter().enumerate() {
            let mut top_score = 0;
            let mut top = true;
            for k in (0..i).rev() {
                top_score += 1;
                if grid[k][j] >= num {
                    top = false;
                    break;
                }
            }
            let mut bottom_score = 0;
            let mut bottom = true;
            for k in (i + 1)..grid.len() {
                bottom_score += 1;
                if grid[k][j] >= num {
                    bottom = false;
                    break;
                }
            }
            let mut left_score = 0;
            let mut left = true;
            for k in (0..j).rev() {
                left_score += 1;
                if grid[i][k] >= num {
                    left = false;
                    break;
                }
            }
            let mut right = true;
            let mut right_score = 0;
            for k in (j + 1)..row.len() {
                right_score += 1;
                if grid[i][k] >= num {
                    right = false;
                    break;
                }
            }
            if top || bottom || left || right {
                count += 1;
            }
            scenic_score.push(top_score*bottom_score*left_score*right_score);
        }
    }

    println!("part 1: {}", count);
    println!("part 2: {}", scenic_score.iter().max().unwrap());
}
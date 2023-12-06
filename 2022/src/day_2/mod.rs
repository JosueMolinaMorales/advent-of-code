/*
A -> Rock
B -> Paper
C -> Scissors

X -> Rock
Y -> Paper
Z -> Scissors

Points:
1 Point -> Choosing rock
2 Points -> Choose Paper
3 POints -> Choosing Scissors

0 -> Points for loss
3 -> For Draw
6 -> For Win
*/
use std::{fs, ops::Index};

enum CHOICES {
    AX, // Elf chose Rock, Need to Lose, Choose Scissors -> 0 + 3 = 3
    AY, // Rock, Need to Draw, Choose Rock -> 3 + 1 = 4
    AZ, // Rock, Need to Win, Choose Paper -> 6 + 2 = 8
    BX, // 0 + 1 -> 1 
    BY, // 3 + 2 -> 5
    BZ, // 6 + 3 -> 9
    CX, // 0 + 2 -> 2
    CY, // 3 + 3 -> 6
    CZ, // 6 + 1 -> 7
}

impl CHOICES {
    fn calc_points(&self) -> i32 {
        match *self {
            CHOICES::AX => 3, 
            CHOICES::AY => 4,
            CHOICES::AZ => 8,
            CHOICES::BX => 1,
            CHOICES::BY => 5,
            CHOICES::BZ => 9,
            CHOICES::CX => 2,
            CHOICES::CY => 6,
            CHOICES::CZ => 7,
        }
    }
}

const FILE_PATH: &str = "./inputs/day_2_input.txt";

pub fn solve_day_two() {
    part_one();
    part_two();
}

/*
    X -> Lose
    Y -> Draw
    Z -> Win
*/
fn part_two() {
    let result: i32 = fs::read_to_string(FILE_PATH)
    .unwrap()
    .split('\n')
    .map(|s| s.to_string())
    .map(|line| {
        // Map a line to points
        let choices = line.split(' ').collect::<Vec<&str>>();
        let elf_choice = *choices.index(0);
        let my_choice = *choices.index(1);

        match elf_choice {
            "A" => {
                match my_choice {
                    "X" => CHOICES::AX, // Need to lose, 
                    "Y" => CHOICES::AY,
                    "Z" => CHOICES::AZ,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            "B" => {
                match my_choice {
                    "X" => CHOICES::BX,
                    "Y" => CHOICES::BY,
                    "Z" => CHOICES::BZ,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            "C" => {
                match my_choice {
                    "X" => CHOICES::CX,
                    "Y" => CHOICES::CY,
                    "Z" => CHOICES::CZ,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            _ => panic!("Expecting A B or C")
        }
    })
    .map(|choice| choice.calc_points())
    .sum();

    println!("Pt2 Result: {}", result);
}

fn part_one() {
    let result: i32 = fs::read_to_string(FILE_PATH)
    .unwrap()
    .split('\n')
    .map(|s| s.to_string())
    .map(|line| {
        // Map a line to points
        let choices = line.split(' ').collect::<Vec<&str>>();
        let elf_choice = *choices.index(0);
        let my_choice = *choices.index(1);
        let mut points_earned = match my_choice {
            "X" => 1,
            "Y" => 2,
            "Z" => 3,
            _ => panic!("Expecting X Y or Z")
        };

        points_earned += match elf_choice {
            "A" => {
                match my_choice {
                    "X" => 3,
                    "Y" => 6,
                    "Z" => 0,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            "B" => {
                match my_choice {
                    "X" => 0,
                    "Y" => 3,
                    "Z" => 6,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            "C" => {
                match my_choice {
                    "X" => 6,
                    "Y" => 0,
                    "Z" => 3,
                    _ => panic!("Expecting X, Y, Z"),
                }
            },
            _ => panic!("Expecting A B or C")
        };
        points_earned
    })
    .sum();

    println!("Result: {}", result);
}
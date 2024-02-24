use std::collections::HashSet;

use crate::utils::{file_loader::FileLoader, point::Point};

pub fn solve_day_fifteen() {
    let lines = FileLoader::new("./inputs/day15.txt").read_lines();
    println!("Day 15 Part 1: {}", part_one(&lines))
}

enum Team {
    Elf,
    Goblin,
}

struct Unit {
    pos: Point,
    hp: i64,
    team: Team,
}

fn part_one(lines: &Vec<String>) -> i64 {
    let mut walls = HashSet::new();
    let mut units = Vec::new();
    for (y, line) in lines.iter().enumerate() {
        for (x, c) in line.chars().enumerate() {
            let pos = Point::new(x as i32, y as i32);
            match c {
                '#' => {
                    walls.insert(pos);
                }
                'E' => {
                    units.push(Unit {
                        pos,
                        hp: 200,
                        team: Team::Elf,
                    });
                }
                'G' => {
                    units.push(Unit {
                        pos,
                        hp: 200,
                        team: Team::Goblin,
                    });
                }
                _ => {}
            }
        }
    }

    let mut rounds = 0;
    loop {
        if !step(&mut units, &walls) {
            return units.iter().map(|u| u.hp).sum::<i64>() * rounds;
        }
        rounds += 1;
    }
    0
}

fn step(units: &mut )

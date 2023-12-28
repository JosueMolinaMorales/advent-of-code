use std::collections::{HashMap, HashSet};

pub fn solve_day_23() {
    let input = include_str!("input.txt");
    println!("Day 23 Part one: {}", part_one(input));
    println!("Day 23 Part two: {}", part_two(input));
}

#[derive(Debug, Hash, Eq, PartialEq, Clone, Copy)]
enum Move {
    North,
    South,
    East,
    West,
}

#[derive(Debug, Hash, Eq, PartialEq, Clone)]
struct Elf(i32, i32);

impl Elf {
    fn first_half_move(
        &self,
        current_move: Move,
        other_elves: &HashSet<Elf>,
    ) -> Option<(i32, i32)> {
        // If no other Elves are in one of those eight positions, the Elf does not do anything during this round.
        let can_move = DIRECTIONS.iter().any(|(i, j)| {
            // If any of the directions are occupied, we dont skip
            let new_pos = (self.0 + i, self.1 + j);
            other_elves.contains(&Elf(new_pos.0, new_pos.1))
        });

        if !can_move {
            return None;
        }

        let direction_check = vec![
            vec![NORTH, NORTH_EAST, NORTH_WEST], // North
            vec![SOUTH, SOUTH_EAST, SOUTH_WEST], // South
            vec![WEST, NORTH_WEST, SOUTH_WEST],  // West
            vec![EAST, NORTH_EAST, SOUTH_EAST],  // East
        ];

        let mut i = match current_move {
            Move::North => 0,
            Move::South => 1,
            Move::West => 2,
            Move::East => 3,
        };

        for _ in 0..4 {
            let dir_check = &direction_check[i % 4];
            // If there are no elves in the same spot, move
            let elf_in_spot = dir_check.iter().any(|(i, j)| {
                let new_pos = (self.0 + i, self.1 + j);
                other_elves.contains(&Elf(new_pos.0, new_pos.1))
            });

            if !elf_in_spot {
                let new_pos = (self.0 + dir_check[0].0, self.1 + dir_check[0].1);
                return Some((new_pos.0, new_pos.1));
            }
            i += 1
        }

        None
    }
}

const NORTH: (i32, i32) = (-1, 0);
const SOUTH: (i32, i32) = (1, 0);
const EAST: (i32, i32) = (0, 1);
const WEST: (i32, i32) = (0, -1);
const NORTH_EAST: (i32, i32) = (-1, 1);
const NORTH_WEST: (i32, i32) = (-1, -1);
const SOUTH_EAST: (i32, i32) = (1, 1);
const SOUTH_WEST: (i32, i32) = (1, -1);
const DIRECTIONS: [(i32, i32); 8] = [
    NORTH, NORTH_EAST, NORTH_WEST, SOUTH, SOUTH_EAST, SOUTH_WEST, EAST, WEST,
];

fn part_one(input: &str) -> i32 {
    let mut elves = parse(input);
    let mut dir_move = vec![Move::North, Move::South, Move::West, Move::East];
    for _ in 0..10 {
        let curr_move = dir_move.get(0).unwrap().clone();
        dir_move.rotate_left(1);
        // First half, for every elf, propose a move
        let proposed_moves = get_proposed_moves(curr_move, &elves);
        // Second half, for every elf, check if any other elf proposed the same move
        move_elves(&mut elves, &proposed_moves)
    }

    let grid = build_grid(&elves);

    // for row in grid.iter() {
    //     for c in row.iter() {
    //         print!("{}", c);
    //     }
    //     println!();
    // }

    let mut empty = 0;
    for row in grid.iter() {
        for c in row.iter() {
            if *c == '.' {
                empty += 1;
            }
        }
    }

    empty
}

fn part_two(input: &str) -> i32 {
    let mut elves = parse(input);

    let dir_move = vec![Move::North, Move::South, Move::West, Move::East];

    let mut i = 1;
    loop {
        let curr_move = dir_move.get(i % dir_move.len()).unwrap().clone();
        // First half, for every elf, propose a move
        let proposed_moves = get_proposed_moves(curr_move, &elves);
        if proposed_moves.is_empty() {
            break;
        }
        // Second half, for every elf, check if any other elf proposed the same move
        move_elves(&mut elves, &proposed_moves);
        i += 1;
    }

    i as i32
}

fn move_elves(elves: &mut HashSet<Elf>, proposed_moves: &HashMap<(i32, i32), Vec<Elf>>) {
    for (elf, pm) in proposed_moves {
        if pm.len() == 1 {
            elves.remove(&pm[0]);
            elves.insert(Elf(elf.0, elf.1));
        }
    }
}

fn get_proposed_moves(curr_move: Move, elves: &HashSet<Elf>) -> HashMap<(i32, i32), Vec<Elf>> {
    let mut proposals: HashMap<(i32, i32), Vec<Elf>> = HashMap::new();
    for elf in elves.iter() {
        let proposed_move = elf.first_half_move(curr_move, &elves);
        if let Some(proposed_move) = proposed_move {
            proposals
                .entry(proposed_move)
                .or_insert(vec![])
                .push(elf.clone());
        }
    }

    proposals
}

fn parse(input: &str) -> HashSet<Elf> {
    let mut elves = HashSet::new();
    for (i, line) in input.lines().enumerate() {
        for (j, c) in line.chars().enumerate() {
            if c == '#' {
                elves.insert(Elf(i as i32, j as i32));
            }
        }
    }

    elves
}

fn build_grid(elves: &HashSet<Elf>) -> Vec<Vec<char>> {
    // Get the min and max of each coordinate
    let mut min_x = std::i32::MAX;
    let mut max_x = std::i32::MIN;
    let mut min_y = std::i32::MAX;
    let mut max_y = std::i32::MIN;

    for elf in elves.iter() {
        if elf.0 < min_x {
            min_x = elf.0;
        }
        if elf.0 > max_x {
            max_x = elf.0;
        }
        if elf.1 < min_y {
            min_y = elf.1;
        }
        if elf.1 > max_y {
            max_y = elf.1;
        }
    }

    let mut grid = vec![];
    for i in min_x..=max_x {
        let mut row = vec![];
        for j in min_y..=max_y {
            if elves.contains(&Elf(i, j)) {
                row.push('#');
            } else {
                row.push('.');
            }
        }
        grid.push(row);
    }

    grid
}

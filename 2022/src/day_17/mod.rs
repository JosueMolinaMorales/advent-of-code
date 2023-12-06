use std::collections::HashSet;

const INPUT: &str = include_str!("day_17_input.txt");
const TRILLION: usize = 1_000_000_000_000;

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
enum Gas {
    Right,
    Left
}
impl Gas {
    fn new(c: char) -> Gas {
        match c {
            '>' => Gas::Right,
            '<' => Gas::Left,
            _ => unreachable!()
        }
    }
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord, Hash)]
enum Rocks {
    HorizontalLine,
    Plus,
    LShape,
    VerticalLine,
    Square
}

impl Rocks {
    fn get_starting_positions(&self, h: usize) -> Vec<(usize, usize)> {
        match *self {
            Rocks::HorizontalLine => vec![
                (2, h),
                (3, h),
                (4, h),
                (5, h)
            ],
            Rocks::Plus => vec![
                (3, h), // Bottom
                (2, h+1),
                (3, h+1),
                (4, h+1),
                (3, h+2) // top
            ],
            Rocks::LShape => vec![
                (2, h),
                (3, h),
                (4, h),
                (4, h+1),
                (4, h+2)
            ],
            Rocks::VerticalLine => vec![
                (2, h),
                (2, h+1),
                (2, h+2),
                (2, h+3),
            ],
            Rocks::Square => vec![
                (2, h),
                (3, h),
                (2, h+1),
                (3, h+1)
            ],
        }
    }

    fn move_rock_right(&self, curr_pos: &mut Vec<(usize, usize)>, grid: &HashSet<(usize, usize)>) {
        // When moving rock right, need to check that its x pos <= 6, also checking for other rocks
        let mut can_move_right = true;
        for (x, y) in curr_pos.iter() {
            if *x + 1 == 7 || grid.contains(&(*x+1, *y)) {
                // if grid contains (x+1, y) then there is already a rock in that position and it cannot move.
                can_move_right = false;
                break;
            }
        }
        if can_move_right {
            for (x, _) in curr_pos {
                *x += 1;
            }
        }
    }

    fn move_rock_left(&self, curr_pos: &mut Vec<(usize, usize)>, grid: &HashSet<(usize, usize)>) {
        // When moving rock right, need to check that its x pos <= 6, also checking for other rocks
        let mut can_move_left = true;
        for (x, y) in curr_pos.iter() {
            if *x as isize - 1 < 0 || grid.contains(&(*x-1, *y)) {
                // if grid contains (x+1, y) then there is already a rock in that position and it cannot move.
                can_move_left = false;
                break;
            }
        }
        if can_move_left {
            for (x, _) in curr_pos {
                *x -= 1;
            }
        }
    }

    fn can_rock_move_down(&self, curr_pos: &Vec<(usize, usize)>, grid: &HashSet<(usize, usize)>) -> bool {
        let mut can_move_down = true;
        for (x, y) in curr_pos.iter() {
            if *y as isize - 1 < 0 || grid.contains(&(*x, *y-1)) {
                // if grid contains (x+1, y) then there is already a rock in that position and it cannot move.
                can_move_down = false;
                break;
            }
        }
        can_move_down
    }

    fn move_rock_down(&self, curr_pos: &mut Vec<(usize, usize)>) {
        // This function is under the assumption that the rock can move down.
        for (_, y) in curr_pos {
            *y -= 1;
        }
    }

    fn place_rock(&self, curr_pos: &Vec<(usize, usize)>, grid: &mut HashSet<(usize, usize)>) {
        // Places all the points in cur_pos in grid
        for (x, y) in curr_pos {
            grid.insert((*x, *y));
        }
    }

}

pub fn solve_day_17() {
    solve_part_one();
    // part 2
    // same current rock, current push, last rock that fell.
    let mut grid: HashSet<(usize, usize)> = HashSet::new(); // Will contain the points of all rocks
    let mut rock_count = 0;
    let mut gases_iter = INPUT.chars().map(|c| Gas::new(c) ).into_iter().cycle();
    let gases_length = INPUT.len();
    let rocks = vec![Rocks::HorizontalLine, Rocks::Plus, Rocks::LShape, Rocks::VerticalLine, Rocks::Square];
    let mut rocks_iter = rocks.iter().cycle();
    let mut total_jets_used = 0;
    let mut cycle: HashSet<(usize, usize, Vec<Rocks>)> = HashSet::new(); // (Current Rock Index, Current Gas Index, The previous 10 rocks)
    let mut previous_ten: Vec<Rocks> = vec![];
    let mut found_cycle: Option<(usize, (usize, usize, Vec<Rocks>), usize)> = None; // Height of rocks, (rock count, gas len, previous 10), rock count
    let mut height_reached: usize = 0;
    let mut limit = 50_455;
    while rock_count < limit {
        let current_rock = rocks_iter.next().unwrap();
        let height = get_height_of_rocks(&grid) + 3 ; // Plus three since height will start 3 above the highest rock
        let mut curr_pos = current_rock.get_starting_positions(height);
        let mut gases_for_rock = vec![];
        if previous_ten.len() < 10 {
            previous_ten.push(current_rock.clone());
        } else {
            let last = previous_ten.last().unwrap();
            previous_ten = vec![last.clone()];
        }
        'falling: loop {
            // Starting dropping rocks
            let gas = gases_iter.next().unwrap();
            if previous_ten.len() == 10 {
                if !cycle.insert((rock_count % 5, total_jets_used % gases_length, previous_ten.clone())) && rock_count > 2022 {
                    if let Some(found_cycle) = found_cycle.as_ref() {
                        if found_cycle.1 == (rock_count % 5, total_jets_used % gases_length, previous_ten.clone()) {
                            // This same formation has been found
                            let curr_height = get_height_of_rocks(&grid);
                            let prev_rock_count = found_cycle.2;
                            let rock_count_diff = rock_count - prev_rock_count;
                            let last_height = found_cycle.0;
                            let height_diff = curr_height - last_height;
                            let iterations_left = TRILLION / rock_count_diff;
                            height_reached = height_diff * iterations_left;
                            let iterations_left = TRILLION - (rock_count_diff * iterations_left);
                            rock_count = 0;
                            limit = iterations_left;
                        }
                    } else {
                        found_cycle = Some((get_height_of_rocks(&grid), (rock_count % 5, total_jets_used % gases_length, previous_ten.clone()), rock_count))
                    }
                }
            }
            match gas {
                Gas::Right => current_rock.move_rock_right(&mut curr_pos, &grid),
                Gas::Left => current_rock.move_rock_left(&mut curr_pos, &grid),
            }
            total_jets_used += 1;
            gases_for_rock.push(gas);
            // Check below rock before stop
            if !current_rock.can_rock_move_down(&mut curr_pos, &grid) {
                break 'falling;
            }

            // Rock can keep falling
            current_rock.move_rock_down(&mut curr_pos);
           
        }
        // Place rock
        current_rock.place_rock(&curr_pos, &mut grid);
        rock_count += 1;
    }
    let heights = get_height_of_rocks(&grid) + height_reached - (2 *found_cycle.unwrap().0);

    println!("Height prt 2: {}", heights);

}

fn solve_part_one() {
    let mut grid: HashSet<(usize, usize)> = HashSet::new(); // Will contain the points of all rocks
    let mut rock_count = 0;
    let mut gases_iter = INPUT.chars().map(|c| Gas::new(c) ).into_iter().cycle();
    let rocks = vec![Rocks::HorizontalLine, Rocks::Plus, Rocks::LShape, Rocks::VerticalLine, Rocks::Square];
    let mut rocks_iter = rocks.iter().cycle();
    while rock_count < 2022 {
        let current_rock = rocks_iter.next().unwrap();
        let height = get_height_of_rocks(&grid) + 3 ; // Plus three since height will start 3 above the highest rock
        let mut curr_pos = current_rock.get_starting_positions(height);
        'falling: loop {
            // Starting dropping rocks
            let gas = gases_iter.next().unwrap();
            match gas {
                Gas::Right => current_rock.move_rock_right(&mut curr_pos, &grid),
                Gas::Left => current_rock.move_rock_left(&mut curr_pos, &grid),
            }

            // Check below rock before stop
            if !current_rock.can_rock_move_down(&mut curr_pos, &grid) {
                break 'falling;
            }

            // Rock can keep falling
            current_rock.move_rock_down(&mut curr_pos);
        }
        // Place rock
        current_rock.place_rock(&curr_pos, &mut grid);
        rock_count += 1;
    }
    
    println!("Height: {}", get_height_of_rocks(&grid));
}
fn get_height_of_rocks(grid: &HashSet<(usize, usize)>) -> usize {
    if grid.is_empty() {
        return 0
    }
    let mut max_y = 0;
    grid.iter().for_each(|(_, y)| {
        if *y > max_y {
            max_y = *y
        }
    });
    max_y + 1
}

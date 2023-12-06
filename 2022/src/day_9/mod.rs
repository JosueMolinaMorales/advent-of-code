use std::collections::HashSet;


const INPUT: &str = include_str!("day_9_input.txt");

enum Direction {
    L, R, U, D
}

impl Direction {
    fn new(direction: &str) -> Direction {
        match direction {
            "U" => Direction::U,
            "R" => Direction::R,
            "L" => Direction::L,
            "D" => Direction::D,
            _ => unreachable!()
        }
    }
}

fn simulate_rope(num_knots: u32) {
    let mut head = (0, 0);
    let mut intermediate_knots = vec![(0,0); (num_knots-1).try_into().unwrap()];
    let mut tail_prev_pos = HashSet::from([*intermediate_knots.last().unwrap()]);

    for line in INPUT.lines() {
        let actions = line.split(' ').collect::<Vec<&str>>();
        let amount = actions[1].parse::<i32>().unwrap();
        let direction = Direction::new(actions[0]);
        for _ in 1..=amount {
            match direction {
                Direction::L => head.0 -= 1,
                Direction::R => head.0 += 1,
                Direction::U => head.1 += 1,
                Direction::D => head.1 -= 1,
            }
            let mut prev_knot = &head;
            for knot in intermediate_knots.iter_mut() {
                let dx: i32 = prev_knot.0 - knot.0;
                let dy: i32 = prev_knot.1 - knot.1;
                if dx.abs() == 2 || dy.abs() == 2 {
                    knot.0 += dx.signum();
                    knot.1 += dy.signum();
                }
                prev_knot = knot;
            }

            
            // After move, add position to hash set
            tail_prev_pos.insert(*intermediate_knots.last().unwrap());
        }
        

    }
    println!("Previous tail count: {}", tail_prev_pos.len());
}

/* Assumption: Head will start at point (0, 0) */
pub fn solve_day_nine() {
    simulate_rope(2);
    simulate_rope(10);
}
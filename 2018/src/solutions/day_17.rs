use std::collections::HashSet;

use regex::Regex;

use crate::utils::{direction::Direction, point::Point};

const SAMPLE_INPUT: &str = r#"x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504"#;

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Water {
    direction: Direction,
    pos: Point,
}
pub fn solve_day_seventeen() {
    let re = Regex::new(r"\d+").unwrap();
    let mut clays = HashSet::new();

    let mut max_y = 0;
    let mut max_x = 0;
    SAMPLE_INPUT.split("\n").for_each(|line| {
        let nums = re
            .find_iter(line)
            .map(|s| s.as_str().parse::<i32>().unwrap())
            .collect::<Vec<i32>>();
        if line.as_bytes()[0] == b"x"[0] {
            if nums[0] > max_y {
                max_y = nums[0] + 1
            }
            if nums[2] > max_x {
                max_x = nums[2] + 1
            }
            for i in nums[1]..=nums[2] {
                clays.insert((nums[0], i));
            }
        } else {
            if nums[0] > max_x {
                max_x = nums[0] + 1
            }
            if nums[2] > max_y {
                max_y = nums[2] + 1
            }
            for i in nums[1]..=nums[2] {
                clays.insert((i, nums[0]));
            }
        }
    });
    let mut grid = Vec::new();

    for i in 0..max_x {
        let mut row = Vec::new();
        for j in 0..max_y {
            if clays.contains(&(j, i)) {
                row.push("#")
            } else {
                row.push(".")
            }
        }
        grid.push(row)
    }

    // Simulate the water falling
    // Goes down until it hits a surface
    // once it hits a surface, it spreads to the left & right
    // once it can no longer spread left, right, it goes up
    let mut curr_water = Water {
        direction: Direction::South,
        pos: Point::new(0, 500),
    };
    let mut moveable_water: HashSet<Water> = HashSet::from([curr_water]);
    let mut i = 0;
    while i < 20 {
        for mut w in moveable_water.clone() {
            moveable_water.extend(move_water(&mut w, &grid, &moveable_water));
        }
        moveable_water = moveable_water
            .iter()
            .filter(|w| grid[w.pos.x as usize][w.pos.y as usize] != "#")
            .cloned()
            .collect();
        i += 1
    }
    println!("{:#?}", moveable_water);
    print_map(max_x, max_y, &grid, &moveable_water);
}

fn move_water(
    water: &mut Water,
    map: &Vec<Vec<&str>>,
    moveable_water: &HashSet<Water>,
) -> Vec<Water> {
    let d: Point = water.direction.into();
    let (dx, dy) = (water.pos.x + d.x, water.pos.y + d.y);

    // Bound Check
    if dx < 0 || dx >= map.len().try_into().unwrap() || dy < 0 || dy >= map[0].len() as i32 {
        return vec![];
    }
    if map[dx as usize][dy as usize] == "#"
        || moveable_water.contains(&Water {
            direction: water.direction,
            pos: Point::new(dx, dy),
        })
    {
        // Expand to the left & right
        return vec![
            Water {
                direction: Direction::West,
                pos: Point::new(water.pos.x, water.pos.y - 1),
            },
            Water {
                direction: Direction::East,
                pos: Point::new(water.pos.x, water.pos.y + 1),
            },
        ];
    }

    return vec![Water {
        direction: water.direction.clone(),
        pos: Point::new(dx, dy),
    }];
}

fn print_map(max_x: i32, max_y: i32, grid: &Vec<Vec<&str>>, water: &HashSet<Water>) {
    let mut grid = grid.clone();
    for w in water {
        grid[w.pos.x as usize][w.pos.y as usize] = "~"
    }
    for i in 0..max_x {
        for j in 450..max_y {
            print!("{}", grid[i as usize][j as usize])
        }
        println!()
    }
}

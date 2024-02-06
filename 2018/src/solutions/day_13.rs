use crate::utils::{direction::Direction, file_loader::FileLoader, point::Point};

pub fn solve_day_thirteen() {
    let input = FileLoader::new("./inputs/day13.txt".into()).read_lines();
    println!("Day 13 part 1: {}", part_one(&input));
    println!("Day 13 part 2: {}", part_two(&input));
}

fn part_one(input: &Vec<String>) -> String {
    solve(input, false)
}

fn part_two(input: &Vec<String>) -> String {
    solve(input, true)
}

fn solve(input: &Vec<String>, remove: bool) -> String {
    let (map, mut carts) = parse(input);
    loop {
        carts.sort_by(|a, b| {
            if a.point.x == b.point.x {
                a.point.y.cmp(&b.point.y)
            } else {
                a.point.x.cmp(&b.point.x)
            }
        });

        for i in 0..carts.len() {
            if carts[i].point.x == -1 {
                continue;
            }

            carts[i].move_cart(&map);

            for j in 0..carts.len() {
                if i != j && carts[i].point == carts[j].point {
                    if remove {
                        carts[i].point.x = -1;
                        carts[j].point.x = -1;
                    } else {
                        return format!("{},{}", carts[i].point.y, carts[i].point.x);
                    }
                }
            }
        }

        if remove {
            carts = carts
                .iter()
                .filter(|cart| cart.point.x != -1)
                .map(|cart| cart.clone())
                .collect();
        }

        if carts.iter().filter(|cart| cart.point.x != -1).count() == 1 {
            let last_cart = carts.iter().find(|cart| cart.point.x != -1).unwrap();
            return format!("{},{}", last_cart.point.y, last_cart.point.x);
        }
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
enum Turn {
    Straight,
    Left,
    Right,
}

#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Ord)]
struct Cart {
    pub point: Point,
    pub direction: Direction,
    pub turn: Turn,
}

impl Cart {
    fn move_cart(&mut self, grid: &Vec<Vec<char>>) {
        match self.direction {
            Direction::East => self.point.y += 1,
            Direction::North => self.point.x -= 1,
            Direction::South => self.point.x += 1,
            Direction::West => self.point.y -= 1,
        }

        // Change direction
        let curr_cell = grid[self.point.x as usize][self.point.y as usize];
        match (curr_cell, self.direction) {
            ('+', dir) => {
                self.direction = match (self.turn, dir) {
                    (Turn::Straight, dir) => dir,
                    (Turn::Left, Direction::East) => Direction::North,
                    (Turn::Left, Direction::North) => Direction::West,
                    (Turn::Left, Direction::South) => Direction::East,
                    (Turn::Left, Direction::West) => Direction::South,
                    (Turn::Right, Direction::East) => Direction::South,
                    (Turn::Right, Direction::North) => Direction::East,
                    (Turn::Right, Direction::South) => Direction::West,
                    (Turn::Right, Direction::West) => Direction::North,
                };
                self.cycle_direction()
            }
            ('/', dir) => {
                self.direction = match dir {
                    Direction::East => Direction::North,
                    Direction::North => Direction::East,
                    Direction::South => Direction::West,
                    Direction::West => Direction::South,
                }
            }
            ('\\', dir) => {
                self.direction = match dir {
                    Direction::East => Direction::South,
                    Direction::North => Direction::West,
                    Direction::South => Direction::East,
                    Direction::West => Direction::North,
                }
            }
            _ => {}
        }
    }

    fn cycle_direction(&mut self) {
        self.turn = match self.turn {
            Turn::Right => Turn::Left,
            Turn::Left => Turn::Straight,
            Turn::Straight => Turn::Right,
        }
    }
}

fn parse(input: &Vec<String>) -> (Vec<Vec<char>>, Vec<Cart>) {
    let mut map = input
        .iter()
        .map(|line| line.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    // Find carts
    let mut carts = Vec::new();
    map.iter().enumerate().for_each(|(x, row)| {
        row.iter().enumerate().for_each(|(y, cell)| match *cell {
            '>' | '<' | 'v' | '^' => {
                carts.push(Cart {
                    point: Point::new(x as i32, y as i32),
                    direction: Direction::from(cell.to_string()),
                    turn: Turn::Left,
                });
            }
            _ => {}
        })
    });

    // Remove carts from map
    carts.iter().for_each(|cart| {
        let (x, y) = (cart.point.x as usize, cart.point.y as usize);
        match cart.direction {
            Direction::East | Direction::West => map[x][y] = '-',
            Direction::North | Direction::South => map[x][y] = '|',
        }
    });

    (map, carts)
}

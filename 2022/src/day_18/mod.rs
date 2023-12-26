use std::collections::{HashMap, HashSet};

const TEST_INPUT: &str = r#"2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5"#;

#[derive(PartialEq, Eq, Debug, Hash, Clone, Copy, Default)]
struct Coord {
    x: isize,
    y: isize,
    z: isize,
}
impl Coord {
    fn neighbours(&self) -> Vec<Coord> {
        let mut neighbours = Vec::new();

        // loop over every dimension in a cube
        for dimension in [Dimension::X, Dimension::Y, Dimension::Z] {
            // add or remove 1 to coordinate in current dimension
            for offset in [-1, 1] {
                // resulting coordinates are from the coord to a side of a cube
                let mut neighbour = self.clone();
                match dimension {
                    Dimension::X => neighbour.x += offset,
                    Dimension::Y => neighbour.y += offset,
                    Dimension::Z => neighbour.z += offset,
                }
                neighbours.push(neighbour);
            }
        }

        neighbours
    }

    fn in_bounds(&self, bounds: &[Self; 2]) -> bool {
        let [mins, maxs] = bounds;
        self.x >= mins.x - 1
            && self.x <= maxs.x + 1
            && self.y >= mins.y - 1
            && self.y <= maxs.y + 1
            && self.z >= mins.z - 1
            && self.z <= maxs.z + 1
    }
}

enum Dimension {
    X,
    Y,
    Z,
}

const INPUT: &str = include_str!("./input.txt");

pub fn solve_day_18() {
    println!("Part 1: {}", part_one(INPUT));
    println!("Part 2: {}", part_two(INPUT))
}

fn parse_input(input: &str) -> HashSet<Coord> {
    let mut cubes = HashSet::new();
    input.lines().for_each(|line| {
        let cube = line
            .split(",")
            .map(|c| c.parse::<isize>().unwrap())
            .collect::<Vec<isize>>();
        cubes.insert(Coord {
            x: cube[0],
            y: cube[1],
            z: cube[2],
        });
    });
    cubes
}

fn bounds(cubes: &HashSet<Coord>) -> [Coord; 2] {
    cubes.iter().fold(
        [Coord::default(), Coord::default()],
        |[mut mins, mut maxs], cube| {
            mins.x = mins.x.min(cube.x);
            mins.y = mins.y.min(cube.y);
            mins.z = mins.z.min(cube.z);
            maxs.x = maxs.x.max(cube.x);
            maxs.y = maxs.y.max(cube.y);
            maxs.z = maxs.z.max(cube.z);
            [mins, maxs]
        },
    )
}

fn exposed(cubes: &HashSet<Coord>) -> HashSet<Coord> {
    let bounds = bounds(cubes);
    let mut exposed = HashSet::new();

    let start = Coord::default();
    let mut stack = Vec::new();
    let mut seen = HashSet::new();

    stack.push(start);
    seen.insert(start);

    while let Some(coord) = stack.pop() {
        for neighbour in coord.neighbours() {
            if cubes.contains(&neighbour) || !neighbour.in_bounds(&bounds) {
                continue;
            }
            if seen.insert(neighbour) {
                stack.push(neighbour);
                exposed.insert(neighbour);
            }
        }
    }

    exposed
}

fn part_two(input: &str) -> usize {
    let cubes = parse_input(input);
    let exposed = exposed(&cubes);

    cubes
        .iter()
        .flat_map(|coord| coord.neighbours())
        .filter(|coord| exposed.contains(coord))
        .count()
}

fn part_one(input: &str) -> usize {
    let cubes = parse_input(input);

    cubes
        .iter()
        .flat_map(|coord| coord.neighbours())
        // only keep neighbours that are not a cube
        .filter(|coord| !cubes.contains(coord))
        .count()
}

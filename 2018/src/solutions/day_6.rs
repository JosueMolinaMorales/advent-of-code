use crate::utils::file_loader::FileLoader;

const TEST_INPUT: &str = r#"1, 1
1, 6
8, 3
3, 4
5, 5
8, 9"#;
pub fn solve_day_six() {
    // let input = FileLoader::new("./inputs/day6.txt".into()).read_lines();
    let input = TEST_INPUT
        .lines()
        .map(|line| line.to_string())
        .collect::<Vec<String>>();
    println!("Day 6 part 1: {}", part_one(input));
    println!("Day 6 part 2: {}", part_two());
}

fn part_one(coords: Vec<String>) -> i32 {
    let mut coords = coords
        .iter()
        .map(|line| {
            let mut split = line.split(", ");
            let x = split.next().unwrap().parse::<i32>().unwrap();
            let y = split.next().unwrap().parse::<i32>().unwrap();
            (x, y)
        })
        .collect::<Vec<(i32, i32)>>();

    let min_x = coords.iter().min_by_key(|(x, _)| x).unwrap().0;
    let max_x = coords.iter().max_by_key(|(x, _)| x).unwrap().0;
    let min_y = coords.iter().min_by_key(|(_, y)| y).unwrap().1;
    let max_y = coords.iter().max_by_key(|(_, y)| y).unwrap().1;

    let mut grid = vec![vec![0; (max_x - min_x + 1) as usize]; (max_y - min_y + 1) as usize];

    for (i, (x, y)) in coords.iter().enumerate() {
        grid[(*y - min_y) as usize][(*x - min_x) as usize] = i as i32;
    }

    // For each point in the grid, find the closest coordinate
    for (y, row) in grid.iter_mut().enumerate() {
        for (x, point) in row.iter_mut().enumerate() {
            let mut min_distance = std::i32::MAX;
            let mut min_coord = 0;
            for (i, (coord_x, coord_y)) in coords.iter().enumerate() {
                let distance = (x as i32 - coord_x).abs() + (y as i32 - coord_y).abs();
                if distance < min_distance {
                    min_distance = distance;
                    min_coord = i;
                } else if distance == min_distance {
                    min_coord = -1;
                }
            }
            *point = min_coord;
        }
    }

    0
}

fn part_two() -> i32 {
    0
}

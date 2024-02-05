use crate::utils::point::Point;

pub fn solve_day_eleven() {
    let input = 9110;
    println!("Day 11 part 1: {}", part_one(input));
    println!("Day 11 part 2: {}", part_two(input));
}

fn part_one(input: i32) -> String {
    let points = create_grid(input);
    let max_top_left_corner = find_max_power(&points, 3);
    format!(
        "{},{}",
        max_top_left_corner.1.x + 1,
        max_top_left_corner.1.y + 1
    )
    .into()
}

fn part_two(input: i32) -> String {
    let points = create_grid(input);
    let mut max = 0;
    let mut max_tlc = Point::new(0, 0);
    let mut max_size = 0;
    for size in 1..=300 {
        let (power, point) = find_max_power(&points, size);
        if power == 0 {
            break;
        }

        if power > max {
            max = power;
            max_tlc = point;
            max_size = size;
        }
    }
    format!("{},{},{}", max_tlc.x + 1, max_tlc.y + 1, max_size)
}

fn create_grid(serial: i32) -> Vec<Vec<i32>> {
    (1..=300)
        .map(|x| {
            (1..=300)
                .map(move |y| {
                    // X coord plus 10
                    let rack_id = x + 10;
                    // power level = rack ID times the Y coord
                    let mut power_level = rack_id * y;
                    // Increase power level by serial number
                    power_level += serial;
                    // Set the power level to itself multiplied by the rack ID
                    power_level *= rack_id;
                    // Keep only the hundres digits of the power level
                    power_level = (power_level % 1000) / 100;
                    // Subtract 5
                    power_level -= 5;
                    power_level
                })
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>()
}

fn find_max_power(points: &Vec<Vec<i32>>, sub_grid_size: usize) -> (i32, Point) {
    let mut max = i32::MIN;
    let mut max_top_left_corner = Point::new(0, 0);
    for row in 0..(points.len() - sub_grid_size - 1) {
        for col in 0..(points.len() - sub_grid_size - 1) {
            let mut power = 0;
            for j in 0..sub_grid_size {
                for k in 0..sub_grid_size {
                    power += points[j + row][k + col]
                }
            }

            if power > max {
                max = power;
                max_top_left_corner = Point::new(row as i32, col as i32)
            }
        }
    }
    (max, max_top_left_corner)
}

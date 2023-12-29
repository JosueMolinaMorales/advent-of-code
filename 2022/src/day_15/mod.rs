use std::collections::HashSet;

const INPUT: &str = include_str!("input_15.txt");

fn merge_range(
    ranges: &Vec<((isize, isize), (isize, isize))>,
    acceptable_ranges: Option<(isize, isize)>,
) -> Vec<isize> {
    let mut range: HashSet<isize> = HashSet::new();
    for r in ranges {
        for i in r.0 .0..=r.1 .0 {
            match acceptable_ranges {
                None => {
                    // print!("{i} ");
                    range.insert(i);
                }
                Some((min, max)) => {
                    if i >= min && i <= max {
                        // print!("{i} ");
                        range.insert(i);
                    }
                }
            }
        }
        // println!()
    }
    let mut range = range.iter().map(|s| *s).collect::<Vec<isize>>();
    range.sort();
    // println!("{:?}", range);
    range
}

fn part_one(points: &Vec<Vec<(isize, isize)>>) -> usize {
    const HORIZONTAL_LINE: isize = 2000000;
    let mut beacons: HashSet<(isize, isize)> = HashSet::new();
    let mut ranges: Vec<((isize, isize), (isize, isize))> = vec![];
    for point in points {
        let scanner = point[0];
        let beacon = point[1];
        let manhattan_dist = (scanner.0.abs_diff(beacon.0) + scanner.1.abs_diff(beacon.1)) as isize;
        let distance_to_line = scanner.1.abs_diff(HORIZONTAL_LINE) as isize;
        let amount_of_hashtags = (2 * manhattan_dist + 1) - 2 * distance_to_line;
        if scanner.1 <= HORIZONTAL_LINE && (scanner.1 + manhattan_dist) >= HORIZONTAL_LINE
            || scanner.1 >= HORIZONTAL_LINE && (scanner.1 - manhattan_dist) <= HORIZONTAL_LINE
        {
            let count_on_either_side = (amount_of_hashtags - 1) / 2;
            let range = (
                (scanner.0 - count_on_either_side, HORIZONTAL_LINE),
                (scanner.0 + count_on_either_side, HORIZONTAL_LINE),
            );
            ranges.push(range);
        }
        if beacon.1 == HORIZONTAL_LINE {
            beacons.insert(beacon);
        }
    }
    let ranges = merge_range(&ranges, None);
    ranges.len() - beacons.len()
}

fn part_two(points: &Vec<Vec<(isize, isize)>>) -> i64 {
    const MAX_COORD: i32 = 4000000;
    // Rectangles (start corner and end corner, inclusive). Begin with one covering whole possible area
    let mut possibilities = vec![([-MAX_COORD, 0], [MAX_COORD, 2 * MAX_COORD])];
    let mut new_possibilities = Vec::new();
    for point in points {
        let sensor = [point[0].0 as i32, point[0].1 as i32];
        let beacon = [point[1].0 as i32, point[1].1 as i32];
        let radius = sensor
            .iter()
            .zip(&beacon)
            .map(|(&a, &b)| (a - b).abs())
            .sum::<i32>();
        let center = [sensor[0] - sensor[1], sensor[0] + sensor[1]];
        let start = [center[0] - radius, center[1] - radius];
        let end = [center[0] + radius, center[1] + radius];

        for &p in &possibilities {
            let (p_start, p_end) = p;
            if !(0..2).all(|i| start[i] <= p_end[i] && p_start[i] <= end[i]) {
                new_possibilities.push(p);
            } else {
                if start[0] > p_start[0] {
                    new_possibilities.push((p_start, [start[0] - 1, p_end[1]]));
                }
                if p_end[0] > end[0] {
                    new_possibilities.push(([end[0] + 1, p_start[1]], p_end));
                }
                if start[1] > p_start[1] {
                    new_possibilities.push((
                        [std::cmp::max(start[0], p_start[0]), p_start[1]],
                        [std::cmp::min(end[0], p_end[0]), start[1] - 1],
                    ));
                }
                if p_end[1] > end[1] {
                    new_possibilities.push((
                        [std::cmp::max(start[0], p_start[0]), end[1] + 1],
                        [std::cmp::min(end[0], p_end[0]), p_end[1]],
                    ));
                }
            }
        }
        possibilities.clear();

        std::mem::swap(&mut possibilities, &mut new_possibilities);
    }

    // Assume there is a 1x1 rectangle somewhere within the allowed area
    let mut ans = 0;
    for (start, end) in possibilities {
        if start == end && (start[0] + start[1]) % 2 == 0 {
            // Transform back into original coordinate system
            let pos = [(start[1] + start[0]) / 2, (start[1] - start[0]) / 2];
            if pos.iter().all(|&x| x >= 0 && x <= MAX_COORD) {
                ans = pos[0] as i64 * MAX_COORD as i64 + pos[1] as i64;
                break;
            }
        }
    }

    ans
}

pub fn solve_day_15() {
    let points = INPUT
        .lines()
        .map(|info| {
            let info = info.trim().replace("Sensor at ", "");
            info.split(": closest beacon is at ")
                .map(|p| {
                    let points = p
                        .split(", ")
                        .map(|s| {
                            let s = s.replace("x=", "");
                            let s = s.replace("y=", "");
                            s.clone()
                        })
                        .collect::<Vec<String>>();
                    (
                        points[0].parse::<isize>().unwrap(),
                        points[1].parse::<isize>().unwrap(),
                    )
                })
                .collect::<Vec<(isize, isize)>>()
        })
        .collect::<Vec<Vec<(isize, isize)>>>();
    println!("Day 15 Part One: {}", part_one(&points));
    println!("Day 15 Part Two: {}", part_two(&points));
}

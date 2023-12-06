use std::collections::VecDeque;

#[derive(Debug, Clone)]
struct Point {
    character: String,
    position: (usize, usize),
    parent: Option<(usize, usize)>,
    explored: bool
}

const INPUT: &str = include_str!("day_12_input.txt");

fn parse_input() -> Vec<Vec<Point>> {
    let mut grid: Vec<Vec<Point>> = vec![];
    let mut row_count = 0;
    let mut column_count = 0;
    let mut lines = INPUT.split("\n").collect::<Vec<&str>>();
    lines.reverse();
    for line in lines {
        let mut row: Vec<Point> = vec![];
        for char in line.chars() {
            row.push(Point { character: char.to_string(), position: (column_count, row_count), parent: None, explored: false});
            column_count += 1;
        }
        grid.push(row);
        row_count += 1;
        column_count = 0;
    }
    grid
}

fn find_item(grid: &Vec<Vec<Point>>, item: String) -> Vec<(usize, usize)> {
    let mut res = vec![];
    for row in grid.iter(){
        for column in row.iter() {
            let col_char = if column.character == "S" && item != "S" { 
                "a".to_string() 
            } else if column.character == "E" && item != "E" { 
                "z".to_string() 
            } else {
                column.character.clone()
            };
            if col_char == item {
                res.push(column.position);
            }
        }
    }
    res
}

/**
 * A vertex, v, is adjacent to u, if and only if
 *  - u is above, to the left, to the right, or below v by one space
 *  - u.char is 1 above or equal to v.char
 *  - u.char is strictly less than v.char
 */
fn get_adjacent_items(grid: &Vec<Vec<Point>>, vertex: (usize, usize)) -> Vec<(usize, usize)> {
    let mut res: Vec<(usize, usize)> = vec![];

    // Top
    if let Some(top_row) = grid.get(vertex.1 + 1) {
        if let Some(top_item) = top_row.get(vertex.0) {
            res.push(top_item.position);
        }
    }

    if let Some(current_row) = grid.get(vertex.1) {
        // Right
        if let Some(right_item) = current_row.get(vertex.0 + 1) {
            res.push(right_item.position);
        }
        // Left
        if vertex.0 != 0 {
            if let Some(left_item) = current_row.get(vertex.0 - 1) {
                res.push(left_item.position)
            }
        }
    }
    
    // Bottom
    if vertex.1 != 0 {
        if let Some(bottom_row) = grid.get(vertex.1 - 1) {
            if let Some(bottom_item) = bottom_row.get(vertex.0) {
                res.push(bottom_item.position);
            }
        }
    }
    // Validate all the adjacent items
    let vertex_point = grid[vertex.1][vertex.0].clone();
    let mut adj_vec = vec![];
    for i in 0..res.len() {
        let adj_point = grid[res[i].1][res[i].0].clone();
        let adj_point_char = if adj_point.character == "E" { "z".to_string() } else { adj_point.character };
        if vertex_point.character == "E" {
            if adj_point_char == "z".to_string() {
                adj_vec.push(adj_point.position);
            }
        } else if vertex_point.character == "S" {
            if adj_point_char == "a".to_string() {
                adj_vec.push(adj_point.position);
            }
        } else {
            if 
                vertex_point.character.as_bytes()[0] + 1 == adj_point_char.as_bytes()[0] ||
                vertex_point.character.as_bytes()[0] >= adj_point_char.as_bytes()[0]
            {
                adj_vec.push(adj_point.position);
            }
        }
    }
    adj_vec
}

fn bfs_climbing(starting_point: String) {
    let mut grid = parse_input();

    let start = find_item(&grid, starting_point);
    let mut count_vec: Vec<u32> = vec![];
    let fresh_grid = grid.clone();

    for starting in start {
        grid = fresh_grid.clone();
        let mut queue = VecDeque::new();
        grid[starting.1][starting.0].explored = true;
        queue.push_front(starting);
        while let Some(v) = queue.pop_front() {
            if grid[v.1][v.0].character == "E" {
                break;
            }
    
            let adj = get_adjacent_items(&grid, v);
            for u in adj {
                let adj_u = &grid[u.1][u.0];
                if !adj_u.explored {
                    grid[u.1][u.0].parent = Some(v);
                    grid[u.1][u.0].explored = true;
                    queue.push_back(u);
                }
            }
        }

        let end = find_item(&grid, "E".to_string()).last().unwrap().clone();
        let mut count = 0;
        let mut current_point = grid[end.1][end.0].clone();
    
        while let Some(parent) = current_point.parent {
            count += 1;
            current_point = grid[parent.1][parent.0].clone();
        }
    
        if count != 0 {
            // println!("count: {}", count);
            count_vec.push(count);
        }

    }

    // get min count in vec
    count_vec.sort();
    println!("Min: {}", count_vec[0])
}

pub fn solve_day_twelve() {
    // Part 1
    bfs_climbing("S".to_string());
    // Part 2
    bfs_climbing("a".to_string());

    
}
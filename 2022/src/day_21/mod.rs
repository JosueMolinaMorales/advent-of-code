use std::{collections::HashMap, time};

const INPUT: &str = include_str!("input.txt");

#[derive(Debug, Clone)]
struct Node<'a> {
    name: &'a str,
    operation: String             // Holds the operation to happen, or number: eg. x * y or 4
}

pub fn solve_day_21() {
    let mut tree: HashMap<&str, Node> = HashMap::new();
    INPUT.lines().for_each(|line| {
        let mut line = line.trim().split(": ");
        let name = line.next().unwrap();
        tree.insert(name.clone(), Node { name, operation: line.next().unwrap().to_string() });
    });

    // Part 1
    let time = time::Instant::now();
    let sum = traverse(&mut tree, "root");
    println!("{sum}");
    println!("time passed: {:?}", time.elapsed());

    let time = time::Instant::now();
    // Part 2
    part_2(tree);
    println!("time passed: {:?}", time.elapsed());
}

fn part_2<'a>(mut tree: HashMap<&str, Node>) {
    // traverse one side of tree
    let node = tree.get("root").unwrap().clone();

    // Operation requires two other nodes
    let mut operations = node.operation.split(" ");
    let first_node = operations.next().unwrap();
    operations.next();
    let second_node = operations.next().unwrap();
    let second_node = traverse(&mut tree, second_node);
    println!("second_node: {}", second_node);

    let mut left_count = traverse(&mut tree, first_node); // First node contains 5.
    let mut humn_count: i64 = 0;
    while left_count != second_node {
        let humn_str = humn_count.to_string();
        tree.entry("humn").and_modify(|n| n.operation = humn_str);
        left_count = traverse(&mut tree, first_node);
        let diff = left_count - second_node;
        if diff > 2_000_000_000 {
            humn_count += 100_000_000;
        } else if diff > 200_000_000 {
            humn_count += 10_000_000;
        } else if diff > 100_000_000 {
            humn_count += 100_000;
        } else if diff > 100_000 {
            humn_count += 100;
        } else {
            humn_count += 1;
        }
    }

    println!("humn_count: {}", humn_count);
    println!("first_node count: {}", left_count);
    println!("second_node count: {}", second_node);

}

fn traverse<'a>(tree: &HashMap<&str, Node>, curr_node: &'a str) -> isize {
    let node = tree.get(curr_node).unwrap();
    if let Some(num) = node.operation.parse::<isize>().ok() {
        return num
    } else {
        // Operation requires two other nodes
        let mut operations = node.operation.split(" ");
        let first_node = operations.next().unwrap();
        let op = operations.next().unwrap();
        let second_node = operations.next().unwrap();

        let first_node = traverse(tree, first_node);
        let second_node = traverse(tree, second_node);

        return match op {
            "+" => first_node + second_node,
            "-" => first_node - second_node,
            "*" => first_node * second_node,
            "/" => first_node / second_node,
            _ => unreachable!()
        }
    }
}
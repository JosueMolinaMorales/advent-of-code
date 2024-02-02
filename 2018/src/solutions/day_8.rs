use std::collections::HashMap;

use crate::utils::file_loader::FileLoader;

pub fn solve_day_eight() {
    let input = FileLoader::new("./inputs/day8.txt".into()).read();
    println!("Day 8 part 1: {}", part_one(input.clone()));
    println!("Day 8 part 2: {}", part_two(input));
}

fn part_one(input: String) -> i32 {
    let mut tree = HashMap::new();
    build_tree(
        &mut 0,
        &input
            .split_whitespace()
            .map(|n| n.parse::<i32>().unwrap())
            .collect(),
        &mut tree,
    );
    tree.iter().fold(0, |acc, (_, node)| {
        acc + node.meta_entries.iter().sum::<usize>() as i32
    })
}

fn part_two(input: String) -> i32 {
    let mut tree = HashMap::new();
    build_tree(
        &mut 0,
        &input
            .split_whitespace()
            .map(|n| n.parse::<i32>().unwrap())
            .collect(),
        &mut tree,
    );

    get_node_value(0, &tree)
}

fn get_node_value(node: usize, tree: &HashMap<usize, Node>) -> i32 {
    let node = match tree.get(&node) {
        Some(node) => node,
        None => return 0,
    };

    if node.children.is_empty() {
        return node.meta_entries.iter().sum::<usize>() as i32;
    }

    let mut sum = 0;
    for child in node.meta_entries.iter() {
        let child = match node.children.get(child - 1) {
            Some(child) => child,
            None => continue,
        };
        sum += get_node_value(*child, tree)
    }

    sum
}

fn build_tree(curr_idx: &mut usize, nums: &Vec<i32>, nodes: &mut HashMap<usize, Node>) {
    let key = curr_idx.clone();
    let child_count = nums[*curr_idx];
    let meta_count = nums[*curr_idx + 1];

    let mut node = Node {
        idx: curr_idx.clone(),
        children: vec![],
        meta_entries: vec![],
    };

    let mut child_idx = *curr_idx + 2;
    for _ in 0..child_count {
        node.children.push(child_idx);
        build_tree(&mut child_idx, nums, nodes);
    }

    // The next points are metas
    for meta in 0..meta_count {
        node.meta_entries
            .push(nums[child_idx + meta as usize] as usize)
    }

    // Move the curr_idx pointer
    *curr_idx = child_idx + meta_count as usize;

    nodes.entry(key).or_insert(node);
}

#[derive(Debug)]
struct Node {
    pub idx: usize,
    pub children: Vec<usize>,
    pub meta_entries: Vec<usize>,
}

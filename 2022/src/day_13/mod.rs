use std::cmp::Ordering;

use serde_json::{Value, json};


const INPUT: &str = include_str!("input_day_13.txt");

fn compare(left: &Value, right: &Value) -> Option<Ordering>{
    match (left, right) {
        (Value::Number(nl), Value::Number(nr)) => {
            match nl.as_u64().unwrap().cmp(&nr.as_u64().unwrap()) {
                Ordering::Equal => None,
                order => Some(order)
            }
        },
        (Value::Array(left), Value::Array(right)) => {
            if left.is_empty() || right.is_empty() {
                match left.len().cmp(&right.len()) {
                    Ordering::Equal => None,
                    order => Some(order),
                }
            } else if let Some(v) = compare(&left[0], &right[0]) {
                Some(v)
            } else {
                compare(&json!(left[1..]), &json!(right[1..]))
            }
        },
        (Value::Number(left), Value::Array(right)) => {
            compare(&json!(vec![left]), &json!(right))
        },
        (Value::Array(left), Value::Number(right)) => {
            compare(&json!(left), &json!(vec![right]))
        },
        (_, _) => unreachable!()
    }
}


pub fn solve_day_13() {
    // Part 1
    let solution = INPUT
        .split("\n\n")
        .enumerate()
        .filter_map(|(i, packets)| {
            let packets = packets.split_whitespace().collect::<Vec<&str>>();
            let left: Value = serde_json::from_str(packets[0]).unwrap();
            let right: Value = serde_json::from_str(packets[1]).unwrap();

            if compare(&left, &right) == Some(Ordering::Less) {
                return Some(i+1)
            }
            None
        }).sum::<usize>();
    println!("solution: {:?}", solution);

    // Part 2
    let mut solution = INPUT
    .split("\n")
    .filter_map(|packets| {
        let packets = packets.trim();
        if packets.is_empty() {
            return None
        }
        let packet: Value = serde_json::from_str(packets).unwrap();
        Some(packet)
    })
    .collect::<Vec<Value>>();
    solution.push(json!([[2]]));
    solution.push(json!([[6]]));
    solution.sort_by(|a, b| compare(a, b).unwrap());
    let product: usize = solution.iter().enumerate().filter_map(|v| {
        if v.1.to_string() == "[[2]]" || v.1.to_string() == "[[6]]" {
            Some(v.0 + 1)
        } else {
            None
        }
    }).product();
    println!("product {:?}", product);
}

#[test]
fn test_packet_left_is_less_than() {
    let left: Value = serde_json::from_str("[1,1,3,1,1]").unwrap();
    let right: Value = serde_json::from_str("[1,1,5,1,1]").unwrap();
    assert!(compare(&left, &right) == Some(Ordering::Less));
}

#[test]
fn test_packet_left_is_less_than2() {
    let left: Value = serde_json::from_str("[1,[2,[3,[4,[5,6,7]]]],8,9]").unwrap();
    let right: Value = serde_json::from_str("[1,[2,[3,[4,[5,6,0]]]],8,9]").unwrap();
    assert!(compare(&left, &right) == Some(Ordering::Greater));
}

#[test]
fn test_packet_left_is_less_than3() {
    let left: Value = serde_json::from_str("[[1],[2,3,4]]").unwrap();
    let right: Value = serde_json::from_str("[[1],4]").unwrap();
    assert!(compare(&left, &right) == Some(Ordering::Less));
}

#[test]
fn test_packet_4() {
    let left: Value = serde_json::from_str("[[[]]]").unwrap();
    let right: Value = serde_json::from_str("[[]]").unwrap();
    assert!(compare(&left, &right) == Some(Ordering::Greater));
}

#[test]
fn test_packet_5() {
    let left: Value = serde_json::from_str("[[4,4],4,4]").unwrap();
    let right: Value = serde_json::from_str("[[4,4],4,4,4]").unwrap();
    assert!(compare(&left, &right) == Some(Ordering::Less));
}
use std::collections::{HashMap, HashSet};

use regex::Regex;

use crate::utils::file_loader;

#[derive(Debug)]
struct Sample {
    old_reg: HashMap<i32, i32>,
    program: Vec<i32>,
    new_reg: HashMap<i32, i32>,
}

pub fn solve_day_sixteen() {
    let binding = file_loader::FileLoader::new("./inputs/day_16.txt".to_string()).read();
    let mut input = binding.split("\n\n\n\n");

    let re = Regex::new(r"\d+").unwrap();
    let samples = input
        .next()
        .unwrap()
        .split("\n\n")
        .map(|line| {
            let mut nums = re
                .find_iter(line)
                .map(|ds| ds.as_str().parse::<i32>().unwrap());
            Sample {
                old_reg: HashMap::from([
                    (0, nums.next().unwrap()),
                    (1, nums.next().unwrap()),
                    (2, nums.next().unwrap()),
                    (3, nums.next().unwrap()),
                ]),
                program: vec![
                    nums.next().unwrap(),
                    nums.next().unwrap(),
                    nums.next().unwrap(),
                    nums.next().unwrap(),
                ],
                new_reg: HashMap::from([
                    (0, nums.next().unwrap()),
                    (1, nums.next().unwrap()),
                    (2, nums.next().unwrap()),
                    (3, nums.next().unwrap()),
                ]),
            }
        })
        .collect::<Vec<Sample>>();
    let program = input
        .next()
        .unwrap()
        .split("\n")
        .map(|line| {
            let mut nums = re
                .find_iter(line)
                .map(|ds| ds.as_str().parse::<i32>().unwrap());
            vec![
                nums.next().unwrap(),
                nums.next().unwrap(),
                nums.next().unwrap(),
                nums.next().unwrap(),
            ]
        })
        .collect::<Vec<Vec<i32>>>();
    println!("{:?}", program);
    // println!("{:?}", samples);
    // Four values; OpCode InA InB OutC
    let actions: Vec<(&str, fn(i32, i32, i32, &mut HashMap<i32, i32>))> = vec![
        ("addr", addr),
        ("addi", addi),
        ("mulr", mulr),
        ("muli", muli),
        ("banr", banr),
        ("bani", bani),
        ("borr", borr),
        ("bori", bori),
        ("setr", setr),
        ("seti", seti),
        ("gtir", gtir),
        ("gtri", gtri),
        ("gtrr", gtrr),
        ("eqir", eqir),
        ("eqri", eqri),
        ("eqrr", eqrr),
    ];

    let mut samples_found = 0;
    let mut possible: HashMap<i32, HashSet<&str>> = HashMap::new();
    let mut ops: HashMap<i32, Vec<&str>> = HashMap::new();
    for i in 0..=15 {
        possible.insert(
            i,
            actions.clone().iter().map(|x| x.0).collect::<HashSet<_>>(),
        );
    }
    for sample in samples.iter() {
        let mut registers = sample.old_reg.clone();
        let mut found = 0;
        for (op, action) in actions.iter() {
            action(
                sample.program[1],
                sample.program[2],
                sample.program[3],
                &mut registers,
            );
            if registers == sample.new_reg {
                // println!("FOUND! {}", op);
                ops.entry(sample.program[0])
                    .and_modify(|v| {
                        v.push(op);
                    })
                    .or_insert(Vec::from([*op]));
                found += 1
            } else if possible[&sample.program[0]].contains(op) {
                possible.entry(sample.program[0]).and_modify(|v| {
                    v.remove(op);
                });
            }
            // Reset registers
            registers = sample.old_reg.clone()
        }
        if found >= 3 {
            samples_found += 1
        }
    }

    let mut mapping = HashMap::new();

    while possible.values().any(|x| x.len() > 0) {
        let keys_to_modify: Vec<i32> = possible
            .iter()
            .filter(|(_, opcodes)| opcodes.len() == 1)
            .map(|(k, _)| *k)
            .collect();
        for k in keys_to_modify {
            let op_code = possible[&k].iter().next().unwrap().to_string();
            mapping.insert(k, op_code.clone());
            for v in possible.values_mut() {
                v.remove(op_code.as_str());
            }
        }
    }

    // P2
    let mut registers = HashMap::from([(0, 0), (1, 0), (2, 0), (3, 0)]);

    for p in program {
        let op_code = p[0];
        let action = mapping[&op_code].clone();
        let f = actions
            .iter()
            .find(|(o, _)| *o == action.as_str())
            .unwrap()
            .1;
        f(p[1], p[2], p[3], &mut registers);
    }
    println!("Day 16 Part 1: {}", samples_found);
    println!("Day 16 Part 2: {}", registers[&0]);
}

fn addr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] + registers[&b]);
}

fn addi(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] + b);
}

fn mulr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] * registers[&b]);
}

fn muli(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] * b);
}

fn banr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] & registers[&b]);
}

fn bani(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] & b);
}

fn borr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] | registers[&b]);
}

fn bori(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a] | b);
}

fn setr(a: i32, _: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, registers[&a]);
}

fn seti(a: i32, _: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, a);
}

fn gtir(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if a > registers[&b] { 1 } else { 0 });
}

fn gtri(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if registers[&a] > b { 1 } else { 0 });
}

fn gtrr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if registers[&a] > registers[&b] { 1 } else { 0 });
}

fn eqir(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if a == registers[&b] { 1 } else { 0 });
}

fn eqri(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if registers[&a] == b { 1 } else { 0 });
}

fn eqrr(a: i32, b: i32, c: i32, registers: &mut HashMap<i32, i32>) {
    registers.insert(c, if registers[&a] == registers[&b] { 1 } else { 0 });
}

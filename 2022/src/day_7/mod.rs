use std::{fs, collections::HashMap };

// const TEST_INPUT: &str = "./input7.txt";
const INPUT_FILE: &str = "./inputs/day_7_input.txt";

#[derive(Debug, Clone)]
struct Directory {
    name: String,
    size: u32,
    children: Vec<String>
}


pub fn solve_day_seven() {
    let input = fs::read_to_string(INPUT_FILE).unwrap();

    let mut directory_structure: HashMap<String, Directory> = HashMap::new();
    let mut stack: Vec<String> = Vec::new(); // Stores the history of directories. 

    for line in input.split('\n') {
        match line.chars().next().unwrap() {
            '$' => {
                let exec = line.split(' ').map(|s| s.trim()).collect::<Vec<&str>>();
                if exec[1] != "cd" {
                    continue
                }
                let dir_name = exec[2].to_string();
                if dir_name == ".." {
                    stack.pop();
                    continue;
                }
                if dir_name == "/" {
                    directory_structure.insert("/".to_string(), Directory { name: "/".to_string(), size: 0, children: Vec::new() });
                }

                stack.push(dir_name);
            },
            _ => {
                // ls output
                let output = line.split(' ').map(|s| s.trim()).collect::<Vec<&str>>();
                match output[0] {
                    "dir" => {
                        let dir_name = output[1].to_string();

                        directory_structure
                        .entry(format!("{}/{}", stack.join("/"), dir_name.clone()))
                        .or_insert(Directory {
                            name: format!("{}/{}", stack.join("/"), dir_name.clone()),
                            size: 0,
                            children: Vec::new()
                        });

                        // add dir as a child of the current directory
                        directory_structure
                            .entry(stack.join("/").to_string())
                            .and_modify(|dir| {
                                dir.children.push(format!("{}/{}", stack.join("/"), dir_name.clone()));
                            });
                        
                    },
                    size => {
                        let size = size.parse::<u32>().unwrap(); 
                        directory_structure
                            .entry(stack.join("/").to_string())
                            .and_modify(|dir| {
                                dir.size += size;
                            });
                    }
                }
            }
        }
    }

    let mut sizes = vec![];
    for dir in directory_structure.keys() {
        let size = directory_size(&directory_structure, dir);
        sizes.push(size);
    }

    let res = sizes.iter().filter_map(|s| if s.1 < 100000 { Some(s.1) } else { None }).sum::<u32>();
    println!("part 1 res: {}", res);

    sizes.sort();
    let total_space_available = 70_000_000;
    let required_unused_space = 30_000_000;
    let total_space_used: u32 = sizes.iter().map(|s| s.1).max().unwrap();
    let current_unused_space = total_space_available - total_space_used;
    let amount_to_free = required_unused_space - current_unused_space;

    let  part_two = sizes.iter().map(|s| s.1).filter(|size| *size >= amount_to_free).min().unwrap();
    println!("part 2: {:?}", part_two);
    
}

fn directory_size(directory_structure: &HashMap<String, Directory>, dir: &str) -> (String, u32) {
    let mut size = 0;

    let directory = directory_structure.get(dir).unwrap();
    for child in &directory.children {
        size += directory_size(directory_structure, child).1;
    }
    size += directory.size;

    (directory.name.clone(), size)
}
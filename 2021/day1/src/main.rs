use std::fs;

fn calculate_window(window: &mut Vec<&str>) -> u64 {
    let mut window_depth = 0;
    for depth in window {
        let depth_int: u64 = depth.parse().unwrap();
        window_depth = window_depth + depth_int;
    }
    return window_depth;
}

fn part_two() {
    let file_input = fs::read_to_string("input").expect("Something went wrong reading the file");

    let depth_list: Vec<&str> = file_input.split("\n").collect();
    let mut last_depth_window = calculate_window(&mut depth_list[0..3].to_vec());
    let mut bigger = 0;
    for (i, _depth) in depth_list.iter().enumerate() {
        if depth_list.len() < i+3 {
            break;
        }
        let depth_window: u64 = calculate_window(&mut depth_list[i..i+3].to_vec());
        if depth_window > last_depth_window {
            bigger+=1;
        }
        last_depth_window = depth_window;
    }
    println!("solution part one: {}", bigger);
}

fn part_one() {
    let file_input = fs::read_to_string("input").expect("Something went wrong reading the file");

    let depth_list: Vec<&str> = file_input.split("\n").collect();
    let mut last_depth: i32 = depth_list[0].parse().unwrap();
    let mut bigger = 0;
    for depth in depth_list {
        let depth_int: i32 = depth.parse().unwrap();
        if depth_int > last_depth {
            bigger+=1;
        }
        last_depth = depth_int;
    }
    println!("solution part two: {}", bigger);
}

fn main(){
    part_two();
    part_one();
}
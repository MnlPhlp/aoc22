use crate::types::Task;

fn parse_stacks(input: &str, test: bool) -> Vec<Vec<u8>> {
    let mut stack_lines = vec![];
    let mut stack_count = 0;

    for l in input.lines() {
        if l.trim_start().starts_with("[") {
            stack_lines.push(l);
        } else {
            stack_count = l.bytes().filter(|b| *b != b' ').count();
            break;
        }
    }
    if test {
        println!("creating {stack_count} stacks from following lines:");
        for line in &stack_lines {
            println!("{line}");
        }
    }
    let mut stacks = vec![];
    for _ in 0..stack_count {
        stacks.push(vec![]);
    }
    for line in stack_lines.iter().rev() {
        for j in 0..stack_count {
            let item = line.as_bytes()[1 + j * 4];
            if item != b' ' {
                stacks[j].push(item)
            }
        }
    }
    return stacks;
}

fn parse_input(input: &str, test: bool) -> (Vec<Vec<u8>>, Vec<[usize; 3]>) {
    let stacks = parse_stacks(input, test);
    if test {
        println!("parsed following stacks:");
        for s in &stacks {
            println!("{s:?}");
        }
    }
    let instructions = parse_instructions(input);
    if test {
        println!("parsed following instructions:\n{instructions:?}");
    }
    return (stacks, instructions);
}

fn parse_instructions(input: &str) -> Vec<[usize; 3]> {
    let mut instructions = vec![];
    for l in input.lines() {
        if l.starts_with("move") {
            let parts: Vec<&str> = l.split(" ").collect();
            let mut inst = [0; 3];
            inst[0] = parts[1].parse::<usize>().unwrap();
            inst[1] = parts[3].parse::<usize>().unwrap();
            inst[1] -= 1;
            inst[2] = parts[5].parse::<usize>().unwrap();
            inst[2] -= 1;
            instructions.push(inst);
        }
    }
    return instructions;
}

fn do_moves(
    stacks: &mut Vec<Vec<u8>>,
    instructions: &Vec<[usize; 3]>,
    move_at_once: bool,
    _test: bool,
) {
    for inst in instructions {
        let count = inst[0];
        let from_stack = inst[1];
        let to_stack = inst[2];
        if move_at_once {
            let mut items = stacks[from_stack][stacks[from_stack].len() - count..].to_owned();
            let stack_len = stacks[from_stack].len();
            stacks[from_stack].truncate(stack_len - count);
            stacks[to_stack].append(&mut items);
        } else {
            for _i in 0..count {
                let item = stacks[from_stack].pop().unwrap();
                stacks[to_stack].push(item);
            }
        }
    }
}

pub fn solve(input: &str, test: bool, task: Task) -> (String, String) {
    let mut res1 = String::from("");
    let mut res2 = String::from("");
    let (mut stacks, instructions) = parse_input(input, test);
    let mut stacks2 = stacks.clone();

    // move one item at a time
    if !matches!(task, Task::Two) {
        do_moves(&mut stacks, &instructions, false, test);
        res1 = "".into();
        for s in stacks {
            if s.len() > 0 {
                res1.push_str(&format!("{}", *s.last().unwrap() as char));
            }
        }
    }

    if !matches!(task, Task::One) {
        // move all items at once
        do_moves(&mut stacks2, &instructions, true, test);
        res2 = "".into();
        for s in stacks2 {
            if s.len() > 0 {
                res2.push_str(&(*s.last().unwrap() as char).to_string());
            }
        }
    }

    return (res1, res2);
}

use crate::types::Task;

fn get_move(m: u8) -> u8 {
    if m < b'D' {
        return m - b'A';
    } else {
        return m - b'X';
    }
}

fn calc_score(line: &str) -> i32 {
    let opponent = get_move(line.bytes().nth(0).unwrap());
    let me = get_move(line.bytes().nth(2).unwrap());
    // score for pick
    let mut score = me as i32 + 1;
    // score for draw
    if me == opponent {
        score += 3
    }
    // score for winning
    if (me + 2) % 3 == opponent {
        score += 6
    }
    score
}

fn decide_move(opponent: u8, me: u8) -> u8 {
    match me {
        b'X' => {
            // pick losing move
            (opponent + 2) % 3
        }
        b'Z' => {
            // pick winning move
            (opponent + 1) % 3
        }
        _ => {
            // pick draw move
            opponent
        }
    }
}

fn calc_score2(line: &str) -> i32 {
    let opponent = get_move(line.bytes().nth(0).unwrap());
    let me = decide_move(opponent, line.bytes().nth(2).unwrap());
    // score for pick
    let mut score = me as i32 + 1;
    // score for draw
    if me == opponent {
        score += 3
    }
    // score for winning
    if (me + 2) % 3 == opponent {
        score += 6
    }
    score
}

pub fn solve(input: &str, _debug: bool, task: Task) -> (String, String) {
    let (mut res1, mut res2) = ("".to_string(), "".to_string());
    let lines: Vec<&str> = input.lines().collect();
    if !matches!(task, Task::Two) {
        let mut score = 0;
        for line in &lines {
            if line.len() == 0 {
                continue;
            }
            score += calc_score(line)
        }
        res1 = score.to_string();
    }

    if !matches!(task, Task::One) {
        let mut score = 0;
        for line in &lines {
            if line.len() == 0 {
                continue;
            }
            score += calc_score2(line)
        }
        res2 = score.to_string();
    }
    (res1, res2)
}

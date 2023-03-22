use crate::types::Task;

fn sum_calories(elves: &Vec<&str>) -> Vec<i32> {
    elves
        .iter()
        .map(|elf| {
            elf.lines()
                .fold(0, |sum, line| sum + line.parse::<i32>().unwrap())
        })
        .collect()
}

pub fn solve(input: &str, _test: bool, task: Task) -> (String, String) {
    let mut res1 = "".to_string();
    let mut res2 = "".to_string();

    let elves = input.split("\n\n").collect();
    let mut calories = sum_calories(&elves);
    calories.sort();

    if !matches!(task, Task::Two) {
        let max_calories = calories.last().unwrap();
        res1 = max_calories.to_string();
    }
    if !matches!(task, Task::One) {
        let top_calories = calories.iter().rev().take(3).sum::<i32>();
        res2 = top_calories.to_string();
    }

    (res1, res2)
}

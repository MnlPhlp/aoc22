use crate::types::Task;

#[derive(Default, Copy, Clone)]
struct SecRange {
    start: i32,
    end: i32,
}

type Pair = [SecRange; 2];

fn fully_overlap(p: &Pair) -> bool {
    p[0].start >= p[1].start && p[0].end <= p[1].end
        || p[1].start >= p[0].start && p[1].end <= p[0].end
}
fn overlap(p: &Pair) -> bool {
    p[0].start >= p[1].start && p[0].start <= p[1].end
        || p[1].start >= p[0].start && p[1].start <= p[0].end
}

fn parse_input(input: &str) -> Vec<Pair> {
    let lines = input.lines();
    let mut pairs: Vec<Pair> = Vec::new();
    for line in lines {
        let ranges = line.split(",");
        let mut pair = [SecRange::default(); 2];
        for (j, r) in ranges.enumerate() {
            let ints: Vec<i32> = r.split("-").map(|e| e.parse::<i32>().unwrap()).collect();
            pair[j] = SecRange {
                start: ints[0],
                end: ints[1],
            };
        }
        pairs.push(pair);
    }
    return pairs;
}

pub fn solve(input: &str, _test: bool, task: Task) -> (String, String) {
    let pairs = parse_input(input);
    let (mut res1, mut res2) = ("".to_string(), "".to_string());

    if !matches!(task, Task::Two) {
        let count = &pairs.iter().filter(|p| fully_overlap(p)).count();
        res1 = count.to_string();
    }
    if !matches!(task, Task::One) {
        let mut count = 0;
        for p in pairs {
            if overlap(&p) {
                count += 1;
            }
        }
        res2 = count.to_string();
    }
    return (res1, res2);
}

use crate::types::Task;

fn find_duplicate(r: &[u8]) -> u8 {
    let mut seen: u64 = 0;
    let len = r.len();
    let (comp1, comp2) = r.split_at(r.len() / 2);
    for i in 0..len / 2 {
        // set bit for character in rucksack 1
        seen |= 1u64 << (comp1[i] - 65);
    }
    for i in 0..len / 2 {
        // check if character in rucksack 2 is already set
        if (seen & (1u64 << (comp2[i] - 65))) > 0 {
            return comp2[i];
        }
    }
    println!("No duplicate found in data:\n  {comp1:?}\n  {comp2:?}");
    return 0;
}

fn get_priority(duplicate: u8) -> i32 {
    if duplicate >= b'a' && duplicate <= b'z' {
        return duplicate as i32 - b'a' as i32 + 1;
    } else {
        return duplicate as i32 - b'A' as i32 + 27;
    }
}

fn find_badges(rucksacks: Vec<&str>) -> Vec<u8> {
    let mut badges = vec![0; rucksacks.len() / 3];
    // loop over groups
    for i in (0..rucksacks.len()).step_by(3) {
        let mut seen = vec![0; 3];
        // create set of seen items for each rucksack in one group
        for j in 0..3 {
            for c in rucksacks[i + j].as_bytes() {
                seen[j] |= 1u64 << (c - 65)
            }
        }
        // search bit (character) set in all 3 rucksacks
        for r in 0..64 {
            if ((1u64 << r) & seen[0] & seen[1] & seen[2]) > 0 {
                badges[i / 3] = r + 65;
                break;
            }
        }
    }
    return badges;
}

pub fn solve(input: &str, test: bool, task: Task) -> (String, String) {
    let rucksacks: Vec<&str> = input.lines().collect();
    if test {
        println!("Rucksacks: {rucksacks:?}");
    }
    let (mut res1, mut res2) = ("".into(), "".into());
    if !matches!(task, Task::Two) {
        let sum: i32 = rucksacks
            .iter()
            .map(|r| find_duplicate(r.as_bytes()))
            .map(|d| get_priority(d))
            .sum();
        println!("Result 1: {sum}");
        res1 = sum.to_string();
    }
    if !matches!(task, Task::One) {
        let badges = find_badges(rucksacks);
        let badge_sum: i32 = badges.iter().map(|b| get_priority(*b)).sum();
        if test {
            for b in badges {
                println!("Badge: {}", String::from_utf8(vec![b]).unwrap());
            }
        }
        println!("\nResult 2: {badge_sum}");
        res2 = badge_sum.to_string();
    }
    return (res1, res2);
}

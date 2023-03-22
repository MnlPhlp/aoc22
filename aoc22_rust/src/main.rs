use rayon::prelude::*;
use std::{
    time::{Duration, Instant},
    vec,
};

use clap::Parser;
use types::Task;

mod day01;
mod day02;
mod day03;
mod types;
mod util;

type SolveFunc = dyn Fn(&str, bool, Task) -> (String, String);

const FUNCS: [&SolveFunc; 3] = [&day01::solve, &day02::solve, &day03::solve];

fn cap_length(msg: &str, length: usize) -> &str {
    if msg.len() <= length {
        return msg;
    }
    &msg[0..length]
}

fn calc_day(
    day: usize,
    result1: &mut String,
    result2: &mut String,
    time: &mut Duration,
    test: bool,
    task: Task,
) {
    println!("\n##################\ncalculating day {day} \n##################\n");
    let start = Instant::now();
    let input = util::read_input(day, test);
    let (res1, res2) = FUNCS[day - 1](&input, test, task);
    *time = Instant::now().duration_since(start);
    *result1 = res1;
    *result2 = res2;
}

#[derive(Parser)]
struct Args {
    #[clap(short, long, default_value = "false")]
    test: bool,
    #[clap(short, long, default_value = "0")]
    day: usize,
    #[clap(long, default_value = "0")]
    task: i32,
    #[clap(short, long, default_value = "false")]
    parallel: bool,
}

fn main() {
    // parse command line arguments
    let args: Args = Args::parse();
    let test = args.test;
    let task = if args.task == 1 {
        Task::One
    } else if args.task == 2 {
        Task::Two
    } else {
        Task::Both
    };
    let day = args.day;
    let parallel = args.parallel;

    let days = if day == 0 {
        (1..=FUNCS.len()).collect::<Vec<usize>>()
    } else {
        vec![day]
    };

    println!("Calculating days: {days:?}");

    let mut results1 = vec![String::new(); days.len()];
    let mut results2 = vec![String::new(); days.len()];
    let mut times = vec![Duration::new(0, 0); days.len()];

    let start = Instant::now();
    if parallel {
        run_parallel(&days, &mut results1, &mut results2, &mut times, test, task);
    } else {
        run_serial(&days, &mut results1, &mut results2, &mut times, test, task);
    }
    let overall = Instant::now().duration_since(start);

    let mut results: String = "## Results:\n".into();
    results += "day | result 1        | result 2        | time (ms) | % of overall time\n";
    results += "--: | :-------------: | :--------------:| --------: | :--------\n";
    for (i, day) in days.iter().enumerate() {
        results += &format!(
            "{:0>3} | {: <15} | {: <15} | {: >9?} | {: >4.2} %\n",
            day,
            cap_length(&results1[i], 15),
            cap_length(&results2[i], 15),
            times[i],
            (times[i].as_micros() as f32 / overall.as_micros() as f32) * 100f32
        );
    }
    results += &format!("\nOverall Time: {overall:?}\n");
    results += &format!(
        "\nSummed Time: {:?}\n",
        times.iter().fold(Duration::new(0, 0), |sum, x| sum + *x)
    );

    println!("{}", results);
}

fn run_serial(
    days: &[usize],
    results1: &mut Vec<String>,
    results2: &mut Vec<String>,
    times: &mut Vec<Duration>,
    test: bool,
    task: Task,
) {
    for (i, day) in days.iter().enumerate() {
        calc_day(
            *day,
            &mut results1[i],
            &mut results2[i],
            &mut times[i],
            test,
            task,
        );
    }
}

fn run_parallel(
    days: &[usize],
    results1: &mut Vec<String>,
    results2: &mut Vec<String>,
    times: &mut Vec<Duration>,
    test: bool,
    task: Task,
) {
    days.par_iter()
        .zip(results1)
        .zip(results2)
        .zip(times)
        .for_each(|(((day, result1), result2), time)| {
            calc_day(*day, result1, result2, time, test, task);
        });
}

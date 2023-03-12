use std::{vec, time::{Duration, Instant}};

use clap::Parser;
use types::Task;

mod util;
mod types;
mod day01;

type SolveFunc = dyn Fn(&str, bool, Task) -> (String,String);

const FUNCS: [&SolveFunc;1] = [
    &day01::solve
];

fn cap_length(msg: &str, length: usize) -> &str {
    if msg.len() <= length {
        return msg;
    }
	&msg[0..length]
}

fn calc_day(day: usize, i: usize, results1: & mut [String], results2:& mut [String], times: & mut [Duration], test: bool, task: Task) {
	println!("\n##################\ncalculating day {day} \n##################\n");
	let start = Instant::now();
	let input = util::read_input(day, test);
	let (res1, res2) = FUNCS[day-1](&input, test, task);
	times[i] = Instant::now().duration_since(start);
	results1[i] = res1;
	results2[i] = res2;
}

#[derive(Parser)]
struct Args{
    #[clap(short, long, default_value = "false")]
    test: bool,
    #[clap(short, long, default_value = "0")]
    day: usize,
    #[clap(long, default_value = "0")]
    task: i32,
}

fn main() {
    // parse command line arguments
    let args: Args = Args::parse();
    let test = args.test;
    let task = if args.task==1 {Task::One} else if args.task==2 {Task::Two} else {Task::Both};
    let day = args.day;

    let days = if day == 0 {
        (1..=FUNCS.len()).collect::<Vec<usize>>()
    } else {
        vec![day]
    };

    println!("Calculating days: {days:?}");

    let mut results1 = vec![String::new(); days.len()];
    let mut results2 = vec![String::new(); days.len()];
    let mut times = vec![Duration::new(0,0); days.len()];

    let start = Instant::now();
    for (i,day) in days.iter().enumerate() {
        calc_day(*day, i, &mut results1, &mut results2, &mut times, test, task);
    }
    let overall = Instant::now().duration_since(start);

    let mut results: String = "## Results:\n".into();
	results += "day | result 1        | result 2        | time (ms) | % of overall time\n";
	results += "--: | :-------------: | :--------------:| --------: | :--------\n";
	for (i, day) in days.iter().enumerate() {
		results += &format!("{:0>3} | {: <15} | {: <15} | {: >9?} | {: >4.2} %\n",
			day,
			cap_length(&results1[i], 15),
			cap_length(&results2[i], 15),
			times[i],
			(times[i].as_micros() as f32 / overall.as_micros() as f32) * 100f32
        );
	}
	results += &format!("\nOverall Time: {overall:?}\n");
	results += &format!("\nSummed Time: {:?}\n", times.iter().fold(Duration::new(0,0), |sum, x| sum + *x));

    println!("{}", results);
}

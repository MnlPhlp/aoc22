use std::{fs::File, io::Read};

pub fn read_input(day: usize, test: bool) -> String {
    let path = format!("../day{day:0>2}/{}input.txt", if test { "test" } else { "" });
    let mut file = File::open(path).expect("File not found");
    let mut contents = String::new();
    file.read_to_string(&mut contents).expect("Something went wrong reading the file");
    contents
}
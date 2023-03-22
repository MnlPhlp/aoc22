use std::{fs::File, io::Read};

pub fn read_input(day: usize, test: bool) -> String {
    let path = format!(
        "../day{day:0>2}/{}nput.txt",
        if test { "testI" } else { "i" }
    );
    let mut file = File::open(path).expect("File not found");
    let mut contents = String::new();
    file.read_to_string(&mut contents)
        .expect("Something went wrong reading the file");
    contents
}

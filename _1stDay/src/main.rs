use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path
};

fn depth_from_input(filename: impl AsRef<Path>) -> Vec<i32> {
    let file = File::open(filename).expect("File doesn't exist");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}

fn main() {
    let lines = depth_from_input("input.txt");
    for line in lines {
        println!("{:?}", line);
    }
}

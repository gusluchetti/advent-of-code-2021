use std::{io, io::prelude::*};
use std::fs::File;

// parse puzzle input in a couple ways:
// default - no separator (whole input as ONE string)
// new line separator (array of lines)
// other separator (any other character or word)

pub fn parse_input(path: String, separator: &str) -> io::Result<()> {
    let mut file = File::open(path)?;
    let mut lines: Vec<String> = Vec::new();
    // let new_line: u8 = '\n' as u8;

    match separator {
        "none" => {
            let mut input = String::new();
            file.read_to_string(&mut input)?;
            println!("{:?}", input);
        },
        "new_line" => {
            let mut input = Vec::new();
            file.read_to_end(&mut input)?;
            println!("{:?}", input);
            for (i, b) in input.iter().enumerate() {
                let mut last_cr: usize = 0;
                let cr: u8 = '\r' as u8;
                if &cr == b {
                    lines.push(input[last_cr..i-1])
                    println!("last CR at index {last_cr}")
                    last_cr = i;
                }
            }
        },
        _ => {
            let mut input = Vec::new();
            file.read_to_end(&mut input)?;
            println!("{:?}", input);
        }
    };

    return Ok(());
}

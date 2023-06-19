use std::{io, io::prelude::*};
use std::fs::File;

// parse puzzle input in a couple ways:
// default - no separator (whole input as ONE string)
// new line separator (array of lines)
// other separator (any other character or word)

pub fn parse_input(path: String, separator: &str) -> io::Result<()> {
    let mut file = File::open(path)?;

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
                let byte: u8 = '\n' as u8;
                if &byte == b {
                    println!("{b} is line end!");
                    println!("{} is last before?", input[i-1]);
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

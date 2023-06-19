use std::{io, io::prelude::*};
use std::fs::File;
use std::str;

// parse puzzle input in a couple ways:
// default - no separator (whole input as ONE string)
// new line separator (array of lines)
// other separator (any other character or word)

enum Separator {
    None,
    NewLine,
    Other
}

pub fn parse_input(path: String, separator: &str) -> io::Result<()> {
    let line_end: u8 = '\n' as u8;
    let mut file = File::open(path)?;
    let mut lines: Vec<&str> = Vec::new();

    let mut input = Vec::new();
    file.read_to_end(&mut input)?;

    let lines = match separator {
        "none" => {
            let single = str::from_utf8(&input).unwrap();
            lines.push(single);
            lines
        },
        "new_line" => {
            let mut last_cr: usize = 0;
            for (i, byte) in input.iter().enumerate() {
                if byte == &line_end {
                    let slice: &[u8] = &input[last_cr+1..i-1];
                    let line: &str = str::from_utf8(slice).unwrap();
                    lines.push(line);
                    last_cr = i;
                }
            }
            lines
        },
        _ => {
            file.read_to_end(&mut input)?;
            lines
        }
    };

    println!("{:?}", lines);
    return Ok(());
}

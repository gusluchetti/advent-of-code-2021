use std::{io, io::prelude::*};
use std::fs::File;

// parse puzzle input in a couple ways:
// default - no separator (whole input as ONE string)
// new line separator (array of lines)
// other separator (any other character or word)


fn parse_input(path: String, separator: &str) -> io::Result<()> {
    let mut file = File::open(path)?;
    let mut input: String = String::from("test");

    match separator {
        "none" => {
            let mut input = String::new();
            file.read_to_string(&mut input)?;
        },
        "new_line" => {
            file.read_to_end(&mut input)?;
        },
        _ => {
            file.read_to_end(&mut input)?;
        }
    };

    println!("{:?}", input);

    return Ok(input);
}


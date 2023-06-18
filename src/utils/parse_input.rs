use std::{io, io::prelude::*};
use std::vec;
use std::fs::File;

// can get input in a couple ways:
// default - no separator (whole input as string)
// new line separator (array of lines)
// other separator (any other character or word)
fn parse_input(path: String) -> Result<Vec<String>, io::Error> {
    let mut file = File::open(path)?;
    let mut input = String::new();
    file.read_to_string(&mut input)?;

    return Ok(input);
}


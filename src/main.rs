mod utils;

use utils::parse_input;

fn main() {
    println!("run parse input with new line param");
    let path = String::from("input.txt");
    parse_input::parse_input(path, "none");
}

mod utils;

use utils::parse_input;

fn main() {
    println!("start test");

    let path = String::from("input.txt");
    parse_input::parse_input(path, "new_line");
}

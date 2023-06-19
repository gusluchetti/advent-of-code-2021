use utils::parse_input::Method;

fn main() -> io::Result<()> {
    let mut file = File::open("input.txt")?;
    let mut input = String::new();
    file.read_to_string(&mut input)?;

    let mut counter = 0;
    for (i, inst) in input.chars().enumerate() {
        if inst == '(' { counter = counter + 1; }
        else if inst == ')' { 
            counter = counter - 1;
        }

        if counter == -1 {
            println!("got in basement at {} instruction", i+1); 
            break;
        }
    }

    println!("final floor: {}", counter);
    Ok(())
}

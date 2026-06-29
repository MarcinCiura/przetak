use przetak::{evaluate, version};
use std::io::{self, Write};

fn main() {
    let v = version();
    print!("Przetak {}> ", v);
    io::stdout().flush().unwrap();

    let mut input = String::new();
    io::stdin().read_line(&mut input).unwrap();
    let result = evaluate(&input);
    println!("{}", result);
}

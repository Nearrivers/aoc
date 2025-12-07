use std::{fs, process};

mod cli;
mod day1;
mod day2;

// Cli usage cargo run -- <day> <part> <flow>
// Where flow is whether the real input or the example is used
fn main() {
    let (day, part, flow) = cli::cli::parse_args();

    let file_path = format!("./src/day{day}/{flow}.txt");
    let contents = fs::read_to_string(&file_path).unwrap_or_else(|err| {
        println!("error: couldn't open file {file_path} : {err}");
        process::exit(1);
    });

    match day {
        1 => {
            match part {
                1 => day1::day1::part1(contents),
                2 => day1::day1::part2(contents),
                _ => {
                    println!("error: only 2 parts.");
                    process::exit(1);
                }
            };
        }
        2 => {
            match part {
                1 => day2::day2::part1(contents),
                2 => day2::day2::part2(contents),
                _ => {
                    println!("error: only 2 parts.");
                    process::exit(1);
                }
            };
        }
        _ => {
            println!("error: day not done yet.");
            process::exit(1);
        }
    }
}

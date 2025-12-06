use std::{env, fmt, process};

pub enum Flow {
    INPUT,
    EXAMPLE,
}

impl fmt::Display for Flow {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Flow::INPUT => write!(f, "input"),
            Flow::EXAMPLE => write!(f, "example"),
        }
    }
}

pub fn print_usage() {
    println!("usage: zig run main.zig -- <day> <part> <flow>\n");
}

pub fn parse_args() -> (u8, u8, Flow) {
    let args: Vec<String> = env::args().collect();

    if args.len() != 4 {
        println!("error: wrong args count");
        print_usage();
        process::exit(1);
    }

    let day_string = &args[1];
    let day = day_string.parse::<u8>().unwrap_or_else(|err| {
        println!("error: day is not a valid number: {err}");
        print_usage();
        process::exit(1);
    });

    let part_string = &args[2];
    let part = part_string.parse::<u8>().unwrap_or_else(|err| {
        println!("error: part is not a valid number: {err}");
        print_usage();
        process::exit(1);
    });

    let flow_arg = &args[3];

    if flow_arg == "input" {
        return (day, part, Flow::INPUT);
    }

    if flow_arg == "example" {
        return (day, part, Flow::EXAMPLE);
    }

    println!("error: wrong flow value. Expected 'example' or 'input'");
    print_usage();
    process::exit(1);
}

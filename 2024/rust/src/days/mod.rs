mod d01;
use crate::utils::inputs::read_input;

type DayPartFunction = fn(Vec<String>) -> Result<String, String>;
type DayEntry = (i32, Option<DayPartFunction>, Option<DayPartFunction>);

pub fn run_days() {
    let days: Vec<DayEntry> = vec![(1, Some(d01::p1), Some(d01::p2))];

    for (d, p1, p2) in days {
        let inputs = match read_input(d) {
            Ok(s) => s,
            Err(_) => {
                println!("failed to read input for day {d}");
                continue;
            }
        };

        println!("Running day {d}");
        println!("part 01: {}", run_part(inputs.clone(), p1));
        println!("part 02: {}", run_part(inputs.clone(), p2));

        println!("\n\n\n");
    }
}

fn run_part(inputs: Vec<String>, part_fn: Option<DayPartFunction>) -> String {
    match part_fn {
        None => "no function provided...".to_string(),
        Some(f) => match f(inputs) {
            Ok(out) => out,
            Err(e) => format!("Error {e:?}").to_string(),
        },
    }
}

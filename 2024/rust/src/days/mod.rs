mod d01;
mod d02;
mod d03;
mod d04;
mod d05;
use crate::utils::inputs::read_input;

use std::time::Instant;

type DayPartFunction = fn(Vec<String>) -> Result<String, String>;
type DayEntry = (i32, Option<DayPartFunction>, Option<DayPartFunction>);

pub fn run_days() {
    let days: Vec<DayEntry> = vec![
        (1, Some(d01::p1), Some(d01::p2)),
        (2, Some(d02::p1), Some(d02::p2)),
        (3, Some(d03::p1), Some(d03::p2)),
        (4, Some(d04::p1), Some(d04::p2)),
        (5, Some(d05::p1), Some(d05::p2)),
        (6, None, None),
        (7, None, None),
        (8, None, None),
        (9, None, None),
        (10, None, None),
        (11, None, None),
        (12, None, None),
        (13, None, None),
        (14, None, None),
        (15, None, None),
        (16, None, None),
        (17, None, None),
        (18, None, None),
        (19, None, None),
        (20, None, None),
        (21, None, None),
        (22, None, None),
        (23, None, None),
        (24, None, None),
        (25, None, None),
    ];

    for (d, p1, p2) in days {
        let inputs = match read_input(d) {
            Ok(s) => s,
            Err(_) => {
                println!("failed to read input for day {d}");
                continue;
            }
        };

        println!("Running day {d}");
        let now = Instant::now();
        let rp1 = run_part(inputs.clone(), p1);
        let elapsed = now.elapsed();
        println!("Part 01 ({}μs): {}", elapsed.as_micros(), rp1);

        let now = Instant::now();
        let rp2 = run_part(inputs.clone(), p2);
        let elapsed = now.elapsed();
        println!("Part 02 ({}μs): {}", elapsed.as_micros(), rp2);

        println!("");
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

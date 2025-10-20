use std::fs::read_to_string;

pub fn read_input(day: i32) -> Result<Vec<String>, std::io::Error> {
    let day_str = day_to_str(day);
    let filename = format!("inputs/{day_str}.txt");

    Ok(read_to_string(filename)?
        .lines()
        .map(String::from)
        .collect())
}

fn day_to_str(day: i32) -> String {
    if day <= 9 {
        return format!("0{day}");
    } else {
        return format!("{day}");
    }
}

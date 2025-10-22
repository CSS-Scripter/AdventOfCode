use std::vec;

#[derive(Debug, Clone)]
enum Operation {
    Mul(i32, i32),
    Do,
    Dont,
}

pub fn p1(inputs: Vec<String>) -> Result<String, String> {
    let mut sum: i32 = 0;
    for line in inputs {
        let ops = get_commands(line);
        for op in ops {
            match op {
                Operation::Mul(x, y) => sum += x * y,
                _ => (),
            }
        }
    }

    Ok(format!("{sum}").to_string())
}

pub fn p2(inputs: Vec<String>) -> Result<String, String> {
    let mut sum: i32 = 0;
    let mut enabled: bool = true;
    for line in inputs {
        let ops = get_commands(line);
        for op in ops {
            match (op, &enabled) {
                (Operation::Mul(x, y), true) => sum += x * y,
                (Operation::Do, false) => enabled = true,
                (Operation::Dont, true) => enabled = false,
                _ => (),
            }
        }
    }
    Ok(format!("{sum}").to_string())
}

fn get_commands(line: String) -> Vec<Operation> {
    let mul_indices = str_find_all(line.clone(), "mul(");
    let do_indices = str_find_all(line.clone(), "do()");
    let dont_indices = str_find_all(line.clone(), "don't()");

    let mut ops: Vec<(usize, Operation)> = vec![];

    for i in mul_indices {
        if let Some(operator) = validate_mul_indice(&line, i) {
            ops.push((i, operator));
        }
    }

    for i in do_indices {
        ops.push((i, Operation::Do));
    }

    for i in dont_indices {
        ops.push((i, Operation::Dont));
    }

    ops.sort_by(|(a, _), (b, _)| a.cmp(b));
    return ops.into_iter().map(|(_, op)| op).collect();
}

fn str_find_all(mut line: String, predicate: &str) -> Vec<usize> {
    let mut indices: Vec<usize> = vec![];
    let mut offset: usize = 0;
    while let Some(i) = line.find(predicate) {
        indices.push(i + offset);
        offset += i + 1;
        line = line[i + 1..].to_string();
    }

    return indices;
}

fn validate_mul_indice(line: &String, indice: usize) -> Option<Operation> {
    let start_indice = &line[indice..].find('(');
    let end_indice = &line[indice..].find(')');
    if start_indice.is_none() || end_indice.is_none() {
        return None;
    }

    let start_indice = start_indice.unwrap() + indice + 1;
    let end_incide = end_indice.unwrap() + indice;

    let mul_section = &line[start_indice..end_incide];
    let numbers: Vec<&str> = mul_section.split(",").collect();
    if numbers.len() != 2 {
        return None;
    }

    let x = numbers[0].parse::<i32>();
    let y = numbers[1].parse::<i32>();

    if x.is_err() || y.is_err() {
        return None;
    }

    Some(Operation::Mul(x.unwrap(), y.unwrap()))
}

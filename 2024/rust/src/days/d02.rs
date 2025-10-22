use std::num::ParseIntError;

pub fn p1(inputs: Vec<String>) -> Result<String, String> {
    let levels = parse_input(inputs).unwrap();
    let mut safe_count = 0;

    for level in levels {
        let safe = is_level_save(level);
        if safe {
            safe_count += 1;
        }
    }

    Ok(format!("{safe_count}").to_string())
}

pub fn p2(inputs: Vec<String>) -> Result<String, String> {
    let levels = parse_input(inputs).unwrap();
    let mut safe_count = 0;

    for level in levels {
        let mut safe = is_level_save(level.clone());
        if !safe {
            for i in 0..level.len() {
                let mut damped = level.clone();
                damped.remove(i);
                safe = is_level_save(damped);
                if safe {
                    break;
                }
            }
        }

        if safe {
            safe_count += 1;
        }
    }

    Ok(format!("{safe_count}").to_string())
}

fn parse_input(inputs: Vec<String>) -> Result<Vec<Vec<i8>>, ParseIntError> {
    let mut outputs: Vec<Vec<i8>> = vec![];
    for line in inputs {
        let nums: Vec<i8> = line.split(' ').map(|c| c.parse::<i8>().unwrap()).collect();
        outputs.push(nums);
    }
    Ok(outputs)
}

fn is_level_save(level: Vec<i8>) -> bool {
    let steps = level_to_steps(level);
    for i in 0..steps.len() {
        let c = steps[i];
        if c > 3 || c < (-3) || c == 0 {
            return false;
        }

        if i != 0 {
            let sc = signum(c);
            let sp = signum(steps[i - 1]);
            if sp != sc {
                return false;
            }
        }
    }

    return true;
}

fn level_to_steps(level: Vec<i8>) -> Vec<i8> {
    let mut steps: Vec<i8> = vec![];
    for i in 1..level.len() {
        let c = level[i];
        let p = level[i - 1];
        steps.push(c - p);
    }
    return steps;
}

fn signum(a: i8) -> i8 {
    if a < 0 {
        return -1;
    }

    if a > 0 {
        return 1;
    }

    return 0;
}

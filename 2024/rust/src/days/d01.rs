use std::collections::HashMap;

pub fn p1(inputs: Vec<String>) -> Result<String, String> {
    let mut l1: Vec<i32> = Vec::new();
    let mut l2: Vec<i32> = Vec::new();

    for line in inputs {
        let nums: Vec<i32> = line
            .split("   ")
            .map(|s| s.parse::<i32>().unwrap())
            .collect();

        l1.push(nums[0]);
        l2.push(nums[1]);
    }

    l1.sort();
    l2.sort();

    let mut total = 0;
    for i in 0..l1.len() {
        total += (l1[i] - l2[i]).abs();
    }

    Ok(format!("{total}").to_string())
}

pub fn p2(inputs: Vec<String>) -> Result<String, String> {
    let mut l1: Vec<i32> = Vec::new();
    let mut counts: HashMap<i32, i32> = HashMap::new();

    for line in inputs {
        let nums: Vec<i32> = line
            .split("   ")
            .map(|s| s.parse::<i32>().unwrap())
            .collect();

        l1.push(nums[0]);
        counts.entry(nums[1]).and_modify(|c| *c += 1).or_insert(1);
    }

    let mut total = 0;
    for n in l1 {
        counts.entry(n).and_modify(|e| total += n * *e);
    }

    Ok(format!("{total}").to_string())
}

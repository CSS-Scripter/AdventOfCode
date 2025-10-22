use std::{collections::HashMap, hash::Hash, num::ParseIntError};

type Updates = Vec<i32>;
type PageOrder = (i32, i32);

pub fn p1(inputs: Vec<String>) -> Result<String, String> {
    let (orders, updates) = parse_input(inputs)?;

    let mut total: i32 = 0;
    for update in updates {
        let mut is_ordered = true;
        for i in 0..update.len() - 1 {
            let order = &orders.get_key_value(&update[i]);
            if order.is_none() {
                break;
            }

            let (_, order) = order.unwrap();
            for j in i + 1..update.len() {
                if !order.should_be_before_page(update[j]) {
                    is_ordered = false;
                    break;
                }
            }
            if is_ordered == false {
                break;
            }
        }
        if is_ordered {
            let middle_i = (update.len() / 2) + 1;
            total += update[middle_i];
        }
    }

    Ok(format!("{}", total).to_string())
}

pub fn p2(inputs: Vec<String>) -> Result<String, String> {
    Err("not implemented".to_string())
}

fn parse_input(
    inputs: Vec<String>,
) -> Result<(HashMap<i32, FormattedPageOrder>, Vec<Updates>), String> {
    let mut input_iterator = inputs.iter();

    let mut page_orders: Vec<PageOrder> = vec![];

    while let Some(l) = input_iterator.next() {
        if l == "" {
            break;
        }
        let num_strs: Vec<Result<i32, ParseIntError>> =
            l.split("|").map(|s| s.parse::<i32>()).collect();
        if num_strs.len() != 2 {
            return Err(format!("line '{l}' does not follow structure '\\d|\\d'"));
        }

        if num_strs[0].is_err() || num_strs[1].is_err() {
            return Err(format!("numbers in '{l}' did not parse correctly"));
        }

        let n1 = num_strs[0].to_owned().unwrap();
        let n2 = num_strs[1].to_owned().unwrap();

        page_orders.push((n1, n2))
    }

    let mut updates: Vec<Updates> = vec![];
    while let Some(l) = input_iterator.next() {
        if l == "" {
            break;
        }
        let num_strs: Vec<i32> = l.split(",").map(|s| s.parse::<i32>().unwrap()).collect();
        updates.push(num_strs);
    }

    Ok((format_page_orders(page_orders), updates))
}

struct FormattedPageOrder {
    pages_before: Vec<i32>,
    pages_after: Vec<i32>,
}

impl FormattedPageOrder {
    pub fn new() -> FormattedPageOrder {
        return FormattedPageOrder {
            pages_before: vec![],
            pages_after: vec![],
        };
    }

    pub fn add_page_before(&mut self, page: i32) {
        self.pages_before.push(page);
    }

    pub fn add_page_after(&mut self, page: i32) {
        self.pages_after.push(page);
    }

    pub fn should_be_before_page(&self, page: i32) -> bool {
        for p in &self.pages_after {
            if p == &page {
                return true;
            }
        }
        return false;
    }
}

fn format_page_orders(orders: Vec<PageOrder>) -> HashMap<i32, FormattedPageOrder> {
    let mut order_map: HashMap<i32, FormattedPageOrder> = HashMap::new();

    for order in orders {
        order_map
            .entry(order.0)
            .or_insert(FormattedPageOrder::new())
            .add_page_after(order.1);
        order_map
            .entry(order.1)
            .or_insert(FormattedPageOrder::new())
            .add_page_before(order.0);
    }

    return order_map;
}

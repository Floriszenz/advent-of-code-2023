use std::{
    collections::HashMap,
    env,
    fs::File,
    io::{self, BufRead, BufReader},
};

type Network = HashMap<String, (String, String)>;

fn build_network(map_document: impl Iterator<Item = io::Result<String>>) -> Network {
    let mut network: Network = HashMap::new();

    for line in map_document.skip(1) {
        let line = line.unwrap();
        let node_id = line.get(0..3).unwrap();
        let left_node = line.get(7..10).unwrap();
        let right_node = line.get(12..15).unwrap();

        network.insert(
            String::from(node_id),
            (String::from(left_node), String::from(right_node)),
        );
    }

    network
}

fn search_path_from_aaa_to_zzz(network: &Network, instructions: &String) -> i32 {
    let mut current_node = "AAA";
    let mut number_of_steps = 0;

    while current_node != "ZZZ" {
        for instruction in instructions.chars() {
            current_node = match instruction {
                'L' => network.get(current_node).map(|node| &node.0).unwrap(),
                'R' => network.get(current_node).map(|node| &node.1).unwrap(),
                _ => unreachable!("There should not be another instruction than L or R"),
            };

            number_of_steps += 1;
        }
    }

    number_of_steps
}

fn calculate_greatest_common_divisor(mut a: i64, mut b: i64) -> i64 {
    let mut t;

    while b != 0 {
        t = b;
        b = a % b;
        a = t;
    }

    a
}

fn calculate_least_common_multiple(a: i64, b: i64) -> i64 {
    (a * b).abs() / calculate_greatest_common_divisor(a, b)
}

fn search_path_from_xxa_to_xxz(network: &Network, instructions: &String) -> i64 {
    let mut current_nodes = Vec::new();
    let mut numbers_of_steps = Vec::new();

    for node in network.keys() {
        if node.ends_with("A") {
            current_nodes.push(String::from(node));
            numbers_of_steps.push(0);
        }
    }

    loop {
        for instruction in instructions.chars() {
            if current_nodes.iter().all(|node| node.ends_with("Z")) {
                return numbers_of_steps
                    .into_iter()
                    .reduce(|acc, e| calculate_least_common_multiple(acc, e))
                    .unwrap();
            }

            for (idx, current_node) in current_nodes.iter_mut().enumerate() {
                if !current_node.ends_with("Z") {
                    *current_node = String::from(match instruction {
                        'L' => network.get(current_node).map(|node| &node.0).unwrap(),
                        'R' => network.get(current_node).map(|node| &node.1).unwrap(),
                        _ => unreachable!("There should not be another instruction than L or R"),
                    });

                    numbers_of_steps[idx] += 1;
                }
            }
        }
    }
}

fn main() {
    if let Some(file_path) = env::args().nth(1) {
        let map_document = File::open(file_path).expect("Failed to open map_document");
        let map_document = BufReader::new(map_document);
        let mut map_document = map_document.lines();
        let instructions = map_document.next().unwrap().unwrap();

        let network = build_network(map_document);

        let number_of_steps_from_aaa_to_zzz = search_path_from_aaa_to_zzz(&network, &instructions);
        let number_of_steps_from_xxa_to_xxz = search_path_from_xxa_to_xxz(&network, &instructions);

        println!(
            "Number of steps required for path from AAA to ZZZ: {number_of_steps_from_aaa_to_zzz}"
        );
        println!(
            "Number of steps required for path from XXA to XXZ: {number_of_steps_from_xxa_to_xxz}"
        );
    } else {
        panic!("Did not provide path to map document file");
    }
}

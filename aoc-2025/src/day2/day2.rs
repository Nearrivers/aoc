pub fn part1(file_content: String) {
    let mut id_sum: u64 = 0;

    for id_range in file_content.split(',') {
        let ids: Vec<&str> = id_range.split('-').collect();

        let first_id: u64 = ids[0].parse::<u64>().unwrap();
        let second_id: u64 = ids[1].replace("\n", "").parse::<u64>().unwrap();

        for id in first_id..=second_id {
            let id_string: String = id.to_string();

            if id_string.len() % 2 != 0 {
                continue;
            }

            let (first_half, second_half) = id_string.split_at(id_string.len() / 2);

            if first_half == second_half {
                id_sum += id;
            }

            println!("{} {}", first_half, second_half);
        }
    }

    println!("id sum is {id_sum}");
}

pub fn part2(file_content: String) {
    println!("{file_content}");
}

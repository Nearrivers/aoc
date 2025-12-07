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
        }
    }

    println!("id sum is {id_sum}");
}

pub fn part2(file_content: String) {
    let mut id_sum: u64 = 0;

    for id_range in file_content.split(',') {
        let ids: Vec<&str> = id_range.split('-').collect();

        let first_id: u64 = ids[0].parse::<u64>().unwrap();
        let second_id: u64 = ids[1].replace("\n", "").parse::<u64>().unwrap();

        for id in first_id..=second_id {
            let id_string: String = id.to_string();

            if sliding_window(&id_string) {
                id_sum += id;
            }
        }
    }

    println!("id sum is {id_sum}");
}

fn sliding_window(id_string: &str) -> bool {
    // longueur de la fenêtre
    let mut n = 1;
    // décalage de la fenêtre de recherche
    let mut m = 1;

    // Tant que la fenêtre est inférieure ou égale à la moitié de
    // la longueur de la chaîne
    while n <= id_string.len() / 2 {
        // Si le décalage pointe vers un index en dehors de la chaîne
        // alors cela veut dire qu'on a pu faire glisser la fenêtre dans toute
        // la string. Cela veut donc dire que la string est valide
        if m >= id_string.len() {
            return true;
        }

        // On cherche l'aiguille dans une botte de fouin
        let needle: &str = &id_string[..n];

        let searched_part: Option<&str> = id_string.get(m..m + n);

        match searched_part {
            Some(s) => {
                if s != needle {
                    n += 1;
                    m = n;
                    continue;
                }

                m += n;
            }
            None => break,
        }
    }

    false
}

#[cfg(test)]
mod tests {
    use crate::day2::day2::sliding_window;

    #[test]
    fn part_2_test() {
        let cases: [(&str, bool); 10] = [
            ("99", true),
            ("95", false),
            ("999", true),
            ("1010", true),
            ("1188511885", true),
            ("1188511886", false),
            ("1188511886", false),
            ("212121212121", true),
            ("212122212121", false),
            ("2121212118", false),
        ];

        for (id_string, expected_outcome) in &cases {
            let got = sliding_window(id_string);

            assert_eq!(
                got, *expected_outcome,
                "got {got}, want {expected_outcome} for id \"{id_string}\""
            );
        }
    }
}

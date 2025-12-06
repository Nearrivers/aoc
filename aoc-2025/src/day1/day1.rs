use std::process;

pub fn part1(file_content: String) {
    let mut dial_pointed_number: i16 = 50;
    let mut password_count: i16 = 0;

    for instruction in file_content.lines() {
        if instruction.len() == 0 {
            break;
        }

        let (direction, number) = instruction.split_at(1);

        const MINIMUM: i16 = 0;
        const MAXIMUM: i16 = 100;

        let line_number = number.parse::<i16>().unwrap_or_else(|err| {
            println!("error: line number is not castable into number: {err}");
            process::exit(1);
        });

        let remainder: i16 = line_number % MAXIMUM;

        println!("{line_number} {remainder} {dial_pointed_number}");

        if direction == "L" {
            dial_pointed_number -= remainder;
        } else {
            dial_pointed_number += remainder;
        }

        if dial_pointed_number == 100 || dial_pointed_number == 0 {
            dial_pointed_number = 0;
            password_count += 1;
            continue;
        }

        if MINIMUM < dial_pointed_number && dial_pointed_number < MAXIMUM {
            continue;
        }

        if dial_pointed_number < 0 {
            dial_pointed_number = MAXIMUM + dial_pointed_number;
            println!("dialPointerNumber has been changed to {dial_pointed_number}");
            continue;
        }

        dial_pointed_number = MINIMUM + dial_pointed_number % MAXIMUM;
    }

    println!("password is {password_count}");
}

pub fn part2(file_content: String) {
    let mut dial_pointed_number: i16 = 50;
    let mut password_count: i16 = 0;

    for instruction in file_content.lines() {
        if instruction.len() == 0 {
            break;
        }

        let (direction, number) = instruction.split_at(1);

        let mut line_number = number.parse::<i16>().unwrap_or_else(|err| {
            println!("error: line number is not castable into number: {err}");
            process::exit(1);
        });

        println!("{direction}{line_number}  {dial_pointed_number}");

        // Si la rotation va à gauche, le chiffre est négatif
        if direction == "L" {
            line_number *= -1;
        }

        turn_dial(&mut dial_pointed_number, &mut password_count, line_number);
    }

    println!("password is {password_count}");
    println!("dial points to {dial_pointed_number}");
}

fn turn_dial(dial_pointed_number: &mut i16, password_count: &mut i16, line_number: i16) {
    let is_dial_pointing_to_zero = *dial_pointed_number == 0;

    const MINIMUM: i16 = 0;
    const MAXIMUM: i16 = 100;

    // Si la rotation absolue est plus grande que le maximum
    // on compte le nombre de tour que le cadran aurait dû faire pour arriver à destination.
    // On incrémente le compte du pwd et on continue la boucle
    if line_number.abs() > MAXIMUM {
        let added = line_number.abs() / MAXIMUM;
        println!("{password_count} + {added}");
        *password_count += added;
    }

    // On décale le cadran d'un nombre < MAXIMUM
    // D'où le modulo
    *dial_pointed_number += line_number % MAXIMUM;

    // Si le cadre pointe entre les bordures, alors il ne se passe rien.
    if MINIMUM < *dial_pointed_number && *dial_pointed_number < MAXIMUM {
        return;
    }

    // Si le cadran pointe sur 0 ou 100
    // On incrémente le compteur du pwd de 1 et on reboucle
    if dial_pointed_number.abs() == MAXIMUM || *dial_pointed_number == MINIMUM {
        *dial_pointed_number = 0;
        *password_count += 1;
        println!("{password_count}");
        return;
    }

    // Si on est en dessous de 0
    // Le cadran pointe vers MAXIMUM + (-cadran)
    // Sachant que le cadran ne peut déjà plus être > MAXIMUM
    if *dial_pointed_number < 0 {
        *dial_pointed_number = MAXIMUM + *dial_pointed_number;

        // Si le cadran ne pointait pas déjà vers 0, alors on incrémente le compteur
        if !is_dial_pointing_to_zero {
            *password_count += 1;
            println!("{password_count}");
        }

        return;
    }

    // Le cadran a dépassé 100 et revient donc à sa position de départ
    // Ici, on est forcément en train d'aller vers la droite
    *dial_pointed_number = MINIMUM + *dial_pointed_number % MAXIMUM;
    *password_count += 1;
    println!("{password_count}");
}

#[cfg(test)]
mod tests {
    use crate::day1::day1::turn_dial;

    #[test]
    fn part_2_fking_works() {
        let cases: [(&str, i16, i16); 16] = [
            ("R1000", 10, 50),
            ("R50,L150,", 2, 50),
            ("R50,", 1, 0),
            ("L50,", 1, 0),
            ("R50,R100,R150,", 3, 50),
            ("R50,L107,R8,L2,", 4, 99),
            ("L250,", 3, 0),
            ("L251,", 3, 99),
            ("R75,L25,R10,", 2, 10),
            ("R30,L130,", 1, 50),
            ("R391,L283,R769,L301", 18, 26),
            ("R78,R165,R900", 11, 93),
            ("L51", 1, 99),
            ("L137,L672,L147", 10, 94),
            ("L49,L1,R1", 1, 1),
            ("R49,R1,L1", 1, 99),
        ];

        for (input, expected_pwd_count, expected_dial) in &cases {
            let mut dial_pointed_number: i16 = 50;
            let mut password_count: i16 = 0;
            for instruction in input.split(',') {
                if instruction.len() == 0 {
                    break;
                }

                let (direction, number) = instruction.split_at(1);

                let mut line_number = number.parse::<i16>().unwrap();

                if direction == "L" {
                    line_number *= -1;
                }

                turn_dial(&mut dial_pointed_number, &mut password_count, line_number);
            }

            assert_eq!(
                dial_pointed_number, *expected_dial,
                "got {dial_pointed_number}, want {expected_dial} for input : {input}"
            );
            assert_eq!(
                password_count, *expected_pwd_count,
                "got {password_count}, want {expected_pwd_count} for input : {input}"
            );
        }
    }
}

pub fn is_armstrong_number(num: u32) -> bool {
    let num_of_digigts = get_number_of_digits(num);
    let mut sum: u64 = 0;

    let mut ii: u64 = num as u64;
    while ii != 0 {
        let number: u64 = ii % 10;
        let mut nn: u64 = 1;
        for _ in 0..num_of_digigts {
            nn *= number;
        }
        sum += nn;
        ii = ii/10;
    }
    sum == num as u64
}

pub fn get_number_of_digits(n: u32) -> u32 {
    let mut sum = 0;
    let mut nn = n;
    while nn != 0 {
        nn = nn /10;
        sum += 1;
    }
    return sum
}
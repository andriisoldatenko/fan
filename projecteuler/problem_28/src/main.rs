fn main() {
    let mut sum = 1; 
    let mut start = 1;
    for n in 1..501 {
        let d1 = start + 2*n;
        let d2 = d1 + 2*n;
        let d3 = d2 + 2*n;
        let d4 = d3 + 2*n;
        start = d4;
        sum = sum + d1 + d2 + d3 + d4;
    }
    println!("{}", sum);
}

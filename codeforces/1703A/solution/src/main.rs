use std::io;
use std::str::FromStr;

#[allow(dead_code)]
fn read_line() -> String {
    let mut buffer = String::new();
    io::stdin()
        .read_line(&mut buffer)
        .expect("failed to read line");
 
    buffer
}

#[allow(dead_code)]
fn read<T : FromStr>() -> Result<T, T::Err>{
    read_line().trim().parse::<T>()
}

#[allow(dead_code)]
fn read_vec<T : FromStr>() -> Result< Vec<T>, T::Err>{
    read_line().split_whitespace().map(|x| x.parse::<T>()).collect()
}

fn solve(){
    let l = read_line().trim().to_lowercase();
     match l.as_ref() {
            "yes" => println!("YES"),
            "10" => (),
            _ => println!("NO"),
        }
}

fn main() {
    let t = read::<i32>().unwrap();
    for _ in 0..t {
        solve();
    }
}

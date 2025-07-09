pub mod merge_sort;
pub mod parallel_merge_sort;

use crate::merge_sort::merge_sort;
use crate::parallel_merge_sort::parallel_merge_sort;

use rand::Rng;
use std::time::Instant;
use std::sync::{Arc, Mutex};

fn main() {
    // change the size of the array (also change the threshold value according to size for better results).
    let size: usize = 10000;

    // this generates the array of required size with random numbers.
    let mut arr: Vec<i32> = (0..size)
        .map(|_| rand::thread_rng().gen_range(1..size as i32))
        .collect();

    // Uncomment the below lines to print the array before sorting (only do this for small size arrays).
    // print!("Array before sorting: ");
    // println!("{:?}", arr);

    // calculating the time for normal merge sort.
    let start_time = Instant::now();
    merge_sort(&mut arr);
    let normalmergesorttime = start_time.elapsed();

    // Uncomment the below lines to print the array after sorting (only do this for small size arrays).
    // print!("Array after sorting with normal merge sort: ");
    // println!("{:?}", arr);

    // calculating the time for parallel merge sort.
    let arr_arc = Arc::new(Mutex::new(arr));
    let start_time = Instant::now();
    parallel_merge_sort(Arc::clone(&arr_arc));
    let parallelmergesorttime = start_time.elapsed(); 

    // Uncomment the below lines to print the array after sorting (only do this for small size arrays).
    // print!("Array after sorting with parallel merge sort: ");
    // println!("{:?}", arr_arc.lock().unwrap());

    println!("Normal Merge Sort completed in {:.2?}", normalmergesorttime);
    println!("Parallel Merge Sort completed in {:.2?}", parallelmergesorttime);
}

// So this is my first time writing any multi threaded code in Rust. So there might be better and faster ways to implement this.
// if you know any better way to implement this or any errors in code, please let me know.
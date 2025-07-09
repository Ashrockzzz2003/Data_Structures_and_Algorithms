use crate::merge_sort::merge_sort;
use std::sync::{Arc, Mutex};
use std::thread;

// adjust the threshold value according to the size of the array.
// if the size of the array is less than the threshold value, then normal merge sort will be used.
const THRESHOLD: usize = 4096;

// merge sort helper function
pub fn merge(arr: &mut [i32], left: &[i32], right: &[i32]) {
    let mut i = 0;
    let mut j = 0;
    let mut k = 0;
    while i < left.len() && j < right.len() {
        if left[i] <= right[j] {
            arr[k] = left[i];
            i += 1;
        }else{
            arr[k] = right[j];
            j += 1;
        }
        k += 1;
    }

    while i < left.len() {
        arr[k] = left[i];
        k += 1;
        i += 1;
    }

    while j < right.len() {
        arr[k] = right[j];
        k += 1;
        j += 1;
    }
}

// parallel merge sort function
pub fn parallel_merge_sort(arr: Arc<Mutex<Vec<i32>>>){
    let mut arr = arr.lock().unwrap();

    if arr.len() <= 1 {
        return;
    }

    // if the size of the array is less than the threshold value, we use normal merge sort will be used.
    if arr.len() <= THRESHOLD {
        merge_sort(&mut arr);
        return;
    }

    // split the array into two halves.
    let mid = arr.len()/2;
    let left_arr = arr[..mid].to_vec();
    let right_arr = arr[mid..].to_vec();

    let left_arc = Arc::new(Mutex::new(left_arr));
    let right_arc = Arc::new(Mutex::new(right_arr));

    // makes a threads to sort the left half.
    let left_handle = thread::spawn({
        let left_arc = Arc::clone(&left_arc);
        move || {
            parallel_merge_sort(left_arc);
        }
    });

    // makes a threads to sort the right half.
    let right_handle = thread::spawn({
        let right_arc = Arc::clone(&right_arc);
        move || {
            parallel_merge_sort(right_arc);
        }
    });

    // waiting for threads to finish.
    left_handle.join().unwrap();
    right_handle.join().unwrap();

    // merge the two halves.
    merge(&mut arr, &left_arc.lock().unwrap(), &right_arc.lock().unwrap());

}




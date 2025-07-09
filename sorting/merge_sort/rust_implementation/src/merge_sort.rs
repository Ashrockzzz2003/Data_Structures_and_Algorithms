// merge sort helper function
pub fn merge(arr: &mut [i32], left: &[i32], right: &[i32]){
    let mut i = 0;
    let mut j = 0;
    let mut k = 0;

    // merge the two arrays.
    while i < left.len() && j < right.len() {
        if left[i] <= right[j] {
            arr[k] = left[i];
            i += 1;
        } else {
            arr[k] = right[j];
            j += 1;
        }
        k += 1;
    }

    // if right array is empty, then merge the left array.
    while i < left.len() {
        arr[k] = left[i];
        i += 1;
        k += 1;
    }

    // if left array is empty, then merge the right array.
    while j < right.len() {
        arr[k] = right[j];
        j += 1;
        k += 1;
    }
}

// merge sort function
pub fn merge_sort(arr: &mut [i32]){
    let len = arr.len();
    if len <= 1 {
        return;
    }

    // split the array into two halves and calls merge sort on both halves.
    let mid = len/2;
    let mut left = arr[..mid].to_vec();
    let mut right =  arr[mid..].to_vec();

    merge_sort(&mut left);
    merge_sort(&mut right); 

    // calls the merge function to merge the two halves.
    merge(arr, &left, &right);
}
# Sleep Sort Algorithm: Traditional and Modified

## Traditional Sleep Sort
**Sleep Sort** is a unique, unconventional sorting algorithm that uses the system's multitasking and timer 
capabilities. It works by:
1. Creating a thread for each number in the array.
2. Each thread "sleeps" for an amount of time proportional to the value it represents.
3. After waking, the thread outputs the value.

### Characteristics
- **Input:** Array of non-negative integers.
- **Output:** A sorted array.
- **Complexity:** Highly dependent on implementation; real-world practicality is limited.
- **Use Case:** Educational or experimental purposes, showcasing creative algorithm design.

### Issues
- Thread overhead.
- Requires non-negative integers.
- Not suitable for real-world applications.

---

## Modified Sleep Sort (Digit-Based)
This project implements a modified version of Sleep Sort. Instead of relying purely on sleep durations, the 
algorithm uses digit-based processing (like radix sort) to enhance performance and produce nearly sorted arrays 
more efficiently.

### Key Modifications:
1. **Digit Bucketing:** Sorts numbers digit by digit, combining Sleep Sort principles with bucket-based sorting.
2. **Handling Negative Numbers:** Directly addresses the issue by converting all numbers to positive values before sorting.

### Results:
Below are the recorded test results during the modified Sleep Sort 
execution for 100000 elements in (0,1000000):
```
Iter-0: Time taken to sort: 54.5770841s, NumUnsorted: 4987
Iter-1: Time taken to sort: 53.6901553s, NumUnsorted: 1733
Iter-2: Time taken to sort: 53.9335724s, NumUnsorted: 1022
Iter-3: Time taken to sort: 52.8005675s, NumUnsorted: 1324
Iter-4: Time taken to sort: 54.2553693s, NumUnsorted: 5977
Iter-5: Time taken to sort: 54.0668559s, NumUnsorted: 4326
Iter-6: Time taken to sort: 53.6720129s, NumUnsorted: 1434
Iter-7: Time taken to sort: 53.8010895s, NumUnsorted: 1634
Iter-8: Time taken to sort: 55.1068059s, NumUnsorted: 1491
Iter-9: Time taken to sort: 53.9399309s, NumUnsorted: 891
Iter-10: Time taken to sort: 55.1479952s, NumUnsorted: 338
Iter-11: Time taken to sort: 53.9148945s, NumUnsorted: 2124
Iter-12: Time taken to sort: 53.7244459s, NumUnsorted: 2734
Iter-13: Time taken to sort: 53.9767424s, NumUnsorted: 828
Iter-14: Time taken to sort: 54.1436026s, NumUnsorted: 1437
Iter-15: Time taken to sort: 53.8498678s, NumUnsorted: 1636
Iter-16: Time taken to sort: 53.8055089s, NumUnsorted: 457
Iter-17: Time taken to sort: 53.830473s, NumUnsorted: 1194
Iter-18: Time taken to sort: 53.6855584s, NumUnsorted: 1748
Iter-19: Time taken to sort: 54.0396626s, NumUnsorted: 1391
Iter-20: Time taken to sort: 53.8572425s, NumUnsorted: 933
Iter-21: Time taken to sort: 53.6339366s, NumUnsorted: 406
Iter-22: Time taken to sort: 53.2707436s, NumUnsorted: 410
Iter-23: Time taken to sort: 53.7752576s, NumUnsorted: 809
Iter-24: Time taken to sort: 53.9908028s, NumUnsorted: 1165
```

### Analysis
1. **Average Time Taken**: ~53.940s. Approximately equal to log10(maximum values in array) * 9s, as radix sort is used
2. **Percentage of Sorting Per Iteration**: On average 1700 elments remain unsorted out of 100000 elements,
which is 1.7% of the array size. 95-99.5 % sorting is achieved. Depends on system load as well.
The loss in accuracy is the trade-off for the increased
3. **Final Sorting of elements**: The final unsorted array is sorted using traditional sorting algorithms which are good for nearly sorted arrays.

---

## Future Work
- Optimize memory usage during digit bucketing.
- Integrate advanced sorting techniques for the final stages.
- Improving accuracy by altering sleep times based on array size and array values.
/*
    Radix Sort Algorithm
    --------------------

    Time Complexity - O(d*(n + b)),
        where d is the max number of digits, n is the size of the input & b is the base of the system.

    Space Complexity - O(n + b),
        where n is the size of the input & b is the base of the system.
*/

#include <iostream>
#include <vector>
#include <algorithm>
#include <cstdlib>
#include <ctime>

using namespace std;

// Count Sort Algorithm.
void countSort(vector<int>& arr, int pos) {
    int size = arr.size();
    vector<int> digit_frequency(10, 0); // Sorting decimal numbers here, hence base is 10.

    // Increments the freq of elements in each bucket.
    for (int num : arr) {
        ++digit_frequency[(num/pos) % 10]; 
    }

    // Calculates the cumulative frequency.
    for (int i = 1; i < 10; i++) {
        digit_frequency[i] += digit_frequency[i - 1];
    }

    // Builds the sorted array from the right to maintain the stability property of Radix Sort.
    vector<int> sorted_arr(size);
    for (int i = size-1; i >= 0; i--) {
        int idx = (arr[i]/pos) % 10;
        sorted_arr[--digit_frequency[idx]] = arr[i]; 
    }

    arr = sorted_arr; // Copies the sorted array to the original array.
}

// Radix Sort Algorithm.
void radixSort(vector<int>& arr) {
    int max = *max_element(arr.begin(), arr.end());
    for(int pos = 1; max/pos > 0; pos *= 10) { 
        countSort(arr, pos); // Calls countSort() for every digit position.
    }
}

// Helper function to print elements of array.
void printArray(vector<int>& arr) {
    for (int num : arr) {
        cout << num << " ";
    }
    cout << endl;
}

void sortArray(vector<int>& arr) {
    if (arr.empty()) { 
        cerr << "The Array to Sort is Empty!" << endl; 
        exit(1);
    }

    vector<int> negative_nums;
    vector<int> positive_nums;

    // Splits the elements of the array into 2 separate arrays.
    for (int i = 0; i < arr.size(); i++) {
        if (arr[i] < 0) {
            negative_nums.push_back(-arr[i]); // Converts negative numbers to positive for Radix Sort.
        } else {
            positive_nums.push_back(arr[i]);
        }
    }

    // Sorts both arrays independently using Radix Sort.
    if (!negative_nums.empty()) {
        radixSort(negative_nums);
    }
    if (!positive_nums.empty()) {
        radixSort(positive_nums);
    }
    
    // Merges the 2 sorted arrays back to original array.
    int index = 0;
    for (int i = negative_nums.size()-1; i >= 0; i--) {
        arr[index++] = -negative_nums[i]; // Converts the numbers back to their negative form.
    }
    for (int num : positive_nums) {
        arr[index++] = num;
    }
}

int main() {
    const int size = 20;
    vector<int> arr(size);

    // Create a worst case unsorted array (descending elements).
    // for (int i = 0; i < size; i++) {
    //     arr[i] = size - i; 
    // }

    // Create an unsorted array with random numbers (negative and positive).
    srand((unsigned) time(0));
    for (int i = 0; i < size; i++) {
        arr[i] = (rand() % 1001) - 500;
    }

    cout << "Unsorted Array: " << endl;
    printArray(arr);

    sortArray(arr);

    cout << "Sorted Array: " << endl;
    printArray(arr);

    return 0;
}

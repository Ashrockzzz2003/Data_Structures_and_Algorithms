#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

// Insertion Sort for individual buckets
void insertionSort(vector<double>& bucket) {
    for (int i = 1; i < bucket.size(); i++) {
        double key = bucket[i];
        int j = i - 1;
        while (j >= 0 && bucket[j] > key) {
            bucket[j + 1] = bucket[j];
            j--;
        }
        bucket[j + 1] = key;
    }
}

// Bucket Sort Function
void bucketSort(vector<double>& arr) {
    int n = arr.size();
    vector<vector<double>> buckets(n);

    for (double num : arr) {
        int bi = n * num;
        buckets[bi].push_back(num);
    }

    for (auto& bucket : buckets) {
        insertionSort(bucket);
    }

    int index = 0;
    for (const auto& bucket : buckets) {
        for (double num : bucket) {
            arr[index++] = num;
        }
    }
}

int main() {
    vector<double> arr = {0.897, 0.565, 0.656, 0.1234, 0.665, 0.3434};
    bucketSort(arr);

    cout << "Sorted array is:\n";
    for (double num : arr) {
        cout << num << " ";
    }
    cout << endl;
    return 0;
}
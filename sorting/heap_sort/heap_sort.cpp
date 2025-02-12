#include <iostream>
#include <vector>
#include <queue>

using namespace std;

// Heap Sort Function
void heapSort(vector<int>& arr) {
    priority_queue<int, vector<int>, greater<int>> minHeap;

    // Insert elements into the min-heap
    for (int num : arr) {
        minHeap.push(num);
    }

    // Extract elements from the min-heap and store them back in the array
    for (size_t i = 0; i < arr.size(); i++) {
        arr[i] = minHeap.top();
        minHeap.pop();
    }
}

int main() {
    vector<int> arr = {12, 11, 13, 5, 6, 7};
    cout << "Original array:" << endl;
    for (int num : arr) {
        cout << num << " ";
    }
    cout << endl;

    heapSort(arr);

    cout << "Sorted array:" << endl;
    for (int num : arr) {
        cout << num << " ";
    }
    cout << endl;
    return 0;
}

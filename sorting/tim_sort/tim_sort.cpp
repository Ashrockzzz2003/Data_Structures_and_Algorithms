/*
    Cpp program to perform Timsort.

    Algorithm:

    Step 1 - Divide the array into the number of blocks known as run.
    Step 2 - Consider the size of run, either 32 or 64.
    Step 3 - Sort the individual elements of every run one by one using insertion sort.
    Step 4 - Merge the sorted runs one by one using the merge function of merge sort.
    Step 5 - Double the size of merged sub-arrays after every iteration.

    Time Complexity Analaysis:

    Best    Case O(n)
    Average Case O(n*log(n))
    Worst   Case O(n*log(n))

    Space O(n)

*/
#include <bits/stdc++.h> 
using namespace std;

const int RUN_SIZE = 32; // For larger data sets we can use 64 as the run size.
/*
    Insertion Sort

    Algorithm for perfInsSort:

    Step 1 - Iterate through the array from the index (start + 1) to stop.
    Step 2 - Select the current element as the key for insertion.
    Step 3 - Compare the key with elements in the sorted part of the array (start to the previous index).
    Step 4 - Shift larger elements one position to the right to make space for the key.
    Step 5 - Insert the key into its correct position in the sorted portion of the array.
    Step 6 - Repeat this process for every element from (start + 1) to stop.

    Time Complexity Analysis:

    Best    Case O(n)
    Average Case O(n^2)
    Worst   Case O(n^2)


    Space O(1)- As the sorting is done in-place.
*/
void perfInsSort(vector<int>&arr, int start, int stop)
{
    for (int idx = start + 1; idx <= stop; ++idx)
    {
        int current = arr[idx];
        int prev = idx - 1;
        while (prev >= start && arr[prev] > current)
        {
            arr[prev + 1] = arr[prev];
            prev--;
        }
        arr[prev + 1] = current;
    }
}

/*
    Merge Sort Helper Function

    Algorithm for mergeSegments:

    Step 1 - Divide the array into two segments (left and right) by finding the middle index.
    Step 2 - Create two temporary arrays to hold the values of the left and right segments.
    Step 3 - Copy the values from the original array to the temporary arrays.
    Step 4 - Merge the two temporary arrays back into the original array by comparing their elements.
    Step 5 - Insert the smaller of the two elements at each step and move the index.
    Step 6 - Once one array is exhausted, insert the remaining elements of the other array.

    Time Complexity Analysis:

    Best    Case O(n)
    Average Case O(n*log(n))
    Worst   Case O(n*log(n))

    Space O(n)- As temporary arrays are created for merging.
*/
void mergeSegments(vector<int>& arr, int left, int middle, int right)
{
    int len1 = middle - left + 1;
    int len2 = right - middle;

    vector<int> leftPart(len1), rightPart(len2);

    for (int i = 0; i < len1; ++i) leftPart[i] = arr[left + i];
    for (int j = 0; j < len2; ++j) rightPart[j] = arr[middle + 1 + j];
    int i = 0, j = 0, k = left;
    while (i < len1 && j < len2)
    {
        if (leftPart[i] <= rightPart[j])
        {
            arr[k] = leftPart[i];
            ++i;
        }
        else
        {
            arr[k] = rightPart[j];
            ++j;
        }
        ++k;
    }

    while (i < len1)
    {
        arr[k] = leftPart[i];
        ++i;
        ++k;
    }

    while (j < len2)
    {
        arr[k] = rightPart[j];
        ++j;
        ++k;
    }
}

/*
    Tim Sort Algorithm

    Algorithm for executeTimsort:

    Step 1 - Split the array into smaller sub-arrays (runs) of size RUN_SIZE.
    Step 2 - Sort each run individually using insertion sort.
    Step 3 - Merge the sorted runs progressively using the mergeSegments function.
    Step 4 - Double the size of the merged runs in each iteration to ensure optimal merging.

    Time Complexity Analysis:

    Best    Case O(n)
    Average Case O(n*log(n))
    Worst   Case O(n*log(n))

    Space O(n)
*/
void executeTimSort(vector<int>& arr, int size)
{
    if(size<=0) return;
    // Sort chunks of size RUN_SIZE using insertion sort.
    for (int i = 0; i < size; i += RUN_SIZE)
        perfInsSort(arr, i, min(i + RUN_SIZE - 1, size - 1));

    // Merge sorted chunks, doubling the merge size in each iteration.
    for (int run = RUN_SIZE; run < size; run *= 2)
    {
        for (int start = 0; start < size; start += 2 * run)
        {
            int mid = start + run - 1;
            int end = min(start + 2 * run - 1, size - 1);
            if (mid < end) mergeSegments(arr, start, mid, end);
        }
    }
}

void displayArray(const vector<int>& arr)
{
    for (int val : arr)
        cout << val << " ";
    cout << endl;
}

int main()
{
    const int sizeOfArray = 128;  
    vector<int> arr(sizeOfArray);

    // Initialize the array with descending values.
    for (int i = 0; i < sizeOfArray; i++)
        arr[i] = sizeOfArray - i;

    cout << "Array before sorting:" << endl;
    displayArray(arr);

    // Sort the array using Timsort.
    executeTimSort(arr, arr.size());

    cout << "Array after sorting:" << endl;
    displayArray(arr);

    return 0;
}
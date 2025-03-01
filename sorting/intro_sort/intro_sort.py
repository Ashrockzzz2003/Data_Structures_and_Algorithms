import math
from random import randint

def insertionSort(A, start , end):
    for j in range(start + 1, end + 1):
        i = j - 1
        key = A[j]

        while i >= start and A[i] > key:
            A[i+1] = A[i]
            i -= 1

        A[i+1] = key


def maxHeapify(A, start, end):
    left = 2 * start + 1
    right = 2 * start + 2

    if left < end and A[left] > A[start]:
        largest = left
    else:
        largest = start

    if right < end and A[right] > A[largest]:
        largest = right

    if largest != start:
        A[start], A[largest] = A[largest], A[start]
        maxHeapify(A, largest, end)


def buildMaxHeap(A, start, end):
    for i in range(start + (end - start) // 2, start, -1):
        maxHeapify(A, i, end)


def heapSort(A, start, end):
    buildMaxHeap(A, start, end)

    for i in range(end-1, start, -1):
        A[start], A[i] = A[i], A[start]
        maxHeapify(A, start , i-1)


def partition(A, start, end):
    pivot = A[end]

    i = start - 1

    for j in range(start, end):
        if A[j] < pivot:
            i += 1
            A[i], A[j] = A[j], A[i]

    A[i+1], A[end] = A[end], A[i+1]
    return i+1


def introHelper(A, start, end, max_depth):

# Helper function for Introsort that combines Quicksort, Heap Sort, and Insertion Sort.
    
    if end-start < 16:                               # Use Insertion Sort for small partitions 
        insertionSort(A, start, end)
    
    elif max_depth == 0:            
        heapSort(A, start, end)                      # Switch to Heap Sort if recursion depth limit reached

    else:
        pivot = partition(A, start, end)             # Quick Sort partitioning
        introHelper(A, start, pivot - 1, max_depth-1)
        introHelper(A, pivot + 1, end, max_depth -1)
        

def introSort(A):
    if len(A) <= 1:                                   # Array is already sorted if it has 0 or 1 element
        return
    max_depth = 2 * math.floor(math.log2(len(A)))     # Maximum allowed recursion depth
    introHelper(A, 0, len(A) - 1, max_depth)


if __name__ == "__main__":
    A = [randint(0, 100) for _ in range(200)]
    print("Unsorted array:", A)
    introSort(A)
    print("Sorted array using Heap Sort:", A)

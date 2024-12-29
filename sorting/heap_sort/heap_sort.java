package Data_Structures_and_Algorithms.sorting.heap_sort;

import java.util.Arrays;

public class heap_sort {
    
    public static void maxheapify(int[] A, int start, int end ){
        int left = 2 * start + 1;
        int right = 2 * start + 2;
        int largest = start;

        if(left < end && A[left] > A[largest]) {
            largest = left;
        }

        if(right < end && A[right] > A[largest]){
            largest = right;
        }

        if(largest != start){
            int temp = A[largest];
            A[largest] = A[start];
            A[start] = temp;

            maxheapify(A, largest , end);
        }     
    }

    public static void buildMaxHeap(int[] A, int start , int end){

        for(int i= (end - 1) / 2; i >= start; i--){
            maxheapify(A, i, end);
        }
    }

    public static void heapSort(int[] A, int start, int end){

        if (A == null || start < 0 || end > A.length || start >= end) {
            throw new IllegalArgumentException("Invalid start or end indices.");
        }

        if( end - start <=1){
            return;
        }

        buildMaxHeap(A, start, end);

        for(int i = end - 1; i > start; i--){
            int temp = A[start];
            A[start] = A[i];
            A[i] = temp;
            maxheapify(A, start, i);
        }
    }

    public static void main(String[] args){
        int[] A={8, 1, 2, 9, 4, 4 , 3, 3, 3 ,3};
        System.out.println("Unsorted array: " + Arrays.toString(A));
        heapSort(A, 0 , A.length);
        System.out.println("Sorted array using Heap Sort: " + Arrays.toString(A));
    }
}

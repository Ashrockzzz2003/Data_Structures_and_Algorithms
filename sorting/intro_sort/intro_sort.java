package Data_Structures_and_Algorithms.sorting.intro_sort;
import Data_Structures_and_Algorithms.sorting.heap_sort.heap_sort;
import Data_Structures_and_Algorithms.sorting.insertion_sort.insertion_sort;
import java.util.Arrays;


public class intro_sort {

    public static int  partition(int[] A, int start, int end){
        end = end-1;
        int pivot = A[end];
        int i = start - 1;

        for(int j = start; j < end; j++){
            if(A[j] < pivot){
                i++;
                int temp = A[i];
                A[i] = A[j];
                A[j] = temp;
            }
        }

        int temp = A[i + 1];
        A[i + 1] = A[end];
        A[end] = temp;

        return i + 1; 
    }

    public static void  introHelper(int[] A, int start , int end , int maxDepth){

    // Helper function for Introsort that combines Quicksort, Heap Sort, and Insertion Sort.

        if(end - start < 16){                        //Use Insertion Sort for small partitions 
            insertion_sort.insertionSort(A, start, end);
        }

        else if (maxDepth == 0){                     // Switch to Heap Sort if recursion depth limit reached
            heap_sort.heapSort(A, start, end);
        }

        else{
            int pivot = partition(A, start, end);    // Quick Sort partitioning
            introHelper(A, start, pivot, maxDepth-1);
            introHelper(A, pivot + 1, end, maxDepth - 1);
        }
    }

    public static void introSort(int[] A){
        if(A.length <= 1){                           // Array is already sorted if it has 0 or 1 element
            return;
        }
        int maxDepth =(int) (2 * Math.log(A.length) / Math.log(2));  // Maximum allowed recursion depth
        introHelper(A, 0, A.length, maxDepth);
    }

    public static void main(String[] args) {
        int[][] testCases = {
            {},                          // Empty array
            {1},                         // Single element
            {5, 3},                      // Two elements unsorted
            {3, 5},                      // Two elements sorted
            {5, 3, 8, 6, 2},             // Multiple elements unsorted
            {8, 6, 5, 3, 2},             // Multiple elements reverse sorted
            {1, 1, 1},                   // All elements identical
            {-1, -5, -3},                // Negative numbers
            {0},                         // Single zero element
            {10, 3, 4, 5, 1, 0, 4, 1}    // normal test case
        };

        for (int[] testCase : testCases) {
            System.out.println("Unsorted array: " + Arrays.toString(testCase));
            introSort(testCase);
            System.out.println("Sorted array using Introsort: " + Arrays.toString(testCase));
            System.out.println();
        }
    }
}

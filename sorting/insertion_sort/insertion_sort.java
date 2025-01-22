package Data_Structures_and_Algorithms.sorting.insertion_sort;

import java.util.Arrays;

public class insertion_sort {

    public static void insertionSort(int[] A, int start, int end){

        if( end - start <=1){
            return;
        }

        for( int i = start + 1; i < end; i++){
            int j = i - 1;
            int key = A[i];

            while(j >= start && A[j] > key){
                A[j + 1] = A[j];
                j--;
            }
            A[j+1] = key;
        }
    }

    public static void main(String[] args){
        int[] A={9,8,7,6,5,4};
        System.out.println("Unsorted array: " + Arrays.toString(A));
        insertionSort(A, 0 , A.length);
        System.out.println("Sorted array using Insertion Sort: " + Arrays.toString(A));
    }
}
public class MergeSort {
    public static void mergeSort(int[] arr) {
        if (arr != null && arr.length > 1) {
            long mid = (long) arr.length / 2;

            int[] left = new int[(int) mid];
            int[] right = new int[arr.length - (int) mid];

            // Dividing the array into two
            System.arraycopy(arr, 0, left, 0, (int) mid);
            System.arraycopy(arr, (int) mid, right, 0, arr.length - (int) mid);

            // Recursively calling mergeSort on both half's of the array
            mergeSort(left);
            mergeSort(right);

            // Merging the sorted left and sorted right array
            merge(left, right, arr);
        }
    }
    public static void merge(int [] left, int[] right, int[] arr) {
        // i -> tracks the index of left, j -> tracks the index of right, j -> tracks the index of array
        int i = 0, j = 0, k = 0;

        // Merging elements from left and right array into a single array while maintaining order
        while (i < left.length && j < right.length) {
                if (left[i] < right[j]) {
                    arr[k++] = left[i++];
                }
                else {
                    arr[k++] = right[j++];
                }
            }

        // if right array is empty, copy the remaining elements from left array
        while (i < left.length) {
                arr[k++] = left[i++];
            }

        // if left array is empty, copy the remaining elements from right array
        while (j < right.length) {
                arr[k++] = right[j++];
            }
    }

    public static void main(String[] args) {
       // use small sizes if printing the array before and after sorting
       int size = 10;

       int[] arr = new int[size]; 

       for (int i = 0; i < size; i++) {
           arr[i] = (int)(Math.random() * size);
       }

       System.out.print("Array before sorting : ");
       for (int i = 0; i < size; i++) {
           System.out.print(arr[i] + " ");
       }

       long startTime = System.nanoTime();
       mergeSort(arr);
       long endTime = System.nanoTime();

       System.out.print("\nArray after sorting : ");
       for (int i = 0; i < size; i++) {
           System.out.print(arr[i] + " ");
       } 

       System.out.println();

       long normalMergeSortTime = endTime - startTime;
       System.out.println("Time taken for parallel merge sort: " + normalMergeSortTime + " ns");
    }
}


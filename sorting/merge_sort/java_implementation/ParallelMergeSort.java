import java.util.concurrent.ForkJoinPool;
import java.util.concurrent.RecursiveAction;

public class ParallelMergeSort {

    // Threshold to switch to normal merge sort
    // change this value according to the size of the array to get the best performance
    private static final int THRESHOLD = 4096;

    public static void parallelMergeSort(int[] arr) {
        SortTask sortTask = new SortTask(arr);
        ForkJoinPool pool = new ForkJoinPool();
        pool.invoke(sortTask);
    }

    public static class SortTask extends RecursiveAction {
        private final int[] arr;

        public SortTask(int[] arr) {
            this.arr = arr;
        }

        @Override
        protected void compute() {
            if (arr.length > 1) {
                // If array is larger than threshold, proceed with parallel merge sort
                if (arr.length > THRESHOLD) {
                    int mid = arr.length / 2;

                    int[] left = new int[mid];
                    int[] right = new int[arr.length - mid];

                    System.arraycopy(arr, 0, left, 0, mid);
                    System.arraycopy(arr, mid, right, 0, arr.length - mid);

                    // Create subtasks for the two halves
                    SortTask first = new SortTask(left);
                    SortTask second = new SortTask(right);

                    // Invoke both tasks in parallel
                    invokeAll(first, second);

                    // Merge the two halves
                    MergeSort.merge(left, right, arr);
                } else {
                    // For arrays smaller than or equal to the threshold, use normal merge sort
                    MergeSort.mergeSort(arr);
                }
            }
        }
    }

    public static void main(String[] args) {
            // use large sizes if measuring the time taken for sorting, also change the threshold value
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
            parallelMergeSort(arr);
            long endTime = System.nanoTime();

            System.out.print("\nArray after sorting : ");
            for (int i = 0; i < size; i++) {
                System.out.print(arr[i] + " ");
            } 

            System.out.println();

            long parallelMergeSortTime = endTime - startTime;
            System.out.println("Time taken for parallel merge sort: " + parallelMergeSortTime + " ns");

    }   
}

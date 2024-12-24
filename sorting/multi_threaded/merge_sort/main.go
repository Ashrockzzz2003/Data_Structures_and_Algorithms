package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arraySize := 100000

	array1 := []int{}

	for i := 0; i < arraySize; i++ {
		array1 = append(array1, rand.Intn(arraySize))
	}

	// Timing normal merge sort
	startMergeSort := time.Now()
	mergeSort(array1)
	elapsedMergeSort := time.Since(startMergeSort)
	fmt.Println("MergeSort for", arraySize, "elements took", elapsedMergeSort)

	// Timing parallel merge sort
	array2 := make([]int, len(array1))
	copy(array2, array1)
	startParallelMergeSort := time.Now()
	ch := make(chan []int)
	go parallelMergeSort(array2, ch)
	<-ch
	elapsedParallelMergeSort := time.Since(startParallelMergeSort)
	fmt.Println("ParallelMergeSort for", arraySize, "elements took", elapsedParallelMergeSort)

	// Compare the two
	speedupPercentage := (float64(elapsedMergeSort-elapsedParallelMergeSort) / float64(elapsedMergeSort)) * 100
	fmt.Println("ParallelMergeSort was", speedupPercentage, "percent faster")
}


// Result:
// ParallelMergeSort was 63.96576490304046 percent faster
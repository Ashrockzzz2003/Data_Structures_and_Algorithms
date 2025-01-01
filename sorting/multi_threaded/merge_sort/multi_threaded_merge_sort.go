package main

import (
	"fmt"
	"math/rand"
	"time"
)

const threshold = 10000

func parallelMergeSort(arr []int, c chan []int) {
	if len(arr) <= 1 {
		c <- arr
		return
	}

	if len(arr) <= threshold {
		timSort(arr)
		c <- arr
		return
	}

	middleIndex := len(arr) / 2
	leftChan := make(chan []int)
	rightChan := make(chan []int)

	go parallelMergeSort(arr[:middleIndex], leftChan)
	go parallelMergeSort(arr[middleIndex:], rightChan)

	leftSorted := <-leftChan
	rightSorted := <-rightChan

	c <- merge(leftSorted, rightSorted)
}

func merge(leftArray, rightArray []int) []int {
	result := make([]int, len(leftArray)+len(rightArray))
	leftIndex, rightIndex := 0, 0
	for i := 0; i < len(result); i++ {
		if leftIndex >= len(leftArray) {
			result[i] = rightArray[rightIndex]
			rightIndex++
			continue
		}
		if rightIndex >= len(rightArray) {
			result[i] = leftArray[leftIndex]
			leftIndex++
			continue
		}
		if leftArray[leftIndex] < rightArray[rightIndex] {
			result[i] = leftArray[leftIndex]
			leftIndex++
		} else {
			result[i] = rightArray[rightIndex]
			rightIndex++
		}
	}
	return result
}

func timSort(array []int) {
	arrayLen := len(array)
	minRun := 32

	// Sort the small runs using insertion sort.
	for i := 0; i < arrayLen; i += minRun {
		end := min(i+minRun, arrayLen)
		insertionSort(array[i:end])
	}

	// Merge the sorted runs using merge function.
	currentRunSize := minRun
	for currentRunSize < arrayLen {
		for leftStart := 0; leftStart < arrayLen; leftStart += 2 * currentRunSize {
			middleIndex := min(leftStart+currentRunSize, arrayLen)
			rightEnd := min(leftStart+2*currentRunSize, arrayLen)

			leftArray := array[leftStart:middleIndex]
			rightArray := array[middleIndex:rightEnd]
			mergedArray := merge(leftArray, rightArray)

			copy(array[leftStart:rightEnd], mergedArray)
		}
		currentRunSize *= 2
	}
}

func insertionSort(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && key < array[j] {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
}

func main() {
	arraySize := 100000

	array1 := []int{}

	for i := 0; i < arraySize; i++ {
		array1 = append(array1, rand.Intn(arraySize))
	}
	array2 := make([]int, len(array1))
	copy(array2, array1)

	// Timing Tim sort.
	startTimSort := time.Now()
	timSort(array1)
	elapsedTimSort := time.Since(startTimSort)
	fmt.Println("TimSort for", arraySize, "elements took", elapsedTimSort)

	// Timing parallel merge sort.
	startParallelMergeSort := time.Now()
	ch := make(chan []int)
	go parallelMergeSort(array2, ch)
	<-ch
	elapsedParallelMergeSort := time.Since(startParallelMergeSort)
	fmt.Println("ParallelMergeSort for", arraySize, "elements took", elapsedParallelMergeSort)

	// Compare the two.
	speedupPercentage := (float64(elapsedTimSort-elapsedParallelMergeSort) / float64(elapsedTimSort)) * 100
	fmt.Println("ParallelMergeSort was", speedupPercentage, "percent faster")
}

/*
    Observation:

    - When the array size is small using goroutines and channels to sort the array is innefficient,
    potentially becoming a constraint for the system.
    - So in such cases we use timsort instead of parallel merge sort.

    Result:

	- ParallelMergeSort was 50.2724112133081 percent faster for 100000
	- ParallelMergeSort was 52.23465286992048 percent faster for 1000000
*/

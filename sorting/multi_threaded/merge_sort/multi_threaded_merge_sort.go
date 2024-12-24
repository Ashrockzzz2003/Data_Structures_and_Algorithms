package main

import "sort"

const threshold = 5000

func parallelMergeSort(arr []int, c chan []int) {
	if len(arr) <= 1 {
		c <- arr
		return
	}

	if len(arr) <= threshold {
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
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
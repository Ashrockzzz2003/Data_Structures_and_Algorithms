package main

func mergeSort(inputArray []int) []int {
	if len(inputArray) <= 1 {
		return inputArray
	}

	middleIndex := len(inputArray) / 2
	leftArray := mergeSort(inputArray[:middleIndex])
	rightArray := mergeSort(inputArray[middleIndex:])

	return merge(leftArray, rightArray)
}
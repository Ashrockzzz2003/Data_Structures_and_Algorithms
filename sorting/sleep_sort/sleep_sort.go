package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

func sleepSort(numbers []int) []int {
	length := len(numbers)
	var wg sync.WaitGroup
	done := make(chan int, length)

	minVal, maxVal := findMinMax(numbers)
	maxVal = maxVal - minVal

	for i := 0; i < length; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sleepTime := min(100000, time.Duration(n-minVal)*time.Millisecond*1000)
			time.Sleep(sleepTime)
			done <- n
		}(numbers[i])
	}
	wg.Wait()

	close(done)

	for i := 0; i < length; i++ {
		numbers[i] = <-done
	}
	return numbers
}

func digitSleepSort(numbers []int, maxDigit int) []int {
	sleepByDigit := func(nums []int, digitPlace int) []int {
		digitChannel := make(chan int, len(nums))
		var digitWG sync.WaitGroup

		count := make([]int, 10)
		currCount := make([]int, 10)
		for _, num := range nums {
			digit := (num / digitPlace) % 10
			count[digit]++

		}

		for index, num := range nums {
			digit := (num / digitPlace) % 10
			currCount[digit]++

			digitWG.Add(1)
			go func(n, d, offset int) {
				defer digitWG.Done()
				delay := time.Duration(d)*1000*time.Millisecond + time.Duration(offset)*time.Microsecond*10000
				time.Sleep(delay)
				digitChannel <- n
			}(num, digit, index)
		}

		digitWG.Wait()
		close(digitChannel)

		results := make([]int, 0, len(nums))
		for range nums {
			results = append(results, <-digitChannel)
		}

		return results
	}

	digitPlace := 1
	for digitPlace <= maxDigit {
		numbers = sleepByDigit(numbers, digitPlace)
		digitPlace *= 10
	}

	return numbers
}

func bucketSortSleep(numbers []int) [][]int {
	if len(numbers) == 0 {
		return [][]int{}
	}
	bucketSize := 100

	minVal, maxVal := findMinMax(numbers)

	// Shift numbers to be positive
	for i := range numbers {
		numbers[i] = numbers[i] - minVal + 1
	}

	maxVal = maxVal - minVal
	buckets := make([][]int, maxVal/bucketSize+1)

	// Place numbers in buckets
	for _, num := range numbers {
		bucketIndex := int(float64(num) / float64(bucketSize))
		buckets[bucketIndex] = append(buckets[bucketIndex], num)
	}

	// Sort each bucket concurrently
	var wg sync.WaitGroup
	for i, _ := range buckets {
		wg.Add(1)
		go func(bucket []int, key int) {
			defer wg.Done()
			buckets[key] = digitSleepSort(bucket, findMaxDigit(buckets[i]))
		}(buckets[i], i)
	}
	wg.Wait()

	return buckets
}

func findMinMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	min, max := numbers[0], numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	return min, max
}

func findMaxDigit(numbers []int) int {
	minNum, maxNum := findMinMax(numbers)
	minNum = minNum + 0
	maxDigit := 1
	for maxNum > 0 {
		maxDigit *= 10
		maxNum = int(maxNum / 10)
	}
	return maxDigit
}

func generateRandomArray(size int, maxValue int) []int {

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(int(maxValue))
	}

	return arr
}

func isSortedArray(numbers []int) int {
	count := 0
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			count++
			// fmt.Println(numbers[i], numbers[i+1])
		}
	}
	// fmt.Println(count)
	return count
}

func isSorted(numbers [][]int) int {
	count := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers[i])-1; j++ {
			if numbers[i][j] > numbers[i][j+1] {
				// fmt.Println(numbers[i])
				// fmt.Println(numbers[i], numbers[i+1])
				count++
			}
		}
	}
	fmt.Println(count)
	return count
}

func insertionSort(arr [][]int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			key := arr[i][j]
			k := j - 1
			// Shift elements that are greater than key
			for k >= 0 && arr[i][k] > key {
				arr[i][k+1] = arr[i][k]
				k--
			}
			arr[i][k+1] = key
		}
	}
	result := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			result = append(result, arr[i][j])
		}
	}

	return result
}

func writeToFile(fileName string, content string) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	file.WriteString(content)
}

func main() {
	for i := 0; i < 25; i++ {
		fmt.Println("Iter-", i)
		numbers := generateRandomArray(100000, 1000000)

		start := time.Now()
		sorted := bucketSortSleep(numbers)
		duration := time.Since(start)

		fmt.Print("Iter-", i, ":", "Time taken to sort:", duration, " ")
		result := isSorted(sorted)
		fmt.Println("NumUnsorted:", result)
		// resultNew := 0
		// if result != 0 {
		// 	sortedFull := insertionSort(sorted)
		// 	resultNew = isSortedArray(sortedFull)
		// }
		// resultNew := resultNew
		// writeToFile("sleepSort.txt", fmt.Sprintf("Iter-%d: Time taken to sort: %v, NumUnsorted: %d\n", i, duration, result))
	}
}

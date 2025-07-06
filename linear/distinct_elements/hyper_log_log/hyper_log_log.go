package main

import (
	"fmt"
	"hash/fnv"
	"math"
	"math/bits"
	"math/rand"
	"sync"
)

// HYPERLOGLOG ALGORITHM :
// Cardinality = constant . m . ( m / sum_{j=1:N} 2^{-R_j} )
// Error percentage = (1.04 / sqrt(m)) * 100

// Calculating cardinality using harmonic mean
func cardinality(leadingZeroes []int) float64 {
	if len(leadingZeroes) == 0 {
		return 0.0
	}
	numBuckets := len(leadingZeroes)
	var harmonicMean float64
	for _, num := range leadingZeroes {
		if num == 0 {
			continue
		}
		harmonicMean += 1.0 / math.Pow(2.0, float64(num))
	}
	if harmonicMean == 0 {
		return 0.0
	}
	alpha := 0.7213 / (1.0 + 1.079/float64(numBuckets))
	return alpha * float64(numBuckets) * (float64(numBuckets) / harmonicMean)
}

// HyperLogLog function for cardinality estimation
func HyperLogLog[T any](values []T, hashFunc func(T) uint64) int {
	numBuckets := 256
	log2NumBuckets := uint64(math.Log2(float64(numBuckets)))
	currCounts := make([]int, numBuckets)

	for _, val := range values {
		hash := hashFunc(val)

		// Find the bucket index (using only the first `log2NumBuckets` bits)
		bucketIndex := (hash & (uint64(numBuckets) - 1))
		consecutiveZeroes := bits.TrailingZeros64((hash >> log2NumBuckets)) + 1
		if consecutiveZeroes > currCounts[bucketIndex] {
			currCounts[bucketIndex] = consecutiveZeroes
		}
	}
	return int(cardinality(currCounts))
}

func HyperLogLogPart[T any](values []T, hashFunc func(T) uint64) []int {
	numBuckets := 256
	log2NumBuckets := uint64(math.Log2(float64(numBuckets)))
	currCounts := make([]int, numBuckets)

	for _, val := range values {
		hash := hashFunc(val)

		// Find the bucket index (using only the first `log2NumBuckets` bits)
		bucketIndex := (hash & (uint64(numBuckets) - 1))
		consecutiveZeroes := bits.TrailingZeros64((hash >> log2NumBuckets)) + 1
		if consecutiveZeroes > currCounts[bucketIndex] {
			currCounts[bucketIndex] = consecutiveZeroes
		}
	}
	return currCounts
}

func parallelHyperLogLog[T any](values []T, hashFunc func(T) uint64) int {
	numParts := 10
	partSize := len(values) / numParts
	parallelCounts := make([][]int, numParts)

	var wg sync.WaitGroup
	for i := 0; i < numParts; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			partStart := i * partSize
			partEnd := (i + 1) * partSize
			if i == numParts-1 {
				partEnd = len(values)
			}
			parallelCounts[i] = HyperLogLogPart(values[partStart:partEnd], hashFunc)
		}(i)
	}
	wg.Wait()

	finalCounts := make([]int, len(parallelCounts[0]))
	for i := 0; i < len(parallelCounts[0]); i++ {
		max := 0
		for j := 1; j < numParts; j++ {
			if parallelCounts[j][i] > max {
				max = parallelCounts[j][i]
			}
		}
		finalCounts[i] = max
	}
	return int(cardinality(finalCounts))
}

func generateRandomArray(size int, maxValue int) ([]int, int) {
	arr := make([]int, size)
	arrSet := make(map[int]bool)
	for i := 0; i < size; i++ {
		val := rand.Intn(maxValue)
		arr[i] = val
		arrSet[val] = true
	}
	return arr, len(arrSet)
}

func fnv1aHash(val int) uint64 {
	hasher := fnv.New64a()
	hasher.Write([]byte(fmt.Sprintf("%d", val)))
	return hasher.Sum64()
}

func runHyperLogLog(verbose bool) float64 {
	values, realCount := generateRandomArray(15000, 10000)
	result := HyperLogLog(values, fnv1aHash)
	if verbose {
		fmt.Printf("Estimated cardinality: %d\n", result)
	}
	error := math.Abs(float64(realCount-result)) / float64(realCount)

	return error
}

func runParallelHyperLogLog(verbose bool) float64 {
	values, realCount := generateRandomArray(15000, 10000)
	result := parallelHyperLogLog(values, fnv1aHash)
	if verbose {
		fmt.Printf("Estimated cardinality: %d\n", result)
	}
	error := math.Abs(float64(realCount-result)) / float64(realCount)
	return error
}

func main() {
	// Hyper log log
	// avgError_hll := 0.0
	// numTrials_hll := 25000
	// for i := 0; i < numTrials_hll; i++ {
	// 	if i%1000 == 0 {
	// 		fmt.Printf("Trial %d\n", i)
	// 	}
	// 	avgError_hll += runHyperLogLog(false)
	// }
	// fmt.Printf("Average error: %f\n", avgError_hll/float64(numTrials_hll))

	// Parallel Hyper log log
	// avgError_phll := 0.0
	// numTrials_phll := 25000
	// for i := 0; i < numTrials_phll; i++ {
	// 	if i%1000 == 0 {
	// 		fmt.Printf("Trial %d\n", i)
	// 	}
	// 	avgError_phll += runParallelHyperLogLog(false)
	// }
	// fmt.Printf("Average error: %f\n", avgError_phll/float64(numTrials_phll))

	runHyperLogLog(true)
	runParallelHyperLogLog(true)
}

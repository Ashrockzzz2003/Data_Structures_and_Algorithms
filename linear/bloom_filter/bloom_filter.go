/*
Bloom Filter:
A Bloom Filter is a probabilistic data structure that is used to test whether an element is in the bloom filter.
It can give false positives but never false negatives. So, it may say an item is present when it is not,
but if it says an item is absent, then it is definitely not there.

How it works:
1. The Bloom Filter uses multiple hash functions to map an input to different positions in a boolean array.
2. To add an element, the corresponding positions in the boolean array are set to true.
3. To check an element, the same hash functions are used to check if all corresponding positions are true.
4. If all positions are true, the item might be in the bloom filter (false positive possible); if any position is false,
   the item is definitely not in the bloom filter.

Usage:
- Used in caching mechanisms like Google BigTable and databases to reduce expensive disk lookups.
  Instead of checking if a data is in the database, which takes more resouces and time,
  the bloom filter is used to check if the data is there or not.
- Used to check if a website is malicious or not.
  The bloom filter checks if the hash of that website address is present in it or not.

References:
- https://medium.com/@meeusdylan/creating-a-bloom-filter-with-go-7d4e8d944cfa
- https://en.wikipedia.org/wiki/Bloom_filter
*/

package main

import (
	"fmt"
	"hash/fnv"
	"math"
)

type bloomFilter struct {
	filter       []bool // Boolean array representing the Bloom Filter
	filterLength int    // Size of the Bloom Filter
	noOfHashes   int    // Number of hash functions to use
	noOfElements int    // Number of elements currently added in the Bloom Filter
}

// Creates a new Bloom Filter with given filter size and estimated number of elements.
func newBloomFilter(filterLength int, totalElements int) *bloomFilter {
	return &bloomFilter{
		filter:       make([]bool, filterLength),
		filterLength: filterLength,
		noOfHashes:   calculateNoOfHashes(filterLength, totalElements),
		noOfElements: 0,
	}
}

// Calculates the optimal number of hash functions to use using the formula: (m/n)*ln(2)
// where m is the size of Bloom Filter, n is the number of elements
func calculateNoOfHashes(filterLength int, noOfElements int) int {
	// To prevent division by zero
	if noOfElements == 0 {
		return 1
	}
	return int(math.Round((float64(filterLength) / (float64(noOfElements)) * math.Ln2)))
}

// Adds an input string to the Bloom Filter by setting multiple bits to true based on
// the indices returned by double hashing.
func (bf *bloomFilter) addInput(input string) {
	hashVals := doubleHashing(input, bf.filterLength, bf.noOfHashes)
	for _, num := range hashVals {
		bf.filter[num] = true
	}
	bf.noOfElements++
}

// Checks if the string is possibly in the Bloom Filter.
func (bf *bloomFilter) checkInput(input string) bool {
	hashVals := doubleHashing(input, bf.filterLength, bf.noOfHashes)
	// If any position is false, then the item is definitely not present
	for _, num := range hashVals {
		if !bf.filter[num] {
			return false
		}
	}

	return true
}

// Uses the FNV-1a hash function(uses less memory and faster) to generate a 32-bit hash value from an input string.
func hashFunc(input string) uint32 {
	h := fnv.New32a()      // Creates new instance of the hash function.
	h.Write([]byte(input)) // Converts the string to bytes and each byte is processed by the hash function
	return h.Sum32()       // Returns the 32-bit hash value
}

// Double hashing done to generate multiple hash values and prevent collision
func doubleHashing(input string, filterLength int, noOfHashes int) []int {
	hashVals := make([]int, noOfHashes)
	hash1 := int(hashFunc(input) % uint32(filterLength))
	hash2 := int((hashFunc(input+"golangrocks") % uint32(filterLength))) // Added a string "golangrocks" to differentiate hash
	for i := 0; i < noOfHashes; i++ {
		hashVals[i] = (hash1 + i*hash2) % filterLength // double hashing formula
	}
	return hashVals
}

// Calculates the false positive probability using the formula: (1 - e^(-kn/m))^k
// where k is the number of hash functions, n is the number of elements and m is the size of Bloom Filter
func falsePositiveProbability(filterLength int, noOfElements int, noOfHashes int) float64 {
	if noOfElements == 0 {
		return 0
	}
	power := -float64(noOfHashes*noOfElements) / float64(filterLength)
	return math.Pow(1-math.Exp(power), float64(noOfHashes))
}

func main() {
	// The inputs are not validated as
	// the focus is on the Bloom Filter implementation.
	var filterLength, totalElements int
	fmt.Print("Enter Bloom Filter length: ")
	fmt.Scanln(&filterLength)
	fmt.Print("Enter the no of elements to add (an estimate is fine): ")
	fmt.Scanln(&totalElements)

	bf := newBloomFilter(filterLength, totalElements)

	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add an item")
		fmt.Println("2. Check an item")
		fmt.Println("3. Show false positive probability")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("\nEnter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter item to add: ")
			var input string
			fmt.Scanln(&input)
			bf.addInput(input)
			fmt.Println("Item added successfully!")

		case 2:
			fmt.Print("Enter item to check: ")
			var input string
			fmt.Scanln(&input)

			if bf.checkInput(input) {
				fmt.Println("Item might be in the set (possible true or false positive).")
			} else {
				fmt.Println("Item is not in the set.")
			}

		case 3:
			probability := falsePositiveProbability(bf.filterLength, bf.noOfElements, bf.noOfHashes)
			fmt.Printf("False positive probability: %.6f\n", probability)

		case 4:
			fmt.Println("Exited.")
			return

		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

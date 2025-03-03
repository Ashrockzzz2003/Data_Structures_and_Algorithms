package main

import "fmt"

// ArraySize is the size of the hash table array
const ArraySize = 7

// hashTable represents the hash table with an array of buckets
type hashTable struct {
	array [ArraySize]*bucket
}

// bucket represents a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// bucketNode represents a node in the linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

// insertKey adds a key to the hash table
func (h *hashTable) insertKey(key string) bool {
	index := hash(key)
	return h.array[index].insert(key)
}

// searchKey returns true if the key exists in the hash table
func (h *hashTable) searchKey(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// deleteKey removes a key from the hash table
func (h *hashTable) deleteKey(key string) bool {
	index := hash(key)
	return h.array[index].delete(key)
}

// insert adds a node with the given key to the bucket
func (b *bucket) insert(k string) bool {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
		return true
	}
	return false
}

// search returns true if the key exists in the bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete removes a node with the given key from the bucket
func (b *bucket) delete(k string) bool {
	if b.head == nil {
		return false
	}

	if b.head.key == k {
		b.head = b.head.next
		return true
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
			return true
		}
		previousNode = previousNode.next
	}
	return false
}

// hash computes a hash value by summing ASCII values of characters and taking modulo ArraySize.
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *hashTable {
	result := &hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	// Initialize the hash table
	hashTable := Init()

	fmt.Println("=== Testing Normal Operations ===")

	// Test case 1: Insert keys
	keys := []string{"Midhunan", "Varun", "Joshua", "Aadesh", "Arjun", "Nishanth", "Akshar"}
	for _, key := range keys {
		success := hashTable.insertKey(key)
		if success {
			fmt.Printf("Inserted: %s\n", key)
		} else {
			fmt.Printf("Failed to insert (duplicate): %s\n", key)
		}
	}

	// Test case 2: Search for a key
	testKey := "Varun"
	if hashTable.searchKey(testKey) {
		fmt.Printf("Search success: %s found in the hash table.\n", testKey)
	} else {
		fmt.Printf("Search failure: %s not found in the hash table.\n", testKey)
	}

	// Test case 3: Delete a key
	if hashTable.deleteKey(testKey) {
		fmt.Printf("Delete success: %s removed from the hash table.\n", testKey)
	} else {
		fmt.Printf("Delete failure: %s not found in the hash table.\n", testKey)
	}

	// Verify deletion
	if hashTable.searchKey(testKey) {
		fmt.Printf("Post-deletion search failure: %s still exists in the hash table.\n", testKey)
	} else {
		fmt.Printf("Post-deletion search success: %s is not in the hash table.\n", testKey)
	}

	fmt.Println("\n=== Testing Edge Cases ===")

	// Edge case 1: Insert duplicate keys
	duplicateKey := "Midhunan"
	success := hashTable.insertKey(duplicateKey)
	if !success {
		fmt.Printf("Duplicate insert prevented: %s\n", duplicateKey)
	}

	// Edge case 2: Delete a non-existent key
	nonExistentKey := "Anush"
	if !hashTable.deleteKey(nonExistentKey) {
		fmt.Printf("Delete failure: %s does not exist in the hash table.\n", nonExistentKey)
	}

	// Edge case 3: Search for a non-existent key
	if !hashTable.searchKey(nonExistentKey) {
		fmt.Printf("Search success: %s is not in the hash table.\n", nonExistentKey)
	}

	// Edge case 4: Insert an empty string
	emptyKey := ""
	if hashTable.insertKey(emptyKey) {
		fmt.Printf("Insert success: Empty key inserted.\n")
	} else {
		fmt.Printf("Insert failure: Empty key was not inserted.\n")
	}

	// Edge case 5: Insert keys with same hash value
	// With ArraySize = 7, "A" and "H" hash to the same index
	collidingKeys := []string{"A", "H"}
	for _, key := range collidingKeys {
		hashTable.insertKey(key)
	}
	fmt.Println("Colliding keys 'A' and 'H' inserted successfully.")

	// Verify collision handling
	for _, key := range collidingKeys {
		if hashTable.searchKey(key) {
			fmt.Printf("Search success: %s found in the hash table.\n", key)
		} else {
			fmt.Printf("Search failure: %s not found in the hash table.\n", key)
		}
	}
}

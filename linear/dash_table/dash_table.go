// Dash Table -> Array of hashtables(mini).
package main

import (
	"errors"
	"fmt"
)

type Slot struct {
	key   string
	value any
	used  bool
}

type Bucket struct {
	slots []*Slot // Each bucket has 14 slots.
}

// Hash table of constant size.
type Segment struct {
	buckets    []*Bucket // 56 regular buckets.
	stash      []*Bucket // 4 stash buckets.
	maxBuckets int
	maxStash   int
}

// The main dash table structure.
type DashTable struct {
	segments []*Segment
	segSize  int // Number of segments.
}

// initializes a DashTable with the given no. of segments.
func CreateDashTable(segSize int) *DashTable {
	segments := make([]*Segment, segSize)
	for i := range segments {
		buckets := make([]*Bucket, 56)
		for j := range buckets {
			buckets[j] = &Bucket{
				slots: make([]*Slot, 14),
			}
		}
		stash := make([]*Bucket, 4)
		for j := range stash {
			stash[j] = &Bucket{
				slots: make([]*Slot, 14),
			}
		}
		segments[i] = &Segment{
			buckets:    buckets,
			stash:      stash,
			maxBuckets: 56,
			maxStash:   4,
		}
	}
	return &DashTable{segments: segments, segSize: segSize}
}

// Primary Hash function.
func (dt *DashTable) hash(key string) int {
	hash := 0
	for _, char := range key {
		hash = (hash*31 + int(char)) % dt.segSize
	}
	return hash
}

// Bucket hash function to find the bucket in which the key should be inserted into.
func (dt *DashTable) bucketHash(key string, bucketCount int) int {
	hash := 0
	for _, char := range key {
		hash = (hash*17 + int(char)) % bucketCount
	}
	return hash
}

// For inserting a new key-value pair.
func (dt *DashTable) Put(key string, value any) error {
	segIndex := dt.hash(key)
	segment := dt.segments[segIndex]
	bucketIndex := dt.bucketHash(key, segment.maxBuckets)

	for i := 0; i < segment.maxBuckets; i++ {
		idx := (bucketIndex + i) % segment.maxBuckets
		bucket := segment.buckets[idx]
		if dt.insertIntoBucket(bucket, key, value) {
			return nil
		}
	}

	// Try stash buckets.
	for _, bucket := range segment.stash {
		if dt.insertIntoBucket(bucket, key, value) {
			return nil
		}
	}

	// Add new stash bucket if needed.
	if len(segment.stash) < segment.maxStash {
		newBucket := &Bucket{slots: make([]*Slot, 14)}
		dt.insertIntoBucket(newBucket, key, value)
		segment.stash = append(segment.stash, newBucket)
		return nil
	}

	// Split segment if stash is full.
	dt.splitSegment(segIndex)
	return dt.Put(key, value)
}

// for inserting a key-value pair into a bucket.
func (dt *DashTable) insertIntoBucket(bucket *Bucket, key string, value any) bool {
	for i, slot := range bucket.slots {
		if slot == nil || !slot.used {
			bucket.slots[i] = &Slot{key: key, value: value, used: true}
			return true
		}
		if slot.used && slot.key == key {
			slot.value = value
			return true
		}
	}
	return false
}

// To get the value of the given key.
func (dt *DashTable) Get(key string) (any, error) {
	segIndex := dt.hash(key)
	segment := dt.segments[segIndex]
	bucketIndex := dt.bucketHash(key, segment.maxBuckets)

	for i := 0; i < segment.maxBuckets; i++ {
		bucket := segment.buckets[(bucketIndex+i)%segment.maxBuckets]
		for _, slot := range bucket.slots {
			if slot != nil && slot.used && slot.key == key {
				return slot.value, nil
			}
		}
	}

	for _, bucket := range segment.stash {
		for _, slot := range bucket.slots {
			if slot != nil && slot.used && slot.key == key {
				return slot.value, nil
			}
		}
	}

	return nil, errors.New("key not found")
}

// To split a segment into two and redistributes keys.
func (dt *DashTable) splitSegment(segIndex int) {
	oldSegment := dt.segments[segIndex]
	newSegment := &Segment{
		buckets:    make([]*Bucket, oldSegment.maxBuckets),
		stash:      []*Bucket{},
		maxBuckets: oldSegment.maxBuckets,
		maxStash:   oldSegment.maxStash,
	}

	for i := range newSegment.buckets {
		newSegment.buckets[i] = &Bucket{slots: make([]*Slot, 14)}
	}

	for _, bucket := range oldSegment.buckets {
		for _, slot := range bucket.slots {
			if slot != nil && slot.used {
				newSegIndex := dt.hash(slot.key)
				if newSegIndex != segIndex {
					dt.Put(slot.key, slot.value)
					slot.used = false
				}
			}
		}
	}

	for _, bucket := range oldSegment.stash {
		for _, slot := range bucket.slots {
			if slot != nil && slot.used {
				newSegIndex := dt.hash(slot.key)
				if newSegIndex != segIndex {
					dt.Put(slot.key, slot.value)
					slot.used = false
				}
			}
		}
	}

	dt.segments = append(dt.segments, newSegment)
	oldSegment.stash = []*Bucket{}
}

// To print the contents of the DashTable.
func (dt *DashTable) PrintTable() {
	for i, segment := range dt.segments {
		fmt.Printf("Segment %d:\n", i)
		for j, bucket := range segment.buckets {
			fmt.Printf("  Bucket %d:\n", j)
			for k, slot := range bucket.slots {
				if slot != nil && slot.used {
					fmt.Printf("    Slot %d: {key: %s, value: %v}\n", k, slot.key, slot.value)
				} else {
					fmt.Printf("    Slot %d: Empty\n", k)
				}
			}
		}
		fmt.Println("  Stash:")
		for j, bucket := range segment.stash {
			fmt.Printf("    Stash Bucket %d:\n", j)
			for k, slot := range bucket.slots {
				if slot != nil && slot.used {
					fmt.Printf("      Slot %d: {key: %s, value: %v}\n", k, slot.key, slot.value)
				} else {
					fmt.Printf("      Slot %d: Empty\n", k)
				}
			}
		}
	}
}

func main() {
	dt := CreateDashTable(4)
	dt.Put("apple", 5)
	dt.Put("apple", 6)
	dt.Put("banana", 10)
	dt.Put("grape", 15)
	dt.Put("orange", 20)
	dt.Put("pineapple", 25)

	fmt.Println("\nDash Table:")
	dt.PrintTable()
}

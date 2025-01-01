package main

import (
	"fmt"
	"math"
	"math/rand"
)
// next allows each node in the skip list to have multiple pointers to other nodes
// at different levels, supporting the layered structure of the skip list.
type Node struct {
	value int
	next  []*Node 
} 				  

type skipList struct {
	head         *Node
	noOfElements int
	maxNoOfLevel int
}

func coinToss(maxNoOfLevel int) int {
	level := 0
	for level < maxNoOfLevel && rand.Intn(2) == 1 {
		level++
	}
	return level
}

func newSkipList() *skipList {
	// Initial maxNoOfLevel can be set to 1 since there are no elements initially.
	initialMaxLevel := 1
	head := &Node{
		value: 0,
		next:  make([]*Node, initialMaxLevel),
	}

	return &skipList{
		head:         head,
		noOfElements: 0,
		maxNoOfLevel: initialMaxLevel,
	}
}

func (sl *skipList) insert(value int) {
	// Calculate max level based on current size.
	newMaxLevel := int(math.Floor(math.Log2(float64(sl.noOfElements + 1))))
	if newMaxLevel < 1 {
		newMaxLevel = 1
	}

	newLevel := coinToss(newMaxLevel)
	newNode := &Node{
		value: value,
		next:  make([]*Node, newLevel+1),
	}

	// Expand head's next slice only if needed for new node.
	if len(sl.head.next) < newLevel+1 {
		newNext := make([]*Node, newLevel+1)
		copy(newNext, sl.head.next)
		sl.head.next = newNext
	}

	current := sl.head
	update := make([]*Node, newLevel+1)
	for i := newLevel; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	for i := 0; i <= newLevel; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}

	sl.noOfElements++
	sl.maxNoOfLevel = newMaxLevel
}

func (sl *skipList) remove(value int) {
	if sl.noOfElements == 0 {
		fmt.Println("Cannot remove from empty list")
		return
	}

	current := sl.head
	update := make([]*Node, sl.maxNoOfLevel)

	// Traverse the list to find the position of the node to be deleted.
	for i := sl.maxNoOfLevel - 1; i >= 0; i-- {
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
		update[i] = current
	}

	current = current.next[0] // Move to the node to be deleted.
	if current == nil || current.value != value {
		fmt.Println("Element", value, "does not exist!")
		return
	}

	for i := 0; i < sl.maxNoOfLevel; i++ {
		if update[i] != nil && i < len(update[i].next) {
			update[i].next[i] = current.next[i]
		}
	}
	sl.noOfElements--

	// Recalculate max level and trim unused levels.
	if sl.noOfElements == 0 {
		sl.head.next = make([]*Node, 1)
		sl.maxNoOfLevel = 1
	} else {
		newMaxLevel := int(math.Log2(float64(sl.noOfElements + 1)))
		if newMaxLevel < 1 {
			newMaxLevel = 1
		}
		sl.maxNoOfLevel = newMaxLevel

		if len(sl.head.next) > sl.maxNoOfLevel {
			sl.head.next = sl.head.next[:sl.maxNoOfLevel]
		}
	}

	fmt.Println("Element", value, "deleted successfully.")
}

func (sl *skipList) search(value int) bool {
	isEmpty := true
	for i := 0; i < sl.maxNoOfLevel; i++ {
		if sl.head.next[i] != nil {
			isEmpty = false
			break
		}
	}
	if isEmpty {
		fmt.Println("Element", value, "not found in empty list.")
		return false
	}

	current := sl.head

	for i := len(sl.head.next) - 1; i >= 0; i-- {
		// Keep moving forward at the current level
		//if the next node's value is less than the search value.
		for current.next[i] != nil && current.next[i].value < value {
			current = current.next[i]
		}
	}

	current = current.next[0]

	if current != nil && current.value == value {
		fmt.Println("Element", value, "found.")
		return true
	} else {
		fmt.Println("Element", value, "not found.")
		return false
	}
}

func (sl *skipList) display() {
	isEmpty := true
	for i := 0; i < sl.maxNoOfLevel; i++ {
		if sl.head.next[i] != nil {
			isEmpty = false
			break
		}
	}
	if isEmpty {
		fmt.Println("Skip List: (empty)")
		return
	}
	fmt.Println("Skip List:")

	for i := len(sl.head.next) - 1; i >= 0; i-- {
		current := sl.head.next[i]

		fmt.Printf("Level %d: ", i)

		for current != nil {
			fmt.Printf("%d ", current.value)
			current = current.next[i]
		}
		fmt.Println()
	}
}

func main() {
	skipList := newSkipList()

	// Test empty list.
	skipList.search(5)
	skipList.remove(5)
	skipList.display()

	// Test single element.
	skipList.insert(10)
	skipList.display()
	skipList.remove(10)
	skipList.display()

	// Test removing non-existent elements.
	skipList.insert(20)
	skipList.remove(10)
	skipList.remove(30)

	// Test duplicate insertions.
	skipList.insert(30)
	skipList.insert(30)
	skipList.display()

	// Test removing all elements.
	skipList.remove(20)
	skipList.remove(30)
	skipList.display()
}

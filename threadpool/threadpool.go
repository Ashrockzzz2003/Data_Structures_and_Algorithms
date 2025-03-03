package main

import (
	"fmt"
	"sync"
	"time"
)

// work is a function that takes
// no arguments and returns nothing
type work func()

// threadPool struct manages a pool of
// threads of a given size and a work queue
type threadPool struct {
	workQueue chan work      // workQueue is a channel through which work is sent
	wg        sync.WaitGroup // wg is a WaitGroup which ensures that all Goroutines are finished before the program exits
}

// newThreadPool function creates a new thread
// pool with a given number of threadCount
func newThreadPool(threadCount int) *threadPool {
	tp := &threadPool{
		workQueue: make(chan work),
	}

	tp.wg.Add(threadCount)

	for range threadCount {
		go func() {
			defer tp.wg.Done()
			for w := range tp.workQueue {
				w()
			}
		}()
	}
	return tp
}

// addWork adds a work function to the work queue
func (tp *threadPool) addWork(w work) {
	tp.workQueue <- w
}

// close closes the work queue and waits for all work to be done
func (tp *threadPool) close() {
	close(tp.workQueue)
	tp.wg.Wait()
}

func main() {
	tp := newThreadPool(4)
	var wg sync.WaitGroup

	for i := range 30 {
		wg.Add(1)
		Work := func() {
			defer wg.Done()
			time.Sleep(500 * time.Millisecond) // Simulates work being done in real life
			fmt.Printf("Work done by thread: %d\n", i)
		}
		tp.addWork(Work)
	}

	tp.close()
	wg.Wait()
}

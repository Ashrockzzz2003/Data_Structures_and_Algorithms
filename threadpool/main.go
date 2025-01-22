package main

import (
	"fmt"
	"sync"
	"time"
)

// work is a function that takes no arguments and returns nothing
type work func()

// threadpool struct manages a pool of threads of a given size and a work queue
type threadPool struct {
	workQueue chan work 		// workQueue is a channel through which work is sent
	wg        sync.WaitGroup	// wg is a WaitGroup which ensures that all Goroutines are finished before the program exits
}

// newThreadPool function creates a new thread pool with a given number of threadCount
func newThreadPool(threadCount int) *threadPool {
	tp := &threadPool{
		workQueue: make(chan work),
	}

	tp.wg.Add(threadCount)

	for i := 0; i < threadCount; i++ {
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

func main() {
	tp := newThreadPool(4)

	for i := 0; i < 30; i++ {
		Work := func() {
			time.Sleep(500 * time.Millisecond) // Simulates work being done in real life
			fmt.Println("Work done by thread")
		}
		tp.addWork(Work)
	}

	close(tp.workQueue)

	tp.wg.Wait()

}
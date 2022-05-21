// This sample program demonstrates how to use a mutex
// to define critical sections of code that need synchronous
// access.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int

	wg    sync.WaitGroup
	mutex sync.Mutex // sync package has lots of good stuff...
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Printf("Final Counter: %d\n", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock() // this is much slower than using atomic package
		{ // parenthesis just to emphasize below block, no functional cause to add them here.
			value := counter
			// Yield the thread and be placed back in queue.
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
		// Release the lock and allow any
		// waiting goroutine through.
	}
}

// This sample program demonstrates how to use the atomic
// package to provide safe access to numeric types.
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Atomic increment
		// why Atomic increment and not just using a mutex ?
		atomic.AddInt64(&counter, 1)
		// Yield the thread and be placed back in queue.
		runtime.Gosched()
	}
}

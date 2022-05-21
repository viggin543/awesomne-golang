// This sample program demonstrates how to create race
// conditions in our programs. We don't want to do this.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter is a variable incremented by all goroutines.
	counter int

	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter()
	go incCounter()

	wg.Wait()
	fmt.Println("Final Counter:", counter) // not 4 !!
	// this is an example of communicating by sharing memory
	// channels allow sharing memory by communicating.. ( TBD )
}

// incCounter increments the package level counter variable.
func incCounter() {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// Capture the value of Counter.
		value := counter
		// Yield the thread and be placed back in queue.
		runtime.Gosched() // yield without blocking go routine

		// Increment our local value of Counter.
		value++

		// Store the value back into Counter.
		counter = value
	}
}

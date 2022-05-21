// This sample program demonstrates how to use the atomic
// package functions Store and Load to provide safe access
// to numeric types.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown is a flag to alert running goroutines to shut down.
	shutdown int64
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	go doWork("A")
	go doWork("B")

	// Give the goroutines time to run.
	time.Sleep(1 * time.Second)

	// Safely flag it is time to shutdown.
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	wg.Wait()
}

// doWork simulates a goroutine performing work and
// checking the Shutdown flag to terminate early.
func doWork(name string) {
	// Schedule the call to Done to tell main we are done.
	defer wg.Done()

	for { // who needs while(true) { ...
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// Do we need to shutdown.
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}

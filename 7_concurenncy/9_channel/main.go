// This sample program demonstrates how to use an unbuffered
// channel to simulate a relay race between four goroutines.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 														DISCUSS SHORTLY
//										SHARING BY COMMUNICATING vs COMMUNICATING BY SHARING
func main() {
	baton := make(chan int) // creating an unbuffered channel ( no buffer, every write will block until someone reads )

	wg.Add(1)

	go Runner(baton)

	baton <- 1 // writing to a channel

	wg.Wait()
}

func Runner(baton chan int) {
	var newRunner int
	runner := <-baton // reading from a channel and declaring a variable in one line

	fmt.Printf("Runner %d Running With Baton\n", runner)

	// New runner to the line.
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track.
	time.Sleep(100 * time.Millisecond)

	// Is the race over.
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// Exchange the baton for the next runner.
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)

	baton <- newRunner
}

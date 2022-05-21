// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

//										 shortly discuss:
//											 - Concurrency versus parallelism
//														- Concurrency can out reform parallelism  ( GoLang's mantra )
// 											 - threads and processes
//											 - os scheduler
//											 - user space \ kernel space context switching
//											 - the 10K problem, MAX THREADS an os can run
//											 - go routines, hundreds of thousands can be scheduled per thread efficiently
//												- logical processors
//												- go runtime scheduler (assigns go routines to threads) - in user space !
//													- what is the benefit of a user space scheduler ?

// 										BLOCKING CODE ( SYSCALL )
//		  disk access for example.
//		  the thread and goroutine are detached from the logical processor and the thread continues to block waiting for the syscall to return.
//		  In the meantime, there’s a logical processor without a thread.
//		  So the scheduler creates a new thread and attaches it to the logical processor.
//		  Then the scheduler will choose another goroutine from the local run queue for execution.
//		  Once the syscall returns, the goroutine is placed back into a local run queue, and the thread is put aside for future use.

// 										BLOCKING CODE  (network I/O call)
//		the goroutine is detached from the logical processor and moved to the runtime integrated network poller.
//		Once the poller indicates a read or write operation is ready,
//		the goroutine is assigned back to a logical processor to handle the operation. (
//		the thread is not blocked and can run some other go routine ( steal from the queue of another thread )
func main() {
	// Allocate 1 logical processor for the scheduler to use.
	runtime.GOMAXPROCS(1)

	// go scheduler does lots of fancy stuff.
	// handles unix epoll system calls (non-blocking io)
	// detects deadlocks
	// work stealing thread scheduling
	//   READ ME !!!!!!!!!  https://morsmachine.dk/go-scheduler
	//   AND ME !!! https://morsmachine.dk/netpoller

	// wg is used to wait for the program to finish.
	// Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	go func() {
		defer wg.Done() // why defer ? why not calling wg.Done() after printing the alphabet ?
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ { // what does 'a'+26 ?
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait() // what will happen if we remove this line ?

	fmt.Println("\nTerminating Program")
}

// above code has three goroutines
// Capital letters are printed first because the first goroutine finishes super fast
// before the scheduler swaps it out for the second goroutine

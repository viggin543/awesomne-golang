// This sample program demonstrates how the goroutine scheduler
// will time slice goroutines on a single thread.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(2)

	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime displays prime numbers for the first 5000 numbers.
func printPrime(prefix string) {
	defer wg.Done()

next:                                        // a label wtf ?? ( exaple with no labels is given below )
	for outer := 2; outer < 50000; outer++ { // my m1 silicon is so fast that 5K first primes don't cause scheduler to time slice the go routine !!!!
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
				//goto next  //--> OMG.. goto ???
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func printPrimesNoLabels(prefix string) {
	defer wg.Done()
	for outer := 2; outer < 5000; outer++ {
		found := false
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				found = true
				break
			}
		}
		if found {
			continue
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

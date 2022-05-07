// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
	"github.com/viggin543/awesomne-golang/chapter5/4_public_private/counters"
)

// main is the entry point for the application.
func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	counter := counters.AlertCounter(10) // change AlertCounter to be private

	fmt.Printf("Counter: %d\n", counter)
}

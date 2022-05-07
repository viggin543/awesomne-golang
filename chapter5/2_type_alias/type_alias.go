// Sample program to show how you can't always get the
// address of a value.
package main

import "fmt"

// duration is a type with a base type of int.
type duration int

// format pretty-prints the duration value.
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	//[idea tip] alt+cmd+n inline a variable
	d := duration(42)
	d.pretty()
	//d = int(1) -> will this compile ?

	// try to inline d
	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}

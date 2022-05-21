package main

import "fmt"

// duration is a type with a base type of int.
type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

func main() {
	//[idea tip] alt+cmd+n inline a variable
	d := duration(42)
	d.pretty()
	//d = int(1)  //will this compile ?
	//duration(42).pretty() //will this compile ?
}

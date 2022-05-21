package main

import (
	"fmt" // imported from stdlib ( GOPATH )
)

//package variable
var greet = "hello"

func greeting() string {
	return greet
}

// init function is implicitly invoked when package is imported
// if a main exists in this package, init is still invoked first
func init() {
	greet = "banana"
}

// the main method must always be inside a main package
func main() {
	fmt.Println("hello hello!", greeting())
}

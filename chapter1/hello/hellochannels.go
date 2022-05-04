package main

import (
	"fmt"
)

func greeting() string {
	return "hello"
}

func main() {
	fmt.Println("hello hello!", greeting())
}

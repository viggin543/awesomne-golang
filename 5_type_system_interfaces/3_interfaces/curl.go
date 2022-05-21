// Sample program to show how to write a simple version of curl using
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

type Writer interface { // interfaces with one method only should have an 'er' suffix
	Write(p []byte) (n int, err error)
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

func main() {
	response, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// io.Copy receives interface params
	_, _ = io.Copy(os.Stdout, response.Body)      // response.Body is a struct pointer that implements the Reader and Closer interfaces
	if err := response.Body.Close(); err != nil { // body.Close will call the method of the struct implementing Closer interface
		fmt.Println(err)
	}
	// notice golang has no "implements" keyword
	// this is called duck typing
	// this makes golang much more flexible than Java,Kotlin,Scala,C# and other compiled statically typed langs
	// quack quack !
}

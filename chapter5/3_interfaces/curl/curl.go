// Sample program to show how to write a simple version of curl using
package main

import (
	"bytes"
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

func main() {
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	// io.Copy receives interface params
	_, _ = io.Copy(os.Stdout, r.Body)      // r.Body is a struct pointer that implements the Reader and Closer interfaces
	if err := r.Body.Close(); err != nil { // body.Close will call the method of the struct implementing Closer interface
		fmt.Println(err)
	}
	// notice golang has no "implements" keyword
	// this is called duck typing
	// this makes golang much more flexible than Java,Kotlin,Scala,C# and other "type safe langs"
	// also notice is statically typed ( unlike Java... )
}

func copySomethingElse() {
	var b bytes.Buffer // why b is not nil ?
	b.Write([]byte("Hello"))
	_, _ = fmt.Fprintf(&b, "World!") // why Fprintf receives a ref to b ?
	_, _ = io.Copy(os.Stdout, &b)    // why io.Copy receives a ref to b ?
}

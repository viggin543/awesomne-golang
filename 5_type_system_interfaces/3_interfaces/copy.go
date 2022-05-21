package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// io.Copy receives Reader and Writer interface implementors
func copyExample() {
	var b bytes.Buffer // why b is not nil ?
	b.Write([]byte("Hello"))
	_, _ = fmt.Fprintf(&b, "World!") // why Fprintf receives a ref to b ?
	_, _ = io.Copy(os.Stdout, &b)    // why io.Copy receives a ref to b ?
	fmt.Println(b.String())
}

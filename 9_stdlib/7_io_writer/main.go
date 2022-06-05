// Sample program to show how different functions from the
// standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	
	fmt.Fprintf(&b, "World!")

	b.WriteTo(os.Stdout)
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// cross package reference
	// b4 go modules, this would have been -> GOPATH/src/github.com/goinaction/code/chapter3/words.
	"github.com/viggin543/awesomne-golang/code/chapter3/words"
)

func main() {
	filename := os.Args[1] // [idea tip] -> run configuration

	contents, err := ioutil.ReadFile(filename) // file to string
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents) // (primitive type) casting

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}

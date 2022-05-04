package main

import (
	"fmt"
	"io/ioutil"
	"os"

	// cross package reference
	"github.com/viggin543/awesomne-golang/code/chapter3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents) // type casting

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}

package main

import (
	//importing from stdlib vs other packages
	"log"
	"os"
	// side effect import ( triggers package init function )
	_ "github.com/viggin543/awesomne-golang/2_example_program/sample/matchers"
	"github.com/viggin543/awesomne-golang/2_example_program/sample/search"
)

func init() {
	// logging package from standard lib
	log.SetOutput(os.Stdout)
}

func main() {
	// [ida tips] -> step into, back caret, forward caret
	search.Run("and")
}

package main

import (
	"log"
	"os"

	_ "github.com/viggin543/awesomne-golang/code/chapter2/sample/matchers"
	"github.com/viggin543/awesomne-golang/code/chapter2/sample/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("and")
}

// - packages
// - import
// - main
// - init ( implicitly called )
// _ side effect import
// - log package from stdlib
// - - importing from stdlib vs other packages
// - - GOROOT and GOPATH env vars
// - - go  env

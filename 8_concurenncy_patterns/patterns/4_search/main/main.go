// This sample program demonstrates how to implement a pattern for
// concurrent requesting results from different systems and either
// wait for all the results to return or just the first one.
package main

import (
	search "github.com/viggin543/awesomne-golang/8_concurenncy_patterns/patterns/4_search"
	"log"
)

func main() {

	results := search.Submit(
		"golang",
		search.OnlyFirst,
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}

	// This time we want to wait for all the results.
	results = search.Submit(
		"golang",
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}
}

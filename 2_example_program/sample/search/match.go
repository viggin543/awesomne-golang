package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) { // write channel
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil || searchResults == nil {
		log.Println(err)
		return
	}

	for _, result := range searchResults {
		results <- result // writing to a channel. this blocks until some-one reads !!
	}
}

func Display(results <-chan *Result) { // read/write channels
	// - iterating over a channel ( this blocks !! until some-one writes to the chanel ! )
	for result := range results {
		// golang string formatting, yes it's ugly... yes everybody hates it
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}

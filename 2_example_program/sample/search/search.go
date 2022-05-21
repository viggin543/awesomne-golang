package search

import (
	"log"
	"sync"
)

// private package variable
// memory allocation
// a map
// zero values in golang & pointer types
var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	// return error pattern
	// error handling
	// public method
	// [idea tip] -> inline struct preview ( caret on feeds var cmd+j, click on Feed definition, then string. use keyboard left key to go back from URI to Feed )
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// channels
	// := syntactic sugar
	// pointer type
	results := make(chan *Result)

	// counting semaphore
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	// range keyword for iterations over arrays,maps and slices
	for _, feed := range feeds {
		// map access, ( bool return arg pattern )
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// calling a goroutine, concurrent execution
		// Matcher interface
		go func(matcher Matcher, feed *Feed) { // anonymous "lambda"
			// a closure
			//[idea tip] -> param list ( cmd + p )
			//[idea tip] -> preview definitions of Match args ( the args of a method often suggest what it does )
			Match(matcher, feed, searchTerm, results) // suggest a more OOP way to doit
			waitGroup.Done()                          // in the closure waitGroup is the same variable as outside the closure ( not a copy and not a pointer )
		}(matcher, feed) // passing an args to the go routine to prevent a race
	}

	go func() {
		waitGroup.Wait()
		close(results)
	}()

	Display(results)
}

//Register [idea tip] -> inline find usages (cmd + b)
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}

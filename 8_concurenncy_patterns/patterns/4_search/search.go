// Package search : search.go manages the searching of results  against Google, Yahoo and Bing.
package search

import "log"

type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

func OnlyFirst(s *searchSession) { s.first = true }

func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	// Perform the searches concurrently. Using a map because
	// it returns the searchers in a random order every time.
	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result

	for search := 0; search < len(session.searchers); search++ {
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		log.Println("search : Submit : Info : Waiting For Results...")
		result := <-session.resultChan

		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		results = append(results, result...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}

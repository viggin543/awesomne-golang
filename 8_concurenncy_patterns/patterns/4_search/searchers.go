package search

import (
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type google struct{}

func (g google) Search(term string, results chan<- []Result) {
	log.Printf("Google : Search : Started : search term[%s]\n", term)

	var r []Result

	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	r = append(r, Result{
		Engine:      "Google",
		Title:       "The Go Programming Language",
		Description: "The Go Programming Language",
		Link:        "https://golang.org/",
	})

	log.Printf("Google : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

type bing struct{}

func (b bing) Search(term string, results chan<- []Result) {
	log.Printf("Bing : Search : Started : search term [%s]\n", term)

	var r []Result

	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	r = append(r, Result{
		Engine:      "Bing",
		Title:       "A Tour of Go",
		Description: "Welcome to a tour of the Go programming language.",
		Link:        "http://tour.golang.org/",
	})

	log.Printf("Bing : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

type yahoo struct{}

func (y yahoo) Search(term string, results chan<- []Result) {
	log.Printf("Yahoo : Search : Started : search term [%s]\n", term)

	var r []Result

	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	r = append(r, Result{
		Engine:      "Yahoo",
		Title:       "Go Playground",
		Description: "The Go Playground is a web service that runs on golang.org's servers",
		Link:        "http://play.golang.org/",
	})

	log.Printf("Yahoo : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

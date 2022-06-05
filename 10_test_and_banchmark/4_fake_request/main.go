// This sample code implement a simple web service.
package main

import (
	"github.com/viggin543/awesomne-golang/10_test_and_banchmark/listing17/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.Routes()
	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
	// what's the difference between :4000 and 127.0.0.1:4000 ?
	// this is important !
}

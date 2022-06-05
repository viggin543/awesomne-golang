package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "WORLD!")
	})
	http.HandleFunc("/hack", func(w http.ResponseWriter, r *http.Request) {
		reverseProxy("https://dummyjson.com/products/1").ServeHTTP(w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("./static"))) // make sure to set correct working directory when running this
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func reverseProxy(u string) *httputil.ReverseProxy {
	address, _ := url.Parse(u)
	p := httputil.NewSingleHostReverseProxy(address)
	p.Director = func(request *http.Request) {
		log.Println("middle man at work - hacking request")
		request.Host = address.Host
		request.URL.Scheme = address.Scheme
		request.URL.Host = address.Host
		request.URL.Path = address.Path
		if address.RawQuery != "" {
			request.URL.RawQuery = address.RawQuery
		}
	}
	p.ModifyResponse = func(response *http.Response) error {
		log.Println("middle man at work - hacking response")
		if response.StatusCode > 399 {
			defer response.Body.Close()
			all, _ := ioutil.ReadAll(response.Body)
			log.Println(string(all))
			response.Body = ioutil.NopCloser(bytes.NewReader([]byte(fmt.Sprintf(string(all)))))
		}
		return nil
	}
	return p
	// JUST TRY TODO THE SAME USING ONLY THE STDLIB OF ANY OTHER PROG LANGUAGE
}

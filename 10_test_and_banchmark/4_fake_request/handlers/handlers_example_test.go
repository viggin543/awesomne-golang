package handlers_test

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func ExampleSendJSON() { // IDEA CAN run this because this function name starts with Example ( a convention )
	r, _ := http.NewRequest("GET", "/sendjson", nil)
	w := httptest.NewRecorder() // NOTICE THERE IS NO CALL TO ListenAndServe
	http.DefaultServeMux.ServeHTTP(w, r)
	var u struct {
		Name  string
		Email string
	}

	if err := json.NewDecoder(w.Body).Decode(&u); err != nil {
		log.Println("ERROR:", err)
	}

	fmt.Println(u)
	// Output:
	// {Bill bill@ardanstudios.com}
}

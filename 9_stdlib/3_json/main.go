package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	gResponse struct {
		ID                 int      `json:"id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		Price              int      `json:"price"`
		Discountpercentage float64  `json:"discountPercentage"`
		Rating             float64  `json:"rating"`
		Stock              int      `json:"stock"`
		Brand              string   `json:"brand"`
		Category           string   `json:"category"`
		Thumbnail          string   `json:"thumbnail"`
		Images             []string `json:"images"`
	}
)

func main() {
	uri := "https://dummyjson.com/products/1"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()

	var gr gResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(gr)

	pretty, err := json.MarshalIndent(gr, "", "    ") // pretty print
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(pretty))
}

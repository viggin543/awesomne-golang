package main

import (
	"database/sql"
	"fmt"

	// side effect import example
	_ "github.com/viggin543/awesomne-golang/code/chapter3/dbdriver/postgres"
)

func main() {
	db, err := sql.Open("postgres", "mydb")
	fmt.Println(db, err)
	err = db.Ping() // this will call our driver implementation
}

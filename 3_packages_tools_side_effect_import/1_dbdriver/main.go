package main

import (
	"database/sql"
	"fmt"
	// side effect import example, dependencies inversion example
	_ "github.com/viggin543/awesomne-golang/code/chapter3/1_dbdriver/postgres"
)

// run with debugger
func main() {
	db, err := sql.Open("postgres", "mydb") // place a breakpoint inside Open ( from 1_dbdriver/postgres" package )
	fmt.Println(db, err)
	err = db.Ping() // this will call our driver implementation
}

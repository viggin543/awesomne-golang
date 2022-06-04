// an example on how to use a channel to
// monitor the amount of time the program is running and terminate
// the program if it runs too long.
package main

import (
	runner "github.com/viggin543/awesomne-golang/8_concurenncy_patterns/patterns/1_runner"
	"log"
	"os"
	"time"
)

const timeout = 3 * time.Second

// try running the main program and interrupt it ( ctrl+c ) b4 timeout hits
func main() {
	log.Println("Starting work.")

	r := runner.New(timeout)

	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Working... - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}

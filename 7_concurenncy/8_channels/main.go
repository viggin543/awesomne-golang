// This sample program demonstrates how to use an unbuffered
// channel to simulate a game of tennis between two goroutines.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Create an unbuffered channel. ( unbuffered -> no spare room, every write will block until someone reads from the channel)
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1

	wg.Wait()
}

// player simulates a person playing the game of tennis.
func player(name string, court chan int) {
	defer wg.Done()

	for {
		// Wait for the ball to be hit back to us.
		ball, ok := <-court
		if !ok {
			// If the channel was closed we won.
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// Pick a random number and see if we miss the ball.
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// Close the channel to signal we lost.
			close(court)
			return
		}

		// Display and then increment the hit count by one.
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// Hit the ball back to the opposing player.
		court <- ball
	}
}

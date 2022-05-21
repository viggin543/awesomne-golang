package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// no password login
// you get an email with a magic link
// clicking on the magic link should log you in
// browser preforms long polling waiting for you to click on the email

func main() {
	verification, _ := LongPollingAwaitEmailVerification(context.Background(), "session_idd")
	fmt.Println("welcome !!")
	fmt.Println(verification)
}

func LongPollingAwaitEmailVerification(ctx context.Context, sessionId string) (*AwaitLoginResponse, error) {
	c, _ := context.WithDeadline(ctx, time.Now().Add(30*time.Second))
	ticker := ImmediateTicker(c, time.Second)
	for {
		select { // A NEW KEYWORD
		case <-c.Done():
			fmt.Println("timed out !")
			return &AwaitLoginResponse{Status: "waiting"}, nil
		case <-ticker:
			status, token := IsEmailVerified(sessionId)
			if status != "waiting" {
				if status == "valid" {
					dropCookie(ctx, token)
				}
				return &AwaitLoginResponse{
					Status:      status,
					AccessToken: token,
				}, nil
			}
			fmt.Println("waiting for user to click email..")
		}
	}
}

func ImmediateTicker(ctx context.Context, interval time.Duration) <-chan time.Time { // return a read channel
	ret := make(chan time.Time, 1) // we create the channel, we are to close it
	ret <- time.Now()
	ticket := time.NewTicker(interval)
	go func() {
		defer close(ret)
		defer ticket.Stop()
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticket.C: // assign a variable and read from a channel
				ret <- t // write to channel
			}
		}
	}()
	return ret
}

func IsEmailVerified(sessionId string) (string, string) {
	if rand.Intn(10) == 5 {
		return "valid", "signed_token"
	}
	return "waiting", ""
}

func dropCookie(ctx context.Context, token string) {
	fmt.Println("dropping a cookie !!")
}

type AwaitLoginResponse struct {
	Status      string
	AccessToken string
}

package main

import (
	"context"
	"time"
)

// no password login
// you get an email with a magic link
// clicking on the magic link should log you in
// browser preforms long polling waiting for you to click on the email

func LongPollingAwaitEmailVerification(ctx context.Context, sessionId string) (*AwaitLoginResponse, error) {
	c, _ := context.WithDeadline(ctx, time.Now().Add(time.Second))
	ticker := ImmediateTicker(c, time.Second)
	for {
		select { // A NEW KEYWORD
		case <-c.Done():
			return &AwaitLoginResponse{Status: "waiting"}, nil
		case <-ticker:
			status, token := IsEmailVerified(sessionId)
			if status != "waiting" {
				if status == "valid" {
					dropCookie(ctx)
				}
				return &AwaitLoginResponse{
					Status:      status,
					AccessToken: token,
				}, nil
			}
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
	panic("implement me")
}

func dropCookie(ctx context.Context) {

}

type AwaitLoginResponse struct {
	Status      string
	AccessToken string
}

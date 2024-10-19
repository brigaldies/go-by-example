package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	ctx := context.Background()

	// trap Ctrl+C and call cancel on the context
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		select {
		case <-c:
			fmt.Printf("Signal received from signal channel\n")
			cancel()
		case <-ctx.Done():
			fmt.Printf("main: Context done\n")
		}
	}()

	doSomethingAwesome(ctx)
}

func doSomethingAwesome(ctx context.Context) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			fmt.Printf("doSomethingAwesome: Context done; Ticker stopped\n")
			return
		case <-ticker.C:
			fmt.Printf("Ticker\n")
		}
	}
}

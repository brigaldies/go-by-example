package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending request", i)
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	fmt.Println("Reading requests...")
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending bursty request", i)
		burstyRequests <- i
	}
	close(burstyRequests)

	fmt.Println("Reading bursty requests...")
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("Bursty request", req, time.Now())
	}
}

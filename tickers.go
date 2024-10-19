package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		defer func() {
			fmt.Println("Picker handler is done")
		}()
		for {
			select {
			case <-done:
				fmt.Println("Request to exit the ticker handler.")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
	fmt.Println("Stopping the ticker handler")
	done <- true
	time.Sleep(2 * time.Second)
	fmt.Println("Main is done")
}

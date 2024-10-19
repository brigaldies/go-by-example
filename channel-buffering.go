package main

import (
	"fmt"
	"time"
)

func main() {

	messages := make(chan string, 3)

	go func() {
		for {
			msg := <-messages
			fmt.Println(fmt.Sprintf("Read message %q", msg))
			fmt.Println(fmt.Sprintf("Processing for 10 secs..."))
			time.Sleep(10 * time.Second)
		}
	}()

	fmt.Println("Writing msg1...")
	messages <- "msg1"
	fmt.Println("Wrote msg1")

	fmt.Println("Writing msg2...")
	messages <- "msg2"
	fmt.Println("Wrote msg2")

	fmt.Println("Writing msg3...")
	messages <- "msg3"
	fmt.Println("Wrote msg3")

	fmt.Println("Wait for 60 secs")
	time.Sleep(60 * time.Second)

	//fmt.Println("Reading...")
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

	fmt.Println("Exit")
}

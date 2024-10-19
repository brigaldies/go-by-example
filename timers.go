package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Waiting for Timer 1 to fire...")
	<-timer1.C
	fmt.Println(fmt.Sprintf("Timer 1 fired, after %v seconds, %v milliseconds, %v millseconds", time.Since(start).Seconds(), time.Since(start).Milliseconds(), time.Since(start)/time.Millisecond))

	timer2 := time.NewTimer(time.Second)
	go func() {
		fmt.Println("Waiting for Timer 2 to fire...")
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}

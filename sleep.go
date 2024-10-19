package main

import (
	"fmt"
	"time"
)

func main() {
	waitTimeMilliseconds := 100
	loopCount := 100
	startTime := time.Now()
	fmt.Printf("Begin: %v\n", startTime)
	for i := 0; i < loopCount; i++ {
		iterationStart := time.Now()
		fmt.Printf("[%d] Time: %v\nWaiting...\n", i, iterationStart)
		// time.Sleep(time.Duration(h.config.ESProductFeaturesThrottleWaitMilliseconds) * time.Millisecond)
		time.Sleep(time.Duration(waitTimeMilliseconds) * time.Millisecond)
		iterationEnd := time.Now()
		iterationDuration := time.Since(iterationStart) / time.Millisecond
		fmt.Printf("[%d] Time: %v, duration=%d ms\n", i, iterationEnd, int(iterationDuration))
	}
	endTime := time.Now()
	duration := time.Since(startTime) / time.Millisecond
	fmt.Printf("Start: %v\n", startTime)
	fmt.Printf("End: %v\n", endTime)
	fmt.Printf("Duration: %d ms\n", int(duration))
}
